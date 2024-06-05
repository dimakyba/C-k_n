package main

import "fmt"

func main() {
	var k, n int
	fmt.Scan(&n)
	fmt.Scan(&k)
	fmt.Println(factorial(n) / (factorial(k) * factorial(n-k)))
}

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
