package site

import (
	"fmt"
	"net/http"
)

//ServeModule provides a convienient way to serve a Go module over HTTP.
//It returns true and serves the module if the given request is a Go module
//request (has the go-get=1 query).
func ServeModule(module, domain, github string, w http.ResponseWriter, r *http.Request) bool {
	if r.URL.Query().Get("go-get") != "" {
		fmt.Fprintf(w,
			`<html><head><meta name="go-import" content="%[2]v%[1]v git https://%[3]v%[1]v"></head></html>`,
			module,
			domain,
			github,
		)
		return true
	}
	return false
}
