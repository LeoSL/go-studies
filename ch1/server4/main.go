// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server4 is a simple web server that renders animated GIFs and accepts URL params
package main

import (
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		c := r.URL.Query()["cycles"][0]
		lissajous(w, c)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}
