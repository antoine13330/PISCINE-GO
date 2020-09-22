package piscine

func Compare(a, b string) int {
	v := 0
	if a > b {
		v = 1
	}
	if a < b {
		v = -1
	}
	if a == b {
		v = 0
	}
	return v
}
