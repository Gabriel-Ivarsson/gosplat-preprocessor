// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package main generated by go-bindata.// sources:
// ../embedding-files-into-app-bindata/assets/css/main.css
// ../embedding-files-into-app-bindata/assets/favicon.ico
// ../embedding-files-into-app-bindata/assets/js/main.js
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

type assetFile struct {
	*bytes.Reader
	name            string
	childInfos      []os.FileInfo
	childInfoOffset int
}

type assetOperator struct{}

// Open implement http.FileSystem interface
func (f *assetOperator) Open(name string) (http.File, error) {
	var err error
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}
	content, err := Asset(name)
	if err == nil {
		return &assetFile{name: name, Reader: bytes.NewReader(content)}, nil
	}
	children, err := AssetDir(name)
	if err == nil {
		childInfos := make([]os.FileInfo, 0, len(children))
		for _, child := range children {
			childPath := filepath.Join(name, child)
			info, errInfo := AssetInfo(filepath.Join(name, child))
			if errInfo == nil {
				childInfos = append(childInfos, info)
			} else {
				childInfos = append(childInfos, newDirFileInfo(childPath))
			}
		}
		return &assetFile{name: name, childInfos: childInfos}, nil
	} else {
		// If the error is not found, return an error that will
		// result in a 404 error. Otherwise the server returns
		// a 500 error for files not found.
		if strings.Contains(err.Error(), "not found") {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
}

// Close no need do anything
func (f *assetFile) Close() error {
	return nil
}

// Readdir read dir's children file info
func (f *assetFile) Readdir(count int) ([]os.FileInfo, error) {
	if len(f.childInfos) == 0 {
		return nil, os.ErrNotExist
	}
	if count <= 0 {
		return f.childInfos, nil
	}
	if f.childInfoOffset+count > len(f.childInfos) {
		count = len(f.childInfos) - f.childInfoOffset
	}
	offset := f.childInfoOffset
	f.childInfoOffset += count
	return f.childInfos[offset : offset+count], nil
}

// Stat read file info from asset item
func (f *assetFile) Stat() (os.FileInfo, error) {
	if len(f.childInfos) != 0 {
		return newDirFileInfo(f.name), nil
	}
	return AssetInfo(f.name)
}

// newDirFileInfo return default dir file info
func newDirFileInfo(name string) os.FileInfo {
	return &bindataFileInfo{
		name:    name,
		size:    0,
		mode:    os.FileMode(2147484068), // equal os.FileMode(0644)|os.ModeDir
		modTime: time.Time{}}
}

// AssetFile return a http.FileSystem instance that data backend by asset
func AssetFile() http.FileSystem {
	return &assetOperator{}
}

var _cssMainCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x28\xc9\xcd\x51\xa8\xe6\xe5\xe2\x4c\xcb\xcf\x2b\xd1\x4d\x4b\xcc\xcd\xcc\xa9\xb4\x52\x28\x4e\xcc\x2b\xd6\x2d\x4e\x2d\xca\x4c\xb3\xe6\xe5\xe2\xd4\x2d\x4f\x4d\xca\xce\x2c\xd1\x2d\x49\xad\x28\xd1\x2d\xce\xac\x4a\xd5\x4d\x4c\xc9\x2a\x2d\x2e\xb1\x52\x30\x34\x30\x50\x05\xab\xc8\x2d\xc6\x21\xcb\xcb\x55\x0b\x08\x00\x00\xff\xff\x32\x4c\x06\xc6\x63\x00\x00\x00")

func cssMainCssBytes() ([]byte, error) {
	return bindataRead(
		_cssMainCss,
		"css/main.css",
	)
}

func cssMainCss() (*asset, error) {
	bytes, err := cssMainCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "css/main.css", size: 99, mode: os.FileMode(438), modTime: time.Unix(1613718750, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _faviconIco = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x0b\x70\x55\xc7\x79\xde\x7b\xce\x45\x12\xb2\x90\xc4\xc3\xe6\x61\x3b\x90\xf8\x31\xc4\x19\x6c\x32\xe3\xc4\x34\xe3\xc6\x34\x6d\xed\xd4\x69\x62\x32\x49\x9a\xa6\x75\xea\x36\x33\xee\xd8\x9e\xb4\x75\xdc\x7a\xa6\xc5\x31\x02\xa6\xd3\x84\x47\x28\x6f\x3b\xc6\x3c\xcc\xeb\x9e\xbd\x08\x21\x04\x92\x28\x12\x0e\x08\x5b\x3c\xec\x02\x92\x78\x08\x21\x09\x41\x25\x24\x10\xe8\x71\xb5\xe7\xdc\xd7\xb9\xf7\xef\xfc\xff\x9e\x73\x74\x25\xdd\x97\x24\x82\x33\x1e\xce\xcc\x3f\x7b\xee\x9e\xdd\xff\xff\xf6\xdf\x7f\xff\xfd\xf7\xdf\xcb\x98\x8b\xa9\x2c\x3f\x1f\xcb\x19\xec\x15\x37\x63\x5f\x67\x8c\xcd\x98\x21\x7f\x6b\xf9\x8c\x6d\x72\x33\x36\x7b\xb6\xf5\xfb\x11\xc6\x9e\xbd\x97\xb1\x99\x8c\xb1\x7c\x6c\xc7\x64\x3d\x3d\x6e\x76\xdb\x1f\xe1\x75\xab\x42\x53\x7e\x22\x3c\x6c\x91\xf0\xb0\x02\x22\xae\x14\x08\x8f\xab\x40\xec\x64\x92\xf0\x5d\x53\xfe\x59\xec\x64\x5f\x15\x1e\xa6\x08\x4f\x7f\x7f\xdf\x7a\x96\xa1\x97\x3c\xb8\x3f\x78\xfa\x0d\x08\x9e\x5d\x08\xc1\xda\x5f\x82\x51\xf6\x65\x30\xca\x67\x41\xf0\xcc\xbf\x11\x19\x07\x1e\x07\xbd\x78\x2a\x18\x65\x8f\xb5\x0b\x4d\x7d\xad\x6f\x07\x73\x0b\xcd\x45\xfd\xbb\x17\xb3\x0c\xa3\x62\xce\xfe\xa8\xd1\x0a\xf8\x98\x6d\x25\x10\x38\xfe\x53\x88\xdc\x3a\x01\x66\x6b\x11\x11\xbe\x07\x8e\xfd\x2d\x84\x2e\x2c\x05\xff\xd1\x79\x3d\xc2\xc3\x7e\x20\x78\x06\x13\x5c\x89\xe9\xdf\x06\xd1\x50\x2f\x04\x4e\xfc\x0c\xcc\x6b\xa5\x60\x1c\xfd\x21\x74\x6c\x1e\x47\x84\xef\x58\x17\xf8\xe4\x65\x30\x6f\x54\x81\x71\xe0\x89\xea\xbe\x2d\x6c\xd2\x80\xfe\xfe\x76\x30\x6f\x1c\x81\x60\xed\x7c\x08\x9e\x5d\x0c\x2d\xeb\xc7\x41\xed\xd2\x29\x50\xbb\x74\x32\xbd\x63\x1d\x7e\xc3\x36\xa1\x73\xff\x19\x14\x1e\x36\x4f\x78\xc7\xc4\xf4\xef\x80\x70\xd3\xfb\x10\x6e\x5c\x0f\xdd\x65\x73\xe1\xdc\xf2\x89\xd0\xb1\x25\x0f\x6e\x7c\x90\x07\xe7\x7f\x33\x81\xea\xf0\x5b\xb8\x79\x13\x61\xd0\x8b\xa7\x2c\x09\x7c\xfc\x57\xac\x7b\x11\xf6\x7f\x6a\x5f\xd4\xb8\x06\xa1\xfa\x15\x10\x6e\x7a\x0f\x3a\xb4\xc7\xa0\x61\x55\x3e\x04\x8f\xfc\x39\x84\x3e\xfa\x4b\x68\x5c\x93\x4f\x75\xf8\x2d\x74\x71\x05\x44\x45\x13\xea\x54\xeb\x7c\x85\xa9\x48\xf8\x1e\x15\xcd\xf4\xcd\xee\x7f\x71\x65\x3e\x04\x3e\x7c\x86\x78\x5c\x5a\x1d\xd3\xbf\x7e\x05\xa0\x2c\x94\x89\xb2\x03\xd5\x3f\x66\x88\x05\x31\x21\x36\x1b\xff\xd9\xe5\x13\xe1\xda\xc6\x5c\x68\xdf\x94\x4b\x63\x71\xf0\x37\xbd\x0f\x38\x56\x1c\x33\x8e\x1d\x75\x80\xba\x40\x9d\x0c\xd6\x5f\xcd\x92\x29\x44\x83\xf5\x87\xba\x76\xfa\x73\x85\xe1\x5c\x18\x07\x9e\x38\x86\x18\x68\x8e\xac\xf9\xbb\xbe\x39\x87\xc8\x99\xbf\x13\x3f\xa3\x39\xc6\xb9\x8e\xed\x4f\xb6\xe0\x61\x3f\x40\xdb\x40\x1b\x41\x5b\x19\x62\x3f\xc7\x7f\x4a\xb6\x85\x0f\xda\x9a\xdd\x9f\xd6\x80\xe6\x62\xd2\x26\xd5\xd7\xd0\x46\xc9\x56\x0f\x3c\xde\x6f\xbf\xe5\xb3\xc8\xa6\xd1\xb6\xc9\xc6\x4f\xbf\x01\x68\xf3\x68\xfb\xce\x3a\xf2\x30\x26\x76\x32\x85\xd6\x08\xae\x95\xc1\xeb\x87\xd6\x14\xb3\x69\x11\xad\x39\xaf\x5b\x1d\xed\xfa\x05\x60\x8c\x65\x32\xc6\x54\x8b\x5c\x31\x64\x3d\x0b\x63\xe8\xb0\x45\x2d\x56\xdf\x99\x96\x8f\x99\x1b\xeb\x67\xf2\x47\x8b\xea\xf3\xf9\x08\xae\x22\xe5\x09\xae\x4e\x17\x5c\xfd\xc2\x28\x69\x8a\xe0\xca\x58\xb4\x3d\xdb\x17\xa6\x29\xff\x5f\x04\x57\x5b\x85\xe6\xba\x92\x90\xb8\x1a\x43\x4a\x4c\xbd\x62\xd7\xb7\x08\xae\xd6\x0b\xae\xfe\x8f\xe0\xea\x9b\x38\x1e\x1f\x1f\x4b\xfe\x34\xb5\x7c\xa5\x40\xec\xca\x02\xbd\x64\x3a\xe8\xfb\xbe\x38\x94\x4a\xa6\x83\xf0\x66\x82\xd0\x14\x10\x9a\x0b\x44\xe1\x38\x5c\x2b\x44\xf8\x4e\x75\x5c\x01\xe1\x75\x03\xf2\x11\x9a\x2b\x2a\xb8\x52\x2b\xb8\x3a\x4f\x68\x8a\x92\x0c\x03\xc9\xf7\xb0\x02\xa3\xfc\x2b\x10\xe9\xae\x85\xa8\x7e\x15\xa2\xfa\x95\x18\xba\x0a\xa1\xb3\x8b\x49\xbe\xbe\xe7\x5e\x08\xfe\xef\x3f\x81\xd9\x51\x09\x91\xbe\x4b\x44\xf8\x8e\x75\xf8\x4d\x70\x37\x04\x3e\xfe\x11\xf9\x14\xbd\x68\x3c\xe2\xe8\x14\x5c\x7d\x49\x78\x99\x2b\x11\x06\x47\xfe\x81\xd9\x10\x0d\xde\x82\xc1\x8f\x79\xe3\x30\xe8\x7b\xa7\x81\x51\xfa\x28\x98\xd7\xca\x00\xa2\xa6\xfc\x80\x65\xcc\x3b\x7e\xc3\x36\xfa\xde\x07\x08\x93\xd9\x5a\x0c\xfa\xfe\x87\x41\x68\xac\x43\x70\xf5\x59\x92\xe3\x1d\xea\x1a\x92\xc9\x8f\x06\x3a\xc1\xff\xe1\x5c\xda\x67\x91\x27\x3e\x91\xde\x0b\xe4\xaf\xfc\x47\x5f\x20\xc2\x77\xac\x23\xac\x1d\x95\xd4\xd6\xff\xe1\x9f\x10\x2f\xfa\x5d\xf2\x05\xc4\x70\x4c\x70\xf5\x7e\x94\x35\x1c\xf9\xa1\xc6\x77\x68\x3e\x43\x17\x57\x4a\xfe\xe8\xc3\xcb\xbe\x8c\xfc\xa0\xcf\x23\x09\xdf\xb1\x0e\xbf\x51\x9f\x8b\x2b\x65\x9f\xc6\x77\xe8\x37\xee\x4f\x62\xd7\x58\xb4\x8f\x05\x3e\x9e\x35\xc4\x1e\x13\xc9\x8f\x06\xbb\xc0\xa8\xfc\x23\xf0\x1f\xfa\x63\x88\x86\xba\x21\x72\xeb\x24\x18\xfb\x1f\x82\xde\x1d\x2e\x68\xdd\x30\x0e\x1a\x56\x4f\x24\xc2\x77\xac\xc3\x6f\xd8\x06\xdb\x62\x1f\xec\x8b\xef\x10\xd6\xc1\xff\xd1\xf7\x11\x67\x83\xe0\xea\x97\x06\xeb\x20\x91\x7c\xd2\xdd\xee\x3c\x08\x37\x6f\x04\x30\x0d\x08\x7c\x34\x8f\xe4\x34\xae\x99\x00\xa7\x7f\x3d\x0d\x4e\xfd\xea\x7e\x22\x7c\xc7\x3a\xfc\x86\x6d\xb0\x6d\xb8\x69\x23\xf5\xb5\xe7\xcc\x6c\x3f\x00\x62\x77\x6e\x54\x68\xae\x57\xad\xf5\x96\x52\x7e\xb0\x6e\x01\xe8\xfb\x66\xd0\x1a\x30\xaf\x95\x83\x28\xcc\x81\x2b\xef\xe6\x92\x3c\x24\x8a\x63\x96\x4d\x71\x7e\xe3\x37\x6c\x83\x6d\xb1\x0f\xf6\x0d\xd6\x15\x48\x5d\x86\x7c\x64\x47\xc2\xc3\xf6\x0a\xae\x66\xc6\xea\x20\xae\xfc\x48\x08\xfc\x55\xdf\x05\x7f\xd5\x77\x00\x22\x41\x08\x7e\xfa\x2a\x74\x6d\x53\xa1\x6e\xf9\x64\x92\x75\x69\xcd\x44\xb8\xb9\x35\x13\x6e\x6d\xcb\x80\xe6\x75\xe3\xa9\x0e\xbf\x61\x1b\x6c\x8b\x7d\xb0\xaf\xbf\xea\x7b\xc4\x8b\xec\xe2\xec\x22\x9c\x83\x66\xc1\xd5\x19\xa9\xe4\xd3\xdc\x1f\x98\x0d\xc1\x33\x6f\x42\x34\xec\x03\x7f\xe5\x1c\x68\xdb\x30\x96\xe4\x9c\xfb\xcd\x7d\xd0\xbd\x63\x0c\xe8\x85\x63\x41\x2f\xbc\x07\x7a\x3d\x6e\xa8\x5f\x39\x89\xbe\x61\x1b\x6c\x8b\x7d\xb0\xaf\xe4\xd9\x65\xd9\x6e\x19\x88\x5d\xd9\xba\xd0\x5c\x73\x53\xca\xd7\xff\x0f\xf4\xfd\x0f\x41\xa8\x61\x35\xf9\x1f\x7d\xdf\x74\x68\x5e\x97\x47\xf3\xdd\xf2\x4e\x1e\xc9\x0d\x9f\xfc\x3b\x30\x4f\xbd\x02\xc6\x9e\x09\xd0\xba\x21\x07\x4e\xfd\x6a\x1a\xb5\xc1\xb6\xe4\xb3\x1a\x56\x13\x0f\xe4\x25\xd7\xed\x39\x5c\x9b\x51\xa1\xb1\x17\xe3\xc8\x7f\x0b\x63\x13\x5c\xef\xd4\xb6\xaf\x91\x7c\x6b\xf8\xf2\x16\x5a\xdb\x62\xcf\x64\xb8\xb4\x7a\x3c\xc9\x68\xdb\x90\x0d\xfe\xb2\x87\x01\x2e\x2d\x07\x68\x5a\x05\x81\xca\x27\xe1\xfa\xe6\x4c\x39\x2f\xab\xc7\x53\xdb\x88\xaf\x9e\xfa\x22\x0f\xe4\xe5\xc4\x47\xa5\x8f\xa2\x0d\xfc\x22\x8e\xfc\x5f\xe0\x37\x3b\x5e\x4f\x2e\xff\x1e\xf0\x97\x7e\x11\xa2\x0d\x4b\x00\x9a\x56\x42\xe0\xe0\x13\x70\x7d\x53\x56\x6a\xf9\x81\x4e\x8a\xbf\x70\xac\x43\xe4\x6b\xec\x45\xd4\x0d\xea\x28\xbe\xfe\x67\x38\xfa\x6f\x5e\x97\x2f\xfd\xcb\xd1\xe7\x21\x74\xec\x87\xa0\xef\xce\x81\xab\xbf\x1d\x67\x7d\xcb\xb3\xd6\xcc\x50\xfd\xe3\xdc\xe2\x1c\xcb\xb3\xd0\x60\xf9\xae\xb9\x68\x1b\xe4\xdf\x53\xd8\x1f\xda\x39\xda\xbe\xee\x55\x41\xf7\xba\xc9\x16\xcf\xaf\xb8\x2f\xa5\xfd\x25\x95\x8f\x6b\x42\x63\xcd\xb8\xcf\x25\x5c\x7f\x5b\xe5\xfa\xc3\x39\xb8\xf0\xdf\xf7\x42\xfb\xc6\x7b\xa0\x63\x53\x36\x5c\x5c\x35\x29\xad\xf5\x97\x42\x7e\x26\xfa\x06\xda\x37\xc2\xbe\xb4\xfc\x0f\xd2\x99\x25\xd3\xd2\xf2\x3f\xc9\xe5\x2b\xf6\x1c\xbc\x86\x3e\x12\x7d\xe5\x50\xff\xeb\x8f\xe3\x7f\x25\x0d\xf5\xbf\xfe\x21\xfe\x37\x99\xfc\x18\x1d\x7c\x09\xf7\x08\xdc\x2b\xc0\xd4\x69\xef\x30\x2a\x46\xb9\xff\x58\x73\x9f\x5a\xbe\xc2\x7c\x1a\xed\x8d\x0b\x70\xaf\xc4\x3d\x53\xee\xbf\xef\x8e\x7a\xff\x4d\x47\x7e\x8c\x0e\xee\xc7\x58\x01\x63\x06\xd4\x1d\xf6\x41\x9b\x18\x51\xfc\x61\xf9\xb2\xb4\xe5\x7b\x55\x1b\xc3\xb3\x18\x33\x61\xec\x84\x31\x14\xf1\xdc\xfb\x40\x5a\xf1\x97\x4e\xf1\xd7\x34\x8a\xd9\x06\x3f\xa9\xe4\x3b\xb6\x48\xb1\xa2\xfa\x12\xc6\x8e\x18\x43\x62\x2c\x89\x31\x25\xc6\x96\x29\xe3\x4f\x6f\x26\xc5\xaa\xf1\x62\x58\x8c\x6d\x31\xc6\x4d\x26\xbf\x7f\x3d\x60\xcc\xac\xce\xa3\x18\x1a\x63\x69\x8c\xa9\xb9\x5b\xc6\xd8\x14\x7f\xe7\xc4\xc4\xdf\x39\xb2\x0e\x63\x73\x8c\x91\x93\xc5\xf0\xc4\x47\x49\x2a\x9f\x30\x68\x0a\xf3\x15\x65\x31\x6b\xaf\x7e\xd3\x3a\x53\xd4\x5b\x67\x8c\x44\xe7\x0f\x49\xc9\xcf\x30\xad\xd6\x59\x27\xa9\xfc\x58\x5d\xf4\x69\x14\x2f\x8d\x95\x67\xab\x51\x9f\xcf\xa6\x5b\x67\xbd\xb4\xe4\x7f\xde\x1f\x7b\x6d\x1c\x66\x2a\x1c\x66\x0c\xe9\x99\xc3\x8c\x4d\xb7\x28\xcf\xa2\x4c\x8b\xd4\x74\xa8\xc5\xa2\x1e\x8b\x02\x16\x99\x8c\xa9\xc0\x48\x90\x6a\xcb\x9d\xc9\x18\x9b\xcd\x18\xfb\xfb\xd8\x3c\xc5\xc3\x9f\xb5\x56\xee\x3e\x77\x9f\x3f\xcc\xc7\xda\x1f\x1f\x16\x5c\x7d\x5b\x70\x75\xa1\xe0\x6a\xc1\xef\x89\x6c\xde\xff\x2a\xb8\xfa\x37\x82\x2b\xb3\x05\x57\x72\x7a\x35\xc6\x0c\x2d\x79\x3e\x29\x0d\xfc\xdf\x16\x5c\x0d\x0a\xae\xc2\x1d\x22\x53\x70\xb5\x53\x70\xf5\x20\xed\xeb\x5c\xc9\xc7\x58\xc3\x48\x33\x3f\x97\x18\xbf\x42\xb1\x57\x5a\x84\x6d\x07\x60\x4a\xd2\x37\x6e\x5b\x97\xfd\x1b\xf5\x56\x2a\xb8\xfa\x24\xec\x67\x4c\xe7\xc3\x1b\x83\x83\x5f\x73\x05\x31\xc6\x0b\xd6\xbe\x45\xe7\xf2\x64\x14\x3c\xbb\x08\x8c\x8a\xa7\xfa\x71\x21\xc6\xc2\x1c\x30\x0e\x7e\x0d\x02\x27\xfe\x81\x78\x20\xe1\x3b\xd6\xc9\x78\x84\x39\xd8\xf5\xe2\x29\x60\x94\xcd\x04\xe1\xcd\x8a\x1d\x47\xa3\xe0\xea\x0b\x7e\xce\x5c\x7a\x8a\xfc\x64\x7c\xfc\x2c\x48\xf1\x75\xb8\x6f\x48\x9c\x37\xf8\x31\xdb\x4a\x64\x9c\x84\xb2\xbd\x19\xe0\x3f\xf2\x17\x60\xb6\xee\x85\x68\xe0\x26\x40\x34\x12\x13\x20\x46\xa8\x0e\xbf\x61\x1b\x6c\x8b\x7d\xf4\xe2\xc9\x10\x6a\x58\x0b\xe1\xcb\x5b\xc1\xa8\xfc\x46\xff\x9c\x70\xb5\x0d\xc7\xd0\xa3\x65\x33\x91\xe6\x3c\x0c\x17\x7f\xa4\xeb\x14\xc5\xb4\xc2\xc3\xe8\x3c\x43\xb1\x65\xb0\x3b\xe5\x98\xb1\x0d\xb6\xc5\x3e\xf2\x7c\x30\x13\x22\xdd\x35\x10\xf5\x5f\xa7\x73\xa9\x28\xbc\xc7\x1e\x43\x93\xe0\xca\x53\x14\x73\x16\xa6\x8e\x89\x86\x83\x9f\xce\x41\x55\xdf\x23\xec\x62\x77\x2e\x9d\xc7\xed\xb3\xa8\x1c\x5c\x10\xa2\xa2\x05\x22\xb7\x3e\x21\xc2\x77\xac\xeb\xff\x1e\xa2\x3e\xd8\x17\x79\x20\xaf\x68\xa8\x87\xda\x84\xea\x97\xcb\xbc\xb3\x1c\x43\xa5\xe0\xea\xe4\x74\x62\xba\xe1\xe0\x0f\x35\xac\x91\x36\xc0\xdd\x10\xac\x7b\xbb\x1f\x7b\x34\x0c\x66\x47\x05\xdd\x3d\xe9\xa5\x8f\x80\x5e\x34\x51\x52\xe9\x23\xf2\x3e\xaa\xa3\x82\xda\xd8\x63\x08\xd6\xfe\x52\xc6\xee\xde\x0c\xe2\x69\x8f\x1d\xcf\x4d\x92\xbf\x1a\x45\x7f\xeb\xf3\x32\x57\x5f\x0a\x3b\x4a\x17\x3f\xea\xd2\x28\x9f\x25\xf5\x76\xe4\x3b\xfd\x79\xad\x70\x1f\x84\xce\xff\x17\xe8\x7b\x26\x59\xfe\xc6\xca\xdb\xdb\x3e\x46\x63\xf4\x0d\xdb\xd8\xbc\xe9\x0c\x78\xe4\x79\xe2\x85\x3c\x69\x9e\x68\x7e\x7b\x21\x50\xfd\x63\x7b\x4d\x5f\x13\x5c\xfd\x7a\xca\x73\x45\x9a\xf8\x69\x7e\xad\xb5\x67\xde\xa8\xea\xd7\x59\xdd\x02\x79\x0f\xc1\xa5\x1f\xe9\xd9\x39\x06\x6e\x6e\xcd\x22\xc2\x77\xb9\x36\x5d\xd4\x06\xdb\xda\xf6\x64\xde\x38\x42\xbc\x90\x27\xf2\x76\x4c\xac\xbb\x86\x72\x44\xd6\x18\xde\xd3\xb9\xcb\x9d\xcc\x1f\xa5\x83\x9f\xf2\x5c\x15\x73\x48\x5f\x81\x4f\x5f\x75\xce\xa8\xe1\xcb\x5b\x1c\x9b\x45\xac\x57\xde\xcd\xa3\x5c\x66\xcd\xd2\x29\x44\xf8\x8e\x75\x72\x1c\x0a\xb5\xc5\x3e\x92\xa9\x49\xbc\x68\x0e\x2a\xe6\x0c\x38\x5b\x87\x2e\x2c\xb1\xce\x86\x6a\xbb\xe0\xea\x57\x93\x9f\x2d\x53\xe3\xa7\x5c\xcf\xae\x6c\xb2\x03\xb3\xb3\x5a\xea\xa9\xaf\x51\x9e\x5d\x35\x06\xdd\xdb\x33\x9c\x9c\x53\x3c\xc2\x6f\xd8\x86\xfc\x4e\xf9\x57\x9c\xfc\x9f\xd9\xf9\xb1\xb4\xbb\xc2\x6c\x30\xdb\xcb\xfb\xf5\xa5\x5f\x95\x79\x41\xb9\x67\x2c\x12\xdc\x9d\xf0\xce\x2a\x1d\xfc\xe4\xdf\xc8\x5f\x7c\x97\x72\xf1\x54\x57\xf3\xef\x64\x17\xbd\x1e\x37\xe5\x6c\x12\x61\xb7\x09\xdb\x60\x5b\xec\x83\x7d\xe5\x00\x0c\xe2\x89\xbc\x51\xc6\x00\x99\x75\x0b\x6d\x1b\x3a\x2e\xb8\x32\x61\xa4\xf8\xf1\x37\xe5\xe9\x35\x06\xa1\x86\x55\x96\x7e\x5a\xc0\x28\x9d\x49\x75\xad\xef\x8d\x1b\x82\xb5\x76\xd9\xc0\x7b\x00\x9b\xb0\xad\xed\xfb\xa3\xfa\x15\x69\x2b\x17\x57\x51\x1d\xca\x88\x95\x8d\xfe\x57\xde\x9f\x29\x3e\xc1\xd5\x6f\x26\xce\x8f\x24\xc7\x1f\xf1\x35\x80\x5e\xf2\x00\xe8\xbb\xf3\x21\x72\xf3\x98\x65\xf7\x1f\x80\xf0\x8e\x81\x9e\x1d\x63\xc8\xc6\x6d\x7c\x35\x4b\xa7\x92\xbd\xdf\xda\x96\x09\x5d\xdb\x33\xa1\x6d\x43\x8e\x73\xbf\x60\xe7\xf9\xb1\x0f\xfa\x48\xdc\x7b\x89\xff\xcd\x6a\xe2\x4d\x79\x65\x5f\xc3\x40\xbd\xfd\xee\x5b\xb6\x0d\x51\x7e\x3b\xde\x3a\x4e\x85\xdf\xec\x38\x44\xb6\x6f\x94\x3d\x46\xff\x67\xc0\x27\x70\xf2\x65\xe2\x7b\xed\xfd\x9c\x98\xbc\xe7\x54\xa9\x5f\x27\x16\x50\x40\xe7\x2a\x5c\xdf\x9c\x4d\xf7\x32\x76\x3b\xec\x83\x7d\x03\x9f\xfc\xa3\xc4\xe9\x6f\x27\xde\x28\x03\x65\x0d\xb0\xa1\x9a\xff\xb0\xf1\x6f\xd3\xb9\x2b\xee\x9d\x69\x2a\xfc\xa4\x6b\x9c\xdf\xc3\xcf\x01\x44\x02\x94\xcb\x35\x2a\x9f\xa6\xba\xa6\xb5\x13\x9c\x7c\x29\xae\x51\x9f\x47\xe6\x93\xf4\xc2\x6c\xd0\x77\x8f\x73\xe2\xcd\xcb\xeb\xf3\x9d\xfc\x2a\xf6\x21\x7e\x87\x9e\x96\x79\xe1\x48\x80\x78\x63\x1d\xca\x1a\x20\xfb\x2a\xb7\xf7\xb3\x6a\xc1\x95\xdc\x91\xe0\x0f\x5d\x58\x26\xfd\xe6\xf1\x97\xa4\xbe\xc4\x65\xca\x5d\xf5\xee\x54\x29\xe7\x6d\xeb\xb5\x0d\xf5\xca\x5d\x60\xec\x9d\x4a\x77\x30\x91\xda\x37\xc0\x5f\x3e\x93\xc6\xd3\xf9\xc1\x58\xb2\x2d\x3b\x4f\xee\xdb\xa9\x12\x0f\xe4\x45\xf3\x79\xfc\x25\x92\x81\xb2\x06\xd8\xee\xcd\x13\xf2\xce\x98\x2b\x4d\x82\xab\x0f\x26\xc9\x91\x3e\x97\x08\x3f\xed\x4f\x3b\x19\x04\x4f\xbd\x2e\x79\xf6\xd4\xd2\xbe\xd3\xbd\x7d\x0c\x9c\xb5\x6c\x1b\xb1\xe1\x7e\xa5\x7b\x55\x08\x1d\xfb\x11\x40\xf3\x5a\x22\xf3\xf4\xcf\x69\x2e\xd0\xff\xa3\xed\x23\x7e\xec\xd3\xbd\xdd\x4d\x31\x74\xa4\xa7\x4e\xca\x38\xf5\xba\x94\x81\xfb\x5b\xac\xef\x10\xcd\x74\xf7\x64\xdd\x79\xcf\x4a\x82\xff\x5b\x42\x63\x86\x71\xf0\xc9\x21\xb1\x24\xf9\x4e\xe4\x5d\x3b\xdf\xd2\xc9\x71\xd0\x8b\xf2\xa1\x6b\x5b\x86\xe3\x63\xb0\xec\xda\x9e\x41\xf7\x6f\x66\xcd\xeb\x00\xcd\x6b\x00\x9a\x56\x43\xb4\x7e\x31\x18\x7b\x27\x83\xcf\x23\xe7\x0a\xf1\x3b\x6d\x8b\xf2\xe9\xff\x3a\x24\xa3\x76\xbe\x94\x31\xc8\x87\x62\x6c\x4a\x31\x8b\xc6\xfa\x04\x57\x9f\x4e\x82\xff\x69\x6c\x43\xb1\x88\xff\xfa\xc8\xf0\x6f\xcb\x00\x7d\x57\x16\x98\x67\x7e\xde\x8f\xff\x42\x01\x18\xc5\x93\x08\xff\xf9\x91\xe0\x0f\x76\x03\xea\x14\x75\x4b\x3a\x4e\x8c\x7f\x16\xe5\x96\xf7\xcd\xa0\x39\x4b\x6e\x3f\x75\x34\xf7\x68\x03\xb6\xfd\xa0\xef\xb9\xb1\x85\xee\xd3\x21\x78\xe4\xcf\xe8\x0e\x10\xed\x27\x7c\xe2\x45\xd0\xbd\x19\xd0\xbd\x23\x83\xda\x0e\xdb\x7e\xc2\x7d\x74\x67\x81\xb6\x4d\x36\x9e\x18\xff\x83\xb8\x46\xf4\xa2\x09\x74\xdf\x31\x92\xf5\x7b\xf5\xb7\xb9\x96\xef\xc9\x81\x60\xd5\x73\x10\xaa\xfe\x3e\xe8\xc5\xf7\x82\xce\x15\xba\xd3\xc2\x31\x9e\xfa\xb5\xbd\x7e\x95\xb4\xd6\xef\x20\xfc\xdf\x8e\x8f\x9f\x72\xeb\xb9\xe4\xa3\x70\x5f\xb9\xca\x07\xf0\x48\xd7\x7f\x9e\x5b\x81\x7b\x53\x46\x7f\xec\x6c\x91\x4f\xeb\x8f\x2f\xb0\x6d\xe3\x30\xfc\x67\xba\xf8\xe5\xde\xa0\x6e\x47\x1e\xc1\x9a\xf9\x03\x78\xa4\xbb\x7f\x21\x35\xad\x1b\x0f\xbd\x9e\x31\xd6\x3d\xa6\x0a\x7d\x9a\x9b\xe6\x05\x75\x3f\x64\xff\x3a\xf9\xb2\xb5\x46\x13\xef\x5f\xe9\xe0\xd7\xed\xbb\x3e\xdc\xa3\x51\x2f\xbf\xfb\xd3\x81\x71\x48\xdc\xf8\x61\x6b\xdc\xf8\xc1\xde\xc7\x30\x6e\x40\x9c\xa8\xeb\x58\xec\xc3\x89\x1f\xd2\xc5\x1f\xb3\x06\xbe\x29\xb8\xd2\xa7\xef\xb9\x0f\x22\x5d\x9f\x0e\xe0\xe1\xc4\x6f\x17\xd3\x8b\xdf\x12\xd1\x70\xe2\xb7\xe1\xe1\xa7\x35\x30\x81\x62\x55\xcd\x45\xe7\xd0\xd8\xe7\xb3\x88\x9f\x87\x85\x5f\x53\x18\xfd\x3f\x95\xab\x8b\xe5\x19\xe3\x71\x3a\x43\x38\x6b\xe0\xf7\x76\x7e\xa9\x8e\x7b\x7e\x19\x2e\xfe\x18\x1b\xc2\xb3\x5a\x07\x9e\xdd\xf0\x0c\xe7\xf0\xb9\x23\xe7\xc7\x9b\xa3\xc4\x4f\x7e\xc8\x8d\x67\x66\x3a\xa7\xef\x7f\x88\xce\xd2\xf6\x93\xf6\xf9\x7d\x47\xcc\xf9\x7d\x47\xb2\xf3\x7b\x55\xdc\xf3\xfb\x48\xf1\xc7\xcc\xc1\xd7\x28\x77\x81\x3e\xae\xfa\xaf\x29\xa7\x01\x43\xf2\x27\xcf\x0f\xcc\x9f\x9c\xbb\x3d\xf9\x93\xd1\xe2\xef\xe3\x2e\xe6\xe3\x4c\x91\x67\x66\x35\x8a\x7e\x8e\xd6\xb2\xa5\xb3\xd4\xf9\xab\x17\xc1\x28\x7d\x04\xf4\x3d\x13\x89\x8c\x44\xf9\xab\xba\xb7\x87\xe6\xaf\x6e\x03\xfe\x98\x39\x98\x2c\xb8\x7a\xc8\xb6\xd9\x50\xfd\x32\x99\x1b\x0c\xf5\x38\xfe\xe2\xf6\xe5\x0f\x13\xe7\x4e\x47\x84\xbf\x50\xb5\xfd\xe9\x1c\x2b\x97\x4a\xb9\x55\xfa\xdf\x50\xe0\x3a\x44\xba\xcf\xc8\xdc\x37\xda\x8a\x9d\xbf\x4d\x82\xc1\xc1\x62\xe5\x6f\xb1\x0f\xe5\x7e\x4b\x1f\xa5\x5c\x70\xd2\x3e\x23\xc0\x4f\x63\xf0\xba\x58\x8f\x27\x1b\xe7\xe1\x05\x99\xd3\x56\x68\xbe\x8d\xca\x6f\xd0\xbe\x19\x6a\x58\xeb\xac\x3d\x27\x7f\xde\x56\x92\x7e\xfe\xbc\xe4\x41\x30\xdb\xf6\xa5\x1e\xf3\x08\xf1\xe3\xa3\x7b\x15\x66\xec\xa2\xff\x6d\xbc\x60\xdd\x2d\x48\xbc\xbb\xb2\xe8\xbf\x31\x7a\xf1\xd4\x24\xf7\x17\xf3\x89\x12\xdd\x5f\x18\x15\x4f\xd1\x1d\x48\xca\x7b\x92\xda\xb7\xac\xff\xda\xba\x86\x8d\x9f\xc6\xc0\x5d\x0c\xbc\xe4\x5b\xd1\x27\x95\x39\x77\x64\x8e\x8f\x19\xcd\xfd\xd1\xb0\xee\xa9\x46\x84\x1f\x1f\x43\x73\xd9\x6b\x7a\xbc\xbc\x6b\xa3\x3b\xb7\x4e\xeb\x0e\xee\x4e\xdd\xf7\x8d\x18\xbf\x33\x0e\x8f\xc2\x7a\x8b\x54\x8c\x35\x72\x04\x57\x67\x0b\xae\xfe\xc4\xba\x0b\x2d\xb8\x03\xf7\xae\x6f\x5b\x77\xbc\x23\xc6\x7f\xf7\xb9\xfb\x7c\x9e\x1e\xb9\x43\x24\x2e\x5b\x18\x63\xcf\x58\x65\x9e\x55\x66\x5a\xa5\x6b\x50\xc9\xec\xb2\xc0\x2a\x9f\x19\x54\x4e\x4f\x50\xe6\x25\x28\x33\x6f\x5f\xd9\x93\xa0\x0c\x24\x28\xcd\x41\x65\xd4\x2a\xc1\x2e\x17\x0e\x2a\x5b\xac\xb2\xc7\x2a\x4d\xab\x4c\xa1\xdf\xff\x0f\x00\x00\xff\xff\xc6\xb9\x24\x2f\xee\x3a\x00\x00")

func faviconIcoBytes() ([]byte, error) {
	return bindataRead(
		_faviconIco,
		"favicon.ico",
	)
}

func faviconIco() (*asset, error) {
	bytes, err := faviconIcoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "favicon.ico", size: 15086, mode: os.FileMode(438), modTime: time.Unix(1612640817, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsMainJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\xce\xcf\x2b\xce\xcf\x49\xd5\xcb\xc9\x4f\xd7\x50\x4a\xad\x48\xcc\x2d\xc8\x49\x55\xd2\xb4\x06\x04\x00\x00\xff\xff\xc8\x9f\xbd\x5f\x17\x00\x00\x00")

func jsMainJsBytes() ([]byte, error) {
	return bindataRead(
		_jsMainJs,
		"js/main.js",
	)
}

func jsMainJs() (*asset, error) {
	bytes, err := jsMainJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/main.js", size: 23, mode: os.FileMode(438), modTime: time.Unix(1613718745, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"css/main.css": cssMainCss,
	"favicon.ico":  faviconIco,
	"js/main.js":   jsMainJs,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"css": {nil, map[string]*bintree{
		"main.css": {cssMainCss, map[string]*bintree{}},
	}},
	"favicon.ico": {faviconIco, map[string]*bintree{}},
	"js": {nil, map[string]*bintree{
		"main.js": {jsMainJs, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = os.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
