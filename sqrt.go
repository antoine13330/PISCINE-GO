package piscine

func Sqrt(n int) int {
	for i := 0; i <= n; i++ {
		if i*i == n {
			return i
		}
	}
	return 0
}
