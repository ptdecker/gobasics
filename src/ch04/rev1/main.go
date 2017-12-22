// rev1
// Reverses an array of ints in place by passing a pointer to the array to a function

package main

import (
    "fmt"
)

const SIZE = 6

func reverse(s *[SIZE]int) {
    for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
        s[i], s[j] = s[j], s[i]
    }
}

// main

func main() {

    a := [SIZE]int{0, 1, 2, 3, 4, 5}

    reverse(&a)

    fmt.Println(a)
}
