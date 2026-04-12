package main

import "fmt"

func ganjil(n int) {
	if n <= 0 {
		return
	}
	ganjil(n - 2)
	fmt.Print(n, " ")
}

func main() {
	var N int
	fmt.Scan(&N)
	if N%2 == 0 {
		N-- 
	}
	ganjil(N)
}