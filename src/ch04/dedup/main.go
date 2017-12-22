// dedup
// An in-place function that eliminates adjacent duplicates strings in a slice of strings

package main

import (
    "fmt"
)

// eliminate adjacent duplicate strings in a slice of strings

// NOTE: index 'i' is not incremented in the for loop as might be expected.
//       Instead, it is only incremented when no adjacent duplicate string was found
//       during a loop iteration.  'Continue' is used after a dupicate string
//       was found and removed from the slice (thus the slice length is one less)
//       which causes the loop comparision to trigger and 'i' is _not_ incremented.
//       If incrementing 'i' was included in the for loop, then continue would cause
//       it to be incremented and adjacent duplicates would not be properly eliminated.

func dedup(strings []string) []string {
    for i := 0; i < (len(strings) - 1); {
        if (strings[i] == strings[i+1]) {
            copy(strings[i:], strings[i+1:])
            strings = strings[:len(strings)-1]
            continue
        }
        i++
    }
    return strings
}

// main

func main() {

    // Test Case: Empty slice of strings
    a := []string{}
    fmt.Println("A:", dedup(a))

    // Test Case: No duplicates
    b := []string{"aaa", "bbb", "ccc", "ddd"}
    fmt.Println("B:", dedup(b))

    // Test Case: First strings are duplicate
    c := []string{"aaa", "aaa", "bbb", "ccc", "ddd"}
    fmt.Println("C:", dedup(c))

    // Test Case: Last strings are duplicates
    d := []string{"aaa", "bbb", "ccc", "ddd", "ddd", "ddd"}
    fmt.Println("D:", dedup(d))

    // Test Case: Duplicate strings throughout
    e := []string{"aaa", "aaa", "aaa", "bbb", "ccc", "ccc", "ddd", "eee", "eee", "eee"}
    fmt.Println("E:", dedup(e))

    // Test Case: All duplicates
    f := []string{"aaa", "aaa", "aaa", "aaa"}
    fmt.Println("F:", dedup(f))

}
