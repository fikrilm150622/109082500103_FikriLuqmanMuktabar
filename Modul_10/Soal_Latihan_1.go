package main

import "fmt"

type arrFloat [1000]float64

func main() {
	var data arrFloat
	var n int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}
	min := data[0]
	max := data[0]
	i := 1
	for i < n {
		if data[i] < min {
			min = data[i]
		}
		if data[i] > max {
			max = data[i]
		}
		i = i + 1
	}
	fmt.Printf("Berat terkecil: %.1f\n", min)
	fmt.Printf("Berat terbesar: %.1f\n", max)
}