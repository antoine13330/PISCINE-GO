package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	count := len(s)

	for i := 0; i < count; i++ {

		z01.PrintRune(rune(s[i]))
	}
}
