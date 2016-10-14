// Popcount2 implements a function which returns the 'population' (i.e. number of bits) which are
// set in a uint64 value. Implementation counts the bits by shifting the argument through all 64
// bit positions testing the rightmost bit each time.  Compare performance to popcount and popcount1
// which leverage a table lookup--popcount with a single expression and popcount1 with a loop
//
// NOTE: This implementation uses a different approach than illustrated in Donovan's & Kernignham's
//       PopCountByShifting() example countained within gopl.io/ch2/popcount/popcount_test.go.  Rather
//       than shifting the bit map each iteration by a growing number of shifts, this approach shifts
//       the mask one place each iteration.  Doing so reduces the benchmark time by 36% (from 85.0 ns/op
//       to 60.6 ns/op in my CPU) although it still does not compete with the lookup table approach nor
//       the bit clearing approach demonstrated in popcount3

package popcount2

// Calculate the population count for each byte in the 64-bit value
func PopCount(x uint64) int {
    var bitcount int
    var bitmask  uint64 = 1
    for i := 0; i < 64; i++ {
        if x & bitmask != 0 {
            bitcount++
        }
        bitmask = (bitmask << 1)
    }
    return bitcount
}

// Processor: 3.1 GHz Intel Core i7
// Memory: 16 GB 1867 MHz DDR3
// OS: OS X El Capitan v10.11.6
//
// BenchmarkPopCount-4  3,000,000,000  5.13 ns/op (table lookup, single expression)
// BenchmarkPopCount1-4 1,000,000,000 13.4  ns/op (table lookup, loop)
// BenchmarkPopCount2-4   200,000,000 60.6  ns/op (bit mask)
