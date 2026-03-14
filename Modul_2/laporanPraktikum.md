# <h1 align="center">Laporan Praktikum Modul 2 - REVIEW ALGORITMA DAN PEMROGRAMAN 1 </h1>
<p align="center">Fikri Luqman Muktabar - 109082500103</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main
import "fmt"
func main() {
 var (
 satu, dua, tiga string
 temp string
 )
 fmt.Print("Masukan input string: ")
 fmt.Scanln(&satu)
 fmt.Print("Masukan input string: ")
 fmt.Scanln(&dua)
 fmt.Print("Masukan input string: ")
 fmt.Scanln(&tiga)
 fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
 temp = satu
 satu = dua
 dua = tiga
 tiga = temp
 fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]

### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

func main() {
	var gelas1, gelas2, gelas3, gelas4 string
	berhasil := true

	for percobaan := 1; percobaan <= 5; percobaan++ {

		fmt.Printf("Percobaan %d : ", percobaan)
		fmt.Scan(&gelas1, &gelas2, &gelas3, &gelas4)

		if gelas1 != "merah" || gelas2 != "kuning" || gelas3 != "hijau" || gelas4 != "ungu" {
			berhasil = false
		}
	}

	fmt.Println("BERHASIL:", berhasil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]

### 3. [Soal]
#### soal1.go

```go
package main

import "fmt"

func main() {
    var berat, totalBerat, sisaBerat, biayaPerKG, biayaSisa int

    fmt.Print("Berat parcel (gram): ")
    fmt.Scan(&berat)

    totalBerat = berat / 1000
    sisaBerat = berat % 1000

    biayaPerKG = totalBerat * 10000

    if sisaBerat > 0 {
        if berat > 10000 { 
            biayaSisa = 0
        } else if sisaBerat <= 500 {
            biayaSisa = sisaBerat * 5
        } else {
            biayaSisa = sisaBerat * 15
        }
    }

    totalBiaya := biayaPerKG + biayaSisa

    fmt.Printf("Detail berat: %d kg + %d gr\n", totalBerat, sisaBerat)
    fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaPerKG, biayaSisa)
    fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]