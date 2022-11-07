package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)

	// this doesn't work as you might expect
	banner("G😂", 6)

	// emojiString are not characters, they are code points
	// length is the number of code points in the string
	// code points are not characters, they are numbers
	// and so the length of the string is 5, not 2, in this case
	// because the G is a single code point, but the emojiString is four
	emojiString := "G😂"
	fmt.Println("Length of", emojiString, "is", len(emojiString))

	// bytes are the raw data that make up a string

	// in go, there are bytes and runes
	// bytes are raw data... uint8
	bytes := emojiString[0]
	fmt.Printf("%c of type %T\n", bytes, bytes)

	for i, r := range emojiString {
		fmt.Println(i, r)
		if i == 0 {
			// runes are code points... int32
			fmt.Printf("%c if of type %T", r, r)
		}
	}

	x, y := 1, "1"
	// prints values
	fmt.Printf("x=%v, y=%v\n", x, y)
	// prints with type, good for debugging
	fmt.Printf("x=%#v, y=%#v\n", x, y)

	fmt.Printf("%20s!", emojiString)

}

func IsPalindrome(s string) bool {
	// for i := 0; i < len(s)/2; i++ {
	// 	if s[i] != s[len(s)-1-i] {
	// 		return false
	// 	}
	// }
	reversed := strings.(s)
	return true
}

func banner(text string, width int) {
	// := means we are declaring a variable and assigning it a value at the same time
	// don't need to specify type, it will be inferred

	// BUG: length is the number of bytes, not the number of characters
	// padding := (width - len(text)) / 2
	padding := (width - utf8.RuneCountInString(text)) / 2

	// the "for" loop is the only loop in Go
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)

	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
