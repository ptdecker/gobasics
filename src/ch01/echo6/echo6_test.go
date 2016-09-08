// test_echo6 is a benchmark test for echo6
package main

import "testing"

func BenchmarkPrintArgs(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PrintArgs()
    }
}
