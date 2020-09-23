package piscine

func IsLower(s string) bool {
	lettre := []rune(s)
	x := 0
	for x <= len(s) && (lettre[x] >= 'a' && lettre[x] <= 'z') {
		x++
		if x == len(s) {
			return true
		}
	}
	return false
}
