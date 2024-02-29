// Package examples Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// examples/capture-stdout-stderr.json
// examples/google-storage.json
// examples/resource-request.json
// examples/input-content.json
// examples/hello-world.json
// examples/log-streaming.json
// examples/md5sum.json
// examples/s3.json
// examples/full-hello.json
package examples

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
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

var _examplesCaptureStdoutStderrJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\x3f\x0f\xda\x30\x10\xc5\xf7\x7c\x8a\x27\xcf\x94\x4c\x5d\x98\xdb\x4a\x48\xdd\x18\x2b\x06\x2b\x3e\x88\xa5\xf8\x8f\xec\x73\x05\xa2\x7c\xf7\xea\x9c\xa4\x01\x54\x10\x2c\x11\xba\x7b\xf7\x78\xbf\x3b\x5f\x1a\x40\x79\xed\x48\x6d\xa0\x76\x6c\x42\x61\x68\x6f\xb0\x63\x43\x29\xa1\xd3\x91\x4b\x22\xb5\x12\x99\xa1\xdc\x25\x1b\xd9\x06\x2f\xea\x6f\xe4\x82\xcf\x9c\x34\x53\x9e\x84\xd6\x1f\x91\x17\x93\x3c\x9a\xc8\xcf\x12\x87\xa0\x8d\xf4\xb9\x27\x07\x9d\x11\x0a\xc7\xc2\x38\xd8\x81\xf2\x7a\xfc\x03\xeb\x63\xe1\xac\x36\xf8\xd5\x00\xc0\xa5\x7e\x6f\xf2\xd5\x7e\x95\xd6\xf2\x43\x9e\xad\x74\xc1\x01\x5d\x70\x4e\x7b\xb3\x46\xcb\x2e\xb6\x75\x08\xae\x64\x06\x9d\x6c\x66\x04\x2f\x19\xd0\x87\xcc\xc8\xe7\xcc\xe4\xd6\x8b\x67\x49\x83\x78\x49\xaa\x4d\xdb\x2e\x06\x8b\x82\xcf\xb1\x86\xf9\xb1\xfd\xf9\x7d\xa9\x46\xcd\xbd\x54\xa7\x09\x55\xeb\xd7\x06\xd8\x57\xb2\x11\xf6\x05\xda\xb8\xb5\xa7\x6c\xd3\x65\xc2\x61\x86\x83\x9d\x77\x4e\xa6\x2e\xb8\x0b\xd1\x92\x11\xfc\x1a\x61\xba\xc2\xa7\xac\x8f\x31\xde\x80\x15\xfd\x48\xbb\x7a\x8a\x46\x29\xbd\x42\x93\x47\xf2\x11\x5a\x1d\xf8\x1c\xed\x2e\xc6\x1b\x68\xa2\x7f\x38\x24\x9d\xa8\x2b\x1c\xd2\x7f\x4e\x69\x9d\x3e\x56\x47\x3d\x44\xeb\x69\xf1\x9c\xb8\x64\x42\xe5\x5e\xad\xa0\xbe\x74\xf2\x75\xe6\x6b\x2e\x6e\x7e\xa5\xf8\x03\x26\x42\x6b\xe8\xf7\x9c\x75\xff\xcf\x62\xba\xcb\xed\xce\x9f\xf5\xee\x43\x37\xd7\xe6\x6f\x00\x00\x00\xff\xff\x94\x21\xf6\x91\xe3\x03\x00\x00")

func examplesCaptureStdoutStderrJsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesCaptureStdoutStderrJson,
		"examples/capture-stdout-stderr.json",
	)
}

func examplesCaptureStdoutStderrJson() (*asset, error) {
	bytes, err := examplesCaptureStdoutStderrJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/capture-stdout-stderr.json", size: 995, mode: os.FileMode(420), modTime: time.Unix(1548721562, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _examplesGoogleStorageJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xcd\x4a\x03\x31\x14\x85\xf7\xf3\x14\x87\xac\x14\x6a\x67\x5c\x54\xb0\x5b\x05\x37\x82\x62\xed\x4a\xba\xb8\x4d\x32\x63\x30\x7f\x4c\x6e\xb0\xb4\xf4\xdd\x25\x29\x75\xf0\x67\x33\x0c\xe4\x7c\xe7\xdc\xef\xd0\x00\xc2\x93\xd3\x62\x09\xf1\x10\xc2\x60\x35\xee\x6c\xc8\x0a\x2b\x0e\x23\x0d\x1a\x7a\x47\x2e\x5a\x2d\x66\x25\xa9\x74\x92\xa3\x89\x6c\x82\x2f\xc0\x2b\xa5\x0f\x18\x1f\x33\x27\x90\x57\x08\x99\xeb\xbf\x24\x8f\xad\xc6\xbf\x7d\xeb\x97\xc7\x34\x3f\xb5\xe9\x9d\x96\x99\xc3\x98\xc4\x12\x6f\x0d\x00\x1c\xea\x17\x10\xc6\xd1\x50\x6f\xca\xdb\xec\x39\xd7\x7c\x7d\x90\xc1\x39\xf2\xaa\x10\xc2\xa9\x45\xca\x4e\xcc\x20\x5a\x76\xb1\xed\x8d\xd5\xf3\xbd\x89\x62\x53\xc3\xc7\x06\xd8\xd4\x9d\xd3\x85\x7f\x47\xce\xde\x7b\x13\x0b\x3b\x8d\xfc\xd2\xbc\x0f\x9f\xde\x06\x52\x20\xc4\xbc\xb5\x46\xa2\xc4\xd1\x8f\xc1\x9d\x1d\xcf\x76\x17\xeb\x15\x9e\x89\xb5\x67\x3c\xf5\xbd\x91\x1a\x8a\x98\x2e\xa7\xea\x3c\xda\x52\x39\xa4\x65\xdb\xe6\x14\x39\x5c\x45\x32\x63\x4b\x31\x5a\x23\xa9\x2c\xa6\xb6\x5b\xdc\x76\x5d\x77\x7d\x53\x65\xbe\xc9\x48\xfc\x5e\xd0\x9f\xae\x93\x6a\x73\x6c\xbe\x02\x00\x00\xff\xff\x40\x00\x1c\x30\xce\x01\x00\x00")

func examplesGoogleStorageJsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesGoogleStorageJson,
		"examples/google-storage.json",
	)
}

func examplesGoogleStorageJson() (*asset, error) {
	bytes, err := examplesGoogleStorageJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/google-storage.json", size: 462, mode: os.FileMode(420), modTime: time.Unix(1548721562, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _examplesResourceRequestJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8f\x39\x6b\x03\x31\x10\x85\x7b\xfd\x8a\xc7\xd4\x22\xec\x82\xdd\x6c\x97\x03\x5c\x05\x82\x21\x55\x30\x61\xa2\x1d\x8c\xb0\x75\x44\xa3\x85\x40\xd8\xff\x6e\xa4\xb5\x4b\xbd\xe3\x7b\x9a\x7f\x03\x50\xe4\x20\x34\x81\x8e\xa2\x69\x29\x4e\x50\xe4\x77\x11\xad\x64\x9b\x3b\x8b\xba\xe2\x73\xf5\x29\xb6\xd0\x9b\x84\x14\xb5\x16\xae\xa2\x60\x54\xd6\xcb\x23\xef\xe3\x19\x3b\xbc\x7e\x7c\xc2\xa5\x22\x6a\x31\xee\x71\x78\xc1\xf1\xf9\xdd\x82\xe3\x8c\x71\x18\xda\x7b\xf6\x7a\x81\x66\x76\xf2\xb4\x2d\x94\xfb\xae\xd2\x84\xf6\x21\x80\x5c\x5e\xbe\x3b\x84\x26\xec\xec\xa6\x15\x0e\x87\x1f\x9a\x30\xee\xef\x42\x03\x6d\xca\x30\x18\x60\xed\x30\xf9\x13\xb7\xd4\x54\x5a\xf3\xab\xe7\x36\x24\x40\x3e\xf0\xb9\x1f\xca\xd7\xec\xa3\x90\x7d\x18\x2e\x85\xc0\x71\x6e\x0d\xd2\xab\x48\x26\x0b\x1a\xe9\xd4\xfd\xd5\x00\x27\xb3\x9a\x5b\x00\x00\x00\xff\xff\xc0\xca\xf1\x83\x2b\x01\x00\x00")

func examplesResourceRequestJsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesResourceRequestJson,
		"examples/resource-request.json",
	)
}

func examplesResourceRequestJson() (*asset, error) {
	bytes, err := examplesResourceRequestJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/resource-request.json", size: 299, mode: os.FileMode(420), modTime: time.Unix(1548721562, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _examplesInputContentJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x51\x3f\x4f\xeb\x30\x10\xdf\xf3\x29\x7e\xca\xd2\xa5\x6a\xa6\xb7\x74\x7e\xef\x89\x4a\x6c\x0c\x0c\x50\x21\x13\x5f\x5a\x4b\xb6\x2f\xb2\x2f\x82\xaa\xea\x77\x47\xbe\x04\x52\xaa\xaa\xb0\x64\xf0\xfd\xfe\xe7\x58\x01\x75\x34\x81\xea\x35\xea\x4d\xec\x07\x41\xe7\x3c\xa1\xe5\x28\x14\x05\x26\x5a\xf0\x20\x9f\xef\xf5\xb2\xe0\x2d\xe5\x36\xb9\x5e\x1c\xc7\x42\xfb\x4b\x81\x63\x96\x64\x84\x32\x86\xec\xe2\x0e\xb2\x27\x2c\x26\x8d\x05\x3a\x47\xde\xa2\xe3\x04\x57\x1c\x32\x84\xd1\x26\x32\x42\x30\xa3\x1d\x47\xa5\xec\x39\x0b\xf2\x21\x0b\x85\xd1\x69\xc4\xd7\x6b\x3c\x55\x00\x70\xd4\xef\x59\xe2\xd6\xc8\xa8\xa9\x70\x3d\x5d\x84\x1b\x3b\x09\x23\xd8\x3f\x79\x08\x2b\x34\x12\xfa\xc6\x45\xbc\x39\xef\xf1\x4a\x53\x10\x7b\x25\xc2\x6a\x16\x95\x43\xaf\x7e\xff\x37\xf7\xff\xe6\xd7\xde\xc8\xbe\xbc\x4e\x92\xf3\x61\x6a\x5e\x6e\x77\xe4\x3d\xe3\x91\x93\xb7\xcf\xb1\x56\xc0\xa9\x02\xb6\x5a\x6f\x5c\xf6\x87\x7e\x59\x2c\xdf\x28\xf8\xa0\x67\x70\x07\x1d\x23\xa3\x35\xbd\x0c\x89\x74\x65\x4d\x26\x94\xe5\x45\x31\xb7\x2a\x0e\xc9\x17\xb9\xf2\x3b\xd6\x4d\xa3\xc4\xd6\x28\xef\xdb\xbc\xbf\x58\xa2\xa4\xbd\x28\x4a\xef\xd4\x0e\xc2\xe9\x4a\x55\x17\xcc\x4e\x15\x8d\xef\x5d\xa4\xf3\x11\x43\x30\xd1\x16\x46\xd9\xa1\x5e\xce\x43\x6f\xbf\x40\xd3\x38\xd7\xad\xab\x53\xf5\x11\x00\x00\xff\xff\x9b\xe3\x40\x40\xe0\x02\x00\x00")

func examplesInputContentJsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesInputContentJson,
		"examples/input-content.json",
	)
}

func examplesInputContentJson() (*asset, error) {
	bytes, err := examplesInputContentJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/input-content.json", size: 736, mode: os.FileMode(420), modTime: time.Unix(1548721562, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _examplesHelloWorldJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\x31\xce\xc2\x30\x0c\x46\xf7\x9c\xe2\x93\xe7\xea\x3f\x40\xe7\x7f\xe0\x0e\xa8\x83\x49\x2d\x1a\x91\xc4\x55\x6c\x04\x52\xd5\xbb\xa3\x84\x85\xc5\xcb\x7b\xcf\xf6\x11\x00\xaa\x5c\x84\x66\xd0\x45\x72\x56\xbc\xb4\xe5\x95\xa6\x0e\x56\xb1\xd8\xd2\xee\x49\x6b\xe7\xff\x52\xb4\x9a\x37\x76\x31\xf8\x26\x28\x6a\x8e\x1b\x5b\x8a\x90\xb8\x29\x9c\xed\xf1\xf7\x4d\xe5\x2d\xf1\xe9\xda\x8c\x66\x5c\x03\x00\x1c\x63\x02\x94\x0a\xdf\xc7\x3d\xce\x7b\xaa\x32\xfc\x01\xa2\x96\xc2\x75\xed\x05\xf5\x7d\x34\x81\xb6\x9f\x9f\x96\x21\x9e\x01\x58\xc2\x19\x3e\x01\x00\x00\xff\xff\x1b\x6b\x00\xfc\xbb\x00\x00\x00")

func examplesHelloWorldJsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesHelloWorldJson,
		"examples/hello-world.json",
	)
}

func examplesHelloWorldJson() (*asset, error) {
	bytes, err := examplesHelloWorldJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/hello-world.json", size: 187, mode: os.FileMode(420), modTime: time.Unix(1548721562, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _examplesLogStreamingJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8f\xb1\x6e\xf3\x30\x10\x83\x77\x3f\x05\xa1\x39\x7f\x82\x7f\x4d\xd6\xa2\x53\x87\xee\x45\x86\x8b\x44\xd8\x02\xa4\x93\xa1\x3b\xb7\x0d\x82\xbc\x7b\x61\xb5\x0b\x87\xfb\x78\x04\xf9\x98\x80\xa0\x52\x19\xce\x08\x6f\x6d\x86\x79\xa7\xd4\xac\x73\x38\xec\x28\xd1\x62\xcf\xab\xe7\xa6\xbb\xe3\x85\xb5\xa9\x79\x17\xa7\xc1\x17\x71\x98\xa7\xb6\xf9\x89\xbd\xa3\xb4\xd9\x20\x9d\x7f\x19\x4c\xb8\xdd\xe1\x0d\xbe\x10\xaf\x9b\x2a\x0b\x92\xd8\x72\x6b\xd2\x93\x1d\xf1\xde\xb3\xba\x0d\x9a\xc4\x79\xf2\x5c\x09\x7e\xb2\xdf\x61\x8c\x4d\xd3\x01\x9b\x7a\x2e\x88\xa2\x91\x85\xe9\xf8\x5b\x89\xdf\x8c\x9b\xb7\x6e\xe1\x8c\x8f\x09\x00\x1e\x43\x81\x90\xab\xcc\x63\x89\x94\x35\x2b\x87\x7f\x80\xd8\x6a\x15\x4d\xfb\x47\xb0\x25\x1c\x10\xfe\xc5\x5d\xbf\x96\x5c\x08\xef\x1b\x2f\x48\x6d\xf4\xb8\xc0\x0a\xb9\xe2\xff\x7e\x51\x86\xeb\xc8\x78\x4e\xc0\x75\x7a\x4e\x3f\x01\x00\x00\xff\xff\x88\xae\x93\xb2\x30\x01\x00\x00")

func examplesLogStreamingJsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesLogStreamingJson,
		"examples/log-streaming.json",
	)
}

func examplesLogStreamingJson() (*asset, error) {
	bytes, err := examplesLogStreamingJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/log-streaming.json", size: 304, mode: os.FileMode(420), modTime: time.Unix(1548721562, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _examplesMd5sumJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\xcf\x6a\xf3\x30\x10\xc4\xef\x7e\x8a\x41\xe7\x10\x9f\xbe\x4b\xce\x5f\x0b\x81\xde\x7a\x2c\x21\x08\x6b\x93\x08\xac\x3f\x78\x57\xe0\x10\xf2\xee\x45\x6b\xa7\x86\x24\xb4\xb4\x17\x63\x6b\x77\x66\xf7\x37\xf2\xa5\x01\x4c\xb4\x81\xcc\x06\x26\xb8\x7f\x5c\x02\x68\xb4\x21\xf7\x64\x56\xb5\xe6\x88\xbb\xc1\x67\xf1\x29\xd6\x96\xff\x14\x52\x64\x19\xac\x10\xc3\xc7\x5c\x04\x36\x3a\xa4\x22\xf5\xf5\xe0\x7b\x62\x14\xf6\xf1\x08\x0b\xf6\xd5\x06\xb3\x6b\x97\x42\xb0\xd1\xad\x27\x5b\x95\xb2\xd9\xe0\xa3\x01\x80\x8b\x3e\x1f\x57\xd1\x36\x55\x68\xf5\x6e\x99\xad\xce\x97\x34\x8f\x58\xa3\x95\x90\xdb\xe9\x63\x3f\x2d\x17\x0a\x0b\x68\xf4\x2c\x48\x11\x72\x22\x9c\x12\x0b\xf8\xcc\x42\x61\xbd\x38\x97\xa1\xaf\x8e\x15\x60\xd3\xb6\x0f\x3e\x4b\xa3\x9c\xb3\x2e\xf8\xba\x7d\x7b\x59\x4e\xb3\x95\x53\x3d\x55\xa1\x8f\x46\xcf\xaf\x0d\xb0\x53\xda\x29\x9e\x9f\x71\x59\x5c\xfa\x86\xf7\x5d\xcb\x48\x87\x5b\xa6\x9e\xd1\xd9\x2c\x65\x20\xa7\xd7\xd0\xa5\xec\xc9\xd5\x44\x6e\x04\xfb\xf9\x66\xfe\x08\x3f\xa9\x7f\x45\x5f\x01\xee\xf0\x69\xa4\xae\x48\x1a\x9e\x04\xe0\x83\x3d\xaa\xa3\xed\xb3\x8f\xb4\x78\xce\xbf\x4b\x55\xcc\xe9\x98\xd5\x92\xef\xee\xab\x6f\x8e\xec\xf9\xf4\xe6\xda\x7c\x06\x00\x00\xff\xff\xb9\xf5\xf2\xe7\xe1\x02\x00\x00")

func examplesMd5sumJsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesMd5sumJson,
		"examples/md5sum.json",
	)
}

func examplesMd5sumJson() (*asset, error) {
	bytes, err := examplesMd5sumJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/md5sum.json", size: 737, mode: os.FileMode(420), modTime: time.Unix(1548721562, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _examplesS3Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\x4b\x6b\x7a\x31\x14\xc4\xf7\xf7\x53\x0c\x59\xfd\xff\xa0\xde\xd8\x5b\x17\x71\xdb\x6e\x0a\x5d\x69\xbb\x2a\x52\x8e\x37\xd1\x86\xe6\x45\x1e\x54\x10\xbf\x7b\x49\xa4\xda\xc7\x26\x04\x66\xce\xcc\xf9\x9d\x63\x07\x30\x47\x56\xb1\x25\xd8\x7a\xc0\x3a\xfb\x48\x7b\x05\x75\x20\x1b\x8c\x62\x93\xaa\x4b\x95\xc6\xa8\x43\xd6\xde\x55\xdb\x13\xa5\x77\x68\x17\x4a\x4e\x20\x27\xe1\x4b\x6e\xff\x91\x1c\xb6\x0a\x77\xc6\x17\x79\x09\x7a\x5e\x3d\xa6\xd9\x39\x46\x1d\xd4\x58\xb2\x8f\x89\x2d\xf1\xd2\x01\xc0\xb1\xbd\x00\xd3\x96\xf6\x6d\x85\xb2\x2d\x2e\x97\xe6\x6f\xc2\xe8\xad\x25\x27\xeb\x04\xb3\x72\x91\x8a\x65\x13\xb0\x3e\xdb\xd0\xef\xb4\x51\xb3\x83\x35\x6c\xd3\xcc\xa7\x0e\xd8\xb4\x9e\xf3\x6a\x7f\x4b\xbe\x30\x9b\x7e\xad\xf8\x45\x77\xef\x3f\x9c\xf1\x24\x41\x08\x65\x6b\xf4\x88\x5a\x84\x5d\xf4\x16\xdf\x0e\xf4\xef\x61\xb5\xc6\xce\x47\x0b\x21\x38\x24\x65\xfa\x7f\x8d\x2c\xd1\xd4\xa8\x34\x2c\xfb\x5e\xc7\x34\xad\xbe\xa9\x10\xbc\xbf\xe1\x5c\x0c\xf3\x41\x0c\xb7\x62\xe0\x9c\xcf\x17\xfc\xf5\x5c\xd2\x40\x2e\xf3\x81\xf2\x5b\x0d\xf8\xc9\x79\xc5\xec\x4e\xdd\x67\x00\x00\x00\xff\xff\x7c\xcd\x05\xad\xb9\x01\x00\x00")

func examplesS3JsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesS3Json,
		"examples/s3.json",
	)
}

func examplesS3Json() (*asset, error) {
	bytes, err := examplesS3JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/s3.json", size: 441, mode: os.FileMode(420), modTime: time.Unix(1548721562, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _examplesFullHelloJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x91\x4f\x6b\xdc\x30\x10\xc5\xef\xfe\x14\x83\xce\x71\x44\x7b\x4c\xa1\x97\x66\x43\x7a\xa9\xc1\x6c\x28\xa5\x84\x32\xd6\x8e\x57\xa2\x92\xc6\xe8\xcf\x66\x4b\xd9\xef\x5e\x34\x5e\x97\xc0\xfa\x20\xec\xa7\xe7\xdf\x3c\xe9\xfd\xed\x00\x54\xc4\x40\xea\x01\xd4\x53\xf5\x1e\x9e\xc9\x7b\x86\x37\x4e\xfe\xa0\xee\xda\xee\x81\xb2\x49\x6e\x29\x8e\x63\x33\x3d\x52\xe0\x98\x4b\xc2\x42\x19\x8a\x25\x08\x9c\x0b\x4c\x98\x9d\x01\x32\x96\xa1\x60\xfe\x0d\x6f\xae\x58\x40\xef\xc5\x41\xe7\x92\x30\xdf\xaf\xb8\x82\xc7\xac\x1e\xa0\x0d\x5e\xbf\xfa\xe1\xdb\xae\x81\xf7\xdf\x07\x71\x5c\xd5\xfd\xf3\xb8\x13\xfd\x69\x78\x19\x55\x07\x70\x91\xdf\xb9\x96\xa5\x96\x46\xf8\x79\x45\xd4\xe4\x9b\x4d\xbf\x64\x4a\x59\x4f\xd5\x58\x8c\x48\x3a\x27\xa3\x8f\xae\xd8\x3a\xdd\x1b\x0e\x9a\x6d\xae\xbd\xe1\xb0\xf4\x93\x63\x3d\xd7\x18\xc9\xeb\x5f\xd9\x24\x2c\xc6\xea\xb9\x7a\xdf\xaf\x68\xbd\x85\x58\xb0\x58\x01\x6f\x23\xb7\x70\x7f\x16\xb9\xad\xc7\xaf\xe3\xee\xcb\x7e\x18\x7f\x48\xb8\x57\x49\x77\x62\x5f\x03\x49\x3a\xa5\x4f\xec\x3f\xa8\x3b\x90\x97\x8f\x6a\x35\xd0\x99\x4c\x2d\x9c\xc4\x22\xbc\xf5\x14\x00\xca\x05\x3c\x0a\x18\xfd\xe2\x22\x5d\xc7\x01\x28\xc3\x21\x60\x3c\x08\x34\xdb\x46\xec\x4d\x5b\xe5\xb6\x67\x66\xf8\x0c\x5b\x46\x3d\x33\x7f\x5a\x6b\x98\x30\xbd\xdf\x98\x30\xad\x11\x04\x49\xf1\xf4\xbf\x03\x11\x6e\x3a\x10\xf5\xa6\x83\xf6\x5c\xba\x6d\x7d\xed\x2e\xdd\xbf\x00\x00\x00\xff\xff\x12\xe1\x02\xb9\x41\x02\x00\x00")

func examplesFullHelloJsonBytes() ([]byte, error) {
	return bindataRead(
		_examplesFullHelloJson,
		"examples/full-hello.json",
	)
}

func examplesFullHelloJson() (*asset, error) {
	bytes, err := examplesFullHelloJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "examples/full-hello.json", size: 577, mode: os.FileMode(420), modTime: time.Unix(1548721562, 0)}
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
	"examples/capture-stdout-stderr.json": examplesCaptureStdoutStderrJson,
	"examples/google-storage.json":        examplesGoogleStorageJson,
	"examples/resource-request.json":      examplesResourceRequestJson,
	"examples/input-content.json":         examplesInputContentJson,
	"examples/hello-world.json":           examplesHelloWorldJson,
	"examples/log-streaming.json":         examplesLogStreamingJson,
	"examples/md5sum.json":                examplesMd5sumJson,
	"examples/s3.json":                    examplesS3Json,
	"examples/full-hello.json":            examplesFullHelloJson,
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
	"examples": {nil, map[string]*bintree{
		"capture-stdout-stderr.json": {examplesCaptureStdoutStderrJson, map[string]*bintree{}},
		"full-hello.json":            {examplesFullHelloJson, map[string]*bintree{}},
		"google-storage.json":        {examplesGoogleStorageJson, map[string]*bintree{}},
		"hello-world.json":           {examplesHelloWorldJson, map[string]*bintree{}},
		"input-content.json":         {examplesInputContentJson, map[string]*bintree{}},
		"log-streaming.json":         {examplesLogStreamingJson, map[string]*bintree{}},
		"md5sum.json":                {examplesMd5sumJson, map[string]*bintree{}},
		"resource-request.json":      {examplesResourceRequestJson, map[string]*bintree{}},
		"s3.json":                    {examplesS3Json, map[string]*bintree{}},
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
