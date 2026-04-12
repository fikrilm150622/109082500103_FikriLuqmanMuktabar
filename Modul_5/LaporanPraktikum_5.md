# <h1 align="center">Laporan Praktikum Modul 5 - REKURSIF - ALGORITMA DAN PEMROGRAMAN 2 </h1>
<p align="center">Fikri Luqman Muktabar - 109082500103</p>

## Unguided 

### 1. [Soal]
#### Soal_Latihan_1.go

```go
package main

import "fmt"

func DeretFibonacci(a, b, n int) {
	if n < 0 {
		return
	}
	fmt.Print(a, " ")
	DeretFibonacci(b, a+b, n-1)
}

func main() {
	var n int
	fmt.Scan(&n)
	DeretFibonacci(0, 1, n)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_4/Output/SS_Soal_Latihan_1.go.png) belom diganti 

### [Penjelasan]
Program ini digunakan untuk menampilkan deret Fibonacci sebanyak N suku. Berikut adalah penjelasannya; package main menandakan bahwa program ini merupakan program utama yang dapat langsung dijalankan. import "fmt" digunakan untuk mengimpor package fmt yang berfungsi untuk melakukan input dan output. func DeretFibonacci(a, b, n int) merupakan prosedur rekursif yang digunakan untuk mencetak deret Fibonacci. Parameter a menyatakan nilai suku saat ini, b menyatakan nilai suku berikutnya, dan n menyatakan jumlah suku yang masih akan ditampilkan. Pada awal fungsi terdapat kondisi if n < 0 sebagai kondisi berhenti (base case), artinya jika nilai n sudah kurang dari 0 maka proses rekursif dihentikan. Jika kondisi tersebut belum terpenuhi, maka fungsi akan mencetak nilai a menggunakan fmt.Print(a, " "). Setelah itu, fungsi akan memanggil dirinya sendiri dengan DeretFibonacci(b, a+b, n-1), di mana nilai a digantikan dengan b, nilai b menjadi hasil penjumlahan a+b, dan n dikurangi satu. Proses ini akan terus berulang sehingga menghasilkan deret Fibonacci secara berurutan. Selanjutnya, func main() merupakan fungsi utama tempat eksekusi program dimulai. Di dalamnya dideklarasikan variabel n bertipe integer yang digunakan untuk menyimpan nilai input melalui fmt.Scan(&n). Setelah itu, fungsi DeretFibonacci(0, 1, n) dipanggil dengan nilai awal a = 0 dan b = 1 sebagai dua suku pertama deret Fibonacci, sehingga program akan menampilkan deret Fibonacci dari suku ke-0 hingga sebanyak N suku sesuai input yang diberikan. 

### 2. [Soal]
#### Soal_Latihan_2.go

```go
package main

import "fmt"

func bintang(k int) {
	if k == 0 {
		return
	}
	fmt.Print("*")
	bintang(k - 1)
}

func pola(i, n int) {
	if i > n {
		return
	}
	bintang(i)
	fmt.Println()
	pola(i+1, n)
}

func main() {
	var N int
	fmt.Scan(&N)
	pola(1, N)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_4/Output/SS_Soal_Latihan_2.go.png) belom diganti

### [Penjelasan]
Program ini digunakan untuk menampilkan pola bintang dari bilangan N, dimulai dari 1 sampai N. Berikut adalah penjelasannya; package main menandakan bahwa program ini merupakan program utama yang dapat langsung dijalankan. import "fmt" digunakan untuk mengimpor package fmt yang berfungsi untuk melakukan input dan output. func bintang(k int) merupakan fungsi rekursif yang digunakan untuk mencetak sejumlah bintang dalam satu baris. Fungsi ini menerima parameter k yang menyatakan jumlah bintang yang akan dicetak. Pada awal fungsi terdapat kondisi if k == 0 sebagai kondisi berhenti (base case), artinya jika nilai k sudah 0 maka proses rekursif dihentikan. Jika belum, fungsi akan mencetak satu bintang menggunakan fmt.Print("*"), kemudian memanggil dirinya sendiri dengan bintang(k - 1) sehingga jumlah bintang akan berkurang satu per satu hingga habis. Selanjutnya, func pola(i, n int) merupakan fungsi rekursif yang digunakan untuk mengatur jumlah baris pola. Parameter i menyatakan baris saat ini, sedangkan n adalah jumlah baris maksimal yang diinputkan. Pada awal fungsi terdapat kondisi if i > n sebagai batas berhenti, sehingga ketika baris sudah melebihi n, proses rekursif akan dihentikan. Di dalam fungsi, bintang(i) dipanggil untuk mencetak bintang sesuai nomor baris, kemudian fmt.Println() digunakan untuk pindah ke baris berikutnya. Setelah itu, fungsi memanggil dirinya sendiri dengan pola(i+1, n) untuk melanjutkan ke baris berikutnya hingga mencapai batas n. func main() {…} merupakan fungsi utama tempat eksekusi program dimulai. Di dalamnya dideklarasikan variabel N bertipe integer yang digunakan untuk menyimpan nilai input melalui fmt.Scan(&N). Setelah itu, fungsi pola(1, N) dipanggil dengan nilai awal i = 1 untuk memulai pencetakan pola bintang dari baris pertama hingga baris ke-n, sehingga terbentuk pola bintang yang bertambah secara bertahap.

### 3. [Soal]
#### Soal_Latihan_3.go

```go
package main

import "fmt"

func faktor(i, n int) {
	if i > n {
		return 
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	faktor(i+1, n) 
}

func main() {
	var N int
	fmt.Scan(&N)
	faktor(1, N)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_4/Output/SS_Soal_Latihan_3.go.png) belom diganti

### [Penjelasan]
Program ini digunakan untuk menampilkan faktor atau bilangan yang apa saja yang habis membagi suatu bilangan N. Berikut adalah penjelasannya; package main menandakan bahwa program ini merupakan program utama yang dapat langsung dijalankan. import "fmt" digunakan untuk mengimpor package fmt yang berfungsi untuk melakukan input dan output. func faktor(i, n int) merupakan sebuah fungsi rekursif yang digunakan untuk mencari dan menampilkan faktor dari bilangan n. Fungsi ini menerima dua parameter, yaitu i sebagai bilangan pembagi yang akan dicek mulai dari 1, dan n sebagai bilangan yang ingin dicari faktornya. Pada awal fungsi terdapat kondisi if i > n yang berfungsi sebagai kondisi berhenti (base case), artinya jika nilai i sudah melebihi n maka proses rekursif akan dihentikan. Selanjutnya terdapat kondisi if n%i == 0 yang digunakan untuk mengecek apakah i merupakan faktor dari n, yaitu ketika n habis dibagi i, maka nilai i akan dicetak menggunakan fmt.Print(i, " "). Setelah itu, fungsi akan memanggil dirinya sendiri dengan faktor(i+1, n) untuk melanjutkan pengecekan ke bilangan berikutnya hingga mencapai n. func main() {…} merupakan fungsi utama tempat eksekusi program dimulai. Di dalamnya dideklarasikan variabel N bertipe integer yang digunakan untuk menyimpan nilai input melalui fmt.Scan(&N). Setelah itu, fungsi faktor(1, N) dipanggil dengan nilai awal i = 1 untuk memulai proses pencarian faktor dari 1 hingga N, sehingga semua faktor dari N akan ditampilkan secara berurutan.

### 4. [Soal]
#### Soal_Latihan_4.go

```go
package main

import "fmt"

func barisan(n int) {
	fmt.Print(n, " ")
	if n > 1 {
		barisan(n - 1)
		fmt.Print(n, " ")
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	barisan(N)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_4/Output/SS_Soal_Latihan_3.go.png) belom diganti

### [Penjelasan]
Program ini digunakan untuk menampilkan barisan bilangan dari N hingga 1, kemudian kembali lagi ke N. Berikut adalah penjelasannya; package main menandakan bahwa program ini merupakan program utama yang dapat langsung dijalankan. import "fmt" digunakan untuk mengimpor package fmt yang berfungsi untuk melakukan input dan output. func barisan(n int) merupakan sebuah fungsi rekursif yang digunakan untuk mencetak barisan bilangan. Fungsi ini tidak mengembalikan nilai, melainkan langsung mencetak hasil menggunakan fmt.Print(n, " "). Pada awal fungsi, nilai n langsung dicetak sehingga membentuk bagian menurun dari N ke 1. Kemudian terdapat kondisi if n > 1 yang berfungsi sebagai pembatas, artinya selama n masih lebih dari 1, fungsi akan terus memanggil dirinya sendiri dengan barisan(n - 1) sehingga nilai akan berkurang satu per satu hingga mencapai 1. Setelah proses rekursif selesai, nilai n akan dicetak kembali menggunakan fmt.Print(n, " "), sehingga membentuk bagian naik dari 1 kembali ke N. func main() {…} merupakan fungsi utama tempat eksekusi program dimulai. Di dalamnya dideklarasikan variabel N bertipe integer yang digunakan untuk menyimpan nilai input melalui fmt.Scan(&N). Setelah itu, fungsi barisan(N) dipanggil untuk menampilkan barisan bilangan dari N ke 1 dan kembali lagi ke N.

### 5. [Soal]
#### Soal_Latihan_5.go

```go
package main

import "fmt"

func ganjil(n int) {
	if n <= 0 {
		return
	}
	ganjil(n - 2)
	fmt.Print(n, " ")
}

func main() {
	var N int
	fmt.Scan(&N)
	if N%2 == 0 {
		N-- 
	ganjil(N)
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_4/Output/SS_Soal_Latihan_3.go.png) belom diganti

### [Penjelasan]
Program ini digunakan untuk menampilkan barisan bilangan ganjil dari 1 hingga N. Berikut adalah penjelasannya; package main menandakan bahwa program ini merupakan program utama yang dapat langsung dijalankan. import "fmt" digunakan untuk mengimpor package fmt yang berfungsi untuk melakukan input dan output. func ganjil(n int) merupakan sebuah fungsi rekursif yang digunakan untuk mencetak bilangan ganjil. Fungsi ini tidak mengembalikan nilai, melainkan langsung mencetak hasil. Di dalam fungsi tersebut terdapat base-case, yaitu jika n <= 0 maka fungsi akan berhenti dan tidak melakukan apa-apa. Jika kondisi tersebut tidak terpenuhi, maka fungsi akan memanggil dirinya sendiri dengan ganjil(n - 2), yang berarti nilai akan dikurangi 2 setiap pemanggilan sehingga hanya memproses bilangan ganjil saja. Setelah proses rekursif selesai, nilai n akan dicetak menggunakan fmt.Print(n, " "), sehingga urutan output tetap dari kecil ke besar. func main() {…} merupakan fungsi utama tempat eksekusi program dimulai. Di dalamnya dideklarasikan variabel N bertipe integer yang digunakan untuk menyimpan nilai input melalui fmt.Scan(&N). Kemudian dilakukan pengecekan if N%2 == 0, jika nilai N adalah genap maka akan dikurangi 1 (N--) agar menjadi bilangan ganjil terdekat. Setelah itu, fungsi ganjil(N) dipanggil untuk menampilkan deret bilangan ganjil dari 1 hingga N.

### 6. [Soal]
#### Soal_Latihan_6.go

```go
package main

import "fmt"

func pangkat(x int, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkat(x, y-1)
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	hasil := pangkat(x, y)
	fmt.Println(hasil)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_4/Output/SS_Soal_Latihan_3.go.png) belom diganti

### [Penjelasan]
Program ini digunakan untuk menghitung hasil pangkat dari dua buah bilangan. Berikut adalah penjelasannya; package main menandakan bahwa program ini merupakan program utama yang dapat langsung dijalankan. import "fmt" digunakan untuk mengimpor package fmt yang berfungsi untuk melakukan input dan output. func pangkat(x int, y int) int merupakan sebuah fungsi rekursif yang digunakan untuk menghitung nilai x dipangkatkan y. Fungsi ini mengembalikan nilai berupa hasil perpangkatan. Di dalam fungsi tersebut terdapat kondisi dasar (base case), yaitu jika y == 0 maka fungsi akan mengembalikan nilai 1, karena setiap bilangan yang dipangkatkan 0 hasilnya adalah 1. Jika kondisi tersebut tidak terpenuhi, maka fungsi akan memanggil dirinya sendiri dengan rumus x * pangkat(x, y-1), yang berarti nilai x akan dikalikan secara berulang sebanyak y kali melalui proses rekursif hingga mencapai kondisi dasar. func main() {…} merupakan fungsi utama tempat eksekusi program dimulai. Di dalamnya dideklarasikan variabel x dan y bertipe integer yang digunakan untuk menyimpan nilai bilangan dan pangkatnya. Nilai tersebut diinputkan menggunakan fmt.Scan(&x, &y). Setelah itu, fungsi pangkat(x, y) dipanggil dan hasilnya disimpan dalam variabel hasil. Nilai hasil tersebut kemudian ditampilkan menggunakan fmt.Println(hasil).