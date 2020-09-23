package piscine

func IsUpper(s string) bool {

	runes := []rune(s)
	for i := 0; i < len(runes); i++ {

		if int(runes[i]) < 65 || int(runes[i]) > 90 {
			return false
		}

	}

	return true
}
