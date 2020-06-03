// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	records := getRecords()

	for _, url := range records {
		go fetch(url, ch) // start a goroutine
	}
	for range records {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// falls here when dial tcp: lookup rmi.org: no such host
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

// func fetchAndWriteFile(url string, ch chan<- string) {
// 	start := time.Now()
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		ch <- fmt.Sprint(err) // send to channel ch
// 		return
// 	}

// 	f, err := os.Create("/Users/leonardodalimasilva/go/src/book/src/gopl.io/ch1/fetchResult.txt")
// 	defer f.Close()

// 	nbytes, err := io.Copy(f, resp.Body)

// 	resp.Body.Close() // don't leak resources
// 	if err != nil {
// 		ch <- fmt.Sprintf("while reading %s: %v", url, err)
// 		return
// 	}

// 	secs := time.Since(start).Seconds()
// 	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
// }

// func withCopy() {
// 	start := time.Now()
// 	for _, url := range os.Args[1:] {
// 		if !(strings.HasPrefix(url, "http://")) {
// 			url = "http://" + url
// 		}

// 		resp, err := http.Get(url)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
// 			os.Exit(1)
// 		}

// 		// writing to os.Stdout is better because it doesn't require
// 		// an initial buffer large enough to hold the entire stream
// 		_, err = io.Copy(os.Stdout, resp.Body)

// 		resp.Body.Close()
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
// 			os.Exit(1)
// 		}
// 		fmt.Printf("%.10fs elapsed\n", time.Since(start).Seconds())
// 		fmt.Println("REQUEST STATUS: " + resp.Status)
// 	}
// }

//!-
