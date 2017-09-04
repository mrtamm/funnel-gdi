// Code generated by go-bindata.
// sources:
// config/default-config.yaml
// config/gridengine-template.txt
// config/htcondor-template.txt
// config/pbs-template.txt
// config/slurm-template.txt
// DO NOT EDIT!

package config

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

var _configDefaultConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x51\x73\xdb\xb8\x11\x7e\xf7\xaf\xd8\xda\xe9\x34\x99\x11\x69\xb9\x99\xeb\xf4\x34\xe3\x07\x5b\x76\x1c\xcf\x39\x89\x6a\x29\x75\xdb\x17\x0f\x48\x2c\x49\x9c\x49\x80\x01\x40\x29\x8c\x9b\xff\xde\x59\x00\x24\x25\x47\xb6\xd5\x9e\x1f\x7a\x79\xc8\x58\xc4\xee\xb7\x8b\xdd\xc5\x07\xec\xce\x51\x2f\x51\x4f\xf6\x00\x0e\x60\xc6\x6c\x01\x56\x81\x2d\x10\x38\xb3\x2c\x61\x06\x21\x13\x25\xc6\x7b\x00\x67\xa7\xb4\x3c\x81\xf8\x30\x6b\xa4\xc4\x32\x5a\x29\x7d\x17\x71\xa1\xc3\xef\x98\x27\x7b\x0e\xe5\xbd\x32\x56\xb2\x0a\x41\x65\x0e\xe9\x9d\x5b\x06\xe3\x0c\xc5\xf0\x81\x09\x59\xb6\x23\xb0\x85\x30\x20\x0c\x34\x06\x39\x24\x2d\xb0\xc6\xaa\xc8\xa4\xac\x44\x6d\x1c\x8e\x55\x90\x2a\x99\x89\xbc\xd1\x08\x64\x0c\xb5\x21\x47\x08\xff\x23\xab\x70\x02\xa5\x4a\x59\x59\x28\x63\xbd\xe1\x99\xd2\xd6\xc3\x65\x4a\xc3\xfb\xc5\x62\x06\xa9\xaa\xaa\x46\x8a\x94\x59\xa1\x24\x30\xc9\x9d\x47\x2b\x4c\x80\x33\x53\x24\x8a\x69\xee\x20\x17\x8b\x19\x69\x4f\xe0\xaf\xe3\xf1\x78\x1b\xda\xf5\x6c\xba\x09\x46\x6a\xd7\xb3\xa9\xd7\xfa\x79\xfc\x73\xd0\xba\xc6\x2f\x8d\xd0\x08\x09\x33\x22\xa5\x3d\x15\x28\x6d\x67\x9f\x80\xc8\xbe\x0f\x05\x9c\xcc\x2e\x69\xfb\x42\xe6\xc0\xa0\x66\xc6\xac\x94\x77\xe7\x00\x2e\x33\x67\x7a\x04\x15\xbb\x43\x30\x14\x01\xab\xa0\xd6\xaa\x46\x5d\xb6\xa0\xd1\x58\x2d\x52\x0b\x2c\x4d\xd1\x98\x2e\x67\x3e\x5c\x2e\x63\x0e\xe5\x35\xc6\x79\x0c\x69\x51\x29\x0e\x7f\x19\x8f\x21\x24\xca\x8b\xc5\x6d\x55\xbe\x09\x69\xf7\xa6\x27\xc0\x92\xf4\xe8\xcf\x6f\xfd\x4e\x2e\x65\x5a\x36\x1c\x81\xc1\xfe\x94\xa5\x05\x46\x53\x25\xad\x56\xe5\x04\xa4\x8a\x8c\x55\x1a\xf7\x7d\x8c\x0b\x64\x1c\x35\x08\x09\x17\x68\x0f\xaf\x84\xb1\xe4\x5f\xad\xa4\xc1\x3e\x91\xb5\xc6\x25\x4a\x0b\x29\x4b\x0b\xda\x6f\xd2\x82\x90\x16\x75\x85\x5c\x30\xdd\xba\x88\x88\x14\x5d\x7e\xcf\x84\x61\x49\x89\x84\xed\x0c\x4f\xc0\xea\x06\xbd\x53\x57\xa2\x12\xd6\xc7\x50\x7c\xf3\x15\xc6\xcc\x1d\xe0\x57\x4c\x1b\xab\x34\x94\x2a\x37\xf0\xda\x58\xae\x1a\x7b\x88\x5a\xbf\x19\x91\x5f\x49\x6b\x3d\xf4\x07\xf6\xf5\x3c\x88\x5e\xa9\x7c\x2e\xbe\xe1\x04\x8e\xc6\xe3\xf1\x18\x0e\xe0\x68\x0c\xbf\x9c\x92\x95\x2b\x95\xe7\xfe\x44\x38\x8b\x2a\xcf\xc9\xe3\x12\x97\x58\x9a\x09\x70\x4c\x9a\x9c\x50\x33\x35\x02\xd4\x5a\x69\x27\x78\x45\xcb\x13\xf7\x39\x28\xde\x68\x61\xd1\x3b\xe4\xd2\x23\x0c\xd4\xcc\x16\x31\xe5\x16\xab\xda\xb6\x23\xbf\xc8\xa8\xba\xb5\xb0\x16\x25\x09\x1a\xcb\x51\xeb\xd8\x81\x7c\x6a\x6c\xdd\xd8\x77\xa2\xc4\x09\xec\xef\xef\xed\xcd\xd3\x02\x79\x53\x76\xc7\xf5\xbd\x5a\x81\xca\x82\x9e\x6e\x24\x30\x30\x9d\x04\x08\x8b\xba\x2f\x54\xca\x26\x48\x26\x95\xc1\x54\x49\xee\x62\xd1\x81\x5d\x33\xdb\x45\x61\xdc\x85\x02\xbc\x5c\x6f\xa5\x62\xb2\x75\x91\x76\x7b\xe9\x8c\x50\x68\x95\xc4\x4d\x53\x1d\xec\xb4\x68\xe4\x1d\xe1\xf6\x20\xa5\x92\x39\xa9\xaf\x98\xb0\x90\xa0\x5d\x21\x4a\x68\x6a\xce\x2c\x1a\x48\x30\x53\x1a\xa1\x62\xfa\xce\x9f\x07\xa9\x38\x02\x47\xc6\x1f\xf3\xff\xa3\xe2\x38\x13\x32\x5f\x88\x0a\x55\x63\x27\x54\xe1\x1b\x7b\xa8\x84\x6c\x2c\x6e\x37\x4f\x27\x31\xd8\x70\x31\x67\xda\x8e\x1e\xfa\x40\x55\xb6\x93\x17\x97\x52\xd8\xde\x8b\xb7\xe3\x0d\x37\x7e\x0a\x6e\x98\x20\xdb\x95\xd5\x50\x04\xc1\x8d\xcb\x33\x58\x89\xb2\x84\x04\x1d\x13\x56\x8c\x58\xa3\x2c\x5b\xc8\x51\x52\x78\x91\xfb\x9a\xb8\x3c\x73\xb5\xe0\x51\x88\x95\x18\xe7\x9a\x38\x60\x1b\xdb\x3a\x31\xcf\xf0\x27\x5e\x6c\x8d\x34\x27\x1d\x6b\xad\x73\x80\x8b\xcc\x40\x5e\x31\xac\x04\xdd\x09\x9b\xa4\x15\x07\x25\x8f\xbc\x8d\x3e\x68\x95\xea\xd6\x40\xaa\x91\x9c\x07\xde\x68\x0a\x6a\xad\x15\x31\x16\xfd\xd9\x6d\xb7\x2b\x7e\x21\xfd\x29\xe1\x42\x63\x6a\x95\x6e\xbd\x99\x1b\xa5\xef\xce\x84\xde\x72\xf1\x0c\xa1\xec\x53\x55\x30\x2a\x25\x02\xe3\x25\xba\xbd\x50\xda\x91\x68\x97\x49\x27\x66\x7d\xa2\x46\x20\xac\xf7\xc0\x14\x8d\x05\xae\x56\xb2\xdb\x55\x74\x04\x15\x32\x69\x48\x5c\x23\x5d\x51\x52\x75\x6a\x31\x8c\xbb\x45\xff\x01\x44\xe5\x38\xcc\x62\xd9\x02\xcb\x2c\x7a\x86\xcf\x84\x36\xd6\x9d\x18\x8f\xda\x97\x47\x74\xd4\x85\xe7\xc4\xd5\x83\xf7\x61\x33\xe3\x56\xb7\x54\x95\x1c\x2d\xa6\x16\x56\x05\x73\x84\xaa\x1a\x9d\xa2\xa7\x0b\xb6\x64\xa2\x24\x8e\x24\x31\x61\x63\x08\x90\x67\x98\x09\x49\xa1\xbd\xee\xc5\x85\xdf\xb5\x33\xd5\xdd\xa6\xfe\x32\x52\x4b\xd4\x5a\x70\x34\x3e\xea\x09\x16\x6c\x29\x54\x60\x9e\x1e\xc0\xd7\x2b\x81\x4f\x67\x9f\xcd\x60\x39\x1e\xbe\xd7\x8d\x99\x40\x28\x24\x57\x94\x27\x1f\x06\x39\xc7\xbf\x17\xa7\x83\xf8\x35\xab\x2e\x92\x09\x8c\xe3\x35\x8d\x33\x61\xee\xc0\xd4\x2c\xc5\x27\x14\x49\xe8\x07\xcd\x77\x2e\xc3\xab\xc8\xb1\x33\xd8\x86\x76\x3f\xa8\x6c\x10\xa4\x69\x65\x3a\x54\xf3\xe6\xa3\xa4\xd7\xf8\xf1\x80\xd3\xbf\xcf\x8e\xa4\x3c\x51\xfe\xb4\x79\xba\x83\xe4\xb0\xbd\xd9\xb4\x2f\x0d\x2a\x3f\xcf\x6f\x87\xce\x38\x65\x77\x37\x53\x7d\xb9\x6c\xa7\xe5\xff\xeb\x4b\x6a\xef\x00\x16\x74\x1c\xd7\x5e\x80\x2c\xb5\x62\x89\x6b\x97\x53\xc2\xd2\x3b\x94\x7c\xef\x00\x4e\xfa\x52\x0e\xdf\x3a\x86\x1a\x51\xb9\x72\xa5\x47\x90\xa7\x38\x02\x55\xa3\x34\x96\xa5\x77\x7b\xa7\x5e\x2e\x88\x91\xb9\xf9\x43\xdc\x50\xe9\x9d\xa8\x2b\xe2\xf7\x8b\xa9\xc3\xf3\x21\x5b\x60\x55\x97\x2e\x9f\xff\x0e\x81\x6f\xa4\x58\xa2\x36\x08\xc7\xb0\x64\x52\x94\x25\x0b\x0b\x39\x5a\x94\x4b\x38\x86\x05\xbd\x42\xfc\x37\xff\xde\x70\x6e\x1f\xc3\xfd\x7d\x7c\xde\xff\xfe\xfe\x3d\x88\x30\x9d\x37\x15\x4a\x6b\xe0\xd8\x53\x13\x5d\xd0\x51\x14\x9e\x68\xf7\xf7\x53\xf7\x47\x2f\x5e\xaa\xdc\x43\x05\xba\xfb\xfe\xfd\xd0\xef\x3f\x72\x0f\xa7\xa8\x54\x79\x67\x9b\xd2\xf9\x50\x36\x30\xa3\xcf\x4b\x10\x54\x2e\x31\x8f\x4b\xaa\xc6\x06\x49\x53\xa8\xa6\xe4\xb7\x56\x33\x69\x32\xd4\xb7\x99\x63\xee\x63\xf8\xe7\xf9\x3c\x48\xac\x0a\x94\xb7\x56\x0d\x22\x3d\xf8\xa7\x8f\xb7\xe7\xff\xb8\x5c\xdc\x7e\xba\xbe\x3d\xff\xfb\xe5\x74\x11\x14\xee\xef\x45\x06\x12\x21\x26\x7a\x80\x31\x44\xfd\x4e\xef\xef\x6b\x2d\xa4\xcd\x60\x5f\xe3\x97\x06\x8d\xbd\x4d\x49\xe4\x18\xfe\xc8\xf7\xbd\xf8\x9a\x68\x04\x28\xf9\xda\xef\x00\xea\x48\x84\x98\xe0\x49\xdc\x0a\x2b\xa5\x5b\x42\x8e\xc7\x19\x5c\x9c\xee\x07\xc5\xe7\xf1\x3d\xd7\x3c\x6b\x80\x13\x6f\xad\xc3\x7b\xbd\x2d\xf8\xe1\xc3\x97\x06\xfd\x4b\x76\x76\x3a\x7f\xac\x10\x0f\xfe\x90\x08\x79\x98\x30\x53\x74\x1f\x66\xa7\x73\x88\x3e\x52\x1a\xdd\x6b\x63\xf0\xd7\xaf\xa8\xe7\x13\xec\x05\xf1\xf9\x9a\xd9\x25\x6d\x1e\xac\x74\x55\x6d\x8e\x8f\x26\x75\x2d\x8f\x5f\x2c\x77\x1d\x78\x85\xd5\x31\xc5\x35\x4f\x5e\x2c\x6b\x1d\x34\x55\xf7\x80\xfd\x5c\xca\x1e\x9c\xee\xa7\xcf\xf2\x1e\xc0\x85\x16\xfc\x5c\xe6\x42\xe2\xee\x09\x7e\xf5\x48\x7a\x5f\xed\x96\xdc\x57\x3b\xa5\x96\xc4\xfa\xa4\xfd\x37\xe9\x7e\x05\x51\x8d\x50\xd5\xe2\xe5\x4e\xa8\xf7\xa5\xb8\x5d\x76\x69\xbe\x78\xb9\x2c\x07\xe8\x8c\x1a\xc4\x1e\xfb\x45\xb3\x0c\xf4\xde\x9a\x5f\x7d\xbe\xfe\xf0\x78\x8a\x0f\x1f\xe6\x78\x7e\x7a\xb2\x98\xbe\x87\x28\xfa\x55\x25\x91\xbb\x1c\xb7\x24\xbc\x17\x92\xbe\xd9\x3a\xfa\x61\xc1\xb3\xff\xf3\xc9\xee\x15\x02\x51\x3f\x5b\x45\x3b\x95\x42\x8f\x4a\x94\x1d\xd5\xa8\x23\xd7\x7e\xbf\x5c\x5d\xf4\x06\x2a\xac\x1c\xaf\xbe\x20\x6b\x0f\xe0\xb6\xaa\x07\xf0\x17\x27\x80\xe9\xf9\xa4\x6f\xaa\xfc\x3c\x8d\xa5\xa9\x6a\xa4\xa5\x5e\x88\xa3\xb4\x82\x95\xa6\x1f\xab\x85\x0e\xa6\x56\xc6\x08\xf7\xec\x75\x8f\xb0\xed\xad\x20\x17\x26\xa5\x67\x7b\xd7\x0b\x9e\x78\xdc\xfe\xed\xe5\xd1\x2e\x94\xca\x4b\x84\x69\xa9\x1a\x4e\xfd\xd6\xaf\xd4\x45\x5c\x9e\xfd\x56\x63\x33\x8f\xf4\x98\xa1\x6f\x4a\xfe\xe6\xfd\xfc\x4b\xc9\x61\x23\x37\x28\xf2\xc2\xae\x35\x20\x33\x8d\x19\x6a\xcf\x60\xd4\xcf\x59\xdf\xb7\x43\x53\xc3\x97\x46\xa4\x77\x65\x3b\x3c\xae\x3f\x0e\x42\xae\x61\x2a\x35\x32\xde\x82\x92\xa5\x90\xd4\x25\x2e\x11\x04\xbd\x27\x65\x00\x69\x6a\xf7\x6a\xef\x00\xbc\xa9\xbf\x11\xea\xdc\x2f\x4f\xe0\x28\x1e\x87\xed\xad\x4f\x13\x52\x96\x16\x48\x39\xa7\xc6\x98\xda\xb4\xa6\xb4\x06\x5e\x57\x6e\xce\x85\x50\x0a\x63\x47\x60\x03\x41\x98\x11\xa0\x4d\xdf\x04\x98\x30\x6e\xd0\x98\x69\x34\x45\xdf\xb8\xb8\x99\xd7\x62\x71\xf5\xe8\x40\x63\xef\xc6\xcd\x42\xfd\x0c\x68\x87\x09\xc0\xb3\xfd\xff\xff\xd0\xfd\x3f\xd5\xfb\xbf\x50\xe7\xff\x44\xdf\xdf\x1b\x01\x63\x95\x66\x39\x82\x69\x8d\xc5\xca\xbb\xe6\x56\x17\xfd\x74\xb9\x71\xa5\x66\xd0\x3e\x28\xbe\xa4\xf5\x1b\xeb\x9a\x87\x11\x24\x8d\x85\x56\x35\x50\x51\xe1\x81\x44\xe4\xce\x2d\x87\x27\x32\x5a\xfa\x93\x46\x5f\x2f\x7e\x46\xe5\x87\xd2\x50\x31\xe9\x8c\xb8\xf1\x97\x77\x68\x68\xcc\x52\xe6\x5f\x1c\xc1\x45\x9f\x63\xf7\x79\xa8\xec\x9b\x42\x58\xa4\x52\xa1\x14\xba\xfc\x0c\xa1\x70\xfd\x98\x81\x55\x21\xd2\xa2\x4b\xad\x30\xc0\xca\x52\xad\xc8\x41\x15\x66\xc0\x5d\xe9\x9e\xf8\x85\x33\xa1\xfb\xa3\x03\x10\x41\x7c\xd8\xcd\x1d\xe6\x6f\x3b\xe7\xe8\xfb\xb9\xe4\xb5\x12\xd2\x0e\xdf\x00\x7e\xc1\x76\xfd\xe7\x1c\x53\x8d\x76\xd2\xe9\x5f\xcc\xd7\x17\x77\x27\x39\xff\xff\xa7\xda\x0a\x25\x59\x19\x6f\x92\xc4\xba\xd2\x73\x5c\xb1\x81\x97\x69\x55\xb9\x3c\xa2\x5c\x0a\xad\x24\xb5\x5c\xf1\xda\xfe\xd6\x29\x72\x43\xf1\x64\x2b\xfa\xa6\xf7\x4f\x62\x03\xbc\xd3\xaa\x3a\x97\xcb\xf5\xf9\xf4\x63\xf3\x88\x87\xb3\x08\x6a\x52\xdd\xd5\x49\x7d\x5f\x37\x03\x0d\x53\xfc\x1f\x46\x13\xdb\x66\x05\x3b\x8d\x24\x76\x1b\x47\x3c\x0e\xff\xcc\x18\x82\x54\x3f\xb0\xaf\xa2\x6a\xaa\x61\x33\xeb\xc3\x77\x37\xa4\x1f\x26\xf0\xb4\xc1\xa4\xc9\x88\xc4\x1f\x8c\x7f\xc9\xe4\xa9\x5b\xf9\xbd\xcd\xe4\xff\x13\x00\x00\xff\xff\xe8\x85\x56\xfe\x3f\x1b\x00\x00")

func configDefaultConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_configDefaultConfigYaml,
		"config/default-config.yaml",
	)
}

func configDefaultConfigYaml() (*asset, error) {
	bytes, err := configDefaultConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/default-config.yaml", size: 6975, mode: os.FileMode(420), modTime: time.Unix(1504897091, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGridengineTemplateTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xbd\x4e\xc3\x30\x14\x85\x77\x3f\xc5\x25\xa5\xa3\x9d\xbc\x00\x53\x8b\x2a\x96\x0e\x2c\x8c\x28\xa9\xaf\xa9\xd5\xe4\x3a\xf2\x0f\x42\x5c\xdd\x77\x47\x71\x2b\x24\xa4\xd0\xed\xe8\xe8\x3b\xdf\x70\x36\x0f\xed\xe0\xa9\x1d\xfa\x74\x56\x9b\x47\xd0\x47\x60\x36\xc7\x60\xf1\xc5\x8a\xd4\x26\x2c\xcd\x5b\x88\x97\xbd\x8f\x22\xad\x2b\x44\x38\xea\x94\x6d\x28\xb9\x02\xf8\x1f\x80\x31\x2a\x66\xef\x80\x10\xcc\x6e\x2e\x09\x3a\xd0\x22\x8a\x79\x8e\x9e\xb2\x83\x66\x99\xcf\x08\xd3\xec\x61\x6b\x9b\x2b\x54\x01\x0d\x48\xb6\xa6\xdb\xfc\xb5\x9f\x0e\x03\x74\x66\xcd\x30\xc2\xf9\xfd\x73\xc2\xe9\x69\x6b\x3a\x77\x68\x6e\xf0\xba\x67\xef\xd3\xe5\xae\xc8\x25\xff\x8d\xbf\xa6\x2b\xfe\x47\xa5\x98\xcd\xf3\x17\x9e\x4a\xee\x87\x11\x45\x80\x82\x45\x88\x85\x40\xeb\x53\x20\xe7\x3f\x96\x3f\x76\x35\x89\xa8\x9f\x00\x00\x00\xff\xff\x5e\x79\x39\x0a\x61\x01\x00\x00")

func configGridengineTemplateTxtBytes() ([]byte, error) {
	return bindataRead(
		_configGridengineTemplateTxt,
		"config/gridengine-template.txt",
	)
}

func configGridengineTemplateTxt() (*asset, error) {
	bytes, err := configGridengineTemplateTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gridengine-template.txt", size: 353, mode: os.FileMode(420), modTime: time.Unix(1504897091, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configHtcondorTemplateTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\xcd\x6e\xea\x30\x10\x85\xf7\x7e\x8a\x11\xd2\x5d\x26\x97\x17\xc8\xa6\x10\x21\x36\x45\xa2\xa8\x3f\xab\x28\xc4\xe3\x60\xe1\x8c\x61\xec\x49\x5b\x45\x7e\xf7\x2a\x80\xa8\xa8\x4a\x77\x67\x74\xbe\xf3\x69\x84\x6c\x8f\x1c\x10\x0a\xe8\x6b\xb2\xce\xd5\xaa\xc5\x88\xd4\x43\x01\x1b\x16\x54\xf8\x81\x8d\xc4\x7a\xeb\x46\x64\x18\xf2\xf2\x7a\xa7\xa4\x6a\x6e\xa5\x43\x8a\x01\x0a\x20\xaf\x11\x58\x08\xb2\xac\xf1\x64\x6c\x3b\xd2\xb3\x53\x4a\x49\x39\xdf\x9e\xf7\x2f\x9e\xf7\x73\xcb\x29\xfd\x6f\x3c\x69\xcf\x19\xf6\x48\x31\x73\xbe\x55\xc8\xec\xf9\x27\x65\x84\x08\x5d\x16\xa2\x46\x66\xe5\x25\x1e\x24\xde\x67\xbc\x44\x15\x76\x5e\x9c\xae\x22\xd7\x14\x0c\x72\x65\xac\xc3\xf1\xc1\xb7\xf2\x49\xbd\xef\x90\xaa\xe8\xbf\xcb\xab\x70\xf5\x58\x95\xaf\xcb\x4d\xb5\x5a\x57\xe5\xf3\x72\xb6\x51\xc3\x60\x0d\x10\x42\x3e\x3b\x48\x80\x29\x64\x29\xa9\x61\x38\xb0\xa5\x68\x60\xc2\x78\x14\x0c\xb1\x6a\xc6\xb2\x80\x7f\x7a\x72\x06\x4f\x50\x06\x48\xfa\x94\x2e\x8a\x75\xdd\x2d\xb6\x30\xcd\xef\x59\x3a\xec\x3c\x7f\x8e\x9e\x7c\x6a\x60\xf1\x30\xb9\x4c\x7e\xb7\xcd\x6d\xd8\xff\xa9\xd3\x36\xec\x6f\x64\xe7\xc5\x8d\x4d\x1d\x05\x05\xd5\x57\x00\x00\x00\xff\xff\xcf\x74\x28\x92\x00\x02\x00\x00")

func configHtcondorTemplateTxtBytes() ([]byte, error) {
	return bindataRead(
		_configHtcondorTemplateTxt,
		"config/htcondor-template.txt",
	)
}

func configHtcondorTemplateTxt() (*asset, error) {
	bytes, err := configHtcondorTemplateTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/htcondor-template.txt", size: 512, mode: os.FileMode(420), modTime: time.Unix(1504897091, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configPbsTemplateTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xd0\xc1\x4a\xc4\x30\x10\xc6\xf1\x7b\x9e\x62\xec\xb2\xc7\xa4\xf5\x2a\xf4\xe2\xae\x88\x97\x45\xf4\xe0\xb9\xd9\x4c\xd6\xb0\xed\xa4\x24\x0d\x08\xc3\xbc\xbb\x6c\x5b\x10\xc1\x7a\x1b\x3e\xfe\xfc\x20\xd9\xdd\xd5\x36\x50\x6d\xbb\xfc\xa9\x76\xaf\x8f\xef\xa0\x4f\xc0\x6c\x4e\xd1\xe1\x8b\x13\x59\xb7\x78\xdb\x3e\x62\xba\x1e\x43\x12\xa9\x7d\x21\xc2\x5e\xe7\xc9\xc5\x32\xad\x09\x6e\x25\x98\x92\x62\x0e\x1e\x08\xc1\x1c\xc6\x92\xa1\x01\x2d\xa2\x98\xc7\x14\x68\xf2\x50\x2d\x40\x0f\x14\x1d\xe6\xf6\xfe\x61\x1c\xa9\xdd\xbb\x6a\xa9\xe7\x52\x03\x92\x9b\xaf\xd5\x79\xeb\x86\x67\x0b\x8d\xd9\xa2\x06\x1c\xda\xbd\x69\xfc\xc5\x56\x6b\xfc\xb7\x73\x0c\xf9\xfa\x2f\xe4\x43\x8f\x3f\xd2\x92\xff\xa2\x14\xb3\x79\xfa\xc2\x73\x99\x3a\xdb\xa3\xc8\xfc\x0a\x48\x85\x40\xeb\x73\x24\x1f\x2e\xb7\x8f\x39\xcc\x97\x88\xfa\x0e\x00\x00\xff\xff\xf7\x12\x22\x2a\x70\x01\x00\x00")

func configPbsTemplateTxtBytes() ([]byte, error) {
	return bindataRead(
		_configPbsTemplateTxt,
		"config/pbs-template.txt",
	)
}

func configPbsTemplateTxt() (*asset, error) {
	bytes, err := configPbsTemplateTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/pbs-template.txt", size: 368, mode: os.FileMode(420), modTime: time.Unix(1504897091, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configSlurmTemplateTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xc1\x6a\xf3\x30\x0c\xc7\xef\x79\x0a\x7d\x29\x3d\x3a\xc9\xf7\x08\x6b\x3a\xba\x5d\x76\xd8\x06\x3b\x27\xb5\xbc\x79\xa9\x65\x23\xdb\x30\x30\x7a\xf7\x91\xb4\xd0\x0c\x56\x76\xfb\x23\xfd\xf4\x43\xd2\xe6\x5f\x3b\x5a\x6a\xc7\x21\x7e\x54\x9b\x97\xdd\xdd\x6b\xff\x00\x4a\x7d\xfa\x51\xd1\xe0\x10\x4a\x69\x9e\xbc\xc6\x47\x2d\xb2\x6a\x53\x1a\xe2\x14\xe1\xff\xaa\x84\xcc\x9e\x67\xfc\xcd\xf3\xb4\xb7\x2c\xd2\x9a\x4c\x84\x27\x15\x93\x46\xe6\x15\xea\x73\x0a\x39\xdd\x62\x7d\x4e\x55\x29\xd6\x00\x21\x34\x7d\xc8\x11\x3a\x50\x22\x55\x29\x81\x2d\x25\x03\xf5\xd5\x74\x0c\x39\xaa\x80\xac\xe6\x7d\x60\xab\xeb\xf3\xc4\x42\x2b\x40\xd2\x4b\xba\xb8\x9e\x07\x77\x18\xa1\x6b\x6e\xeb\x1c\x3a\xd8\x36\x9d\x39\xec\xea\x0b\xfe\xbb\x69\x6f\xe3\xf4\x87\x2a\xb9\x70\x55\x9d\xf9\x1f\xae\xaa\x94\xe6\xfe\x0b\x8f\x39\x0d\xe3\x09\x45\x80\xbc\x46\xe0\x4c\xf3\x55\x9e\x8c\x7d\x9f\xff\xd3\x2f\x49\xa4\xfa\x0e\x00\x00\xff\xff\xa8\xf7\x70\x32\xa6\x01\x00\x00")

func configSlurmTemplateTxtBytes() ([]byte, error) {
	return bindataRead(
		_configSlurmTemplateTxt,
		"config/slurm-template.txt",
	)
}

func configSlurmTemplateTxt() (*asset, error) {
	bytes, err := configSlurmTemplateTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/slurm-template.txt", size: 422, mode: os.FileMode(420), modTime: time.Unix(1504897091, 0)}
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
	"config/default-config.yaml":     configDefaultConfigYaml,
	"config/gridengine-template.txt": configGridengineTemplateTxt,
	"config/htcondor-template.txt":   configHtcondorTemplateTxt,
	"config/pbs-template.txt":        configPbsTemplateTxt,
	"config/slurm-template.txt":      configSlurmTemplateTxt,
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
	"config": {nil, map[string]*bintree{
		"default-config.yaml":     {configDefaultConfigYaml, map[string]*bintree{}},
		"gridengine-template.txt": {configGridengineTemplateTxt, map[string]*bintree{}},
		"htcondor-template.txt":   {configHtcondorTemplateTxt, map[string]*bintree{}},
		"pbs-template.txt":        {configPbsTemplateTxt, map[string]*bintree{}},
		"slurm-template.txt":      {configSlurmTemplateTxt, map[string]*bintree{}},
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