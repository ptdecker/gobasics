// Echo8 [named Echo4 in the text] prints its command line arguments and supports
// two optional flags:  '-n' omits the trailing new line and '-s sep' separates the 
// argumentes with the contents of string 'sep' instead of a space
package main

import (
    "flag"
    "fmt"
    "strings"
)

var n   = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
    flag.Parse()
    fmt.Print(strings.Join(flag.Args(), *sep))
    if !*n {
        fmt.Println()
    }
}

