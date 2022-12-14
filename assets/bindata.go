// Code generated by go-bindata. DO NOT EDIT.
// sources:
// service/app.service
// service/env
// code/index.html
package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _serviceAppService = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x92\x41\x6f\xdb\x3e\x0c\xc5\xef\xfa\x14\x04\x7c\xfd\xff\x8d\x14\x43\x76\xf3\x61\x5b\xb2\x2d\x58\x9b\x16\x4d\xda\x1e\x8a\x1e\x54\x9b\x76\x84\xda\xa2\x47\x52\xce\x84\xc2\xdf\x7d\x90\x13\x63\xed\x0e\x3b\xd9\xef\x47\x3f\x51\xe6\xe3\xe3\x9d\x77\xfa\x64\x56\x28\x25\xbb\x5e\x1d\xf9\xe2\xf5\x35\x7f\x23\xc7\xd1\x7c\xaa\x15\xb9\x90\x28\x2d\x35\xb9\x5a\x6e\x50\xcf\xcc\xa3\x1e\x89\x5f\xde\xc3\x2e\xca\xcf\xb6\xca\x05\x79\x70\x25\x02\x63\xe5\x64\x56\xc6\x3c\xee\x4e\x6f\x4f\x26\x83\x2b\xaa\x5c\x1d\x41\x0f\x28\x08\x7a\x24\x18\x6c\x1b\x50\xc0\xfa\x0a\x82\x2f\xa9\xeb\xd0\x6b\x2a\x77\xe0\x6a\x88\x14\xe0\x60\x07\x34\x19\x30\xf6\x24\x70\x74\x7a\x80\x96\x54\x80\x6a\xa8\x5d\x7b\x76\x36\xa8\x60\x3d\x7c\xdf\xef\x6f\x00\x99\x89\x61\xb9\x58\xc0\x33\x96\x36\x48\x32\x53\x0d\x7a\xb0\x6a\xb2\x2c\x33\x19\x5c\xba\xce\xe9\xd5\xfa\xea\xf2\xfa\xcb\x8f\xc2\xf9\xda\x79\xa7\x71\xe6\xdb\xeb\xaf\x9b\xcb\x75\xf1\x71\xb9\xfc\xb0\x34\xfb\xd8\x63\x21\xae\xeb\x5b\x34\x77\x82\x9c\x06\x95\x9e\xde\x76\x38\x8e\xe6\x1b\x53\xe8\x67\xd6\x24\x31\x8e\xe6\x81\xf8\xc5\xf9\x66\xe5\x18\x4b\x25\x8e\xd3\x70\x67\x31\x8e\x66\xed\x07\xc7\xe4\xd3\x7f\x16\xab\xcd\xed\xbf\xca\xeb\xed\x7d\xd1\x33\x55\xa1\x4c\xb1\xbc\x2b\xdd\xaf\x6f\x77\x9b\xeb\x6d\x31\x5c\xe4\x8b\x7c\x61\x32\xf8\xcb\x77\x51\x4c\x83\xbd\x30\xeb\x5f\x58\xee\xd4\xb2\xa6\x46\x49\x04\x4d\x57\xbf\x45\x99\xa0\x6d\x8f\x36\x8a\x31\x26\x83\x1d\x75\x08\x95\x13\x65\xf7\x1c\x52\x43\x81\xce\x46\xf0\xa4\x20\xa1\xef\x89\xf5\x1c\xdb\xc1\x72\x85\xde\xf9\x06\xaa\xe9\xe2\x6e\x40\xc9\x61\x73\x8a\xab\xb4\x7e\x72\xa4\xd3\xd3\xf7\x30\x2f\x45\x15\x52\x12\x4a\x29\xa8\xe0\x5f\x3c\x1d\x3d\xd0\xb4\x6f\xff\xc1\x9c\x3b\x85\x93\x87\x3c\xca\xdb\xc6\x58\xc1\x73\x4c\xc7\x33\x0c\xc8\xe2\xc8\xa7\x48\x25\x8a\x62\x57\xe5\x26\x83\x1b\x26\xc5\x52\x77\x13\x29\xea\xd0\xb6\x13\x74\x83\x55\x5c\x61\xea\x2f\x45\x44\xf9\x03\xf7\x5d\x7f\x06\x5b\xda\xe2\x31\x51\xd7\x62\x83\x52\x28\x87\xb4\xb3\x1b\x2f\x6a\xdb\xf6\xc9\x3c\x58\xaf\x58\x7d\x8e\x45\x17\x5a\x75\xff\x07\x41\x9e\x97\xff\x77\x00\x00\x00\xff\xff\xf7\x73\xf1\x65\x4a\x03\x00\x00")

func serviceAppServiceBytes() ([]byte, error) {
	return bindataRead(
		_serviceAppService,
		"service/app.service",
	)
}

func serviceAppService() (*asset, error) {
	bytes, err := serviceAppServiceBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "service/app.service", size: 842, mode: os.FileMode(420), modTime: time.Unix(1669379573, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _serviceEnv = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\x5f\x4f\x1b\xc7\x17\x7d\x9f\x4f\x71\x1f\x7e\x7e\xc4\x5e\xfe\xfc\xd2\x08\x6b\x1f\xd2\xb8\x55\x2b\x11\x29\x72\xe0\xa1\x42\x3c\x2c\xf6\x18\x56\xda\xdd\x49\x77\xc6\x21\x51\x14\x09\x52\xc0\x24\xaa\x63\x52\xe2\x18\x02\x24\xa4\x4d\x30\x55\x60\x71\xd5\x16\x83\xc1\xf6\x97\xf1\xec\xec\x7e\x8b\x6a\x76\xd6\x98\x56\x89\xd2\x87\x5a\x7e\xb1\xee\xb9\xf7\x9e\x73\xe6\xf8\xa6\xd3\x69\x10\xdb\xcf\xf9\xfa\xa1\xf8\xb0\x25\x1e\x9f\xa1\xe9\x3b\xcc\x60\x26\x71\x66\x50\xbf\xd0\x3b\x3d\xf7\x77\x5a\xe8\x1b\x42\x19\xe8\x30\x3c\xf2\x45\x52\x4b\x6a\xc9\x61\x34\x68\x3c\xe6\x95\x5f\xd0\x6d\xe2\xca\xfa\x75\x4d\xd3\x50\x1a\xf8\xca\x52\xe0\x9d\x06\xcd\x03\xfe\xe6\x0d\xba\x61\x59\x64\xe1\xa6\x4b\x28\xcd\x10\xdb\x30\x1d\xd0\x81\x38\x08\xc9\xdd\x7e\xed\x3d\xef\xd6\xc2\x95\xb2\x68\x7b\x68\x7a\x82\xcc\xc9\xbd\x41\x67\x83\x97\x5a\x7c\xad\xe1\x3f\xab\xf3\xb5\x13\x5e\x69\xa0\x3b\x2c\x4f\x8a\x0c\x00\x64\x6f\xa1\x20\x41\x5e\x83\xb7\xab\x12\x14\x8d\xf0\x5f\x96\x7a\xe7\x27\x28\x8b\x73\xc4\xcd\xc7\x38\x07\xf5\x17\xf0\xa3\x4d\xfe\xf8\x40\x6c\x7b\xbc\x5d\x45\x19\xd3\xc5\x39\x46\xdc\x07\xa0\xc3\xff\x1e\x66\xbe\xcd\x3e\x4a\x59\x64\x8e\xc6\x84\xaa\x0d\xbf\xec\xf1\xd6\x46\x9f\x53\xc6\x60\xc6\xac\x41\x71\x44\xac\xfb\xda\x7f\xf6\xfe\xd3\x86\xa8\xfa\xdf\x0d\x19\x1d\xd5\xae\xa1\x2b\x73\xf9\x7a\x59\xd4\x1b\x28\x33\x1b\x71\x64\x98\xb2\x41\xe3\x8b\x03\x7f\xad\xc9\xd7\xcb\x68\x8a\x62\xd7\x31\x6c\x0c\x3a\xb8\x84\x0c\x10\xfc\x78\x55\xec\x2d\xa1\xdb\x06\xa5\x0b\x52\xa7\x5c\x3f\x9a\xcc\x11\x7b\x80\x38\xaa\x89\xc3\xfd\x70\x7b\x15\xdd\x9c\x37\x5c\x8a\x99\xdc\x52\x64\x85\xeb\xf6\xec\xd8\xe7\x5d\xd3\xc1\x7e\x40\xbf\xb7\x86\x12\xdf\x0d\x25\xec\xa1\x44\x3e\x69\x91\xb9\xab\xe4\x55\x97\x38\x7a\x22\x5a\x75\x18\x1e\xa7\xa6\x85\x1d\x06\x23\xe3\xd8\x75\x89\x0b\xa3\xe3\x0b\x86\xeb\xc0\xd8\xb8\xe9\x14\x08\x9a\xc0\xf7\xb0\x05\x3a\x8c\x29\x63\xb3\x38\x6f\xd2\xbe\xa9\xd1\x8f\xff\xca\x51\xf1\xc7\xcf\xfc\xa2\x0a\x79\x5c\x30\x8a\x16\x83\x69\x6d\x68\xf8\xff\x33\x7d\x83\xb5\x7f\x7a\x77\xa3\xc8\xe6\x41\x47\x69\x08\xcf\x37\x03\xef\x9d\xb8\xd8\xe0\x47\x9b\x7e\xed\x24\xac\x76\xa7\x45\xfd\xa7\x19\x34\x39\x39\x01\x3a\x5c\x1b\x93\x49\xd6\xe2\x50\x1c\xbc\xf5\x77\xbb\x2a\x3e\x7d\x09\x93\xd8\xbe\x6b\x19\x2c\xca\xc5\xd5\x7a\xaf\x1d\xd5\x3f\x12\x32\x17\x53\x52\x74\x73\x98\xa6\x58\xdc\x7b\xd9\xaa\xde\xc1\x7f\xf2\x2b\xff\xad\x1a\x07\xe4\xab\xfb\x0c\x3b\xd4\x24\xf2\xdf\x92\x9c\x67\xb6\x25\xc1\xc7\x15\xff\xf0\x6d\xe0\x75\xc3\x9a\x17\xfe\xd0\x0e\x4b\x65\xff\x65\x83\x3f\xdd\x0b\xda\x6d\x35\x07\x65\xb1\x45\x8c\x38\xff\xcc\x2d\xe2\xab\xfc\x79\x65\x33\x2c\x55\x82\xce\x99\xa8\xfe\x28\x0e\xf7\x79\xa5\x39\x50\x91\xcc\x60\xcb\xb4\x4d\x26\xd5\xf0\xe6\x7e\xd0\x39\x53\x68\xbe\xb6\x1a\xbe\x7a\x11\xa3\x27\x70\x21\x8a\xd3\xc3\x84\x44\x55\x7e\xff\x38\x2a\x6b\xce\xcd\xcb\x67\x4a\x3c\x42\x6a\x7b\x74\x51\x72\x10\xee\x6e\xf9\x8b\x4b\xc1\x9f\xcb\x7e\x6b\x5d\xe9\xed\x5b\xa9\x00\xd1\xc5\x69\x3f\x17\x1f\xb6\xf8\x76\xc7\xdf\x2b\xa1\xaf\x8d\x7b\x66\x8e\x38\x91\x96\x4f\x5a\x98\x2a\x28\x54\xd2\xcc\x91\x28\xdf\xd2\x1b\xbe\xd3\xe0\xbb\x8b\x68\xca\xb5\x40\x7d\x74\x48\xd1\x68\x89\x24\xee\xbd\x0e\xb7\x56\xd4\x45\xf8\xfc\x63\x0d\x16\xc5\x03\x22\x49\xbd\xd3\xa7\xbd\x8b\xbd\x3e\xff\xa9\xbb\xd2\xf4\x99\xcb\xb3\xa7\xaa\xe2\xd5\x72\xfc\xac\x3b\x8b\xfc\x5d\x5d\x7e\x1b\x15\xe0\x65\x99\x10\xb8\xf5\x25\xba\x65\xdc\x37\xed\xa2\x0d\x3a\x8c\xc8\x90\xaa\x26\xc5\x3f\x68\x1e\xf3\xce\x72\xc4\x5f\x87\x54\x31\x1a\x7f\x09\x51\x31\x53\x2a\xfe\x4d\xd8\x54\xfb\x5f\x01\x00\x00\xff\xff\x98\xad\x40\x8e\xe9\x05\x00\x00")

func serviceEnvBytes() ([]byte, error) {
	return bindataRead(
		_serviceEnv,
		"service/env",
	)
}

func serviceEnv() (*asset, error) {
	bytes, err := serviceEnvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "service/env", size: 1513, mode: os.FileMode(420), modTime: time.Unix(1669380808, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _codeIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x6d\x6f\xa3\xc6\x13\x7f\x9f\x4f\xb1\x7f\xfe\x95\xc0\x95\x59\xec\xeb\xe9\x7a\x22\xc6\x6a\x7b\x4d\xd5\x9c\x72\x97\xaa\x49\xd4\x56\x6d\xa5\xac\xd9\x31\x6c\xb2\xcc\x72\xbb\x0b\x8e\x63\xf1\xdd\xab\x05\x1c\x43\x1e\xd4\x4a\xe1\xcd\xc2\x3c\xed\x6f\x66\x7e\x33\xf6\xe2\x7f\x3f\x9e\x7f\xb8\xfc\xe3\x97\x13\x92\xdb\x42\x2e\x8f\x16\xee\x20\x92\x61\x96\x78\xf7\x79\xf8\x33\x43\xe3\x39\x29\x30\xbe\x3c\x22\x84\x90\x45\x01\x96\x91\x34\x67\xda\x80\x4d\xbc\xab\xcb\x9f\xc2\xf7\x5e\xaf\xb2\xc2\x4a\x58\xee\x76\xf4\xd2\xbd\x34\xcd\x22\xea\x24\x03\xc7\xdc\xda\x32\x84\x2f\x95\xa8\x13\xef\xf7\xf0\xea\xfb\xf0\x83\x2a\x4a\x66\xc5\x4a\x82\x47\x52\x85\x16\xd0\x26\xde\xe9\x49\x02\x3c\x03\x6f\xe8\x89\xac\x80\xc4\xab\x05\x6c\x4a\xa5\xed\xc0\x78\x23\xb8\xcd\x13\x0e\xb5\x48\x21\x6c\x3f\xa6\x44\xa0\xb0\x82\xc9\xd0\xa4\x4c\x42\x32\xa7\xb3\x7d\x28\x29\xf0\x96\x68\x90\x89\x67\xec\x56\x82\xc9\x01\xac\x47\x72\x0d\xeb\xc4\x73\xd8\x4c\x1c\x45\x15\x96\xb7\x19\x4d\x55\x11\x49\xb6\xad\xc4\x77\x6f\xe8\x3b\xfa\x3e\xe2\xc2\xd8\x28\x35\xa6\x13\xd2\xd4\x98\x7d\x4c\x93\x6a\x51\xda\xee\xc3\x3d\x35\xd3\xc4\xb0\xa2\x94\xf0\xd1\x28\x24\x09\xf9\x78\x71\xfe\x99\x96\xae\x62\xc1\x6e\x47\x9d\xb0\x69\x26\xc7\x23\x7b\xce\x2c\xdb\x5b\x1a\xab\x05\x66\x62\xbd\x0d\x0e\x51\xa6\x04\x2b\x29\xa7\xc4\xfb\xcb\x7a\xbd\xeb\x22\x1a\x5e\xbc\x68\xf3\x39\x80\x58\x29\xbe\x25\xbb\x87\x4f\xf7\x94\x8c\x73\x81\x59\x4c\x66\xc7\x23\x79\xc1\x74\x26\xf0\x89\x38\x07\x91\xe5\x36\x26\xf3\xd9\xac\xce\xc7\xaa\xb6\xca\x9d\x66\x33\xd6\xa8\x1a\xf4\x5a\xaa\x4d\x4c\x72\xc1\x39\xe0\x41\xdb\x3c\xbc\xfd\xbf\xad\x60\x98\x2a\x0e\x8f\x10\xbe\x1c\xf7\x05\x30\xcd\xbe\x12\x5d\xf2\x8b\xa8\x23\xea\xc2\x65\xbf\x3c\x5a\x70\x51\x13\xc1\x13\xef\x70\xa1\xb7\x5c\x44\x5c\xd4\xce\xb4\xb7\xe9\x8a\x48\x8c\x4e\x0f\x0c\x48\x39\xd2\x95\x52\xd6\x9d\x08\x36\x62\x37\xec\x2e\x92\x62\x65\xa2\x54\x8a\x72\xa5\x98\xe6\xf4\xc6\x44\x6f\xe8\x8c\xce\xe7\x03\x59\x21\x90\xde\x18\x77\xc7\xbe\x35\xcf\x86\x7f\x4a\xb0\x6f\xe9\xbb\x8e\x60\x1d\xb9\xfe\x43\x8c\x97\x21\x16\x0a\x59\xaa\x42\xe0\xc2\x2a\x1d\xcd\xe8\x37\x6f\xe9\x3c\x2a\x04\x46\xb5\x89\xa4\x62\x1c\xf4\xcb\x38\x3b\xfe\x68\x37\xa0\x1a\x68\xaa\x70\x2d\xb2\xe0\xd0\xa3\x92\xd9\xdc\xc4\x64\x47\x6a\x13\x13\xff\x75\x50\x7c\xd2\x4c\x1f\x02\xfb\xb5\x89\x50\x1a\x3f\x26\x3b\x56\x33\x21\xd9\x4a\xc2\x19\xc3\xac\x62\x19\xb8\x0b\xfd\xaf\xfd\xd8\xbf\xcf\xc3\x14\xfd\xa6\x77\xdb\x4f\xd0\x46\x20\x57\x1b\xfa\xa9\xbd\xea\x04\x6b\xa1\x15\x16\x80\x96\x24\x03\x72\x65\x60\x7f\x53\xfa\x16\xf4\x95\x96\x31\x59\x57\x98\x5a\xa1\x90\x04\x9b\x56\x78\xca\xa7\x44\xb2\x15\xc8\xc9\x23\x3e\x6a\xb0\x95\x46\x72\xed\xa6\x33\xb6\x70\x67\xa3\x1b\x56\xb3\xae\x54\xc7\xfb\x1d\x58\xd9\x75\xf8\x7e\xfa\xd5\x0e\xd0\x11\xec\xea\xd7\x53\xb7\xd2\x14\x02\xda\xe0\x7a\x14\xcd\x80\x5c\xff\x0b\xce\x7e\x76\x99\x81\x16\xe8\x2b\x4a\xec\x8f\x62\x36\xe3\x69\x12\x85\xdb\xa1\x17\x6d\x1e\x26\x78\x65\x23\x1d\xda\xa8\x2b\x64\x7f\x7c\x62\x02\x7b\x8e\xf9\x93\xe3\xeb\x49\x73\x3d\x98\xd8\xbe\x7b\xc7\x43\xa6\x05\x7f\x3a\x02\xf4\xd1\xbb\x83\x16\x4c\xa0\xff\xf7\x74\xd0\xac\x61\x77\xfa\xae\x77\xb8\x68\xef\x92\x6a\x60\x16\x02\xae\xd2\xca\x55\x96\x7e\xa9\x40\x6f\x2f\x40\x42\x6a\x95\x0e\xbc\xc1\xde\xf1\x26\xd3\x47\x55\xaf\x99\xac\x20\x6e\xf7\xf0\x74\xa4\x90\x3d\x0d\x63\xe2\xdf\x18\x85\xfe\x58\x6b\x73\x28\x20\x26\xee\x37\xcf\xbd\x35\xcd\x58\xbd\x56\x68\x2f\xc4\xbd\x73\x9e\xbf\x2d\xef\x1e\x39\x4b\x81\xf0\xb9\x2a\x56\xa0\xdd\x38\x3d\x89\xad\x55\x85\x1c\x78\x97\x80\x50\x18\x93\x35\x93\x06\xc6\x46\x26\xd5\x4a\xca\x1f\x60\xab\x90\x9f\x31\x63\xcf\x04\xc2\xb3\x86\x1a\x18\x3f\x47\xb9\x7d\x56\x69\xd9\xaa\x83\xf9\xf6\x20\xdf\x0f\xd8\xa3\x41\x63\x9c\x9f\xd4\x80\xf6\x4c\x18\x0b\x08\x3a\xf0\x35\x18\x71\x0f\xfe\x94\x40\xdd\xf2\x79\xf9\xb4\x4f\x52\xa5\xcc\xa5\x40\x35\xb8\x15\x14\x4c\xfa\xc8\x47\x83\x0d\x14\xb5\xff\x40\xfe\x09\x00\x00\xff\xff\xf0\xc1\x25\xcb\x91\x08\x00\x00")

func codeIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_codeIndexHtml,
		"code/index.html",
	)
}

func codeIndexHtml() (*asset, error) {
	bytes, err := codeIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "code/index.html", size: 2193, mode: os.FileMode(420), modTime: time.Unix(1671523337, 0)}
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
	"service/app.service": serviceAppService,
	"service/env": serviceEnv,
	"code/index.html": codeIndexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
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
	"code": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{codeIndexHtml, map[string]*bintree{}},
	}},
	"service": &bintree{nil, map[string]*bintree{
		"app.service": &bintree{serviceAppService, map[string]*bintree{}},
		"env": &bintree{serviceEnv, map[string]*bintree{}},
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
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
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

