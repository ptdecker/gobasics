// Popcount implements a function which returns the 'population' (i.e. number of bits) which are
// set in a uint64 value

package popcount

var pc [256]byte // A lookup table for the population count of a byte

// Initialize the population count table
func init() {
    for i := range pc {
        pc[i] = pc[i/2] + byte(i & 1)
    }
}

// Calculate the population count for each byte in the 64-bit value
func PopCount(x uint64) int {
    return int(
        pc[byte(x >> (0 * 8))] +
        pc[byte(x >> (1 * 8))] +
        pc[byte(x >> (2 * 8))] +
        pc[byte(x >> (3 * 8))] +
        pc[byte(x >> (4 * 8))] +
        pc[byte(x >> (5 * 8))] +
        pc[byte(x >> (6 * 8))] +
        pc[byte(x >> (7 * 8))])
}
