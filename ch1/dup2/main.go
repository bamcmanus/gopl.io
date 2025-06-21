// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	counts := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for _, files := range counts {
		if len(files) > 1 {
			fmt.Println(strings.Join(files, " "))
		}
	}
}

func countLines(f *os.File, lineCounts map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		if _, ok := lineCounts[text]; !ok {
			lineCounts[text] = []string{}
		}
		lineCounts[text] = append(lineCounts[text], f.Name())
	}
	if input.Err() != nil {
		log.Fatalln("input failed to scan")
	}
}

