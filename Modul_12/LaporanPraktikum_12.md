# <h1 align="center">Laporan Praktikum Modul 12 - SEARCHING - ALGORITMA DAN PEMROGRAMAN 2 </h1>
<p align="center">Fikri Luqman Muktabar - 109082500103</p>

## Unguided 

### 1. [Soal]
#### Soal_Latihan_1.go

```go
package main

import "fmt"

func main() {
	var suara [21]int
	var x int
	var totalMasuk, suaraSah int
	fmt.Scan(&x)
	for x != 0 {
		totalMasuk++
		if x >= 1 && x <= 20 {
			suara[x]++
			suaraSah++
		}
		fmt.Scan(&x)
	}
	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", suaraSah)
	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_10/Output/SS_Soal_Latihan_1.png) belom diganti

### [Penjelasan]
Program ini digunakan untuk membaca, memvalidasi, dan menghitung suara pada pemilihan ketua RT. Data suara dimasukkan dalam bentuk angka yang mewakili nomor calon ketua RT. Kemudan program akan menghitung jumlah seluruh suara yang masuk, jumlah suara sah, serta menampilkan jumlah suara yang diperoleh masing-masing calon.
package main berarti program ini merupakan program utama yang dapat langsung dijalankan.
import "fmt" digunakan untuk memanggil package fmt yang berfungsi untuk proses input dan output, seperti Scan, Println, dan Printf.
func main() {...} merupakan fungsi utama tempat seluruh proses program dijalankan.
Pada awal program dibuat beberapa variabel, yaitu:
suara [21]int yang digunakan untuk menyimpan jumlah suara setiap calon ketua RT. Indeks array digunakan mulai dari 1 sampai 20 sesuai nomor calon. Lalu, x digunakan untuk menyimpan data suara yang diinput, kemudian totalMasuk digunakan untuk menghitung jumlah seluruh data suara yang masuk, dan suaraSah digunakan untuk menghitung jumlah suara yang valid atau sah.
Program kemudian membaca data suara pertama menggunakan fmt.Scan(&x).
Selanjutnya program melakukan proses pembacaan data menggunakan perulangan: for x != 0
Perulangan akan terus berjalan selama data yang dimasukkan bukan 0. Angka 0 digunakan sebagai tanda bahwa proses input telah selesai.
Pada setiap perulangan, variabel totalMasuk akan bertambah satu untuk menghitung seluruh data suara yang diterima, baik valid maupun tidak valid.
Kemudian dilakukan proses validasi data menggunakan percabangan: if x >= 1 && x <= 20
Kondisi tersebut digunakan untuk memeriksa apakah nomor calon berada pada rentang 1 sampai 20. Jika kondisi bernilai benar, maka:
suara[x]++ digunakan untuk menambahkan jumlah suara calon sesuai nomor yang dipilih dan suaraSah++ digunakan untuk menambah jumlah suara sah.
Jika data berada di luar rentang tersebut, maka data dianggap tidak valid dan tidak dihitung sebagai suara sah.
Setelah proses validasi selesai, program kembali membaca data berikutnya menggunakan fmt.Scan(&x) hingga ditemukan angka 0.
Pada bagian akhir, program menampilkan:
jumlah seluruh suara yang masuk menggunakan fmt.Println, jumlah suara sah,serta daftar calon yang memperoleh suara beserta jumlah suaranya.
Untuk menampilkan jumlah suara tiap calon digunakan perulangan: for i := 1; i <= 20; i++
Perulangan ini digunakan untuk memeriksa seluruh calon dari nomor 1 sampai 20.
Di dalam perulangan dilakukan pengecekan: if suara[i] > 0
Kondisi tersebut digunakan agar hanya calon yang memperoleh suara saja yang ditampilkan.
Terakhir, fmt.Printf("%d: %d\n", i, suara[i]) digunakan untuk menampilkan nomor calon dan jumlah suara.

### 2. [Soal]
#### Soal_Latihan_2.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_10/Output/SS_Soal_Latihan_2.png) belom diganti

### [Penjelasan]
Program ini digunakan untuk membaca, memvalidasi, dan menghitung suara pada pemilihan ketua RT. Selain menghitung jumlah suara yang masuk dan suara sah, program juga menentukan calon yang menjadi ketua RT dan wakil ketua RT berdasarkan jumlah suara terbanyak.
package main berarti program ini merupakan program utama yang dapat langsung dijalankan.
import "fmt" digunakan untuk memanggil package fmt yang berfungsi untuk proses input dan output, seperti Scan dan Println.
func main() {...} merupakan fungsi utama tempat seluruh proses program dijalankan.
Pada awal program dibuat beberapa variabel, yaitu:
suara [21]int yang digunakan untuk menyimpan jumlah suara setiap calon ketua RT. Indeks array digunakan mulai dari 1 sampai 20 sesuai nomor calon. Lalu, x digunakan untuk menyimpan data suara yang diinput, kemudian totalMasuk digunakan untuk menghitung jumlah seluruh suara yang masuk, suaraSah digunakan untuk menghitung jumlah suara valid atau sah. Di program yang kedua ini ada tambahan variabel yaitu; ketua yang digunakan untuk menyimpan nomor calon yang memperoleh suara terbanyak, kemudian wakil digunakan untuk menyimpan nomor calon yang memperoleh suara terbanyak kedua, lalu suaraTerbanyak digunakan untuk menyimpan jumlah suara terbesar sementara, dan suaraKedua digunakan untuk menyimpan jumlah suara terbesar kedua sementara.
Program kemudian membaca data pertama menggunakan: fmt.Scan(&x)
Selanjutnya program melakukan proses pembacaan data menggunakan perulangan: for x != 0
Perulangan akan terus berjalan selama data yang dimasukkan bukan 0. Angka 0 digunakan sebagai tanda bahwa proses input telah selesai.
Pada setiap perulangan, variabel totalMasuk akan bertambah satu untuk menghitung seluruh data suara yang diterima.
Kemudian dilakukan proses validasi menggunakan percabangan: if x >= 1 && x <= 20
Kondisi tersebut digunakan untuk memeriksa apakah nomor calon berada pada rentang 1 sampai 20. Jika kondisi bernilai benar, maka:
suara[x]++ digunakan untuk menambahkan jumlah suara calon sesuai nomor yang dipilih dan suaraSah++ digunakan untuk menambah jumlah suara sah.
Jika data berada di luar rentang tersebut, maka data dianggap tidak valid dan tidak dihitung sebagai suara sah.
Setelah seluruh data selesai dibaca, program mencari calon dengan suara terbanyak dan suara terbanyak kedua menggunakan perulangan: for i := 1; i <= 20; i++
Perulangan ini digunakan untuk memeriksa seluruh jumlah suara setiap calon.
Pada bagian: if suara[i] > suaraTerbanyak
program akan memeriksa apakah jumlah suara calon saat ini lebih besar dari jumlah suara terbesar sementara. Jika benar, maka:
nilai suaraTerbanyak sebelumnya dipindahkan ke suaraKedua, lalu nomor ketua sebelumnya dipindahkan menjadi wakil, kemudian jumlah suara terbesar diperbarui dengan suara calon saat ini, dan nomor calon disimpan sebagai ketua RT.
Sedangkan pada bagian: else if suara[i] > suaraKedua
program memeriksa apakah jumlah suara calon lebih besar dari suara terbesar kedua sementara. Jika benar, maka nilai tersebut disimpan sebagai suara terbesar kedua dan nomor calon disimpan sebagai wakil ketua RT.
Pada bagian akhir, program menampilkan:
jumlah seluruh suara yang masuk, jumlah suara sah, nomor calon yang menjadi ketua RT, dan nomor calon yang menjadi wakil ketua RT.
Terakhir, fmt.Println digunakan untuk menampilkan hasil.

### 3. [Soal]
#### Soal_Latihan_3.go

```go
package main

import "fmt"

const NMAX = 1000000
var data [NMAX]int

func main() {
	var n, k int
	var hasil int
	fmt.Scan(&n, &k)
	isiArray(n)
	hasil = posisi(n, k)
	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	var kiri, kanan, tengah int
	kiri = 0
	kanan = n - 1
	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if data[tengah] == k {
			return tengah
		} else if k < data[tengah] {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	return -1
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_10/Output/SS_Soal_Latihan_3.png) belom diganti

### [Penjelasan]
Program ini digunakan untuk mencari suatu bilangan pada kumpulan data integer yang sudah terurut membesar menggunakan metode Binary Search. Program akan menampilkan posisi data yang dicari jika ditemukan, atau menampilkan tulisan “TIDAK ADA” jika data tidak ditemukan.
package main berarti program ini merupakan program utama yang dapat langsung dijalankan.
import "fmt" digunakan untuk memanggil package fmt yang berfungsi untuk proses input dan output, seperti Scan dan Println.
const NMAX = 1000000 digunakan untuk menentukan kapasitas maksimum array, yaitu sebanyak 1.000.000 data.
var data [NMAX]int digunakan untuk membuat array global bernama data yang berfungsi menyimpan seluruh bilangan yang diinput.
func main() {...} merupakan fungsi utama tempat seluruh proses program dijalankan.
Pada fungsi main, dibuat beberapa variabel, yaitu:
n yang digunakan untuk menyimpan jumlah data, lalu k digunakan untuk menyimpan bilangan yang akan dicari, kemudaian hasil digunakan untuk menyimpan posisi data hasil pencarian.
Program kemudian membaca nilai n dan k menggunakan: fmt.Scan(&n, &k)
Setelah itu program memanggil prosedur: isiArray(n)
Prosedur ini digunakan untuk mengisi array data sebanyak n elemen.
Selanjutnya program memanggil fungsi: hasil = posisi(n, k)
Fungsi posisi digunakan untuk melakukan pencarian data menggunakan algoritma Binary Search.
Pada prosedur: func isiArray(n int)
program melakukan pengisian array menggunakan perulangan: for i := 0; i < n; i++
Perulangan tersebut digunakan untuk membaca data satu per satu, kemudian menyimpannya ke dalam array data.
Pada fungsi: func posisi(n, k int) int
dibuat beberapa variabel, yaitu:
kiri yang digunakan untuk menyimpan batas kiri pencarian, kanan digunakan untuk menyimpan batas kanan pencarian, dan tengah digunakan untuk menyimpan indeks tengah array.
Nilai awal pencarian diatur dengan: kiri = 0 dan kanan = n - 1
Kemudian proses pencarian dilakukan menggunakan perulangan: for kiri <= kanan
Perulangan akan terus berjalan selama batas kiri masih lebih kecil atau sama dengan batas kanan.
Di dalam perulangan, indeks tengah dihitung menggunakan: tengah = (kiri + kanan) / 2
Selanjutnya dilakukan pengecekan:
Jika data[tengah] == k, maka data ditemukan dan fungsi mengembalikan posisi tengah. Lalu, jika k < data[tengah], maka pencarian dilanjutkan ke bagian kiri array dengan mengubah nilai kanan. Dan jika tidak, maka pencarian dilanjutkan ke bagian kanan array dengan mengubah nilai kiri.
Jika seluruh proses pencarian selesai tetapi data tidak ditemukan, maka fungsi mengembalikan nilai -1.
Kembali pada fungsi main, program akan memeriksa hasil pencarian:
Jika hasil bernilai -1, maka program akan menampilkan tulisan TIDAK ADA dan jika tidak, program akan menampilkan posisi data yang ditemukan.