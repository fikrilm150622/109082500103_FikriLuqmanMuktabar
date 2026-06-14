package main

import "fmt"

const NMAX = 1000

type arrInt [NMAX]int

func selectionSortAsc(T *arrInt, n int) {
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
	var data, ganjil, genap arrInt
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		for j := 0; j < m; j++ {
			fmt.Scan(&data[j])
		}
		nGanjil := 0
		nGenap := 0
		for j := 0; j < m; j++ {
			if data[j]%2 == 1 {
				ganjil[nGanjil] = data[j]
				nGanjil++
			} else {
				genap[nGenap] = data[j]
				nGenap++
			}
		}
		selectionSortAsc(&ganjil, nGanjil)
		selectionSortAsc(&genap, nGenap)
		pertama := true
		for j := 0; j < nGanjil; j++ {
			if !pertama {
				fmt.Print(" ")
			}
			fmt.Print(ganjil[j])
			pertama = false
		}
		for j := nGenap - 1; j >= 0; j-- {
			if !pertama {
				fmt.Print(" ")
			}
			fmt.Print(genap[j])
			pertama = false
		}
		fmt.Println()
	}
}