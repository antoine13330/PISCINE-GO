package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args[0:]
	for a := 1; a <= len(argument); a++ {
		for b := a + 1; b < len(argument); b++ {
			if argument[a] > argument[b] {
				argument[a], argument[b] = argument[b], argument[a]
			}
		}
	}
	for i := 1; i < len(argument); i++ {
		for _, letter := range argument[i] {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
