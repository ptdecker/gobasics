// Exercise 4.2
// Calculates the SHA hash of its standard input (SHA256 by default, 384 and 512 as options)

// TODO: http://learngowith.me/a-better-way-of-handling-stdin/
// c.f. https://www.reddit.com/r/golang/comments/4su0js/a_better_way_of_handling_stdin/?st=j5vrysv3&sh=65d80256

package main

import (
	"io"
	"os"
	"bytes"
	"bufio"
	"fmt"
)

var buf []bytes.Buffer

// main

func main() {
	var err error
	var r rune
	reader := bufio.NewReader(os.Stdin)
	for err != io.EOF {
		r, _, err = reader.ReadRune()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(r)
	}
}
