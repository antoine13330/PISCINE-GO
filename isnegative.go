package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	var aRune string = "TF"
	if nb < 0 {
		z01.PrintRune(rune(aRune[0]))
	} else {
		z01.PrintRune(rune(aRune[1]))
	}

	z01.PrintRune('\n')

}
