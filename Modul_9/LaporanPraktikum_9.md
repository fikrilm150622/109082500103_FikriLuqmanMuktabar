# <h1 align="center">Laporan Praktikum Modul 9 - ARRAY - ALGORITMA DAN PEMROGRAMAN 2 </h1>
<p align="center">Fikri Luqman Muktabar - 109082500103</p>

## Unguided 

### 1. [Soal]
#### Soal_Latihan_1.go

```go
package main

import (
	"fmt"
	"math"
)

type titik struct {
	x int
	y int
}

type lingkaran struct {
	pusat titik
	r int
}

func jarak(p, q titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

func didalam(c lingkaran, p titik) bool {
	return jarak(c.pusat, p) <= float64(c.r)
}

func main() {
	var c1, c2 lingkaran
	var p titik
	fmt.Scan(&c1.pusat.x, &c1.pusat.y, &c1.r)
	fmt.Scan(&c2.pusat.x, &c2.pusat.y, &c2.r)
	fmt.Scan(&p.x, &p.y)
	dalam1 := didalam(c1, p)
	dalam2 := didalam(c2, p)
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
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_9/Output/SS_Soal_Latihan_1.png) 

### [Penjelasan]
Program ini dipakai untuk menentukan posisi sebuah titik terhadap dua lingkaran, apakah titik tersebut ada di dalam lingkaran 1, lingkaran 2, keduanya, atau justru di luar keduanya.
package main artinya program ini adalah program utama yang bisa langsung dijalankan.
import ("fmt", "math") digunakan untuk memanggil package fmt (buat input dan output) dan math (buat perhitungan, khususnya akar kuadrat dengan math.Sqrt).
type titik struct { x int; y int } adalah tipe data bentukan untuk menyimpan koordinat sebuah titik, yaitu x dan y.
type lingkaran struct { pusat titik; r int } adalah tipe data untuk lingkaran, yang terdiri dari titik pusat (bertipe titik) dan jari-jari (r).
func jarak(p, q titik) float64 adalah fungsi untuk menghitung jarak antara dua titik. Rumus yang dipakai adalah rumus jarak biasa, yaitu akar dari (x1 - x2)² + (y1 - y2)². Hasilnya bertipe float64 karena menggunakan akar.
func didalam(c lingkaran, p titik) bool adalah fungsi untuk mengecek apakah suatu titik p berada di dalam lingkaran c. Caranya dengan menghitung jarak titik ke pusat lingkaran, lalu dibandingkan dengan jari-jari. Kalau jaraknya lebih kecil atau sama dengan jari-jari, berarti titiknya ada di dalam lingkaran.
func main() {…} adalah bagian utama program. Di sini dibuat variabel c1 dan c2 untuk menyimpan dua lingkaran, serta p untuk menyimpan titik yang akan dicek.
Data dimasukkan lewat fmt.Scan, mulai dari lingkaran 1 (pusat dan radius), lalu lingkaran 2, dan terakhir koordinat titik.
Setelah itu, program memanggil fungsi didalam untuk masing-masing lingkaran dan hasilnya disimpan di variabel dalam1 dan dalam2.
Terakhir, program menentukan hasilnya pakai if-else. Kalau titik ada di kedua lingkaran, akan ditampilkan "Titik di dalam lingkaran 1 dan 2". Kalau hanya di lingkaran 1, tampil "Titik di dalam lingkaran 1". Kalau hanya di lingkaran 2, tampil "Titik di dalam lingkaran 2". Dan kalau tidak masuk keduanya, tampil "Titik di luar lingkaran 1 dan 2".

### 2. [Soal]
#### Soal_Latihan_2.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_9/Output/SS_Soal_Latihan_2.png)

### [Penjelasan]
Program ini digunakan untuk mengolah data dalam array. Jadi nanti program bisa menampilkan isi array, memisahkan data berdasarkan indeks ganjil dan genap, menampilkan elemen dengan indeks kelipatan tertentu, menghapus data pada indeks tertentu, menghitung rata-rata dan standar deviasi, serta mencari berapa kali suatu angka muncul.
package main artinya ini adalah program utama yang bisa langsung dijalankan.
import ("fmt", "math") dipakai untuk memanggil package fmt (buat input dan output) dan math (buat perhitungan seperti pangkat dan akar).
func main() {…} adalah bagian utama program, semua proses dijalankan di sini.
Di awal program, dibuat variabel n untuk menyimpan jumlah data yang akan dimasukkan. Setelah itu dibuat array data dengan ukuran sesuai n menggunakan make.
Selanjutnya program meminta user memasukkan nilai satu per satu ke dalam array lewat perulangan for, mulai dari indeks ke-0 sampai ke-(n-1).
Bagian a menampilkan semua isi array dengan cara looping dan mencetak setiap elemen.
Bagian b menampilkan elemen dengan indeks ganjil. Program ngecek pakai i % 2 != 0, jadi yang muncul cuma indeks 1, 3, 5, dan seterusnya.
Bagian c menampilkan elemen dengan indeks genap. Kondisinya i % 2 == 0, jadi termasuk indeks 0, 2, 4, dan seterusnya.
Bagian d menampilkan elemen dengan indeks kelipatan tertentu. User diminta memasukkan nilai kelipatan. Kalau user memasukkan 0, program langsung kasih pesan error karena tidak boleh dibagi 0. Kalau valid, program menampilkan elemen dengan indeks yang memenuhi i % kelipatan == 0.
Bagian e digunakan untuk menghapus elemen pada indeks tertentu. User memasukkan indeks yang mau dihapus. Kalau indeksnya valid, elemen tersebut dihapus dengan cara menggeser semua elemen setelahnya ke kiri. Setelah itu jumlah data (n) dikurangi 1.
Bagian f menghitung rata-rata. Program menjumlahkan semua isi array ke variabel total, lalu dibagi dengan n. Hasilnya disimpan di rata dalam bentuk desimal.
Bagian g menghitung standar deviasi. Caranya, setiap data dikurangi rata-rata, lalu hasilnya dikuadratkan. Semua nilai itu dijumlahkan, dibagi n, lalu diakar pakai math.Sqrt.
Bagian h digunakan untuk menghitung frekuensi suatu angka. User memasukkan angka yang ingin dicari, lalu program menghitung berapa kali angka itu muncul di array menggunakan perulangan.
Di akhir, program menampilkan hasil frekuensi tersebut.

### 3. [Soal]
#### Soal_Latihan_3.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_9/Output/SS_Soal_Latihan_3.png) 

### [Penjelasan]
Program ini dipakai buat mencatat dan menampilkan hasil pertandingan antara dua klub bola. Jadi nanti user masukin nama klub dulu, terus masukin skor pertandingan satu per satu. Setelah input selesai, program baru menampilkan hasil tiap pertandingan (siapa yang menang atau draw).
package main artinya ini adalah program utama yang bisa langsung dijalankan.
import ("fmt") dipakai buat manggil package fmt yang fungsinya untuk input dan output (misalnya Print dan Scan).
func main() {…} adalah bagian inti program, semua proses dijalankan di sini.
Di awal program, dibuat beberapa variabel:
klubA dan klubB buat nyimpan nama klub
skorA dan skorB buat nyimpan skor tiap pertandingan
skorPertandingan buat nyimpan semua skor yang diinput (array 2 dimensi)
hasil buat nyimpan nama klub yang menang
pertandingan buat ngitung jumlah pertandingan
Setelah itu program minta user masukin nama klub A dan klub B.
Lanjut ke bagian input skor pakai perulangan for.
User bakal masukin skor dari pertandingan 1, 2, dan seterusnya.
Setiap input dicek:
Kalau ada skor yang negatif, program langsung berhenti (ini tanda input selesai)
Kalau masih valid, skor disimpan ke array skorPertandingan
Nah setelah semua input selesai, baru deh program mulai ngolah datanya.
Program nge-loop semua pertandingan yang sudah diinput:
Kalau skor klub A lebih besar, berarti yang menang klub A
Kalau skor klub B lebih besar, berarti klub B yang menang
Kalau sama, berarti hasilnya Draw
Setiap hasil ditampilkan dengan format seperti:
Hasil 1 : MU
Hasil 2 : Inter
Kalau ada yang menang, nama klubnya juga disimpan ke array hasil.
Di bagian akhir, program menampilkan:
Pertandingan selesai

### 4. [Soal]
#### Soal_Latihan_4.go

```go
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0
	for {
		fmt.Scanf("%c", &ch)
		if ch == '.' || *n >= NMAX {
			break
		}
		if ch != '\n' && ch != ' ' {
			t[*n] = ch
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	var temp rune
	for i := 0; i < n/2; i++ {
		temp = t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	var rev tabel
	for i := 0; i < n; i++ {
		rev[i] = t[i]
	}
	balikanArray(&rev, n)
	for i := 0; i < n; i++ {
		if t[i] != rev[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int
	fmt.Print("Teks         : ")
	isiArray(&tab, &m)

	var asli tabel
	for i := 0; i < m; i++ {
		asli[i] = tab[i]
	}	
	balikanArray(&tab, m)
	fmt.Print("Reverse teks : ")
	cetakArray(tab, m)
	fmt.Print("\nTeks         : ")
	cetakArray(asli, m)
	fmt.Print("Palindrom    ? ")
	if palindrom(asli, m) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_9/Output/SS_Soal_Latihan_4.png) 

### [Penjelasan]
Program ini dipakai buat membaca sekumpulan karakter dari user, terus karakter itu dibalik urutannya (reverse), dan dicek juga apakah termasuk palindrom atau nggak. User masukin huruf satu per satu, dan berhenti kalau sudah ketemu tanda titik (.). Setelah itu program bakal nampilin hasil kebalikannya dan hasil pengecekan palindrom (true atau false).
package main artinya ini adalah program utama yang bisa langsung dijalankan.
import "fmt" dipakai buat manggil package fmt, yang gunanya buat input dan output (misalnya Print dan Scan).
Di program ini nggak cuma ada main, tapi juga ada beberapa prosedur dan fungsi, yaitu:
isiArray buat input karakter
cetakArray buat nampilin isi array
balikanArray buat ngebalik isi array
palindrom buat ngecek palindrom
Di awal program ada:
NMAX buat batas maksimal jumlah karakter (127)
tabel tipe data array buat nyimpen karakter
Masuk ke bagian prosesnya:
isiArray dipakai buat baca input dari user.
User masukin karakter satu per satu, terus disimpen ke array.
Input berhenti kalau ketemu titik (.) atau sudah penuh.
cetakArray dipakai buat nampilin isi array ke layar.
Karakternya ditampilkan satu-satu dengan spasi biar rapi.
balikanArray dipakai buat ngebalik isi array.
Jadi karakter depan ditukar sama belakang, terus maju ke tengah.
palindrom dipakai buat ngecek apakah teks itu palindrom.
Caranya:
Array disalin dulu
Terus dibalik
Lalu dibandingin sama yang asli
Kalau sama berarti palindrom (true)
Kalau beda bukan palindrom (false)
Di bagian main, semua dijalanin:
Pertama, program minta input:
Teks         :
Terus isi array dibalik dan ditampilkan:
Reverse teks :
Setelah itu, teks aslinya ditampilkan lagi:
Teks         :
Terakhir, dicek palindromnya:
Palindrom    ? true / false