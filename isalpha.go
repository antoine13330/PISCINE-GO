package piscine

func IsAlpha(s string) bool {
	lettre := []rune(s)
	x := 0
	for x <= len(s) && ((lettre[x] >= 'a' && lettre[x] <= 'z') ||
		(lettre[x] >= 'A' && lettre[x] <= 'Z') ||
		(lettre[x] >= '0' && lettre[x] <= '9')) {
		x++
		if x == len(s) {
			return true
		}
	}
	return false
}
