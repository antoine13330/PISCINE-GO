package piscine

func IsPrime(nb int) bool {
	i := 2
	if nb <= 1 {
		return false
	}
	for i <= nb {
		if nb%i == 0 && i != nb {
			return false
		}

		i = i + 1
	}
	return true
}
