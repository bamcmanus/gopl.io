// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	var s string
	sep := " "
	for i := 0; i < len(os.Args); i++ {
		//fmt.Printf("index: %d; value: %q\n", i, os.Args[i])
		s += os.Args[i] + sep
	}
	fmt.Println(s)
	fmt.Printf("dumb for loop took %dns\n", time.Since(start).Nanoseconds())

	start = time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Printf("join took %dns\n", time.Since(start).Nanoseconds())

}

//!-
