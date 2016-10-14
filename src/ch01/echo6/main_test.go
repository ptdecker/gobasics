// main_test is a benchmark test for echo6
package main

import (
    "bytes"
    "testing"
)

func BenchmarkPrintArgs(b *testing.B) {
    out = new(bytes.Buffer) // capture output
    for i := 0; i < b.N; i++ {
        PrintArgs([]string{"one", "two", "three"})
    }
}
