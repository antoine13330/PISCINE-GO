package piscine

func Any(f func(string) bool, a []string) bool {
	for i, v := range a {
		a[i] = f(a)
	}
	return f(v)

}
