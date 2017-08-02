// Declares KB to YB in a compact manner

// c.f. https://en.wikipedia.org/wiki/Orders_of_magnitude_(data)

package main

import "fmt"

const (
	KB = 1000    // 1e3  = 1,000
	MB = KB * KB // 1e6  = 1,000,000
	GB = MB * KB // 1e9  = 1,000,000,000
	TB = GB * KB // 1e12 = 1,000,000,000,000
	PB = TB * KB // 1e15 = 1,000,000,000,000,000
	EB = PB * KB // 1e18 = 1,000,000,000,000,000,000
	ZB = EB * KB // 1e21 = 1,000,000,000,000,000,000,000
	YB = ZB * KB // 1e24 = 1,000,000,000,000,000,000,000,000
)

func main() {
	fmt.Printf("%e\n", float64(KB))
	fmt.Printf("%e\n", float64(MB))
	fmt.Printf("%e\n", float64(GB))
	fmt.Printf("%e\n", float64(TB))
	fmt.Printf("%e\n", float64(PB))
	fmt.Printf("%e\n", float64(EB))
	fmt.Printf("%e\n", float64(ZB))
	fmt.Printf("%e\n", float64(YB))
}

// NOTE: A slick shorter way (67 characters), but not as easy to the eyes, is
// provided in a StackOverflow article:
//
// const ( KB, MB, GB, TB, PB, EB, ZB, YB = 1e3, 1e6, 1e9, 1e12, 1e15, 1e18, 1e21, 1e24 )