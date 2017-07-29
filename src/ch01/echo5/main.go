// Exercise 1-2
//
// Modify the 'echo' program [echo3] to print the index and value of each of its
// arguments one per line.

// Echo5 prints its command-line arguments one per line prepended by the argument number
package main

import (
    "fmt"
    "os"
)

func main() {
    for i, arg := range os.Args[1:] { 
        // Note: Printf is introduced in section 1.3 of the book so using
        //       it here for a section 1.2 exercise is a forward reference
        fmt.Printf("%d:\t%s\n", i + 1, arg)
    }
}
