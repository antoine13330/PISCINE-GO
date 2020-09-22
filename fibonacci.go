package piscine

func Fibonacci(index int) int {
	fibonacci := 0

	if index < 0 {
		fibonacci = -1
	}
	if index == 0 {
		fibonacci = 0
	}
	if index == 1 {
		fibonacci = 1
	}
	if index > 1 {
		fibonacci = Fibonacci(index-1) + Fibonacci(index-2)
	}
	return fibonacci
}
