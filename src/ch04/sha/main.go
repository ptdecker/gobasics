// Exercise 4.2
// Calculates the SHA hash of its standard input (SHA256 by default, 384 anrd 512 as options)

// TODO: http://learngowith.me/a-better-way-of-handling-stdin/
// c.f. https://www.reddit.com/r/golang/comments/4su0js/a_better_way_of_handling_stdin/?st=j5vrysv3&sh=65d80256

// For text data, see files in 'data' folder referring to README.md as a guide.
// c.f. https://www.di-mgt.com.au/sha_testvectors.html

package main

import (
	"bufio"
    "bytes"
    "crypto/sha256"
    "crypto/sha512"
	"fmt"
    "io"
    "os"
    "strconv"
)

// main

func main() {

	var err error
	var r rune
    var bits int64 = 256 // Default to 256 bits

    buff := bytes.NewBuffer(nil)
    reader := bufio.NewReader(os.Stdin)

    // Read arguments

    for i := 1; i < len(os.Args); i++ {
        if os.Args[i][0] == '-' {
            bits, err = strconv.ParseInt(os.Args[i][1:], 0, 64)
            if err != nil {
                fmt.Printf("Error: Could not convert '%s' to a value: %s\n", os.Args[i][1:], err)
                os.Exit(1);
            }
        }
    }

    // Read buffer

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

    // Calculate SHA hash value

    switch bits {
        case 256:
            fmt.Printf("%x\n", sha256.Sum256(buff.Bytes()))
        case 384:
            fmt.Printf("%x\n", sha512.Sum384(buff.Bytes()))
        case 512:
            fmt.Printf("%x\n", sha512.Sum512(buff.Bytes()))
        default:
            fmt.Printf("Error: '%d' bits is an invalid choice--256, 384, and 512 are supported\n", bits)
            os.Exit(1);
    }

}
