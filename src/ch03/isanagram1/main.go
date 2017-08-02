// Test to see if two strings are anagrams

// This approach uses a "checking off" algorithm which is not efficient [O(n^2)] but
// avoids needing language features not yet introduced in the text.
// c.f. http://interactivepython.org/runestone/static/pythonds/AlgorithmAnalysis/AnAnagramDetectionExample.html

package main

import (
	"fmt"
	"bytes"
)

// 0xFF is an invalid UTF-8 byte thus a good choice for a sentinel character
// c.f. https://softwareengineering.stackexchange.com/questions/190409/a-unicode-sentinel-value-i-can-use
const sentinel = '\xFF'

// isAnagram()

func isAnagram(s1, s2 string) bool {

	// Neither empty nor differing length strings are anagrams

	if len(s1) == 0 || len(s2) == 0 || len(s1) != len (s2) {
		return false
	}

	// Store the first of the strings in a byte slice so we can modify it

	bs := []byte(s1)

	// For each rune in the second string, search the byte slice for it. If not
	// found, then they are not anagrams.  If found, replace that instance of
	// the rune with a sentinel value, e.g. mark it as used.

	for _, r := range s2 {
		i := bytes.IndexRune(bs, r)
		if i == -1 {
			return false
		}
		bs[i] = sentinel
	}

	return true
}

// main()

func main() {
	fmt.Println(isAnagram("",""))
	fmt.Println(isAnagram("heart", "earth"))
	fmt.Println(isAnagram("this","that"))
	fmt.Println(isAnagram("HI\uD09C\uD09F\uE3869B\uE386A1\uF09F9880", "\uF09F9880\uE386A1\uD09FH\uE3869B\uD09CI"))
}
