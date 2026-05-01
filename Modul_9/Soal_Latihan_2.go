package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)
	data := make([]int, n)
	fmt.Println("Masukkan nilai ke dalam array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Data ke-%d: ", i)
		fmt.Scan(&data[i])
	}

	fmt.Println("\na. Isi array:")
	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	fmt.Println("\nb. Indeks ganjil:")
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Print(data[i], " ")
		}
	}
	fmt.Println()

	fmt.Println("\nc. Indeks genap:")
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(data[i], " ")
		}
	}
	fmt.Println()

	var kelipatan int
	fmt.Print("\nd. Masukkan nilai kelipatan indeks: ")
	fmt.Scan(&kelipatan)
	fmt.Println("Elemen dengan indeks kelipatan", kelipatan, ":")
	if kelipatan == 0 {
		fmt.Println("Tidak valid (tidak boleh 0)")
	} else {
		for i := 0; i < n; i++ {
			if i%kelipatan == 0 {
				fmt.Print(data[i], " ")
			}
		}
		fmt.Println()
	}

	var hapus int
	fmt.Print("\ne. Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&hapus)
	if hapus >= 0 && hapus < n {
		for i := hapus; i < n-1; i++ {
			data[i] = data[i+1]
		}
		n--
		fmt.Println("Array setelah dihapus:")
		for i := 0; i < n; i++ {
			fmt.Print(data[i], " ")
		}
		fmt.Println()
	} else {
		fmt.Println("Indeks tidak valid")
	}

	var total int
	for i := 0; i < n; i++ {
		total += data[i]
	}
	rata := float64(total) / float64(n)
	fmt.Printf("\nf. Rata-rata: %.2f\n", rata)

	var jumlahKuadrat float64
	for i := 0; i < n; i++ {
		jumlahKuadrat += math.Pow(float64(data[i])-rata, 2)
	}
	stdDev := math.Sqrt(jumlahKuadrat / float64(n))
	fmt.Printf("g. Standar deviasi: %.2f\n", stdDev)

	var cari int
	fmt.Print("\nh. Masukkan angka yang dicari: ")
	fmt.Scan(&cari)
	hitung := 0
	for i := 0; i < n; i++ {
		if data[i] == cari {
			hitung++
		}
	}
	fmt.Println("Frekuensi", cari, "=", hitung)
}