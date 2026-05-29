# <h1 align="center">Laporan Praktikum Modul 10 - PENCARIAN NILAI MAX/MIN - ALGORITMA DAN PEMROGRAMAN 2 </h1>
<p align="center">Fikri Luqman Muktabar - 109082500103</p>

## Unguided 

### 1. [Soal]
#### Soal_Latihan_1.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_9/Output/SS_Soal_Latihan_1.png) belom diganti 

### [Penjelasan]
Program ini digunakan untuk mencari berat anak kelinci yang paling kecil dan paling besar dari data yang diinputkan. User akan memasukkan jumlah anak kelinci terlebih dahulu, kemudian memasukkan berat masing-masing anak kelinci satu per satu. Setelah seluruh data selesai diinput, program akan menampilkan berat terkecil dan berat terbesar.
package main berarti program ini merupakan program utama yang dapat langsung dijalankan.
import "fmt" digunakan untuk memanggil package fmt yang berfungsi untuk proses input dan output, seperti Print, Scan, dan Printf.
type arrFloat [1000]float64 digunakan untuk membuat tipe data array bernama arrFloat yang mampu menyimpan hingga 1000 data bertipe float64 atau bilangan desimal.
func main() {...} merupakan bagian utama program tempat seluruh proses dijalankan.
Pada awal program dibuat beberapa variabel:
- data digunakan untuk menyimpan seluruh berat anak kelinci
- n digunakan untuk menyimpan jumlah anak kelinci yang akan diinput
Setelah itu program meminta user memasukkan jumlah anak kelinci menggunakan fmt.Print dan fmt.Scan.
Selanjutnya program melakukan input berat anak kelinci menggunakan perulangan for. Perulangan dimulai dari i = 0 dan akan terus berjalan hingga jumlah data yang dimasukkan terpenuhi.
Setiap berat anak kelinci yang diinput akan disimpan ke dalam array data sesuai indeksnya masing-masing.
Setelah seluruh data selesai dimasukkan, program menentukan nilai awal minimum dan maksimum dengan:
min := data[0]
max := data[0]
Artinya, data pertama dijadikan sebagai nilai terkecil sementara dan nilai terbesar sementara.
Kemudian program mulai mencari nilai minimum dan maksimum menggunakan perulangan:
for i < n
Di dalam perulangan dilakukan pengecekan:
- Jika data[i] lebih kecil dari min, maka nilai min diperbarui dengan data tersebut.
- Jika data[i] lebih besar dari max, maka nilai max diperbarui dengan data tersebut.
Proses ini dilakukan terus hingga seluruh data selesai diperiksa satu per satu.
Pada bagian akhir, program akan menampilkan hasil pencarian berupa:
- berat anak kelinci paling kecil
- berat anak kelinci paling besar
fmt.Printf digunakan untuk menampilkan angka desimal dengan format satu angka di belakang koma menggunakan %.1f.

### 2. [Soal]
#### Soal_Latihan_2.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_9/Output/SS_Soal_Latihan_2.png) belom diganti

### [Penjelasan]
Program ini digunakan untuk menghitung total berat ikan di setiap wadah dan mencari rata-rata berat ikan per wadah. User akan memasukkan jumlah ikan terlebih dahulu, kemudian memasukkan jumlah ikan yang dapat dimasukkan ke dalam satu wadah. Setelah itu user memasukkan berat masing-masing ikan satu per satu. Setelah seluruh data selesai diinput, program akan menampilkan total berat ikan di setiap wadah dan rata-rata berat setiap wadah.
package main berarti program ini merupakan program utama yang dapat langsung dijalankan.
import "fmt" digunakan untuk memanggil package fmt yang berfungsi untuk proses input dan output, seperti Print, Scan, dan Printf.
type arrIkan [1000]float64 digunakan untuk membuat tipe data array bernama arrIkan yang mampu menyimpan hingga 1000 data bertipe float64 atau bilangan desimal.
func main() {...} merupakan bagian utama program tempat seluruh proses dijalankan.
Pada awal program dibuat beberapa variabel:
- ikan digunakan untuk menyimpan seluruh berat ikan
- x digunakan untuk menyimpan jumlah ikan yang akan diinput
- y digunakan untuk menyimpan jumlah ikan dalam setiap wadah
- rata digunakan untuk menyimpan hasil rata-rata berat setiap wadah
Setelah itu program meminta user memasukkan jumlah ikan dan jumlah ikan per wadah menggunakan fmt.Print dan fmt.Scan.
Selanjutnya program melakukan input berat ikan menggunakan perulangan for. Perulangan dimulai dari i = 0 dan akan terus berjalan hingga jumlah data ikan yang dimasukkan terpenuhi.
Setiap berat ikan yang diinput akan disimpan ke dalam array ikan sesuai indeksnya masing-masing.
Kemudian program mulai menghitung total berat ikan di setiap wadah menggunakan perulangan:
for i < x
Di dalam perulangan tersebut dibuat variabel:
total = 0
Variabel total digunakan untuk menyimpan jumlah berat ikan pada setiap wadah.
Lalu program menggunakan perulangan lagi:
for j := 0; j < y && i < x; j++
Perulangan ini digunakan untuk mengelompokkan ikan ke dalam wadah sesuai jumlah yang ditentukan oleh user.
Di dalam perulangan dilakukan proses:
- total = total + ikan[i] digunakan untuk menjumlahkan berat ikan pada wadah saat ini.
- i++ digunakan untuk berpindah ke data ikan berikutnya.
Setelah seluruh ikan dalam satu wadah selesai dijumlahkan, program menampilkan total berat wadah menggunakan:
fmt.Printf("%.1f ", total)
Kemudian nilai total tersebut ditambahkan ke variabel rata menggunakan:
rata = rata + total
Proses ini dilakukan terus hingga seluruh data ikan selesai diproses.
Setelah semua total berat wadah diperoleh, program menghitung rata-rata berat setiap wadah dengan:
rata = rata / float64((x+y-1)/y)
Bagian float64 digunakan agar hasil pembagian menjadi bilangan desimal.
Pada bagian akhir, program akan menampilkan:
- total berat ikan di setiap wadah
- rata-rata berat setiap wadah
fmt.Printf digunakan untuk menampilkan angka desimal dengan format satu angka di belakang koma menggunakan %.1f.


### 3. [Soal]
#### Soal_Latihan_3.go

```go
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total = total + arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var arrBerat arrBalita
	var n int
	var bMin, bMax, rata float64
	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBerat[i])
	}
	hitungMinMax(arrBerat, n, &bMin, &bMax)
	rata = rerata(arrBerat, n)
	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_9/Output/SS_Soal_Latihan_3.png) belom diganti 

### [Penjelasan]
Program ini digunakan untuk mencari berat balita paling kecil, paling besar, dan menghitung rerata berat balita dari data yang diinputkan. User akan memasukkan jumlah data berat balita terlebih dahulu, kemudian memasukkan berat masing-masing balita satu per satu. Setelah seluruh data selesai diinput, program akan menampilkan berat minimum, berat maksimum, dan rerata berat balita.
package main berarti program ini merupakan program utama yang dapat langsung dijalankan.
import "fmt" digunakan untuk memanggil package fmt yang berfungsi untuk proses input dan output, seperti Print, Scan, dan Printf.
type arrBalita [100]float64 digunakan untuk membuat tipe data array bernama arrBalita yang mampu menyimpan hingga 100 data bertipe float64 atau bilangan desimal.
Program ini memiliki dua subprogram, yaitu:
- func hitungMinMax()
- func rerata()
func hitungMinMax() digunakan untuk mencari nilai minimum dan maksimum dari data berat balita.
Di dalam function hitungMinMax(), nilai awal minimum dan maksimum ditentukan dengan:
*bMin = arrBerat[0]
*bMax = arrBerat[0]
Artinya, data pertama dijadikan sebagai nilai minimum sementara dan nilai maksimum sementara.
Kemudian program melakukan perulangan: for i := 1; i < n; i++
Perulangan dimulai dari data kedua karena data pertama sudah dijadikan nilai awal minimum dan maksimum.
Di dalam perulangan dilakukan pengecekan:
- Jika arrBerat[i] lebih kecil dari *bMin, maka nilai minimum diperbarui dengan data tersebut.
- Jika arrBerat[i] lebih besar dari *bMax, maka nilai maksimum diperbarui dengan data tersebut.
Proses ini dilakukan terus hingga seluruh data selesai diperiksa satu per satu.
Function rerata() digunakan untuk menghitung rata-rata berat balita.
Di dalam function rerata(), dibuat variabel: total
Yang digunakan untuk menyimpan jumlah seluruh berat balita.
Kemudian program melakukan perulangan: for i := 0; i < n; i++
Di dalam perulangan, seluruh data berat balita dijumlahkan menggunakan: total = total + arrBerat[i]
Setelah semua data selesai dijumlahkan, function mengembalikan nilai rata-rata dengan: return total / float64(n)
Bagian float64(n) digunakan agar hasil pembagian menjadi bilangan desimal.
func main() {...} merupakan bagian utama program tempat seluruh proses dijalankan.
Pada awal program dibuat beberapa variabel:
- arrBerat digunakan untuk menyimpan seluruh data berat balita
- n digunakan untuk menyimpan jumlah data balita yang akan diinput
- bMin digunakan untuk menyimpan berat minimum balita
- bMax digunakan untuk menyimpan berat maksimum balita
- rata digunakan untuk menyimpan rerata berat balita
Setelah itu program meminta user memasukkan jumlah data berat balita menggunakan fmt.Print dan fmt.Scan.
Selanjutnya program melakukan input berat balita menggunakan perulangan for. Perulangan dimulai dari i = 0 dan akan terus berjalan hingga jumlah data yang dimasukkan terpenuhi.
Setiap berat balita yang diinput akan disimpan ke dalam array arrBerat sesuai indeksnya masing-masing.
Setelah seluruh data selesai dimasukkan, program memanggil function: hitungMinMax(arrBerat, n, &bMin, &bMax)
Function tersebut digunakan untuk mencari nilai minimum dan maksimum.
Kemudian program memanggil function: rata = rerata(arrBerat, n)
Function tersebut digunakan untuk menghitung rerata berat balita.
Pada bagian akhir, program akan menampilkan:
- berat balita minimum
- berat balita maksimum
- rerata berat balita
fmt.Printf digunakan untuk menampilkan angka desimal dengan format dua angka di belakang koma menggunakan %.2f.