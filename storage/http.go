package storage

import (
	"context"
	"fmt"
	"io"
	"net/http"
	urllib "net/url"
	"os"
	"strings"
	"time"

	"github.com/ohsu-comp-bio/funnel/config"
	"github.com/ohsu-comp-bio/funnel/util/fsutil"
)

// HTTP provides read access to public URLs.
type HTTP struct {
	client *http.Client
}

// NewHTTP creates a new HTTP instance.
func NewHTTP(conf config.HTTPStorage) (*HTTP, error) {
	client := &http.Client{
		Timeout: time.Duration(conf.Timeout),
	}
	return &HTTP{client}, nil
}

// Stat returns information about the object at the given storage URL.
func (b *HTTP) Stat(ctx context.Context, url string) (*Object, error) {
	resp, cleanUrl, requestURI, err := b.doRequest(ctx, "HEAD", url)
	if err != nil {
		return nil, err
	}

	modtime, _ := http.ParseTime(resp.Header.Get("Last-Modified"))

	return &Object{
		URL:          cleanUrl,
		Name:         requestURI,
		Size:         resp.ContentLength,
		LastModified: modtime,
		ETag:         resp.Header.Get("ETag"),
	}, nil
}

// Get copies a file from a given URL to the host path.
func (b *HTTP) Get(ctx context.Context, url, path string) (*Object, error) {
	resp, cleanUrl, requestURI, err := b.doRequest(ctx, "GET", url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dest, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("httpStorage: creating host file: %s", err)
	}

	_, copyErr := io.Copy(dest, fsutil.Reader(ctx, resp.Body))
	closeErr := dest.Close()

	if copyErr != nil {
		return nil, fmt.Errorf("httpStorage: copying file: %s", copyErr)
	}
	if closeErr != nil {
		return nil, fmt.Errorf("httpStorage: closing file: %s", closeErr)
	}

	modtime, _ := http.ParseTime(resp.Header.Get("Last-Modified"))

	return &Object{
		URL:          cleanUrl,
		Name:         requestURI,
		Size:         resp.ContentLength,
		LastModified: modtime,
		ETag:         resp.Header.Get("ETag"),
	}, nil
}

// Put is not supported by HTTP storage.
func (b *HTTP) Put(ctx context.Context, url string, hostPath string) (*Object, error) {
	return nil, fmt.Errorf("httpStorage: Put operation is not supported")
}

// Join joins the given URL with the given subpath.
func (b *HTTP) Join(url, path string) (string, error) {
	return strings.TrimSuffix(url, "/") + "/" + path, nil
}

// List is not supported by HTTP storage.
func (b *HTTP) List(ctx context.Context, url string) ([]*Object, error) {
	return nil, fmt.Errorf("httpStorage: List operation is not supported")
}

// UnsupportedOperations describes which operations (Get, Put, etc) are not
// supported for the given URL.
func (b *HTTP) UnsupportedOperations(url string) UnsupportedOperations {
	if err := b.supportsPrefix(url); err != nil {
		return AllUnsupported(err)
	}

	ops := UnsupportedOperations{
		List: fmt.Errorf("httpStorage: List operation is not supported"),
		Put:  fmt.Errorf("httpStorage: Put operation is not supported"),
	}

	return ops
}

func (b *HTTP) supportsPrefix(url string) error {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return &ErrUnsupportedProtocol{"httpStorage"}
	}
	return nil
}

func (b *HTTP) doRequest(ctx context.Context, method, url string) (resp *http.Response, cleanUrl string, requestURI string, err error) {
	u, err := urllib.Parse(url)
	if err != nil {
		err = fmt.Errorf("httpStorage: parsing URL: %s", err)
		return
	}

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		err = fmt.Errorf("httpStorage: creating %s request: %s", method, err)
		return
	}

	if u.User != nil && u.User.Username() == "bearer" {
		if token, ok := u.User.Password(); ok {
			req.Header.Set("Authorization", "Bearer "+token)
		}
		u.User = nil
	}

	resp, err = b.client.Do(req)
	if err != nil {
		err = fmt.Errorf("httpStorage: executing %s request: %s", method, err)
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("httpStorage: %s request returned status code: %d", method, resp.StatusCode)
		resp.Body.Close()
		return
	}

	u.User = nil
	cleanUrl = u.String()
	requestURI = u.RequestURI()
	return
}
