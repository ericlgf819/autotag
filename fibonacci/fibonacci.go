package fibonacci

func Fibonacci(i int) int {
	if i == 0 || i == 1 {
		return i
	}
	return i + Fibonacci(i-1)
}
