// A non-recursive version of comma that leverages bytes.Buffer instead of string concatination

package main

import (
	"fmt"
	"bytes"
	"unicode/utf8"
	"unicode"
)

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
	i := 1
	for len(s) > 0 {
		r, size := utf8.DecodeLastRuneInString(s)
		s = s[:len(s)-size]
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
			if len(s) > 0 && i%3 == 0 {
				buf.WriteString(",")
			}
			i++
		}
	}
	return reverse(buf.String())
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
