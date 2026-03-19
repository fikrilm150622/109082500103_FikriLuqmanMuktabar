package main

import "fmt"

func factorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	if a < c || b < d {
		fmt.Println("Input tidak valid (harus a >= c dan b >= d)")
		return
	}

	fmt.Println(permutation(a, c), combination(a, c))
	fmt.Println(permutation(b, d), combination(b, d))
}