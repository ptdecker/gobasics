// Echo6 prints its command-line arguments leveraging 'range' with timings
//
// NOTE: Arguments are passed to PrintArgs() to support the benchmarking tests
//       in main_test.go.  See section 11.2.2 for explaniation as to how to
//       test a command.  Echo is revisited more extensively in that section.
package main

import (
    "fmt"
    "io"
    "os"
    "time"
)

var out io.Writer = os.Stdout // modified during testing

func PrintArgs(args []string) {
    s, sep := "", "" 
    for _, arg := range args { 
        s += sep + arg 
        sep = " "
    }
    fmt.Fprintf(out, s)
}

func main() {
    start := time.Now()
    PrintArgs(os.Args[1:]);
    fmt.Printf("%.8fs elapsed\n", time.Since(start).Seconds())
}
