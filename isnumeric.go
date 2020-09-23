package piscine

func IsNumeric(s string) bool {
	lettre := []rune(s)
	x := 0
	for x <= len(s) && (lettre[x] >= '0' && lettre[x] <= '9') {
		x++
		if x == len(s) {
			return true
		}
	}
	return false
}
