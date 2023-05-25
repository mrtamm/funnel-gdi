package storage

import (
	"bytes"
	"context"
	"fmt"
	"strings"
        "os/exec"
	"os"
	"github.com/ohsu-comp-bio/funnel/config"
)

// HTSGET provides read access to public URLs.
type HTSGET struct {
	conf config.HTSGETStorage
}

// NewHTSGET creates a new HTSGET instance.
func NewHTSGET(conf config.HTSGETStorage) (*HTSGET, error) {
	return &HTSGET{conf: conf}, nil
}

// Join a directory URL with a subpath.
func (b *HTSGET) Join(url, path string) (string, error) {
        return "", nil
}

// Stat returns information about the object at the given storage URL.
func (b *HTSGET) Stat(ctx context.Context, url string) (*Object, error) {
        return nil, nil
}

// List a directory. Calling List on a File is an error.
func (b *HTSGET) List(ctx context.Context, url string) ([]*Object, error) {
        return nil, nil
}

func (b *HTSGET) Put(ctx context.Context, url, path string) (*Object, error) {
        return nil, nil
}

// Get copies a file from a given URL to the host path.
func (b *HTSGET) Get(ctx context.Context, url, path string) (*Object, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var cmd *exec.Cmd
	if strings.HasPrefix(url, "htsget://bearer:") {
		_bearer_start := len("htsget://bearer:")
		_bearer_stop := strings.Index(url, "@")
		if _bearer_stop < 1 {
			return nil, fmt.Errorf("Bearer token not terminated by @")
		}
		bearer := url[_bearer_start:_bearer_stop]
		url = "htsget://"+url[_bearer_stop+1:]
		cmd = exec.Command("htsget", "--bearer-token", bearer, strings.Replace(url, "htsget://", "https://", 1), "--output", path)
	} else {
		cmd = exec.Command("htsget", strings.Replace(url, "htsget://", "https://", 1), "--output", path)
	}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("Error running htsget: %v %s %s", err, stdout.String(), stderr.String())
	}
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return &Object{
		URL:	      url,
		Size:         info.Size(),
		LastModified: info.ModTime(),
		Name:	      path,
	}, nil
}

// UnsupportedOperations describes which operations (Get, Put, etc) are not
// supported for the given URL.
func (b *HTSGET) UnsupportedOperations(url string) UnsupportedOperations {
	if err := b.supportsPrefix(url); err != nil {
		return AllUnsupported(err)
	}

	ops := UnsupportedOperations{
                List: fmt.Errorf("htsgetStorage: List operation is not supported"),
                Put:  fmt.Errorf("htsgetStorage: Put operation is not supported"),
		Join: fmt.Errorf("htsgetStorage: Join operation is not supported"),
		Stat: fmt.Errorf("htsgetStorage: Stat operation is not supported"),
        }
	return ops
}

func (b *HTSGET) supportsPrefix(url string) error {
	if !strings.HasPrefix(url, "htsget://")  {
		return &ErrUnsupportedProtocol{"htsgetStorage"}
	}
	return nil
}

// htsgetclient exists implements the storage API and reuses an HTSGET client
// for recursive calls.
type htsgetclient struct {
}

// Get copies a file from a given URL to the host path.
func (b *htsgetclient) Get(ctx context.Context, url, path string) (*Object, error) {
	return nil, nil
}
