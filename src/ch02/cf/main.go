// Cf converts its numeric argument to Celsius and Fahrenheit
package main

import (
    "fmt"
    "os"
    "strconv"
    "ch02/tempconv"
)

func main() {

    for _, arg := range os.Args[1:] {

        t, err := strconv.ParseFloat(arg, 64)
 
        if err != nil {
            fmt.Fprintf(os.Stderr, "cf: %v\n", err)
        }

        f := tempconv.Fahrenheit(t)
        c := tempconv.Celsius(t)

        fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))

    }
}
