# <h1 align="center">Laporan Praktikum Modul 14 - SORTING - ALGORITMA DAN PEMROGRAMAN 2 </h1>
<p align="center">Fikri Luqman Muktabar - 109082500103</p>

## Unguided 

### 1. [Soal]
#### Soal_Latihan_1.go

```go
package main

import "fmt"

const NMAX = 1000

type arrInt [NMAX]int

func selectionSort(T *arrInt, n int) {
	var idx_min, temp int
	i := 1
	for i <= n-1 {
		idx_min = i - 1
		j := i
		for j < n {
			if T[idx_min] > T[j] {
				idx_min = j
			}
			j++
		}
		temp = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = temp
		i++
	}
}

func main() {
	var n, m int
	var rumah arrInt
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}
		selectionSort(&rumah, m)
		for j := 0; j < m; j++ {
			fmt.Print(rumah[j])
			if j < m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_14/Output/SS_Soal_Latihan_1.png)

### [Penjelasan]
Program ini digunakan untuk mengurutkan nomor rumah kerabat Hercules pada setiap daerah secara menaik (ascending) menggunakan algoritma Selection Sort. Data masukan berupa sejumlah daerah, kemudian untuk setiap daerah dimasukkan daftar nomor rumah kerabat. Setelah proses pengurutan selesai, program akan menampilkan nomor rumah yang telah terurut dari yang terkecil hingga terbesar.
package main menunjukkan bahwa program ini merupakan program utama yang dapat langsung dijalankan.
import "fmt" digunakan untuk memanggil package fmt yang berfungsi untuk melakukan proses input dan output, seperti Scan, Print, dan Println.
const NMAX = 1000 digunakan untuk mendefinisikan ukuran maksimum array yang dapat digunakan untuk menyimpan data nomor rumah.
type arrInt [NMAX]t digunakan untuk membuat tipe data array bernama arrInt yang berisi bilangan bulat sebanyak NMAX elemen.
func selectionSort(T *arrInt, n int) merupakan subprogram yang digunakan untuk mengurutkan data dalam array menggunakan metode Selection Sort secara ascending.
Pada subprogram selectionSort terdapat variabel:idx_min yang digunakan untuk menyimpan indeks dari nilai terkecil yang ditemukan pada proses pencarian dan juga temp yang digunakan sebagai variabel sementara saat melakukan pertukaran (swap) data.
Proses pengurutan dilakukan menggunakan perulangan: 
i := 1 
for i <= n-1
Perulangan ini digunakan untuk menentukan posisi data yang akan diisi dengan nilai terkecil berikutnya.
Pada setiap iterasi, nilai awal indeks terkecil disimpan pada: idx_min = i - 1
Kemudian dilakukan pencarian nilai terkecil pada bagian array yang belum terurut menggunakan perulangan: 
j := i 
for j < n
Di dalam perulangan tersebut dilakukan pengecekan: if T[idx_min] > T[j]
Apabila ditemukan nilai yang lebih kecil, maka indeks nilai terkecil diperbarui dengan: idx_min = j
Setelah seluruh data pada bagian yang belum terurut diperiksa, dilakukan proses pertukaran data menggunakan variabel sementara:
temp = T[idx_min]
T[idx_min] = T[i-1]
T[i-1] = temp
Proses ini bertujuan untuk menempatkan nilai terkecil pada posisi yang seharusnya.
Selanjutnya nilai i ditambah satu sehingga proses pengurutan dapat dilanjutkan ke posisi berikutnya sampai seluruh data terurut.
func main() {...} merupakan fungsi utama yang menjalankan seluruh proses program.
Pada awal fungsi utama dibuat beberapa variabel, yaitu:
n digunakan untuk menyimpan banyaknya daerah.
m digunakan untuk menyimpan banyaknya nomor rumah pada setiap daerah.
rumah bertipe arrInt yang digunakan untuk menyimpan data nomor rumah kerabat.
Program terlebih dahulu membaca nilai n menggunakan: fmt.Scan(&n)
Nilai tersebut menyatakan jumlah daerah yang akan diproses.
Selanjutnya dilakukan perulangan: for i := 0; i < n; i++
Perulangan ini digunakan untuk memproses setiap daerah satu per satu.
Pada setiap iterasi, program membaca nilai m yang menyatakan banyaknya nomor rumah pada daerah tersebut.
Kemudian dilakukan perulangan: for j := 0; j < m; j++
Perulangan ini digunakan untuk membaca seluruh nomor rumah dan menyimpannya ke dalam array rumah.
Setelah seluruh data nomor rumah berhasil dibaca, program memanggil subprogram: selectionSort(&rumah, m)
Pemanggilan ini bertujuan untuk mengurutkan seluruh nomor rumah pada daerah tersebut secara menaik menggunakan algoritma Selection Sort.
Setelah proses pengurutan selesai, program menampilkan hasilnya menggunakan perulangan: for j := 0; j < m; j++
Perulangan ini digunakan untuk mencetak seluruh elemen array yang telah terurut.
Di dalam perulangan dilakukan pengecekan: if j < m-1
Kondisi tersebut digunakan agar spasi hanya dicetak di antara angka dan tidak muncul setelah angka terakhir.
Terakhir, fmt.Println() digunakan untuk berpindah ke baris berikutnya sehingga hasil pengurutan setiap daerah ditampilkan pada baris yang berbeda.

### 2. [Soal]
#### Soal_Latihan_2.go

```go
package main

import "fmt"

const NMAX = 1000

type arrInt [NMAX]int

func selectionSortAsc(T *arrInt, n int) {
	var idx_min, temp int
	i := 1
	for i <= n-1 {
		idx_min = i - 1
		j := i
		for j < n {
			if T[idx_min] > T[j] {
				idx_min = j
			}
			j++
		}
		temp = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = temp
		i++
	}
}

func main() {
	var n, m int
	var data, ganjil, genap arrInt
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		for j := 0; j < m; j++ {
			fmt.Scan(&data[j])
		}
		nGanjil := 0
		nGenap := 0
		for j := 0; j < m; j++ {
			if data[j]%2 == 1 {
				ganjil[nGanjil] = data[j]
				nGanjil++
			} else {
				genap[nGenap] = data[j]
				nGenap++
			}
		}
		selectionSortAsc(&ganjil, nGanjil)
		selectionSortAsc(&genap, nGenap)
		pertama := true
		for j := 0; j < nGanjil; j++ {
			if !pertama {
				fmt.Print(" ")
			}
			fmt.Print(ganjil[j])
			pertama = false
		}
		for j := nGenap - 1; j >= 0; j-- {
			if !pertama {
				fmt.Print(" ")
			}
			fmt.Print(genap[j])
			pertama = false
		}
		fmt.Println()
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_14/Output/SS_Soal_Latihan_2.png)

### [Penjelasan]
Program ini digunakan untuk mengelompokkan nomor rumah kerabat Hercules menjadi dua kategori, yaitu nomor rumah ganjil dan genap. Selanjutnya, nomor rumah ganjil diurutkan secara menaik (ascending), sedangkan nomor rumah genap diurutkan secara menaik terlebih dahulu kemudian ditampilkan secara terbalik sehingga menghasilkan urutan menurun (descending). Setelah itu, program menampilkan seluruh nomor rumah ganjil terlebih dahulu, diikuti nomor rumah genap.
package main menunjukkan bahwa program ini merupakan program utama yang dapat langsung dijalankan.
import "fmt" digunakan untuk memanggil package fmt yang berfungsi untuk proses input dan output, seperti Scan, Print, dan Println.
const NMAX = 1000 digunakan untuk menentukan kapasitas maksimum array yang dapat digunakan dalam program.
type arrInt [NMAX]int digunakan untuk mendefinisikan tipe data array bernama arrInt yang berisi bilangan bulat sebanyak NMAX elemen.
func selectionSortAsc(T *arrInt, n int) merupakan subprogram yang digunakan untuk mengurutkan data dalam array secara menaik menggunakan algoritma Selection Sort.
Pada subprogram selectionSortAsc terdapat dua variabel, yaitu: idx_min yang digunakan untuk menyimpan indeks dari nilai terkecil yang ditemukan dan juga temp yang digunakan sebagai variabel sementara saat proses pertukaran data (swap).
Proses pengurutan dimulai dengan perulangan:
i := 1
for i <= n-1
Perulangan ini digunakan untuk menentukan posisi yang akan diisi oleh nilai terkecil berikutnya.
Pada setiap iterasi, indeks nilai terkecil sementara disimpan pada: idx_min = i - 1
Kemudian dilakukan pencarian nilai terkecil pada bagian array yang belum terurut menggunakan perulangan:
j := i
for j < n
Di dalam perulangan tersebut dilakukan pengecekan: if T[idx_min] > T[j]
Apabila ditemukan nilai yang lebih kecil, maka indeks nilai terkecil diperbarui dengan: idx_min = j
Setelah seluruh data diperiksa, dilakukan proses pertukaran nilai menggunakan variabel sementara:
temp = T[idx_min]
T[idx_min] = T[i-1]
T[i-1] = temp
Proses ini bertujuan untuk menempatkan nilai terkecil pada posisi yang seharusnya. Langkah tersebut diulangi sampai seluruh data terurut secara ascending.
func main() {...} merupakan fungsi utama yang menjalankan seluruh proses program.
Pada awal fungsi utama dibuat beberapa variabel, yaitu:
n digunakan untuk menyimpan banyaknya daerah.
m digunakan untuk menyimpan banyaknya nomor rumah pada setiap daerah.
data digunakan untuk menyimpan seluruh nomor rumah yang dibaca.
ganjil digunakan untuk menyimpan nomor rumah ganjil.
genap digunakan untuk menyimpan nomor rumah genap.
Program terlebih dahulu membaca nilai n menggunakan: fmt.Scan(&n)
Nilai tersebut menyatakan jumlah daerah yang akan diproses.
Selanjutnya dilakukan perulangan: for i := 0; i < n; i++
Perulangan ini digunakan untuk memproses setiap daerah satu per satu.
Pada setiap iterasi, program membaca nilai m yang menyatakan banyaknya nomor rumah pada daerah tersebut. Setelah itu dilakukan perulangan: for j := 0; j < m; j++
Perulangan ini digunakan untuk membaca seluruh nomor rumah dan menyimpannya ke dalam array data.
Kemudian dibuat dua variabel pencacah, yaitu: nGanjil := 0 dan nGenap := 0
Variabel nGanjil digunakan untuk menghitung banyaknya nomor rumah ganjil, sedangkan nGenap digunakan untuk menghitung banyaknya nomor rumah genap.
Selanjutnya dilakukan proses pemisahan data menggunakan perulangan: for j := 0; j < m; j++
Pada setiap data dilakukan pengecekan: if data[j]%2 == 1
Kondisi tersebut digunakan untuk menentukan apakah suatu nomor rumah merupakan bilangan ganjil.
Jika kondisi bernilai benar, maka nomor rumah disimpan ke dalam array ganjil dan nilai nGanjil ditambah satu.
Sebaliknya, jika kondisi bernilai salah, maka nomor rumah disimpan ke dalam array genap dan nilai nGenap ditambah satu.
Setelah proses pemisahan selesai, kedua array diurutkan menggunakan subprogram:
selectionSortAsc(&ganjil, nGanjil) dan selectionSortAsc(&genap, nGenap)
Array ganjil diurutkan secara menaik, begitu pula array genap.
Kemudian dibuat variabel: pertama := true
Variabel ini digunakan untuk mengatur pencetakan spasi agar tidak muncul sebelum angka pertama.
Nomor rumah ganjil ditampilkan terlebih dahulu menggunakan perulangan: for j := 0; j < nGanjil; j++
Karena array ganjil sudah diurutkan secara ascending, maka data akan tercetak dari yang terkecil ke yang terbesar.
Setelah seluruh bilangan ganjil dicetak, program menampilkan nomor rumah genap menggunakan perulangan: for j := nGenap - 1; j >= 0; j--
Perulangan dilakukan dari indeks terakhir menuju indeks pertama sehingga data genap yang sebelumnya telah diurutkan menaik akan tampil dalam urutan menurun (descending).
Di dalam kedua perulangan tersebut terdapat pengecekan: if !pertama
Kondisi ini digunakan agar spasi hanya dicetak di antara angka dan tidak dicetak sebelum angka pertama.
Terakhir, fmt.Println() digunakan untuk berpindah ke baris berikutnya sehingga hasil pengolahan setiap daerah ditampilkan pada baris yang berbeda.

### 3. [Soal]
#### Soal_Latihan_3.go

```go
package main

import "fmt"

const NMAX = 1000

type arrInt [NMAX]int

func insertionSort(T *arrInt, n int) {
	var temp int
	i := 1
	for i <= n-1 {
		j := i
		temp = T[j]

		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j--
		}
		T[j] = temp
		i++
	}
}

func main() {
	var T arrInt
	var n, x, jarak int
	var tetap bool
	n = 0
	fmt.Scan(&x)
	for x >= 0 {
		T[n] = x
		n++
		fmt.Scan(&x)
	}
	insertionSort(&T, n)
	for i := 0; i < n; i++ {
		fmt.Print(T[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	if n <= 1 {
		fmt.Println("Data berjarak 0")
	} else {
		jarak = T[1] - T[0]
		tetap = true
		i := 1
		for i < n-1 && tetap {
			if T[i+1]-T[i] != jarak {
				tetap = false
			}
			i++
		}
		if tetap {
			fmt.Println("Data berjarak", jarak)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_14/Output/SS_Soal_Latihan_3.png)

### [Penjelasan]
Program ini digunakan untuk membaca sekumpulan bilangan bulat non-negatif, kemudian mengurutkannya menggunakan metode Insertion Sort secara menaik (ascending). Setelah data terurut, program akan memeriksa apakah selisih antara setiap dua data yang berurutan selalu sama. Jika selisihnya tetap, program akan menampilkan nilai jaraknya. Jika tidak, program akan menampilkan informasi bahwa data memiliki jarak yang tidak tetap.
package main menunjukkan bahwa program ini merupakan program utama yang dapat langsung dijalankan.
import "fmt" digunakan untuk memanggil package fmt yang berfungsi untuk proses input dan output, seperti Scan, Print, dan Println.
const NMAX = 1000 digunakan untuk menentukan kapasitas maksimum array yang digunakan untuk menyimpan data.
type arrInt [NMAX]int digunakan untuk mendefinisikan tipe data array bernama arrInt yang berisi bilangan bulat sebanyak NMAX elemen.
func insertionSort(T *arrInt, n int) merupakan subprogram yang digunakan untuk mengurutkan data dalam array menggunakan algoritma Insertion Sort secara menaik (ascending).
Pada subprogram insertionSort terdapat variabel temp yang digunakan untuk menyimpan sementara nilai yang sedang akan disisipkan ke posisi yang tepat.
Proses pengurutan dimulai dengan:
i := 1
for i <= n-1
Perulangan ini digunakan untuk memilih elemen yang akan disisipkan ke bagian array yang sudah terurut.
Pada setiap iterasi, nilai yang akan disisipkan disimpan terlebih dahulu ke dalam variabel: temp = T[j]
Kemudian dilakukan pencarian posisi yang sesuai menggunakan perulangan: for j > 0 && temp < T[j-1]
Kondisi tersebut digunakan untuk memeriksa apakah nilai temp lebih kecil daripada elemen sebelumnya.
Jika kondisi bernilai benar, maka elemen sebelumnya digeser satu posisi ke kanan menggunakan: T[j] = T[j-1]
Setelah itu indeks j dikurangi satu agar proses pencarian posisi dapat dilanjutkan.
Setelah posisi yang sesuai ditemukan, nilai temp ditempatkan pada posisi tersebut menggunakan: T[j] = temp
Proses ini diulangi hingga seluruh data tersusun secara ascending.
func main() {...} merupakan fungsi utama yang menjalankan seluruh proses program.
Pada awal fungsi utama dibuat beberapa variabel, yaitu:
T bertipe arrInt yang digunakan untuk menyimpan data bilangan bulat.
n digunakan untuk menghitung banyaknya data yang tersimpan.
x digunakan untuk menyimpan data yang sedang dibaca dari masukan.
jarak digunakan untuk menyimpan selisih antar data yang berurutan.
tetap bertipe boolean yang digunakan untuk menandai apakah jarak antar data selalu sama atau tidak.
Nilai awal n diisi dengan: n = 0
yang menunjukkan bahwa belum ada data yang tersimpan di dalam array.
Program kemudian membaca data pertama menggunakan: fmt.Scan(&x)
Selanjutnya dilakukan perulangan: for x >= 0
Perulangan ini akan terus berjalan selama data yang dimasukkan bernilai non-negatif.
Di dalam perulangan, data disimpan ke dalam array menggunakan: T[n] = x
Kemudian nilai n ditambah satu untuk menghitung jumlah data yang berhasil disimpan.
Setelah itu program kembali membaca data berikutnya menggunakan: fmt.Scan(&x)
Perulangan akan berhenti ketika pengguna memasukkan bilangan negatif.
Setelah seluruh data selesai dibaca, program memanggil subprogram: insertionSort(&T, n)
untuk mengurutkan seluruh data yang tersimpan secara menaik menggunakan algoritma Insertion Sort.
Selanjutnya hasil pengurutan ditampilkan menggunakan perulangan: for i := 0; i < n; i++
Perulangan ini digunakan untuk mencetak seluruh elemen array yang telah terurut.
Di dalam perulangan dilakukan pengecekan: if i < n-1
Kondisi tersebut digunakan agar spasi hanya dicetak di antara angka dan tidak muncul setelah angka terakhir.
Setelah seluruh data ditampilkan, program memeriksa apakah jumlah data kurang dari atau sama dengan satu menggunakan: if n <= 1
Jika kondisi tersebut terpenuhi, maka program menampilkan: Data berjarak 0
karena tidak ada pasangan data yang dapat digunakan untuk menghitung selisih.
Jika jumlah data lebih dari satu, program menghitung selisih awal menggunakan: jarak = T[1] - T[0]
Nilai tersebut dijadikan acuan untuk membandingkan selisih data-data berikutnya.
Kemudian variabel: tetap = true
digunakan sebagai penanda bahwa sementara semua data dianggap memiliki jarak yang sama.
Selanjutnya dilakukan perulangan:
i := 1
for i < n-1 && tetap
Perulangan ini digunakan untuk memeriksa selisih setiap pasangan data yang berurutan.
Di dalam perulangan dilakukan pengecekan: if T[i+1]-T[i] != jarak
Apabila ditemukan selisih yang berbeda dari nilai jarak, maka: tetap = false
yang menandakan bahwa data tidak memiliki jarak yang tetap.
Setelah seluruh data diperiksa, dilakukan percabangan: if tetap
Jika nilai tetap masih bernilai true, maka program menampilkan: Data berjarak x
dengan x adalah nilai selisih yang sama pada seluruh data.
Sebaliknya, jika tetap bernilai false, maka program menampilkan: Data berjarak tidak tetap
yang menunjukkan bahwa terdapat perbedaan selisih antar data yang berurutan.

### 4. [Soal]
#### Soal_Latihan_4.go

```go
package main

import "fmt"

const NMAX = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating int
}

type DaftarBuku [NMAX]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(
			&pustaka[i].id,
			&pustaka[i].judul,
			&pustaka[i].penulis,
			&pustaka[i].penerbit,
			&pustaka[i].eksemplar,
			&pustaka[i].tahun,
			&pustaka[i].rating,
		)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var idxMax int
	idxMax = 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[idxMax].rating {
			idxMax = i
		}
	}
	fmt.Println("Buku Terfavorit:")
	fmt.Println("Judul :", pustaka[idxMax].judul)
	fmt.Println("Penulis :", pustaka[idxMax].penulis)
	fmt.Println("Penerbit :", pustaka[idxMax].penerbit)
	fmt.Println("Tahun :", pustaka[idxMax].tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var temp Buku
	i := 1
	for i <= n-1 {
		j := i
		temp = pustaka[j]
		for j > 0 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j--
		}
		pustaka[j] = temp
		i++
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	var batas int
	batas = 5
	if n < 5 {
		batas = n
	}
	fmt.Println("5 Buku Dengan Rating Tertinggi:")
	for i := 0; i < batas; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n, r int) {
	var kiri, kanan, tengah int
	var ketemu bool
	kiri = 0
	kanan = n - 1
	for kiri <= kanan && !ketemu {
		tengah = (kiri + kanan) / 2
		if pustaka[tengah].rating == r {
			ketemu = true
		} else if r > pustaka[tengah].rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	if ketemu {
		fmt.Println("Data Buku Ditemukan:")
		fmt.Println("Judul :", pustaka[tengah].judul)
		fmt.Println("Penulis :", pustaka[tengah].penulis)
		fmt.Println("Penerbit :", pustaka[tengah].penerbit)
		fmt.Println("Tahun :", pustaka[tengah].tahun)
		fmt.Println("Eksemplar :", pustaka[tengah].eksemplar)
		fmt.Println("Rating :", pustaka[tengah].rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var n, ratingCari int
	fmt.Scan(&n)
	DaftarkanBuku(&pustaka, n)
	CetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)
	fmt.Scan(&ratingCari)
	CariBuku(pustaka, n, ratingCari)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_14/Output/SS_Soal_Latihan_4.png)

### [Penjelasan]
Program ini digunakan untuk mengelola data buku dalam sebuah perpustakaan. Program dapat membaca data buku, menampilkan buku dengan rating tertinggi, mengurutkan data buku berdasarkan rating secara menurun menggunakan metode Insertion Sort, menampilkan lima buku dengan rating tertinggi, serta mencari buku berdasarkan rating menggunakan metode Binary Search.
package main menunjukkan bahwa program ini merupakan program utama yang dapat langsung dijalankan.
import "fmt" digunakan untuk memanggil package fmt yang berfungsi untuk proses input dan output, seperti Scan, Print, dan Println.
const NMAX = 7919 digunakan untuk menentukan jumlah maksimum data buku yang dapat disimpan dalam array.
type Buku struct {...} digunakan untuk mendefinisikan tipe data terstruktur (struct) bernama Buku. Struct ini memiliki beberapa field, yaitu:
id untuk menyimpan kode buku.
judul untuk menyimpan judul buku.
penulis untuk menyimpan nama penulis buku.
penerbit untuk menyimpan nama penerbit buku.
eksemplar untuk menyimpan jumlah eksemplar buku.
tahun untuk menyimpan tahun terbit buku.
rating untuk menyimpan nilai rating buku.
type DaftarBuku [NMAX]Buku digunakan untuk mendefinisikan array yang berisi data bertipe Buku.
Subprogram DaftarkanBuku
func DaftarkanBuku(pustaka *DaftarBuku, n int) digunakan untuk membaca dan menyimpan data buku ke dalam array pustaka.
Proses pembacaan data dilakukan menggunakan perulangan: for i := 0; i < n; i++
Perulangan tersebut digunakan untuk membaca sebanyak n data buku.
Di dalam perulangan, setiap field buku dibaca menggunakan:
fmt.Scan(
    &pustaka[i].id,
    &pustaka[i].judul,
    &pustaka[i].penulis,
    &pustaka[i].penerbit,
    &pustaka[i].eksemplar,
    &pustaka[i].tahun,
    &pustaka[i].rating,
)
Data yang dibaca kemudian disimpan ke dalam array pustaka.
Subprogram CetakTerfavorit
func CetakTerfavorit(pustaka DaftarBuku, n int) digunakan untuk mencari dan menampilkan buku yang memiliki rating tertinggi.
Pada awal subprogram dibuat variabel: idxMax = 0
Variabel idxMax digunakan untuk menyimpan indeks buku yang memiliki rating tertinggi.
Selanjutnya dilakukan perulangan: for i := 1; i < n; i++
Perulangan ini digunakan untuk membandingkan rating setiap buku dengan rating tertinggi sementara.
Di dalam perulangan dilakukan pengecekan: if pustaka[i].rating > pustaka[idxMax].rating
Jika ditemukan rating yang lebih tinggi, maka nilai idxMax diperbarui sesuai indeks buku tersebut.
Setelah proses pencarian selesai, program menampilkan judul, penulis, penerbit, dan tahun terbit dari buku yang memiliki rating tertinggi.
Subprogram UrutBuku
func UrutBuku(pustaka *DaftarBuku, n int) digunakan untuk mengurutkan data buku berdasarkan rating secara menurun (descending) menggunakan algoritma Insertion Sort.
Pada subprogram ini dibuat variabel: var temp Buku
Variabel temp digunakan untuk menyimpan sementara data buku yang sedang diproses.
Proses pengurutan dimulai dengan:
i := 1
for i <= n-1
Perulangan ini digunakan untuk memilih elemen yang akan disisipkan ke posisi yang sesuai.
Pada setiap iterasi, data buku yang akan dipindahkan disimpan terlebih dahulu ke dalam: temp = pustaka[j]
Kemudian dilakukan pencarian posisi yang sesuai menggunakan perulangan: for j > 0 && temp.rating > pustaka[j-1].rating
Kondisi tersebut digunakan untuk mengurutkan data berdasarkan rating dari yang terbesar ke yang terkecil.
Jika kondisi bernilai benar, maka data sebelumnya digeser satu posisi ke kanan menggunakan: pustaka[j] = pustaka[j-1]
Setelah posisi yang sesuai ditemukan, data yang tersimpan pada temp ditempatkan pada posisi tersebut menggunakan: pustaka[j] = temp
Proses ini diulangi sampai seluruh data buku terurut berdasarkan rating secara menurun.
Subprogram Cetak5Terbaru
func Cetak5Terbaru(pustaka DaftarBuku, n int) digunakan untuk menampilkan lima buku dengan rating tertinggi.
Pada awal subprogram dibuat variabel: batas = 5
Variabel batas digunakan untuk menentukan jumlah buku yang akan ditampilkan.
Kemudian dilakukan pengecekan: if n < 5
Jika jumlah buku kurang dari lima, maka nilai batas diubah menjadi n agar tidak terjadi akses indeks di luar array.
Selanjutnya program menampilkan judul buku menggunakan perulangan: for i := 0; i < batas; i++
Karena array telah diurutkan secara menurun berdasarkan rating, maka lima elemen pertama merupakan lima buku dengan rating tertinggi.
Subprogram CariBuku
func CariBuku(pustaka DaftarBuku, n, r int) digunakan untuk mencari buku berdasarkan rating menggunakan metode Binary Search.
Pada awal subprogram dibuat beberapa variabel:
kiri digunakan untuk menyimpan batas kiri pencarian.
kanan digunakan untuk menyimpan batas kanan pencarian.
tengah digunakan untuk menyimpan indeks tengah.
ketemu digunakan untuk menandai apakah data ditemukan atau tidak.
Nilai awal diberikan sebagai berikut:
kiri = 0
kanan = n - 1
Selanjutnya dilakukan perulangan: for kiri <= kanan && !ketemu
Perulangan akan terus berjalan selama rentang pencarian masih ada dan data belum ditemukan.
Pada setiap iterasi dihitung indeks tengah menggunakan: tengah = (kiri + kanan) / 2
Kemudian dilakukan pengecekan: if pustaka[tengah].rating == r
Jika rating yang dicari ditemukan, maka: ketemu = true
Jika rating yang dicari lebih besar dari rating pada posisi tengah: r > pustaka[tengah].rating
maka pencarian dilanjutkan ke bagian kiri array.
Sebaliknya, jika rating yang dicari lebih kecil, maka pencarian dilanjutkan ke bagian kanan array.
Apabila data ditemukan, program akan menampilkan seluruh informasi buku yang meliputi judul, penulis, penerbit, tahun terbit, jumlah eksemplar, dan rating.
Jika data tidak ditemukan, program akan menampilkan pesan: Tidak ada buku dengan rating seperti itu
Fungsi Main
func main() {...} merupakan fungsi utama yang menjalankan seluruh proses program.
Pada awal fungsi utama dibuat beberapa variabel, yaitu:
pustaka bertipe DaftarBuku yang digunakan untuk menyimpan seluruh data buku.
n digunakan untuk menyimpan jumlah data buku.
ratingCari digunakan untuk menyimpan rating yang akan dicari.
Program terlebih dahulu membaca jumlah buku menggunakan: fmt.Scan(&n)
Kemudian program memanggil: DaftarkanBuku(&pustaka, n)
untuk membaca seluruh data buku.
Setelah itu program memanggil: CetakTerfavorit(pustaka, n)
untuk menampilkan buku dengan rating tertinggi.
Selanjutnya data buku diurutkan menggunakan: UrutBuku(&pustaka, n)
yang mengurutkan data berdasarkan rating secara menurun menggunakan algoritma Insertion Sort.
Setelah proses pengurutan selesai, program memanggil: Cetak5Terbaru(pustaka, n)
untuk menampilkan lima buku dengan rating tertinggi.
Kemudian program membaca rating yang ingin dicari menggunakan: fmt.Scan(&ratingCari)
Terakhir, program memanggil: CariBuku(pustaka, n, ratingCari)
untuk mencari dan menampilkan data buku yang memiliki rating sesuai dengan masukan pengguna menggunakan metode Binary Search.