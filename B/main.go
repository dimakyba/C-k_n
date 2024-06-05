package main

import "fmt"

func main() {
	var n, k uint64
	fmt.Scan(&n)
	fmt.Scan(&k)
	fmt.Println(binomialCoef(uint64(n), uint64(k)) % 9929)
}

func binomialCoef(n, k uint64) uint64 {
	if k > n {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}
	if k > n-k {
		k = n - k
	}

	var nom, denom uint64 = 1, 1
	for i := uint64(0); i < k; i++ {
		nom *= (n - i)
		denom *= (i + 1)
	}
	return nom / denom
}
