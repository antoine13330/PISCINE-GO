package main

import "github.com/01-edu/z01"

func main() {

	i := 0
	var aRune int = "0123456789\n"

	for i <= 11 {

		z01.PrintRune(rune(aRune[i]))
		i = i - 1

	}

}
