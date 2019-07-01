package api

import (
	"fmt"
	"net/http"
)

// HelloWorld is a handler that will print a hello world message
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world - this is a go program!!\n")
}

// Healthz is a handler used by health endpoints
func Healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK!\n")
}
