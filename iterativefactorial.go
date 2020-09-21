package piscine

func IterativeFactorial(nb int) int {
	A := nb
	B := 1
	if nb >= 0 && nb <= 25 {
		if nb == 0 {
			B = 1
		} else {
			for A >= 1 {
				B = B * A
				A = A - 1
			}
		}
	} else {
		B = 0
	}
	return B
}
