// Demonstration of popcount3
package main

import (
    "fmt"

    "ch02/popcount3"
)

func main() {
    fmt.Printf("PopCount(0x1234567890ABCDEF) = %d\n", popcount3.PopCount(0x1234567890ABCDEF))
    fmt.Println("\nTo compare performance with 'popcount', 'popcount1', and 'popcount2' run:\n\n\tgo test -bench=. ch02/popcount3\n");
}
