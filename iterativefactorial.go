package piscine

func IterativeFactorial(nb int) int {
	B := nb - 1
	if nb <= 25 && nb > 0 {
		for B > 0 {
			nb = nb * B
			B = B - 1
		}
	}
	return nb
}
