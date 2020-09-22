package piscine

func FindNextPrime(nb int) int {
	i := 2
	if nb <= 1 {
		return NextPrime(nb + 1)
	}
	for i <= nb {
		if nb%i == 0 && i != nb {
			return NextPrime(nb + 1)
		}

		i = i + 1
	}
	return nb
}

func NextPrime(nb int) int {
	i := 2
	if nb <= 1 {
		nb = 2
	}
	for i <= nb {
		if nb%i == 0 && i != nb {
			nb = nb + 1
			i = 2
		}

		i = i + 1
	}
	return nb
}
