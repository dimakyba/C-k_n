package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n, k uint64
		fmt.Scan(&n)
		fmt.Scan(&k)
		fmt.Println(binomialCoeff(n, k))
	}
}

func binomialCoeff(n, k uint64) uint64 {
	C := make([][]uint64, n+1)
	for i := range C {
		C[i] = make([]uint64, k+1)
	}

	var i, j uint64

	// Calculate value of Binomial Coefficient in bottom-up manner
	for i = 0; i <= n; i++ {
		for j = 0; j <= min(i, k); j++ {
			// Base Cases
			if j == 0 || j == i {
				C[i][j] = 1
			} else {
				// Calculate value using previously stored values
				C[i][j] = C[i-1][j-1] + C[i-1][j]
			}
		}
	}

	return C[n][k]
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}
