// Echo6 prints its command-line arguments leveraging 'range' with timingsi
package main

import (
    "fmt"
    "os"
    "time"
)

func PrintArgs() {
    s, sep := "", "" 
    for _, arg := range os.Args[1:] { 
        s += sep + arg 
        sep = " "
    }
    fmt.Println(s)
}

func main() {
    start := time.Now()
    PrintArgs();
    fmt.Printf("%.8fs elapsed\n", time.Since(start).Seconds())
}
