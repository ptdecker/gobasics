// rotate1
// Rotates an array left one position moving the first array entry to the last

package main

import (
    "fmt"
)

func rotateleft(s []int) {
    if len(s) < 2 { // if either a single or no elements, then no work to do
        return
    }
    temp := s[0]
    copy(s[0:], s[1:])
    s[len(s)-1] = temp
}

// main

func main() {

    // Test case:  Empty array
    a := [...]int{}
    rotateleft(a[:])
    fmt.Println(a)

    // Test case:  One element array
    b := [...]int{1}
    rotateleft(b[:])
    fmt.Println(b)

    // Test case:  Multiple array elements
    c := [...]int{0, 1, 2, 3, 4, 5}
    rotateleft(c[:])
    fmt.Println(c)
}
