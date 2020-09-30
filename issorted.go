package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	first := false
	second := false
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			first = true
		} else if f(a[i], a[i+1]) < 0 {
			second = true
		}
	}
	if first && second {
		return false
	} else {
		return true
	}
}
