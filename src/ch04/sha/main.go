// Exercise 4.2
// Calculates the SHA hash of its standard input (SHA256 by default, 384 anrd 512 as options)

// TODO: http://learngowith.me/a-better-way-of-handling-stdin/
// c.f. https://www.reddit.com/r/golang/comments/4su0js/a_better_way_of_handling_stdin/?st=j5vrysv3&sh=65d80256

package main

import (
	"bufio"
    "bytes"
    "crypto/sha256"
	"fmt"
    "io"
    "os"
)

// main

func main() {
	var err error
	var r rune
    buff := bytes.NewBuffer(nil)
	reader := bufio.NewReader(os.Stdin)
	for { 
		r, _, err = reader.ReadRune()
		if err == io.EOF {
            break
        }
        if err != nil {
			fmt.Println(err)
            continue
		}
        _, err = buff.WriteRune(r)
        if err != nil {
            fmt.Println(err)
            continue
        }
	}
    sum := sha256.Sum256(buff.Bytes())
    fmt.Printf("%x\n", sum)
}
