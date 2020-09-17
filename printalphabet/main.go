package main

import "github.com/01-edu/z01"

func main() {

	i := 0
	var aRune string = "abcdefghijklmnopqrstuvwxyz"

	for i < 26 {

		z01.PrintRune(rune(aRune[i]))
		i = i + 1 
		\n

	}

}
