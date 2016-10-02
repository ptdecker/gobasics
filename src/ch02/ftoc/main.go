// Ftoc prints two Fahrenheit-to-Celsius conversions
package main

import "fmt"

func main() {
    const freezingF, boilingF = 32.0, 212.0
    fmt.Printf("%3g\u00b0F = %3g\u00b0C\n", freezingF, fToC(freezingF))
    fmt.Printf("%3g\u00b0F = %3g\u00b0C\n", boilingF,  fToC(boilingF))
}

func fToC(f float64) float64 {
    return (f - 32) * 5 / 9
}

