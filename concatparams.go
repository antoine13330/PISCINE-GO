package piscine

func ConcatParams(args []string) string {

	str := ""
	a := 0

	for _, b := range args {
		if a != 0 {
			str += "\n"
		}
		str = str + b
		a = 1

	}
	return str
}
