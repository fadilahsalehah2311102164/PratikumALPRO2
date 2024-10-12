package main

import "fmt"

func main() {
    var sisi float64

    // Input nilai sisi
    fmt.Print("Masukkan panjang sisi persegi: ")
    fmt.Scan(&sisi)

    // Menghitung luas dan keliling
    luas := sisi * sisi
    keliling := 4 * sisi

    // Menampilkan hasil
    fmt.Printf("Luas persegi: %.2f\n", luas)
    fmt.Printf("Keliling persegi: %.2f\n", keliling)
}