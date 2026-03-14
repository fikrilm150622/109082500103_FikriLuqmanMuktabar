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
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_2/Output/SS%20Soal_Latihan_A.png)
[penjelasan]
Program tersebut digunakan untuk menukar posisi tiga buah string yang diinputkan dengan metode rotasi menggunakan variabel sementara (temp) dan menampilkan hasil sebelum serta sesudah pertukaran. Berikut adalah penjelasannya; package main menandakan bahwa program ini merupakan program utama yang dapat langsung dijalankan. import "fmt" mengimpor package fmt yang digunakan untuk melakukan input dan output, seperti membaca data dari pengguna dan menampilkan hasil ke layar. func main() {…} merupakan fungsi utama tempat eksekusi program dimulai. var ( satu, dua, tiga string temp string ) mendeklarasikan beberapa variabel bertipe string. Variabel satu, dua, dan tiga digunakan untuk menyimpan tiga input string dari pengguna, sedangkan temp digunakan sebagai variabel sementara untuk membantu proses pertukaran nilai. fmt.Print("Masukan input string: ") menampilkan teks “Masukan input string:” ke layar agar pengguna mengetahui bahwa program meminta input. fmt.Scanln(&satu) membaca input pertama dari pengguna dan menyimpannya ke dalam variabel satu. fmt.Print("Masukan input string: ") kembali menampilkan teks yang sama untuk meminta input berikutnya. fmt.Scanln(&dua) membaca input kedua dan menyimpannya ke variabel dua. fmt.Print("Masukan input string: ") menampilkan kembali pesan input. fmt.Scanln(&tiga) membaca input ketiga dan menyimpannya ke variabel tiga. fmt.Println("Output awal = " + satu + " " + dua + " " + tiga) menampilkan nilai ketiga string sesuai dengan urutan awal saat dimasukkan oleh pengguna. temp = satu menyimpan nilai variabel satu ke dalam variabel sementara temp agar nilainya tidak hilang saat proses pertukaran dilakukan. satu = dua mengganti nilai satu dengan nilai dari dua. dua = tiga mengganti nilai dua dengan nilai dari tiga. tiga = temp mengganti nilai tiga dengan nilai yang sebelumnya disimpan di temp (nilai awal dari satu). fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga) menampilkan hasil akhir setelah posisi string ditukar. Dengan proses ini, urutan data berubah dari (satu, dua, tiga) menjadi (dua, tiga, satu). Program ini pada dasarnya melakukan rotasi atau pertukaran posisi tiga buah string menggunakan variabel sementara.

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
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_2/Output/SS%20Soal_Latihan_B.png)
[penjelasan]
Program ini digunakan untuk mengecek apakah susunan warna cairan pada 4 tabung reaksi selama 5 percobaan sesuai dengan urutan yang ditentukan, yaitu merah, kuning, hijau, ungu. Jika semua percobaan memiliki urutan yang benar maka program menampilkan BERHASIL: true, dan jika ada satu saja yang tidak sesuai maka program menampilkan BERHASIL: false. Berikut adalah penjelasannya; package main menandakan bahwa program ini merupakan program utama yang dapat langsung dijalankan. import "fmt" digunakan untuk mengimpor package fmt, yang berfungsi untuk input dan output, seperti membaca data dari pengguna dan menampilkan hasil ke layar. func main() { ... } merupakan fungsi utama tempat program mulai dijalankan. var gelas1, gelas2, gelas3, gelas4 string digunakan untuk mendeklarasikan empat variabel bertipe string yang akan menyimpan warna cairan dari masing-masing gelas atau tabung reaksi. berhasil := true membuat variabel berhasil dengan nilai awal true. Variabel ini digunakan untuk menandai apakah seluruh percobaan memiliki urutan warna yang benar. for percobaan := 1; percobaan <= 5; percobaan++ { ... } merupakan perulangan yang dijalankan sebanyak 5 kali, karena pada soal disebutkan bahwa percobaan dilakukan sebanyak 5 kali. fmt.Printf("Percobaan %d : ", percobaan) digunakan untuk menampilkan tulisan Percobaan ke-berapa sesuai dengan nomor percobaan yang sedang berlangsung. fmt.Scan(&gelas1, &gelas2, &gelas3, &gelas4) digunakan untuk membaca input dari pengguna, yaitu empat warna yang dimasukkan secara berurutan, lalu menyimpannya ke dalam variabel gelas1, gelas2, gelas3, dan gelas4. if gelas1 != "merah" || gelas2 != "kuning" || gelas3 != "hijau" || gelas4 != "ungu" { berhasil = false } digunakan untuk memeriksa apakah urutan warna sesuai dengan yang ditentukan. Jika salah satu warna tidak sesuai dengan urutan merah, kuning, hijau, ungu, maka variabel berhasil akan diubah menjadi false. fmt.Println("BERHASIL:", berhasil) digunakan untuk menampilkan hasil akhir program setelah semua percobaan selesai dilakukan. Jika semua percobaan memiliki urutan yang benar maka hasilnya true, dan jika ada yang salah maka hasilnya false.

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
![Screenshot Output Unguided 1_1](https://github.com/fikrilm150622/109082500103_FikriLuqmanMuktabar/blob/main/Modul_2/Output/SS%20Soal_Latihan_C.png)
[penjelasan]
Program ini digunakan untuk menghitung biaya pengiriman parcel berdasarkan beratnya. Berikut adalah penjelasannya; package main menandakan bahwa program ini merupakan program utama (bisa langsung dijalankan). import "fmt" mengimpor package fmt yang digunakan untuk input dan output, seperti membaca data dan menampilkan hasil. func main() {…} fungsi utama tempat eksekusi program dimulai. var berat, totalBerat, sisaBerat, biayaPerKG, biayaSisa int mendeklarasikan beberapa variabel dalam bentuk integer. fmt.Print("Berat parcel (gram): ") menampilkan teks: “Berat parcel (gram)”. fmt.Scan(&berat) membaca input, dan menyimpan nilainya ke dalam variabel berat. totalBerat = berat / 1000 sisaBerat = berat % 1000 mengubah berat ke formatnya; totalBerat = bagian KG sisaBerat = sisa gramnya. biayaPerKG = totalBerat * 10000 harga per kilogram adalah Rp 10.000. Jadi jumlah kilogram dikali harga per kilo. if sisaBerat > 0 { if berat > 10000 { biayaSisa = 0 } else if sisaBerat <= 500 { biayaSisa = sisaBerat * 5} else { biayaSisa = sisaBerat * 15 } } jika ada sisa berat, lanjut cek; jika total berat lebih dari
10 kg (10000 gr), maka sisa berat gratis maka biayaSisa = 0. Jika sisa berat ≤ 500 gr, maka biaya dihitung sisaBerat × 5. Jika sisa berat > 500 gr, maka biaya dihitung sisaBerat × 15. totalBiaya := biayaPerKG + biayaSisa total biaya dikumpulkan dari biaya per kg + biaya sisa. fmt.Printf("Detail berat: %d kg + %d gr\n", totalBerat, sisaBerat)fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaPerKG, biayaSisa) fmt.Printf("Total biaya: Rp. %d\n", totalBiaya) menampilkan hasilnya.
