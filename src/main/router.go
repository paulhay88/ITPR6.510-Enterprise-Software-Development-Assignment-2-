package main

import (
	"fmt"
	"net/http"
)

// MyRouter used to call ServeHTTP
type MyRouter struct {
}

func (p *MyRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		meetings(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

// Any Routing functions...

func meetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}
