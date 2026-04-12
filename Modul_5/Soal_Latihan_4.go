package main

import "fmt"

func barisan(n int) {
	fmt.Print(n, " ")
	if n > 1 {
		barisan(n - 1)
		fmt.Print(n, " ")
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	barisan(N)
}