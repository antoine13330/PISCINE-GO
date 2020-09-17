package main

import "github.com/01-edu/z01"

func main() {

	i := 26
	var aRune string = "\nabcdefghijklmnopqrstuvwxyz"

	for i >= 0 {

		z01.PrintRune(rune(aRune[i]))
		i = i - 1

	}

}
