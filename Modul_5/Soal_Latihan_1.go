package main

import "fmt"

func DeretFibonacci(a, b, n int) {
	if n < 0 {
		return
	}
	fmt.Print(a, " ")
	DeretFibonacci(b, a+b, n-1)
}

func main() {
	var n int
	fmt.Scan(&n)
	DeretFibonacci(0, 1, n)
}