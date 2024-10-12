package main

import "fmt"

// Fungsi untuk menghitung luas persegi
func hitungLuas(sisi float64) float64 {
	return sisi * sisi
}

// Fungsi untuk menghitung keliling persegi
func hitungKeliling(sisi float64) float64 {
	return 4 * sisi
}

func main() {
	var sisi float64

	// Tanya pengguna untuk memasukkan sisi persegi
	fmt.Print("Masukkan sisi persegi: ")
	fmt.Scanln(&sisi)

	// Hitung dan tampilkan luas dan keliling
	luas := hitungLuas(sisi)
	keliling := hitungKeliling(sisi)

	fmt.Printf("Luas persegi adalah: %.2f\n", luas)
	fmt.Printf("Keliling persegi adalah: %.2f\n", keliling)
}