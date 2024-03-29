/*
 * Copyright 2023 Dgraph Labs, Inc. and Contributors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dgraphtest

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/dgraph-io/dgraph/x"
)

const (
	waitDurBeforeRetry = time.Second
)

var (
	errNotImplemented = errors.New("NOT IMPLEMENTED")
)

type Cluster interface {
	Client() (*dgo.Dgraph, error)
	AdminPost(body []byte) ([]byte, error)
	AlphasHealth() ([]string, error)
	AssignUids(num uint64) error
}

func SetupSchema(c Cluster, dbSchema string) error {
	client, err := c.Client()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	return client.Alter(ctx, &api.Operation{Schema: dbSchema})
}

func Mutate(c Cluster, rdfs string) (*api.Response, error) {
	client, err := c.Client()
	if err != nil {
		return nil, err
	}

	txn := client.NewTxn()
	defer func() { _ = txn.Discard(context.Background()) }()

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	mu := &api.Mutation{SetNquads: []byte(rdfs), CommitNow: true}
	return txn.Mutate(ctx, mu)
}

func Query(c Cluster, query string) (*api.Response, error) {
	client, err := c.Client()
	if err != nil {
		return nil, err
	}

	txn := client.NewTxn()
	defer func() { _ = txn.Discard(context.Background()) }()

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	return txn.Query(ctx, query)
}

type graphQLParams struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type graphQLResponse struct {
	Data       json.RawMessage        `json:"data,omitempty"`
	Errors     x.GqlErrorList         `json:"errors,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

func RunAdminQuery(c Cluster, params graphQLParams) ([]byte, error) {
	reqBody, err := json.Marshal(params)
	if err != nil {
		return nil, errors.Wrap(err, "error while marshalling params")
	}

	respBody, err := c.AdminPost(reqBody)
	if err != nil {
		return nil, errors.Wrap(err, "error while running admin query")
	}

	var gqlResp graphQLResponse
	if err := json.Unmarshal(respBody, &gqlResp); err != nil {
		return nil, errors.Wrap(err, "error unmarshalling GQL response")
	}
	if len(gqlResp.Errors) > 0 {
		return nil, errors.Wrapf(gqlResp.Errors, "error while running admin query")
	}
	return gqlResp.Data, nil
}

func Backup(c Cluster, forceFull bool, backupPath string) error {
	const query = `mutation backup($dst: String!, $ff: Boolean!) {
		backup(input: {destination: $dst, forceFull: $ff}) {
			response {
				code
			}
			taskId
		}
	}`
	params := graphQLParams{
		Query:     query,
		Variables: map[string]interface{}{"dst": backupPath, "ff": forceFull},
	}
	resp, err := RunAdminQuery(c, params)
	if err != nil {
		return err
	}

	var backupResp struct {
		Backup struct {
			Response struct {
				Code string `json:"code,omitempty"`
			} `json:"response,omitempty"`
			TaskID string `json:"taskId,omitempty"`
		} `json:"backup,omitempty"`
	}
	if err := json.Unmarshal(resp, &backupResp); err != nil {
		return errors.Wrap(err, "error unmarshalling backup response")
	}
	if backupResp.Backup.Response.Code != "Success" {
		return fmt.Errorf("backup failed")
	}
	return WaitForTask(c, backupResp.Backup.TaskID)
}

func WaitForTask(c Cluster, taskId string) error {
	const query = `query task($id: String!) {
		task(input: {id: $id}) {
			status
		}
	}`
	params := graphQLParams{
		Query:     query,
		Variables: map[string]interface{}{"id": taskId},
	}

	for {
		time.Sleep(waitDurBeforeRetry)

		resp, err := RunAdminQuery(c, params)
		if err != nil {
			return err
		}

		var statusResp struct {
			Task struct {
				Status string `json:"status,omitempty"`
			} `json:"task,omitempty"`
		}
		if err := json.Unmarshal(resp, &statusResp); err != nil {
			return errors.Wrap(err, "error unmarshalling status response")
		}
		switch statusResp.Task.Status {
		case "Success":
			return nil
		case "Failed", "Unknown":
			return fmt.Errorf("task failed with status: %s", statusResp.Task.Status)
		}
	}
}

func Restore(c Cluster, backupPath string, backupId string, incrFrom, backupNum int, encKey string) error {
	query := `mutation restore($location: String!, $backupId: String,
			$incrFrom: Int, $backupNum: Int, $encKey: String) {
		restore(input: {location: $location, backupId: $backupId, incrementalFrom: $incrFrom,
				backupNum: $backupNum, encryptionKeyFile: $encKey}) {
			code
			message
		}
	}`
	params := graphQLParams{
		Query: query,
		Variables: map[string]interface{}{"location": backupPath, "backupId": backupId,
			"incrFrom": incrFrom, "backupNum": backupNum, "encKey": encKey},
	}
	resp, err := RunAdminQuery(c, params)
	if err != nil {
		return err
	}

	var restoreResp struct {
		Restore struct {
			Code    string
			Message string
		}
	}
	if err := json.Unmarshal(resp, &restoreResp); err != nil {
		return errors.Wrap(err, "error unmarshalling restore response")
	}
	if restoreResp.Restore.Code != "Success" {
		return fmt.Errorf("restore failed, response: %+v", restoreResp.Restore)
	}
	return nil
}

func WaitForRestore(c Cluster) error {
loop:
	for {
		time.Sleep(waitDurBeforeRetry)

		resp, err := c.AlphasHealth()
		if err != nil {
			return err
		}
		for _, hr := range resp {
			if strings.Contains(hr, "opRestore") {
				continue loop
			}
		}
		return nil
	}
}
