// Demonstration of popcount1
package main

import (
    "fmt"

    "ch02/popcount1"
)

func main() {
    fmt.Printf("PopCount(0x1234567890ABCDEF) = %d\n", popcount1.PopCount(0x1234567890ABCDEF))
    fmt.Println("\nTo compare performance with 'popcount' run:\n\n\tgo test -bench=. ch02/popcount1\n")
}
