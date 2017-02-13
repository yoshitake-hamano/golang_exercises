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
    "os"
)

type LineRecord struct {
  counts    int
  filenames []string
}

func main() {
    records := make(map[string]LineRecord)
    files   := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, "-", records)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, arg, records)
            f.Close()
        }
    }
    for line, record := range records {
        if record.counts > 1 {
            fmt.Printf("%d\t%s\n", record.counts, line)
            for _,filename := range record.filenames {
                fmt.Printf("\t%s\n", filename)
            }
        }
    }
}

func countLines(f *os.File, filename string, records map[string]LineRecord) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        record := records[input.Text()]
        record.counts++
        record.filenames = append(record.filenames, filename)
        records[input.Text()] = record
    }
    // NOTE: ignoring potential errors from input.Err()
}

//!-
