package piscine

func NRune(s string, n int) rune {
	number := 0
	for range s {
		number++
	}

	runes := []rune(s)

	if n > number || n <= 0 {
		return 0
	} else {
		return runes[n-1]
	}
}
