// Popcount3 implements a function which returns the 'population' (i.e. number of bits) which are
// set in a uint64 value. Implementation clears the rightmost non-zero bit of x and uses its
// effect to count the bits.  Compare performance to popcount, popcount1, and popcount2 which
// each use different alternate approaches.

package popcount3

// Calculate the population count for each byte in the 64-bit value
func PopCount(x uint64) int {
    bitcount := 0
    for x != 0 {
        x = x & (x - 1) // Clears the rightmost non-zero bit
        bitcount++
    }
    return bitcount
}

// Processor: 3.1 GHz Intel Core i7
// Memory: 16 GB 1867 MHz DDR3
// OS: OS X El Capitan v10.11.6
//
// BenchmarkPopCount-4  3,000,000,000  5.22 ns/op (table lookup, single expression)
// BenchmarkPopCount1-4 1,000,000,000 13.7  ns/op (table lookup, loop)
// BenchmarkPopCount2-4   200,000,000 60.0  ns/op (bit mask)
// BenchmarkPopCount3-4 1,000,000,000 21.9  ns/op (bit clearing)
