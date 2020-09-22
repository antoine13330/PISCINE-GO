package piscine

func RecursivePower(nb int, power int) int {
	if power >= 0 {
		arg := 1
		if i := 1; i < power+1 {
			return nb * RecursivePower(nb, power-1)
		}
		return arg
	} else {
		return 0
	}
}
