package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		return
	}
	a1 := arg[0]
	a2 := arg[1]

	for x := range a1 {
		for y := range a2 {
			if a1[x] == a2[y] {
				z01.PrintRune(rune(a1[x]))
			}
		}
	}
}
