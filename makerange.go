package piscine

func MakeRange(min, max int) []int {
	if max > min {
		len := max - min
		result := []int{}
		for i := 0; i < len; i++ {
			result = append(result, min+i)
		}
		return result
	} else {
		return nil
	}
}
