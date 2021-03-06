package popcount1_test

import (
    "testing"

    "ch02/popcount"
    "ch02/popcount1"
)

func BenchmarkPopCount(b *testing.B) {
    for i := 0; i < b.N; i++ {
        popcount.PopCount(0x1234567890ABCDEF)
    }
}

func BenchmarkPopCount1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        popcount1.PopCount(0x1234567890ABCDEF)
    }
}
