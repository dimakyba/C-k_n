package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)
	var sum int64

	for i := int64(0); i <= n; i++ {
		var res = binomialCoef(n, i)
		sum += res
	}
	fmt.Println(sum)
}

func binomialCoef(n, k int64) int64 {
	if k == 0 || k == n {
		return 1
	}
	if k > n-k {
		k = n - k
	}
	result := int64(1)
	for i := int64(0); i < k; i++ {
		result *= (n - i)
		result /= (i + 1)
	}
	return result
}
