package main

import "fmt"

func main() {
	var suara [21]int
	var x int
	var totalMasuk, suaraSah int
	var ketua, wakil, suaraTerbanyak, suaraKedua int
	fmt.Scan(&x)
	for x != 0 {
		totalMasuk++
		if x >= 1 && x <= 20 {
			suara[x]++
			suaraSah++
		}
		fmt.Scan(&x)
	}
	for i := 1; i <= 20; i++ {
		if suara[i] > suaraTerbanyak {
			suaraKedua = suaraTerbanyak
			wakil = ketua
			suaraTerbanyak = suara[i]
			ketua = i
		} else if suara[i] > suaraKedua {
			suaraKedua = suara[i]
			wakil = i
		}
	}
	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", suaraSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}