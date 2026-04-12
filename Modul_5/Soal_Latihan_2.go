package main

import "fmt"

func bintang(k int) {
	if k == 0 {
		return
	}
	fmt.Print("*")
	bintang(k - 1)
}

func pola(i, n int) {
	if i > n {
		return
	}
	bintang(i)
	fmt.Println()
	pola(i+1, n)
}

func main() {
	var N int
	fmt.Scan(&N)
	pola(1, N)
}