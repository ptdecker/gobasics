// Demonstration of popcount
package main

import (
    "fmt"

    "ch02/popcount"
)

func main() {
    fmt.Printf("PopCount(0x1234567890ABCDEF) = %d\n", popcount.PopCount(0x1234567890ABCDEF))
}
