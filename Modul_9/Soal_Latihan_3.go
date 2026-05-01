package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var skorPertandingan [100][2]int
	var hasil []string
	var pertandingan int = 0
	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)
	for {
		fmt.Printf("Pertandingan %d : ", pertandingan+1)
		fmt.Scan(&skorA, &skorB)
		if skorA < 0 || skorB < 0 {
			break
		}
		skorPertandingan[pertandingan][0] = skorA
		skorPertandingan[pertandingan][1] = skorB
		pertandingan++
	}
	for i := 0; i < pertandingan; i++ {
		if skorPertandingan[i][0] > skorPertandingan[i][1] {
			hasil = append(hasil, klubA)
			fmt.Printf("Hasil %d : %s\n", i+1, klubA)
		} else if skorPertandingan[i][1] > skorPertandingan[i][0] {
			hasil = append(hasil, klubB)
			fmt.Printf("Hasil %d : %s\n", i+1, klubB)
		} else {
			fmt.Printf("Hasil %d : Draw\n", i+1)
		}
	}
	fmt.Println("Pertandingan selesai")
}