package main

import "github.com/01-edu/z01"

func main() {
	for i := 0; i < len("x = 42, y = 21\n"); i++ {
		z01.PrintRune(rune("x = 42, y = 21\n"[i]))
	}
}
