package piscine

func RecursiveFactorial(nb int) int {
	A := nb
	B := 1
	if nb >= 0 && nb <= 25 {
		if nb == 0 {
			B = 1
		} else {
			B = A * RecursiveFactorial(A-1)
		}
	} else {
		B = 0
	}
	return B
}
