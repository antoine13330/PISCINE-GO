package piscine

func Map(f func(int) bool, a []int) []bool {
	var Abool = make([]bool, len(a))
	for i, b := range a {
		f[i] = Abool[b]
		return Abool
	}
}
