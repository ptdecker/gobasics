// Echo7 prints its command-line arguments leveraging 'Join' and prints timings
package main

import (
    "fmt"
    "os"
    "strings"
    "time"
)

func PrintArgs() {
    fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
    start := time.Now()
    PrintArgs()
    fmt.Printf("%.8fs elapsed\n", time.Since(start).Seconds())
}
