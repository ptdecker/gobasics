// main_test is a benchmark test for surface1
package main

import (
    "bytes"
    "testing"
)

func BenchmarkGenSVG(b *testing.B) {
    out = new(bytes.Buffer) // capture output
    for i := 0; i < b.N; i++ {
        genSVG()
    }
}
