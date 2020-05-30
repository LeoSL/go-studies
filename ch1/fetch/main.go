// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	// withReadAll()
	withCopy()
}

// func withReadAll() {
// 	start := time.Now()
// 	for _, url := range os.Args[1:] {
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
// 			os.Exit(1)
// 		}
// 		b, err := ioutil.ReadAll(resp.Body)

// 		resp.Body.Close()
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
// 			os.Exit(1)
// 		}
// 		fmt.Printf("%s", b)
// 		fmt.Printf("%.10fs elapsed\n", time.Since(start).Seconds())
// 	}
// }

func withCopy() {
	start := time.Now()
	for _, url := range os.Args[1:] {
		if !(strings.HasPrefix(url, "http://")) {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// writing to os.Stdout is better because it doesn't require
		// an initial buffer large enough to hold the entire stream
		_, err = io.Copy(os.Stdout, resp.Body)

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%.10fs elapsed\n", time.Since(start).Seconds())
		fmt.Println("REQUEST STATUS: " + resp.Status)
	}
}

//!-
