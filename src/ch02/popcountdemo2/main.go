// Demonstration of popcount2
package main

import (
    "fmt"

    "ch02/popcount2"
)

func main() {
    fmt.Printf("PopCount(0x1234567890ABCDEF) = %d\n", popcount2.PopCount(0x1234567890ABCDEF))
    fmt.Println("\nTo compare performance with 'popcount' and 'popcount1' run:\n\n\tgo test -bench=. ch02/popcount2\n")
}
