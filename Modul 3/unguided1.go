// FadIlah Salehah
// 2311102164
// S1 IF 11 5

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("================================")
	fmt.Println("Program untuk Menghitung Permutasi & Kombinasi")
	fmt.Println("================================")

	// Mendapatkan input dari pengguna
	var a, b, c, d int
	a, b, c, d = ambilInput()

	// Validasi masukan
	if !validasiInput(a, b, c, d) {
		fmt.Println("Input tidak valid. Pastikan a >= c dan b >= d.")
		return
	}

	// Menampilkan hasil perhitungan
	tampilkanHasil(a, b, c, d)
}

// Fungsi untuk mengambil input dari pengguna
func ambilInput() (int, int, int, int) {
	var inputA, inputB, inputC, inputD string
	fmt.Print("Masukkan nilai a (bilangan asli): ")
	fmt.Scanln(&inputA)
	fmt.Print("Masukkan nilai b (bilangan asli): ")
	fmt.Scanln(&inputB)
	fmt.Print("Masukkan nilai c (bilangan asli): ")
	fmt.Scanln(&inputC)
	fmt.Print("Masukkan nilai d (bilangan asli): ")
	fmt.Scanln(&inputD)

	// Konversi input string menjadi integer
	a, _ := strconv.Atoi(inputA)
	b, _ := strconv.Atoi(inputB)
	c, _ := strconv.Atoi(inputC)
	d, _ := strconv.Atoi(inputD)

	return a, b, c, d
}

// Fungsi untuk memvalidasi apakah input sesuai syarat
func validasiInput(a, b, c, d int) bool {
	return a >= c && b >= d
}

// Fungsi untuk menghitung faktorial dari sebuah bilangan
func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Fungsi untuk menghitung nilai permutasi P(n, r)
func hitungPermutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

// Fungsi untuk menghitung nilai kombinasi C(n, r)
func hitungKombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

// Fungsi untuk menampilkan hasil perhitungan permutasi dan kombinasi
func tampilkanHasil(a, b, c, d int) {
	fmt.Println("--------------------------------")
	fmt.Println("Hasil Perhitungan:")
	fmt.Printf("Permutasi P(%d, %d) = %d!/(%d-%d)! = %d\n", a, c, a, a, c, hitungPermutasi(a, c))
	fmt.Printf("Kombinasi C(%d, %d) = %d!/((%d!)*(%d-%d)!) = %d\n", a, c, a, c, a, c, hitungKombinasi(a, c))
	fmt.Printf("Permutasi P(%d, %d) = %d!/(%d-%d)! = %d\n", b, d, b, b, d, hitungPermutasi(b, d))
	fmt.Printf("Kombinasi C(%d, %d) = %d!/((%d!)*(%d-%d)!) = %d\n", b, d, b, d, b, d, hitungKombinasi(b, d))
	fmt.Println("--------------------------------")
}
