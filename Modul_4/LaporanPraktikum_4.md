# <h1 align="center">Laporan Praktikum Modul 4 - PROSEDUR - ALGORITMA DAN PEMROGRAMAN 2 </h1>
<p align="center">Fikri Luqman Muktabar - 109082500103</p>

## Unguided 

### 1. [Soal]
#### Soal_Latihan_1.go

```go
package main

import "fmt"

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var fn, fnr int
	factorial(n, &fn)
	factorial(n-r, &fnr)
	*hasil = fn / fnr
}

func combination(n, r int, hasil *int) {
	var fn, fr, fnr int
	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}

func main() {
	var a, b, c, d int
	var hasilP1, hasilC1, hasilP2, hasilC2 int

	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &hasilP1)
	combination(a, c, &hasilC1)

	permutation(b, d, &hasilP2)
	combination(b, d, &hasilC2)

	fmt.Println(hasilP1, hasilC1)
	fmt.Println(hasilP2, hasilC2)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_4/Output/SS_Soal_Latihan_1.go.png)
### [Penjelasan]
Program ini digunakan untuk menghitung nilai permutasi dan kombinasi dari dua pasang bilangan, yaitu P(n,r) dan C(n,r). Berikut adalah penjelasannya; package main menandakan bahwa program ini merupakan program utama yang dapat langsung dijalankan. import "fmt" digunakan untuk mengimpor package fmt yang berfungsi untuk input dan output. func factorial(n int, hasil *int) merupakan prosedur untuk menghitung nilai faktorial dari suatu bilangan n. Variabel hasil menggunakan pointer (*int) agar nilainya bisa dikembalikan ke pemanggil prosedur. Di dalamnya, *hasil = 1 sebagai nilai awal, lalu dilakukan perulangan for dari 1 sampai n sehingga diperoleh nilai n!. func permutation(n, r int, hasil *int) adalah prosedur untuk menghitung nilai permutasi P(n,r). Di dalamnya dideklarasikan variabel fn dan fnr untuk menyimpan hasil faktorial dari n dan (n−r). Prosedur factorial dipanggil untuk menghitung keduanya, kemudian hasil permutasi dihitung dengan rumus P(n,r) = n! / (n−r)! dan disimpan ke dalam *hasil. func combination(n, r int, hasil *int) adalah prosedur untuk menghitung nilai kombinasi C(n,r). Di dalamnya digunakan variabel fn, fr, dan fnr untuk menyimpan nilai faktorial n, r, dan (n−r). Setelah itu hasil kombinasi dihitung menggunakan rumus C(n,r) = n! / (r!(n−r)!) dan disimpan ke dalam *hasil. func main() {…} merupakan fungsi utama tempat eksekusi program dimulai. Variabel a, b, c, d digunakan untuk menyimpan input bilangan, sedangkan hasilP1, hasilC1, hasilP2, hasilC2 digunakan untuk menyimpan hasil perhitungan permutasi dan kombinasi. fmt.Scan(&a, &b, &c, &d) membaca empat input bilangan. Selanjutnya dilakukan pemanggilan prosedur untuk menghitung P(a,c), C(a,c), P(b,d), dan C(b,d). Terakhir, fmt.Println(hasilP1, hasilC1) dan fmt.Println(hasilP2, hasilC2) digunakan untuk menampilkan hasil dalam dua baris.

### 2. [Soal]
#### Soal_Latihan_2.go

```go
package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	var waktu int
	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)

		if waktu <= 300 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var soal, skor int

	maxSoal := -1
	minSkor := 999999

	for {
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}

		hitungSkor(&soal, &skor)

		if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			maxSoal = soal
			minSkor = skor
			pemenang = nama
		}
	}

	fmt.Println(pemenang, maxSoal, minSkor)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_4/Output/SS_Soal_Latihan_2.go.png)
### [Penjelasan]
Program ini digunakan untuk menentukan pemenang dalam kompetisi pemrograman berdasarkan jumlah soal yang berhasil diselesaikan dan total waktu yang dibutuhkan. Berikut adalah penjelasannya; package main menandakan bahwa program ini merupakan program utama yang dapat langsung dijalankan. import "fmt" digunakan untuk mengimpor package fmt yang berfungsi untuk input dan output. func hitungSkor(soal *int, skor *int) merupakan prosedur yang digunakan untuk menghitung jumlah soal yang berhasil diselesaikan dan total waktu pengerjaan. Parameter soal dan skor menggunakan pointer (*int) agar nilainya dapat dikembalikan ke pemanggil prosedur. Di dalam prosedur ini dideklarasikan variabel waktu untuk menyimpan waktu pengerjaan tiap soal, kemudian *soal dan *skor diinisialisasi dengan nilai 0. Selanjutnya dilakukan perulangan sebanyak 8 kali untuk membaca waktu pengerjaan setiap soal menggunakan fmt.Scan(&waktu). Jika waktu kurang dari atau sama dengan 300 menit, maka soal dianggap berhasil diselesaikan, sehingga *soal ditambah 1 dan *skor ditambah dengan waktu tersebut. func main() {…} merupakan fungsi utama tempat eksekusi program dimulai. Di dalamnya dideklarasikan variabel nama dan pemenang bertipe string untuk menyimpan nama peserta dan nama pemenang, serta variabel soal, skor, maxSoal, dan minSkor bertipe integer untuk menyimpan jumlah soal, total waktu, serta nilai maksimum soal dan minimum waktu terbaik. Variabel maxSoal diinisialisasi dengan -1 dan minSkor dengan nilai besar (999999) sebagai nilai awal pembanding. Program kemudian melakukan perulangan tanpa batas menggunakan for {…} untuk membaca nama peserta dengan fmt.Scan(&nama). Jika nama yang dimasukkan adalah "Selesai", maka perulangan dihentikan dengan break. Jika tidak, prosedur hitungSkor dipanggil untuk menghitung jumlah soal dan skor peserta tersebut. Selanjutnya dilakukan proses seleksi pemenang, yaitu jika jumlah soal peserta lebih besar dari maxSoal, atau jumlah soalnya sama tetapi total waktunya lebih kecil dari minSkor, maka data pemenang akan diperbarui. Setelah seluruh data peserta selesai dibaca, program menampilkan hasil akhir berupa nama pemenang, jumlah soal yang diselesaikan, dan total waktu menggunakan fmt.Println(pemenang, maxSoal, minSkor). 

### 3. [Soal]
#### Soal_Latihan_3.go

```go
package main

import "fmt"

func cetakDeret(n int) {
	for n != 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Print(1)
}

func main() {
	var n int

	fmt.Scan(&n)

	cetakDeret(n)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_4/Output/SS_Soal_Latihan_3.go.png)
### [Penjelasan]
Program ini digunakan untuk mencetak deret bilangan berdasarkan aturan Skiena, yaitu jika suatu bilangan genap maka dibagi 2, dan jika ganjil maka dikalikan 3 lalu ditambah 1, hingga mencapai nilai 1. Berikut adalah penjelasannya; package main menandakan bahwa program ini merupakan program utama yang dapat langsung dijalankan. import "fmt" digunakan untuk mengimpor package fmt yang berfungsi untuk input dan output. func cetakDeret(n int) merupakan sebuah prosedur yang digunakan untuk mencetak setiap suku dari deret berdasarkan nilai awal n. Prosedur ini tidak mengembalikan nilai, melainkan langsung menampilkan hasilnya. Di dalam prosedur ini digunakan perulangan for n != 1 yang berarti proses akan terus berjalan selama nilai n belum sama dengan 1. Pada setiap iterasi, nilai n akan dicetak menggunakan fmt.Print(n, " "). Kemudian dilakukan pengecekan kondisi, jika n adalah bilangan genap (n%2 == 0), maka nilai n diperbarui menjadi n / 2. Jika n adalah bilangan ganjil, maka nilai n diperbarui menjadi 3*n + 1. Proses ini terus berulang hingga n bernilai 1. Setelah perulangan selesai, angka terakhir yaitu 1 dicetak menggunakan fmt.Print(1) sebagai suku penutup deret. func main() {…} merupakan fungsi utama tempat eksekusi program dimulai. Di dalamnya dideklarasikan variabel n bertipe integer yang digunakan untuk menyimpan nilai awal deret. Nilai n diinputkan menggunakan fmt.Scan(&n). Setelah itu, prosedur cetakDeret(n) dipanggil untuk mencetak seluruh deret bilangan sesuai aturan yang telah ditentukan. Hasil akhirnya akan ditampilkan dalam satu baris dengan setiap angka dipisahkan oleh spasi hingga mencapai nilai 1.