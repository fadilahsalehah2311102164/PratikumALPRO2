// FadIlah Salehah
// 2311102164
// S1 IF 11 5

package main

import (
	"fmt"
)

// Fungsi f(x) = x^2
func f(x int) int {
	return x * x
}

// Fungsi g(x) = x - 2
func g(x int) int {
	return x - 2
}

// Fungsi h(x) = x + 1
func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int

	// Menerima input nilai a dan menghitung f(g(h(a)))
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Println("Hasil (fogoh)(a) =", f(g(h(a))))

	// Menerima input nilai b dan menghitung g(h(f(b)))
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Println("Hasil (gohof)(b) =", g(h(f(b))))

	// Menerima input nilai c dan menghitung h(f(g(c)))
	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)
	fmt.Println("Hasil (hofog)(c) =", h(f(g(c))))
}