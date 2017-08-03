// Exercise 4.1
// Function that counts the number of bits that are different in two SHA256 hashes

package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte // A lookup table for the population count of a byte

// Initialize the population count table

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i & 1)
	}
}

// Count the bits in a slice of bytes

func countbits(sb [sha256.Size]uint8) int {
	var bits int = 0
	for i := range sb {
		bits += int(pc[sb[i]])
	}
	return bits
}

// Determines count of different bits between two SHA256 hashes

func SHA256bitdiff(sha1, sha2 [sha256.Size]uint8) int {
	b1 := countbits(sha1)
	b2 := countbits(sha2)
	if b1 > b2 {
		return b1 - b2
	}
	return b2 - b1
}

// main

func main() {
	c1 := sha256.Sum256([]byte("X"))
	c2 := sha256.Sum256([]byte("x"))
	fmt.Printf("%x\n%x\n%d bits are different\n", c1, c2, SHA256bitdiff(c1, c2))
}