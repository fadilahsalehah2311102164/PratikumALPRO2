// FadIlah Salehah
// 2311102164
// S1 IF 11 5

package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik
func jarak(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}

// Fungsi untuk menentukan apakah titik (x, y) berada di dalam lingkaran
func diDalamLingkaran(cx, cy, r, x, y int) bool {
	return jarak(cx, cy, x, y) <= float64(r)
}

func main() {
	var cx1, cy1, r1 int // Koordinat pusat dan radius lingkaran 1
	var cx2, cy2, r2 int // Koordinat pusat dan radius lingkaran 2
	var x, y int         // Koordinat titik yang akan dicek

	// Input data lingkaran 1
	fmt.Print("Masukkan koordinat pusat (cx1, cy1) dan radius r1 untuk lingkaran 1: ")
	fmt.Scan(&cx1, &cy1, &r1)

	// Input data lingkaran 2
	fmt.Print("Masukkan koordinat pusat (cx2, cy2) dan radius r2 untuk lingkaran 2: ")
	fmt.Scan(&cx2, &cy2, &r2)

	// Input titik yang akan diperiksa
	fmt.Print("Masukkan koordinat titik (x, y): ")
	fmt.Scan(&x, &y)

	// Memeriksa apakah titik berada di dalam lingkaran 1 atau lingkaran 2
	diDalamL1 := diDalamLingkaran(cx1, cy1, r1, x, y)
	diDalamL2 := diDalamLingkaran(cx2, cy2, r2, x, y)

	// Menampilkan hasil sesuai kondisi
	if diDalamL1 && diDalamL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalamL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalamL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}