// Exercise 1-1
//
// "Modify the 'echo' program [echo3] to also print os.Args[0], the name of the command
// that invoked it." 

// Echo4 prints its command-line arguments including the name of the command itself
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println(strings.Join(os.Args, " "))
}
