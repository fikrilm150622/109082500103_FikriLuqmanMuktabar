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