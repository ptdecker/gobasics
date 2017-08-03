// A non-recursive version of comma that leverages bytes.Buffer instead of string concatination

package main

import (
	"fmt"
	"bytes"
	"unicode/utf8"
	"unicode"
	"strings"
)

//c.f. https://stackoverflow.com/questions/29038314/determining-whitespace-in-go
const white_space = "\t\n\v\f\r \u0085\u00A0\u1680" // basic whitespace - could be expanded

// reverse()
// c.f. https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go#10030772

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// comma()

func comma(s string) string {

	var buf bytes.Buffer
	var decpart string
	i := 1

	// Remove leading white space

	s = strings.Trim(s, white_space)

	// Carve off first number using white space as a delimiter

	lastrune := strings.IndexFunc(s, unicode.IsSpace)
	if lastrune > 0 {
		s = s[:lastrune]
	}

	// Carve off decimal portion if it exists

	decpos := strings.LastIndexAny(s, ".")
	if decpos > 0 {
		decpart = s[decpos:]
		s = s[:len(s)-decpos]
	}

	// Add commas by walking through digits in reverse

	for len(s) > 0 {
		r, size := utf8.DecodeLastRuneInString(s)
		s = s[:len(s)-size]
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
			if len(s) > 0 && i%3 == 0 {
				buf.WriteString(",")
			}
			i++
		} else if r == '+' || r == '-' {
			buf.WriteRune(r)
		}
	}

	// Reverse buffer and return

	return reverse(buf.String()) + decpart
}

// main()

func main() {
	fmt.Println(comma("123456789"))
	fmt.Println(comma(""))
	fmt.Println(comma("ABCD"))
	fmt.Println(comma("ABC1234DEF"))
	fmt.Println(comma("1234.56"))
	fmt.Println(comma("1234.5678"))
	fmt.Println(comma("1234 5678"))
	fmt.Println(comma("1234.5678 9012"))
	fmt.Println(comma("-1234"))
	fmt.Println(comma("-1234.3456"))
	fmt.Println(comma("+1234"))
	fmt.Println(comma("  1234 567 8901"))
}
