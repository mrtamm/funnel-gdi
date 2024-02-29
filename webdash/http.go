package webdash

import (
	"fmt"
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
)

var fs = &assetfs.AssetFS{
	Asset:     Asset,
	AssetDir:  AssetDir,
	AssetInfo: AssetInfo,
	//Prefix:    "webdash/build",
}

var index = MustAsset("index.html")

// FileServer provides access to the bundled web assets (HTML, CSS, etc)
// via an http.Handler
func FileServer() http.Handler {
	return http.FileServer(fs)
}

// RootHandler returns an http handler which always responds with the /index.html file.
func RootHandler() http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if _, err := resp.Write(index); err != nil {
			fmt.Printf("Detected error while writing HTTP response: %s\n", err)
		}
	})
}
