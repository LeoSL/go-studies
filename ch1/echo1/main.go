// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func echo() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + strconv.Itoa(i) + os.Args[i]
		sep = "\n"
	}
	fmt.Println(s)
}

func main() {
	start := time.Now()
	echo()
	fmt.Printf("%.10fs elapsed\n", time.Since(start).Seconds())
}

//!-
