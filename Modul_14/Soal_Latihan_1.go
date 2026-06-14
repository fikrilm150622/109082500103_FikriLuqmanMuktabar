package main

import "fmt"

const NMAX = 1000

type arrInt [NMAX]int

func selectionSort(T *arrInt, n int) {
	var idx_min, temp int
	i := 1
	for i <= n-1 {
		idx_min = i - 1
		j := i
		for j < n {
			if T[idx_min] > T[j] {
				idx_min = j
			}
			j++
		}
		temp = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = temp
		i++
	}
}

func main() {
	var n, m int
	var rumah arrInt
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}
		selectionSort(&rumah, m)
		for j := 0; j < m; j++ {
			fmt.Print(rumah[j])
			if j < m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}