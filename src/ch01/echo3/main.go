// Echo3 prints its command-line arguments leveraging strings.Join
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println(strings.Join(os.Args[1:], " "))
}
