package main

import "fmt"

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var fn, fnr int
	factorial(n, &fn)
	factorial(n-r, &fnr)
	*hasil = fn / fnr
}

func combination(n, r int, hasil *int) {
	var fn, fr, fnr int
	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}

func main() {
	var a, b, c, d int
	var hasilP1, hasilC1, hasilP2, hasilC2 int

	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &hasilP1)
	combination(a, c, &hasilC1)

	permutation(b, d, &hasilP2)
	combination(b, d, &hasilC2)

	fmt.Println(hasilP1, hasilC1)
	fmt.Println(hasilP2, hasilC2)
}