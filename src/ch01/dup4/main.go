// Dup4 prints names of all files in which a duplicate line appears.
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {

    counts := make(map[string]int)
    files  := os.Args[1:]

    if len(files) == 0 {
        countLines(os.Stdin, counts)
	    containsDups(nil, counts)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts)
            containsDups(f, counts)
            f.Close()
        }
    }
}

func countLines(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
    }
    // NOTE: Potential errors from input.Err() are ignored
}

func containsDups(f *os.File, counts map[string]int) {
    for value := range counts {
        if counts[value] > 1 {
            if f == nil {
                fmt.Println("stdio")
            } else {
                fmt.Println(f.Name())
            }
            return
        }
    }
}
