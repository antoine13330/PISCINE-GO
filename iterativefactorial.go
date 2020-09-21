package piscine

func IterativeFactorial(nb int) int {

	if nb <= 20 && nb > 0 {
		B := nb - 1
		for B > 0 {
			nb = nb * B
			B = B - 1
		}
	}
	if nb == 0 {
		nb = nb + 1
	}
	if nb < 0 {
		nb = 0
	}
	return nb
}
