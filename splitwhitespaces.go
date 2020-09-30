package piscine

func SplitWhiteSpaces(s string) []string {

	var answer []string
	rep := ""
	for _, c := range s {
		if c != rune(32) {
			rep += string(rune(c))
		}
		if c == rune(32) {
			if rep != "" {
				answer = append(answer, rep)
				rep = ""
			}
		}
	}
	answer = append(answer, rep)
	return answer
}
