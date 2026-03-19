# <h1 align="center">Laporan Praktikum Modul 3 - FUNGSI - </h1>
<p align="center">Fikri Luqman Muktabar - 109082500103</p>

## Unguided 

### 1. [Soal]
#### Soal_Latihan_1.go

```go
package main

import "fmt"

func factorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

    if a < c || b < d {
		fmt.Println("Input tidak valid (harus a >= c dan b >= d)")
		return
	}

    fmt.Println(permutation(a, c), combination(a, c))
	fmt.Println(permutation(b, d), combination(b, d))
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_3/Output/SS%20Soal_Latihan_1%20.png)
### [Penjelasan]
Program ini digunakan untuk menghitung nilai permutasi dan kombinasi dari dua pasang bilangan, yaitu a terhadap c dan b terhadap d, dengan syarat a >= c dan b >= d. `package main` menandakan bahwa program ini merupakan program utama yang dapat langsung dijalankan. `import "fmt"` digunakan untuk mengimpor package fmt yang berfungsi untuk melakukan input dan output. `func factorial(n int) int { ... }` merupakan fungsi yang digunakan untuk menghitung nilai faktorial dari suatu bilangan n, dengan cara mengalikan semua bilangan dari 1 sampai n, dan hasilnya disimpan dalam variabel `hasil`. `func permutation(n, r int) int { ... }` adalah fungsi yang digunakan untuk menghitung permutasi P(n, r) dengan rumus (n! / (n-r)!), sedangkan `func combination(n, r int) int { ... }` digunakan untuk menghitung kombinasi C(n, r) dengan rumus (n! / (r!(n-r)!)). Pada bagian `func main() { ... }`, program mulai dijalankan. `var a, b, c, d int` digunakan untuk mendeklarasikan empat variabel bertipe integer yang akan menyimpan nilai input. `fmt.Scan(&a, &b, &c, &d)` berfungsi untuk membaca empat buah bilangan yang dimasukkan. Selanjutnya, terdapat kondisi `if a < c || b < d { ... }` yang digunakan untuk melakukan validasi input sesuai dengan syarat soal, yaitu jika (a < c) atau (b < d), maka program akan menampilkan pesan `"Input tidak valid (harus a >= c dan b >= d)"` dan menghentikan program dengan `return`. Jika input valid, maka program akan langsung menghitung dan menampilkan hasil menggunakan `fmt.Println(permutation(a, c), combination(a, c))` untuk baris pertama, dan `fmt.Println(permutation(b, d), combination(b, d))` untuk baris kedua.

### 2. [Soal]
#### Soal_Latihan_2.go

```go
package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int

    fmt.Scan(&a, &b, &c)

	hasil1 := f(g(h(a)))

    hasil2 := g(h(f(b)))

	hasil3 := h(f(g(c)))

    fmt.Println(hasil1)
	fmt.Println(hasil2)
	fmt.Println(hasil3)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_3/Output/SS%20Soal_Latihan_2%20.png)
### [Penjelasan]
Program ini digunakan untuk menghitung nilai dari tiga komposisi fungsi, yaitu (fogoh)(a) = f(g(h(a))), (gohof)(b) = g(h(f(b))), dan (hofog)(c) = h(f(g(c))). `package main` menandakan bahwa program ini merupakan program utama yang dapat langsung dijalankan. `import "fmt"` digunakan untuk mengimpor package fmt yang berfungsi untuk melakukan input dan output. `func f(x int) int { return x * x }` merupakan fungsi untuk menghitung kuadrat suatu bilangan sesuai dengan f(x) = x². `func g(x int) int { return x - 2 }` merupakan fungsi untuk mengurangi suatu bilangan dengan 2 sesuai dengan g(x) = x - 2. `func h(x int) int { return x + 1 }` merupakan fungsi untuk menambahkan suatu bilangan dengan 1 sesuai dengan h(x) = x + 1. Pada bagian `func main() { ... }`, program mulai dijalankan. `var a, b, c int` digunakan untuk mendeklarasikan tiga variabel bertipe integer. `fmt.Scan(&a, &b, &c)` digunakan untuk membaca tiga buah bilangan sebagai input. Selanjutnya, `hasil1 := f(g(h(a)))` digunakan untuk menghitung (fogoh)(a), `hasil2 := g(h(f(b)))` digunakan untuk menghitung (gohof)(b), dan `hasil3 := h(f(g(c)))` digunakan untuk menghitung (hofog)(c). Terakhir, hasil ditampilkan menggunakan `fmt.Println(hasil1)`, `fmt.Println(hasil2)`, dan `fmt.Println(hasil3)` dalam tiga baris sebagai output.

### 3. [Soal]
#### Soal_Latihan_3.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

    fmt.Scan(&cx1, &cy1, &r1)

	fmt.Scan(&cx2, &cy2, &r2)

	fmt.Scan(&x, &y)

    dalam1 := didalam(cx1, cy1, r1, x, y)
	dalam2 := didalam(cx2, cy2, r2, x, y)

    if dalam1 && dalam2 {
	    fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_3/Output/SS%20Soal_Latihan_3%20.png)
### [Penjelasan]
Program ini digunakan untuk menentukan posisi sebuah titik terhadap dua lingkaran, apakah titik tersebut berada di dalam lingkaran 1, lingkaran 2, keduanya, atau di luar keduanya. `package main` menandakan bahwa program ini merupakan program utama yang dapat langsung dijalankan. `import "fmt"` digunakan untuk mengimpor package fmt yang berfungsi untuk melakukan input dan output, sedangkan `import "math"` digunakan untuk mengakses fungsi matematika seperti akar kuadrat dan perpangkatan. `func jarak(a, b, c, d float64) float64 { ... }` merupakan fungsi yang digunakan untuk menghitung jarak antara dua titik, yaitu titik (a, b) dan (c, d), dengan menggunakan rumus jarak Euclidean √((a−c)² + (b−d)²). Perhitungan dilakukan dengan `math.Pow` untuk perpangkatan dan `math.Sqrt` untuk akar kuadrat. `func didalam(cx, cy, r, x, y float64) bool { ... }` adalah fungsi yang digunakan untuk mengecek apakah titik (x, y) berada di dalam suatu lingkaran yang memiliki pusat (cx, cy) dan radius r. Fungsi ini akan mengembalikan nilai boolean `true` jika jarak antara titik dan pusat lingkaran kurang dari atau sama dengan radius (berarti titik berada di dalam atau tepat pada lingkaran), dan `false` jika sebaliknya. Pada bagian `func main() { ... }`, program mulai dijalankan. `var cx1, cy1, r1 float64` dan `var cx2, cy2, r2 float64` digunakan untuk mendeklarasikan variabel yang menyimpan koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2. `var x, y float64` digunakan untuk menyimpan koordinat titik yang akan diuji. `fmt.Scan(&cx1, &cy1, &r1)` berfungsi untuk membaca input lingkaran 1, kemudian `fmt.Scan(&cx2, &cy2, &r2)` untuk membaca input lingkaran 2, dan `fmt.Scan(&x, &y)` untuk membaca koordinat titik. Selanjutnya, program memanggil fungsi `didalam` untuk masing-masing lingkaran dan menyimpan hasilnya ke dalam variabel `dalam1` dan `dalam2`. Setelah itu, dilakukan pengecekan menggunakan struktur kondisi `if-else`. Jika `dalam1` dan `dalam2` bernilai `true`, maka program akan menampilkan "Titik di dalam lingkaran 1 dan 2". Jika hanya `dalam1` yang bernilai `true`, maka ditampilkan "Titik di dalam lingkaran 1". Jika hanya `dalam2` yang bernilai `true`, maka ditampilkan "Titik di dalam lingkaran 2". Namun jika keduanya bernilai `false`, maka program akan menampilkan "Titik di luar lingkaran 1 dan 2".
