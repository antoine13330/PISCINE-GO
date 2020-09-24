package piscine

func AppendRange(min, max int) []int {
	var Slice []int
	for max < min {
		return nil
	}
	for i := min; i < max; i++ {
		Slice = append(Slice, i)
	}

	return Slice
}
