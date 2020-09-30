package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	var monTableau []int
	if n < 0 {
		z01.PrintRune('-')
		n *= -1
	}
	for {
		unit := n % 10
		n /= 10
		monTableau = append(monTableau, unit)
		if n == 0 {
			break
		}
	}
	for i := len(monTableau) - 1; i >= 0; i-- {
		z01.PrintRune(rune(monTableau[i] + '0'))
	}
}
