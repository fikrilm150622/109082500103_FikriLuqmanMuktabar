package main

import "fmt"

type arrIkan [1000]float64

func main() {
var ikan arrIkan
var x, y int
var rata float64
fmt.Print("Masukkan jumlah ikan: ")
fmt.Scan(&x)
fmt.Print("Masukkan jumlah ikan per wadah: ")
fmt.Scan(&y)
for i := 0; i < x; i++ {
	fmt.Printf("Masukkan berat ikan ke-%d: ", i+1)
	fmt.Scan(&ikan[i])
}
i := 0
fmt.Print("Total berat setiap wadah: ")
for i < x {
	var total float64 = 0
	for j := 0; j < y && i < x; j++ {
		total = total + ikan[i]
		i++
	}
	fmt.Printf("%.1f ", total)
	rata = rata + total
}
fmt.Println()
rata = rata / float64((x+y-1)/y)
fmt.Printf("Rata-rata berat setiap wadah: %.1f\n", rata)
}