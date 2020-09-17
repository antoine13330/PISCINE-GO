package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	var nombre string = string("0123456789")

	for A := 0; A <= 9; A++ {
		for B := A + 1; B <= 9; B++ {
			for C := B + 1; C <= 9; C++ {
				z01.PrintRune(rune(nombre[A]))
				z01.PrintRune(rune(nombre[B]))
				z01.PrintRune(rune(nombre[C]))

				if A == 7 && B == 8 && C == 9 {
					z01.PrintRune(10)
				} else {
					z01.PrintRune(44)
					z01.PrintRune(32)
				}

			}
		}
	}
}
