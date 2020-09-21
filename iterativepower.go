package piscine

func IterativePower(nb int, power int) int {
	result := nb
	if power > 1 {
		for o := nb - 2; o <= power; o++ {
			result = result * nb

		}
	}
	if power == 0 {
		result = 1

	}
	if power < 0 {
		result = 0
	}
	return result
}
