package piscine

func IterativeFactorial(nb int) int {
	B := nb - 1
	for nb > 0 && B > 0 {
		nb = nb * B
		B = B - 1

	}
	return nb
}
