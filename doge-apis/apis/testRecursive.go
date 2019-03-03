package apis

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func TestRecursive() {
	fmt.Println(factorial(7))
}
