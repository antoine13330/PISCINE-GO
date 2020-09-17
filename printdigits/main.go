package main

import "github.com/01-edu/z01"

func main() {

	i := 0
	var aRune string = "0123456789\n"

	for i <= 10 {

		z01.PrintRune(rune(aRune[i]))
		i = i + 1

	}

}
