// Cfk converts its numeric argument to Celsius, Fahrenheit, and Kelvin
package main

import (
    "fmt"
    "os"
    "strconv"
    "ch02/tempconv1"
)

func main() {

    for _, arg := range os.Args[1:] {

        t, err := strconv.ParseFloat(arg, 64)
 
        if err != nil {
            fmt.Fprintf(os.Stderr, "cf: %v\n", err)
        }

        f := tempconv1.Fahrenheit(t)
        c := tempconv1.Celsius(t)
        k := tempconv1.Kelvin(t)

        fmt.Printf("%s = %s, %s = %s, ", c, tempconv1.CToF(c), c, tempconv1.CToK(c))
        fmt.Printf("%s = %s, %s = %s, ", f, tempconv1.FToC(f), f, tempconv1.FToK(f))
        fmt.Printf("%s = %s, %s = %s\n", k, tempconv1.KToC(k), k, tempconv1.KToF(k))

    }
}
