package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := os.Args[0]
	slashindex := -1
	for i, r := range s {

		if r == '/' {
			slashindex = i
		}
	}
	for _, letter := range s[slashindex+1:] {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
