// Echo7 prints its command-line arguments leveraging 'Join' and prints timings
//
// NOTE: Arguments are passed to PringArgs() to support the benchmarking tests
//       in main_test.go.  Se section 11.2.2 for explanation as to how to
//       test a command.  Echo is revisited more extensively in that section
package main

import (
    "fmt"
    "io"
    "os"
    "strings"
    "time"
)

var out io.Writer = os.Stdout // modified during testing

func PrintArgs(args []string) {
    fmt.Fprintln(out, strings.Join(args, " "))
}

func main() {
    start := time.Now()
    PrintArgs(os.Args[1:])
    fmt.Printf("%.8fs elapsed\n", time.Since(start).Seconds())
}
