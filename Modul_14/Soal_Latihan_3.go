package main

import "fmt"

const NMAX = 1000

type arrInt [NMAX]int

func insertionSort(T *arrInt, n int) {
	var temp int
	i := 1
	for i <= n-1 {
		j := i
		temp = T[j]

		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j--
		}
		T[j] = temp
		i++
	}
}

func main() {
	var T arrInt
	var n, x, jarak int
	var tetap bool
	n = 0
	fmt.Scan(&x)
	for x >= 0 {
		T[n] = x
		n++
		fmt.Scan(&x)
	}
	insertionSort(&T, n)
	for i := 0; i < n; i++ {
		fmt.Print(T[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	if n <= 1 {
		fmt.Println("Data berjarak 0")
	} else {
		jarak = T[1] - T[0]
		tetap = true
		i := 1
		for i < n-1 && tetap {
			if T[i+1]-T[i] != jarak {
				tetap = false
			}
			i++
		}
		if tetap {
			fmt.Println("Data berjarak", jarak)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}