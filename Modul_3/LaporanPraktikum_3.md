# <h1 align="center">Laporan Praktikum Modul 3 - FUNGSI - ALGORITMA DAN PEMROGRAMAN 2 </h1>
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
[penjelasan]

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
[penjelasan]

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
[penjelasan]
