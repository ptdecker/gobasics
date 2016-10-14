// Popcount1 implements a function which returns the 'population' (i.e. number of bits) which are
// set in a uint64 value. Implementation leverages a loop instead of a single expression.  Compare
// performance to popcount which usies inline lookups instead of a loop.

package popcount1

var pc [256]byte // A lookup table for the population count of a byte

// Initialize the population count table
func init() {
    for i := range pc {
        pc[i] = pc[i/2] + byte(i & 1)
    }
}

// Calculate the population count for each byte in the 64-bit value
func PopCount(x uint64) int {
    var bitcount byte
    for bytenum := uint(0); bytenum < 8; bytenum++ {
        bitcount += pc[byte(x >> (bytenum * 8))]
    }
    return int(bitcount)
}

// Processor: 3.1 GHz Intel Core i7
// Memory: 16 GB 1867 MHz DDR3
// OS: OS X El Capitan v10.11.6
//
// BenchmarkPopCount-4  3,000,000,000  5.15 ns/op (table lookup, single expression)
// BenchmarkPopCount1-4 1,000,000,000 13.8  ns/op (table lookup, loop)
