package piscine

func CountIf(f func(string) bool, tab []string) int {
	y := 0
	for _, v := range tab {
		if f(v) == true {
			y++
		}
	}
	return y
}
