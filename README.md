<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL II</strong></h2>
<h2 align="center"><strong> REVIEW STRUKTUR KONTROL </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Fadilah salehah / 2311102164<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh,S.Kom.,M.Kom.
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong><br>
  <strong>FAKULTAS INFORMATIKA</strong><br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong><br>
  <strong>2024</strong>
</p>

------

## Daftar Isi
1. [Tujuan Praktikum](#tujuan-praktikum)
2. [Dasar Teori](#dasar-teori)
3. [Guided](#guided)
4. [Unguided](#unguided)
5. [Kesimpulan](#kesimpulan)
6. [Daftar Pustaka](#daftar-pustaka)

## Tujuan Praktikum
1. Mahasiswa diharapkan memahami berbagai bentuk struktur kontrol seperti perulangan (`for`), pengkondisian (`if, else if, else, dan switch`), serta penggunaan kontrol lainnya seperti `break`, `continue`, dan `goto`.
2. Mahasiswa mampu mengimplementasikan struktur kontrol tersebut dalam program Go, serta mengkombinasikannya untuk membangun alur logika program yang kompleks.
3. Mahasiswa dapat mengenali dan menangani kondisi khusus (edge case) yang mungkin muncul dalam program, sehingga program dapat berjalan dengan benar untuk berbagai situasi input.
   
## Dasar Teori
Review struktur kontrol dalam bahasa Go (Golang) berkaitan dengan berbagai elemen yang digunakan untuk mengontrol aliran eksekusi program. 
Struktur kontrol mengatur bagaimana blok kode tertentu dieksekusi berdasarkan kondisi atau perulangan. [1]
Berikut beberapa dasar teori dari struktur kontrol dalam bahasa Go:

**1. Pernyataan Kondisional (`if, else if, else`)**

- Digunakan untuk mengeksekusi blok kode tertentu jika kondisi tertentu terpenuhi.[1]
- Bentuk dasar:
```go
if kondisi {
    // blok kode dieksekusi jika kondisi true
} else if kondisi_lain {
    // blok kode dieksekusi jika kondisi_lain true
} else {
    // blok kode dieksekusi jika tidak ada kondisi yang terpenuhi
}
```
- Go juga mendukung pendeklarasian variabel dalam pernyataan if :
```go
if nilai := hitungNilai(); nilai > 10 {
    fmt.Println("Nilai lebih besar dari 10")
}
```

**2. Pernyataan Switch (`Switch`)**

- Alternatif dari beberapa pernyataan if else, mempermudah pembacaan kode.[1]
- Tidak seperti bahasa lain, switch Go secara otomatis tidak membutuhkan break untuk menghentikan eksekusi.[1]
- Bentuk dasar
```go
switch ekspresi {
case nilai1:
    // blok kode jika ekspresi == nilai1
case nilai2:
    // blok kode jika ekspresi == nilai2
default:
    // blok kode jika tidak ada nilai yang cocok
}
```
- `switch` juga bisa digunakan tanpa ekspresi untuk mengevaluasi kondisi :
```go
switch {
case x > 10:
    fmt.Println("x lebih besar dari 10")
case x < 5:
    fmt.Println("x lebih kecil dari 5")
default:
    fmt.Println("x berada di antara 5 dan 10")
}
```
**3. Perulangan (`for`)**
   
- Bahasa Go hanya memiliki satu jenis perulangan, yaitu for. Namun, for ini fleksibel dan dapat digunakan untuk berbagai bentuk perulangan.[1]
- Bentuk dasar for
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```
- Perulangan tanpa kondisi (`infinite loop`)
```go
for {
    fmt.Println("Perulangan tanpa akhir")
}
```
- Digunakan untuk mengiterasi elemen dalam slice, array, map, atau channel :
```go
for index, value := range collection {
    fmt.Println(index, value)
}
```

**4. Pernyataan `break` dan `continue`**

- `break` digunakan untuk menghentikan eksekusi dari perulangan atau `switch`[1]
```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // keluar dari loop saat i == 5
    }
}
```
- `continue` digunakan untuk melewati eksekusi kode yang tersisa dalam iterasi dan langsung lanjut ke iterasi berikutnya.
```go
for i := 0; i < 10; i++ {
    if i % 2 == 0 {
        continue // loncat ke iterasi berikutnya jika i genap
    }
    fmt.Println(i) // hanya akan mencetak bilangan ganjil
}
```
**5. Pernyataan `goto`**
- Go mendukung `goto`, meskipun jarang digunakan karena dapat membuat kode sulit diikuti. `goto` memindahkan eksekusi program ke label yang ditentukan.[1]
```go
func main() {
    i := 0
loop:
    if i < 5 {
        fmt.Println(i)
        i++
        goto loop
    }
}
```

**6. Pernyataan `Defer`**
- `defer` digunakan untuk menunda eksekusi fungsi sampai fungsi yang menampung defer selesai.[1]
```go
func contoh() {
    defer fmt.Println("Ini dieksekusi terakhir")
    fmt.Println("Ini dieksekusi pertama")
}
```
- `defer`sering digunakan untuk menutup resource seperti file atau koneksi jaringan, memastikan resource tersebut ditutup meskipun terjadi error.

**7.  `Panic` dan `Recover`**
- `panic` digunakan untuk menghentikan eksekusi program secara tiba-tiba. Biasanya digunakan untuk menangani kesalahan yang tidak bisa ditangani di runtime.[1]
```go
panic("Terjadi kesalahan!")
```
- `recover` digunakan untuk mengontrol dan mengatasi panic, sehingga program bisa kembali berjalan.
```go
func contohPanic() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recover dari panic:", r)
        }
    }()
    panic("Terjadi panic")
}
```

## Guided 

### 1. Program Sederhana untuk Membaca dan Menampilkan Nama

### Source Code :
```go
package main

import (
	"fmt" // Mengimpor paket fmt yang menyediakan fungsi untuk input/output standar
)

func main() {
	var nama string // Deklarasi variabel nama dengan tipe string

	// Menggunakan fmt.Scanln untuk membaca input dari pengguna.
	// Harus menggunakan &nama karena Scanln mengharapkan pointer agar bisa mengisi nilai ke variabel nama
	fmt.Scanln(&nama)

	// Mencetak nilai variabel nama ke layar setelah input diambil dari pengguna
	fmt.Println(nama)
}
```
### Output:
![image](https://github.com/user-attachments/assets/a20a46ea-41e8-4d10-8cbb-52fa13195046)

### Full code Screenshot:
![carbon](https://github.com/user-attachments/assets/c359c2f5-d6fb-4357-aca2-412823f4b83a)

### Deskripsi Program : 
Program ini mengimplementasikan proses dasar untuk membaca input teks dari pengguna dan kemudian menampilkan input tersebut di layar. 
Ini bisa berguna dalam berbagai aplikasi, misalnya sebagai langkah awal untuk menerima nama pengguna atau data lain.

### Algoritma Program

- Inisialisasi:

Import paket fmt untuk melakukan input dan output.

- Deklarasi Variabel:

Buat variabel nama untuk menyimpan input yang akan diberikan oleh pengguna.

- Input:

Gunakan fmt.Scanln(&nama) untuk mengambil input dari pengguna. Fungsi ini akan menunggu pengguna untuk mengetikkan sesuatu dan menekan Enter.

- Output:

Setelah mendapatkan input, gunakan fmt.Println(nama) untuk mencetak nilai dari variabel nama ke layar.

### Cara Kerja Program

- Membaca Input:

Program dimulai dengan menunggu input dari pengguna. Ketika pengguna mengetikkan nama mereka dan menekan Enter, input tersebut akan disimpan ke dalam variabel nama.

- Menampilkan Output:

Setelah input diterima, program mencetak kembali nama tersebut menggunakan fmt.Println, sehingga pengguna bisa melihat apa yang mereka masukkan.

### 2. Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. 
**Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. 
Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. 
Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’, ‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang.
Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. 
Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.
Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read)**

![Screenshot 2024-10-02 193437](https://github.com/user-attachments/assets/6427a80b-5677-4ec7-91d0-b5de64bea11d)

### Source Code :
```go
package main

import (
	"bufio"   // Mengimpor paket bufio untuk membaca input pengguna secara efisien
	"fmt"     // Mengimpor paket fmt untuk fungsi input/output seperti Println
	"os"      // Mengimpor paket os untuk menggunakan fungsi input dari sistem operasi
	"strings" // Mengimpor paket strings untuk manipulasi string, seperti split dan trim
)

func main() {
	// Urutan warna yang benar yang harus dimasukkan oleh pengguna
	correctOrder := []string{"merah", "kuning", "hijau", "ungu"}

	// Membaca input pengguna menggunakan bufio
	reader := bufio.NewReader(os.Stdin)

	// Variabel untuk melacak apakah semua percobaan berhasil
	success := true

	// Melakukan loop untuk 5 percobaan
	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i) // Menampilkan nomor percobaan ke layar

		// Membaca input dari pengguna dan memotong karakter newline di akhir
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // Menghapus spasi dan newline dari input

		// Memisahkan input berdasarkan spasi, menghasilkan array warna
		colors := strings.Split(input, " ")

		// Mengecek apakah setiap warna yang dimasukkan sesuai dengan urutan yang benar
		for j := 0; j < 4; j++ {
			if colors[j] != correctOrder[j] { // Membandingkan dengan urutan warna yang benar
				success = false // Jika salah satu warna tidak sesuai, set flag success menjadi false
				break           // Keluar dari loop jika ada kesalahan
			}
		}

		// Jika ada satu percobaan yang gagal, keluar dari loop utama
		if !success {
			break
		}
	}

	// Menampilkan hasil akhir, apakah semua percobaan berhasil
	if success {
		fmt.Println("BERHASIL: true") // Jika semua percobaan benar
	} else {
		fmt.Println("BERHASIL: false") // Jika ada percobaan yang salah
	}
}
```

### Output:
Jika pengguna berhasil memasukkan urutan yang tepat untuk setiap percobaan, program akan menampilkan hasil BERHASIL: true. 
![Screenshot 2024-10-02 183842](https://github.com/user-attachments/assets/d17310d2-f4a5-464a-8ffb-85aa86415a3c)

Jika ada satu percobaan yang salah, program akan langsung menghentikan percobaan dan menampilkan hasil BERHASIL: false.
![Screenshot 2024-10-02 184014](https://github.com/user-attachments/assets/25b6cb62-7bb6-4fca-b065-1321dcf1239d)

### Full code Screenshot:
![carbon (1)](https://github.com/user-attachments/assets/12c03350-9f26-4e2d-a2f7-0f1fe58f29ce)

### Deskripsi Program :
Program ini meminta pengguna untuk memasukkan urutan warna yang benar selama 5 percobaan. 
Urutan warna yang diharapkan adalah merah, kuning, hijau, dan ungu. Program akan membandingkan input pengguna dengan urutan yang benar. 
Jika pengguna berhasil memasukkan urutan yang tepat untuk setiap percobaan, program akan menampilkan hasil BERHASIL: true. 
Namun, jika ada satu percobaan yang salah, program akan langsung menghentikan percobaan dan menampilkan hasil BERHASIL: false.

### Algoritma Program

- Inisialisasi:

Buat sebuah slice correctOrder yang berisi urutan warna yang benar: "merah", "kuning", "hijau", "ungu".
Gunakan bufio.NewReader untuk memudahkan pembacaan input dari pengguna.
Inisialisasi variabel success untuk melacak apakah semua percobaan berhasil.

- Loop untuk Percobaan:

Lakukan loop sebanyak 5 kali untuk menerima input dari pengguna.
Tampilkan nomor percobaan ke layar.
Baca input dari pengguna, hapus whitespace, dan pisahkan input menjadi array warna.

- Verifikasi Input:

Periksa apakah setiap warna dalam input sesuai dengan urutan yang benar.
Jika ada kesalahan, set success menjadi false dan keluar dari loop.

- Output Hasil:

Setelah semua percobaan, tampilkan apakah semua percobaan berhasil atau tidak.

### Cara Kerja Program
- Program mulai dengan mendefinisikan urutan warna yang benar.
- Menggunakan bufio.Reader, program meminta pengguna untuk memasukkan warna dalam format yang benar (spasi antara warna).
- Setelah menerima input, program memotong spasi yang tidak perlu dan memisahkan warna berdasarkan spasi.
- Program membandingkan setiap warna dengan urutan yang benar.
- Jika semua warna cocok dalam semua percobaan, program menampilkan "BERHASIL: true". Jika ada kesalahan dalam satu percobaan, program menampilkan "BERHASIL: false".


### 3. Penjumlahan 5 Angka dari Input Pengguna

### Source Code :
```go
package main

import (
	"fmt" // Mengimpor paket fmt untuk fungsi input/output
)

func main() {
	// Deklarasi variabel a, b, c, d, e untuk menyimpan input dari pengguna
	var a, b, c, d, e int
	var hasil int // Variabel untuk menyimpan hasil penjumlahan

	// Menggunakan fmt.Scanln untuk membaca 5 angka integer dari input pengguna
	fmt.Scanln(&a, &b, &c, &d, &e)

	// Melakukan penjumlahan dari semua variabel a, b, c, d, e
	hasil = a + b + c + d + e

	// Mencetak hasil penjumlahan ke layar
	fmt.Println("Hasil Penjumlahan ", a, b, c, d, e, "adalah = ", hasil)
}
```
### Output:
![image](https://github.com/user-attachments/assets/61a8cb27-b0f5-4d37-a32b-17c9a7d6a07b)

### Full code Screenshot:
![carbon (2)](https://github.com/user-attachments/assets/211ce685-a1a9-4737-abb5-aa7a793def53)

### Deskripsi Program :
Program ini meminta pengguna untuk memasukkan 5 angka integer, kemudian menghitung dan menampilkan hasil penjumlahan dari kelima angka tersebut. 
Program menggunakan fungsi fmt.Scanln untuk membaca input angka dan menyimpan hasil penjumlahan dalam variabel hasil, lalu mencetak hasilnya.

### Algoritma Program

- Inisialisasi:

Deklarasikan variabel a, b, c, d, dan e untuk menyimpan input dari pengguna.
Deklarasikan variabel hasil untuk menyimpan hasil penjumlahan.

- Input:

Gunakan fmt.Scanln untuk membaca lima angka integer dari pengguna.

- Proses:

Jumlahkan semua angka yang telah diinput ke dalam variabel hasil.

- Output:

Tampilkan hasil penjumlahan ke layar dengan format yang jelas.

### Cara Kerja Program
- Program dimulai dengan mendeklarasikan variabel untuk menyimpan input pengguna.
- Menggunakan fmt.Scanln, program meminta pengguna untuk memasukkan lima angka. Pengguna dapat memasukkan angka secara bersamaan, dipisahkan oleh spasi.
- Setelah angka dibaca, program menjumlahkan kelima angka tersebut dan menyimpannya dalam variabel hasil.
- Program kemudian mencetak hasil penjumlahan ke layar dengan format yang mudah dibaca.

### 4.Diberikan sebuah nilai akhir mata kuliah (NAM) [0..100] dan standar penilaian nilai mata kuliah (NMK) sebagai berikut:

![Screenshot 2024-10-02 185847](https://github.com/user-attachments/assets/62da3edc-4b69-4cda-addf-6491bbc0ff6b)

![Screenshot 2024-10-02 185904](https://github.com/user-attachments/assets/5a7367c6-eb53-4400-91d0-4c4699f48aa3)

**Program berikut menerima input sebuah bilangan riil yang menyatakan NAM. Program menghitung NMK dan menampilkannya.**

**Jawablah pertanyaan-pertanyaan berikut:**

a.	Jika nam diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesuai spesifikasi soal?

b.	Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program seharusnya!

c.	Perbaiki program tersebut! Ujilah dengan masukan: 93.5; 70.6; dan 49.5. Seharusnya keluaran yang diperoleh adalah ‘A’, ‘B’, dan ‘D’.

**Jawaban :**

a. Keluaran jika nam adalah 80.1

Jika nilai nam adalah 80.1, program akan mengecek kondisi apakah nilai tersebut lebih besar dari 80. Karena 80.1 memang lebih besar dari 80, maka nilai indeks yang akan diberikan adalah "A". 
Dengan demikian, keluaran dari program tersebut adalah:

![Screenshot 2024-10-03 095027](https://github.com/user-attachments/assets/8c848f91-63e2-43a5-986e-2b67e5240368)

Hasil keluaran ini sudah sesuai dengan spesifikasi soal, karena nilai 80.1 memang berada di rentang nilai untuk indeks "A".

b. Kesalahan pada program
Meskipun program menghasilkan keluaran yang benar dalam beberapa kasus, terdapat beberapa kekeliruan yang harus diperbaiki:

- Rentang nilai yang kurang tepat:

Pada bagian pengecekan nilai, rentang yang diberikan tidak sepenuhnya jelas. 
Sebagai contoh, tidak ada pengecekan untuk nilai yang persis 72.5. 
Seharusnya rentang nilai perlu didefinisikan dengan lebih jelas, misalnya menggunakan kondisi >= untuk memastikan semua nilai tertangani dengan baik.

- Urutan logika kondisi:

Urutan pengecekan kondisi else if harus diperbaiki. 
Kondisi harus diatur dari nilai tertinggi ke nilai terendah, agar pengecekan dilakukan secara berurutan dan program tidak melewatkan pengecekan nilai yang seharusnya.

c. Perbaikan Program

Untuk memperbaiki program ini, rentang nilai harus didefinisikan lebih jelas dan urutan pengecekan kondisi perlu diperbaiki agar sesuai dengan spesifikasi soal. 

Berikut adalah program yang telah diperbaiki:

```go
package main

import "fmt"

func main() {
	// Deklarasi variabel untuk menyimpan nilai numerik dan nilai huruf
	var nam float32 // Variabel 'nam' untuk menyimpan nilai numerik dengan tipe float32
	var nmk string  // Variabel 'nmk' untuk menyimpan nilai huruf (indeks)

	// Meminta input nilai dari pengguna
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nam) // Membaca input nilai numerik dari pengguna

	// Logika penentuan nilai huruf berdasarkan nilai numerik
	if nam >= 80 {
		nmk = "A" // Jika nilai 80 atau lebih, indeksnya A
	} else if nam >= 70 {
		nmk = "B" // Jika nilai 70 atau lebih, indeksnya B
	} else if nam >= 60 {
		nmk = "C" // Jika nilai 60 atau lebih, indeksnya C
	} else if nam >= 50 {
		nmk = "D" // Jika nilai 50 atau lebih, indeksnya D
	} else if nam >= 40 {
		nmk = "E" // Jika nilai 40 atau lebih, indeksnya E
	} else {
		nmk = "F" // Jika nilai di bawah 40, indeksnya F
	}

	// Menampilkan hasil penentuan nilai huruf (indeks)
	fmt.Printf("Nilai Indeks untuk nilai %.2f adalah %s\n", nam, nmk)
}
```
### Source Code :

**Sebelum Perbaikan**

```go
package main

import "fmt"

func main() {
	// Deklarasi variabel untuk menyimpan nilai numerik dan nilai huruf
	var nam float32 // Variabel 'nam' untuk menyimpan nilai numerik dengan tipe float32
	var nmk string  // Variabel 'nmk' untuk menyimpan nilai huruf (indeks)

	// Meminta input nilai dari pengguna
	fmt.Print("Masukkan nilai : ")
	fmt.Scan(&nam) // Membaca input nilai numerik dari pengguna

	// Logika penentuan nilai huruf berdasarkan nilai numerik
	if nam > 80 {
		nmk = "A" // Jika nilai lebih dari 80, indeksnya A
	} else if nam > 72.5 {
		nmk = "B" // Jika nilai lebih dari 72.5, indeksnya B
	} else if nam > 65 {
		nmk = "C" // Jika nilai lebih dari 65, indeksnya C
	} else if nam > 50 {
		nmk = "D" // Jika nilai lebih dari 50, indeksnya D
	} else if nam > 40 {
		nmk = "E" // Jika nilai lebih dari 40, indeksnya E
	} else {
		nmk = "F" // Jika nilai di bawah atau sama dengan 40, indeksnya F
	}

	// Menampilkan hasil penentuan nilai huruf (indeks)
	fmt.Printf("Nilai Indeks untuk nilai %.2f adalah %s\n", nam, nmk)
}
```

**Setelah Perbaikan**

```go
package main

import "fmt"

func main() {
	// Deklarasi variabel untuk menyimpan nilai numerik dan nilai huruf
	var nam float32 // Variabel 'nam' untuk menyimpan nilai numerik dengan tipe float32
	var nmk string  // Variabel 'nmk' untuk menyimpan nilai huruf (indeks)

	// Meminta input nilai dari pengguna
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nam) // Membaca input nilai numerik dari pengguna

	// Logika penentuan nilai huruf berdasarkan nilai numerik
	if nam >= 80 {
		nmk = "A" // Jika nilai 80 atau lebih, indeksnya A
	} else if nam >= 70 {
		nmk = "B" // Jika nilai 70 atau lebih, indeksnya B
	} else if nam >= 60 {
		nmk = "C" // Jika nilai 60 atau lebih, indeksnya C
	} else if nam >= 50 {
		nmk = "D" // Jika nilai 50 atau lebih, indeksnya D
	} else if nam >= 40 {
		nmk = "E" // Jika nilai 40 atau lebih, indeksnya E
	} else {
		nmk = "F" // Jika nilai di bawah 40, indeksnya F
	}

	// Menampilkan hasil penentuan nilai huruf (indeks)
	fmt.Printf("Nilai Indeks untuk nilai %.2f adalah %s\n", nam, nmk)
}
```

### Output:

**Sebelum Perbaikan**

![Screenshot 2024-10-02 190824](https://github.com/user-attachments/assets/436979d0-a574-4d9f-8f43-2857192e1d3a)

**Setelah Perbaikan**

![Screenshot 2024-10-02 190645](https://github.com/user-attachments/assets/aa9ec8cd-5603-46a3-bf6a-f4ce712fe18a)

### Full code Screenshot:

**Sebelum Perbaikan**

![carbon (4)](https://github.com/user-attachments/assets/34e0a35a-58f3-4103-bf25-e285f27528bc)

**Setelah Perbaikan**

![carbon (3)](https://github.com/user-attachments/assets/22a979c2-0790-4e98-a935-4fdde0c334b7)

### Deskripsi Program :
Program ini menerima input berupa nilai numerik (float) dari pengguna, kemudian menentukan indeks nilai (A, B, C, D, E, atau F) berdasarkan rentang nilai yang sudah ditentukan. 
Program ini menggunakan kondisi if-else untuk mengevaluasi nilai dan menampilkan indeks yang sesuai dengan input pengguna.

### Algoritma Program

- Inisialisasi:

Deklarasikan variabel nam untuk menyimpan nilai numerik dengan tipe float32.
Deklarasikan variabel nmk untuk menyimpan indeks huruf (string).

- Input:

Tampilkan pesan untuk meminta pengguna memasukkan nilai.
Gunakan fmt.Scan untuk membaca input nilai dari pengguna.

- Logika Penentuan Indeks:

Gunakan serangkaian kondisi if-else untuk menentukan indeks huruf berdasarkan nilai yang dimasukkan:
Jika nilai ≥ 80, indeksnya adalah "A".
Jika nilai ≥ 70, indeksnya adalah "B".
Jika nilai ≥ 60, indeksnya adalah "C".
Jika nilai ≥ 50, indeksnya adalah "D".
Jika nilai ≥ 40, indeksnya adalah "E".
Jika nilai < 40, indeksnya adalah "F".

- Output:

Tampilkan hasil indeks huruf berdasarkan nilai yang dimasukkan.

### Cara Kerja Program
- Program dimulai dengan mendeklarasikan dua variabel: satu untuk menyimpan nilai numerik dan satu lagi untuk menyimpan indeks huruf.
- Program meminta pengguna untuk memasukkan nilai dengan menggunakan fmt.Print dan membaca input menggunakan fmt.Scan.
- Setelah menerima input, program mengevaluasi nilai tersebut dengan serangkaian kondisi if-else untuk menentukan indeks huruf yang sesuai.
- Program kemudian mencetak hasil indeks huruf ke layar.

## Unguided 

### 1.Buatlah sebuah program yang menerima input sebuah bilangan bulat positif (dan tidak nol) N, kemudian program akan meminta input berupa nama bunga secara berulang sebanyak N kali dan nama tersebut disimpan ke dalam pita.
**Suatu pita (string) berisi kumpulan nama-nama bunga yang dipisahkan oleh spasi dan ‘–‘, contoh pita diilustrasikan seperti berikut ini.
Pita: mawar – melati – tulip – teratai – kamboja – anggrek**
**Buatlah sebuah program yang menerima input sebuah bilangan bulat positif (dan tidak nol) N, kemudian program akan meminta input berupa nama bunga secara berulang sebanyak N kali dan nama tersebut disimpan ke dalam pita.
(Petunjuk: gunakan operasi penggabungan string dengan operator “+” ). Tampilkan isi pita setelah proses input selesai.**

**Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):**

![Screenshot 2024-10-02 193546](https://github.com/user-attachments/assets/bd92d792-f592-42ce-9b37-0a2a409d15ab)

**Modifikasi program sebelumnya, proses input akan berhenti apabila user mengetikkan ‘SELESAI’. Kemudian tampilkan isi pita beserta banyaknya bunga yang ada di dalam pita
Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):**

![Screenshot 2024-10-02 194044](https://github.com/user-attachments/assets/cc66b37d-2def-4765-a588-a3eaf9468da5)

### Source Code :
```go
package main

import "fmt"

func main() {
	// Deklarasi variabel untuk menyimpan pita dan banyaknya bunga
	var pita string     // Variabel 'pita' digunakan untuk menggabungkan nama bunga
	var banyakBunga int // Variabel 'banyakBunga' untuk menghitung jumlah bunga yang dimasukkan

	// Menggunakan loop untuk meminta input bunga dari pengguna secara berulang
	for {
		var bunga string
		// Meminta input nama bunga
		fmt.Print("Bunga (ketik 'SELESAI' untuk berhenti): ")
		fmt.Scanln(&bunga) // Membaca input dari pengguna

		// Jika pengguna mengetik "SELESAI", maka keluar dari loop
		if bunga == "SELESAI" {
			break
		}

		// Jika pita sudah memiliki isi, tambahkan pemisah " – "
		if pita != "" {
			pita += " – "
		}
		// Menambahkan nama bunga ke pita
		pita += bunga

		// Meningkatkan jumlah bunga yang dimasukkan
		banyakBunga++
	}

	// Menampilkan hasil akhir
	fmt.Println("Pita:", pita)                   // Menampilkan pita yang berisi nama-nama bunga
	fmt.Println("Banyaknya bunga:", banyakBunga) // Menampilkan jumlah bunga yang dimasukkan
}
```
### Output:
![Screenshot 2024-10-03 100321](https://github.com/user-attachments/assets/89d51921-46cc-4e7e-a824-32a83a28f141)

### Full code Screenshot:
![carbon (5)](https://github.com/user-attachments/assets/492fde39-be45-4471-b04f-633967ab8f6a)

### Deskripsi Program :
Program ini berfungsi untuk menerima input dari pengguna berupa nama-nama bunga, menggabungkan nama-nama tersebut dalam satu string yang disebut "pita", serta menghitung jumlah bunga yang dimasukkan. 
Pengguna bisa memasukkan nama bunga secara berulang hingga mengetik "SELESAI", yang menandakan akhir dari input.

### Algoritma Program

- Inisialisasi:

Deklarasikan variabel pita untuk menyimpan gabungan nama-nama bunga.
Deklarasikan variabel banyakBunga untuk menghitung jumlah bunga yang dimasukkan.

- Loop untuk Input:

Gunakan loop for tanpa batas untuk terus meminta input dari pengguna.
Minta pengguna untuk memasukkan nama bunga, dengan pesan yang menunjukkan bahwa mereka dapat mengetik "SELESAI" untuk menghentikan input.

- Pengecekan Input:

Jika pengguna mengetik "SELESAI", keluar dari loop.
Jika pita tidak kosong, tambahkan pemisah " – " sebelum menambahkan nama bunga baru.
Tambahkan nama bunga ke variabel pita dan tingkatkan banyakBunga setiap kali bunga baru dimasukkan.

- Output:

Setelah loop selesai, tampilkan hasil akhir, yaitu nama-nama bunga yang dimasukkan dan jumlah total bunga.

### Cara Kerja Program
- Program mulai dengan mendeklarasikan dua variabel untuk menyimpan nama bunga dan jumlah bunga.
- Dalam loop, program meminta pengguna untuk memasukkan nama bunga. Jika pengguna mengetik "SELESAI", loop akan berhenti.
- Setiap nama bunga yang dimasukkan akan ditambahkan ke variabel pita. Jika pita sudah berisi nama bunga sebelumnya, program akan menambahkan pemisah " – ".
- Setelah keluar dari loop, program mencetak daftar nama bunga dan jumlah bunga yang dimasukkan.

### 2. Buatlah program Pak Andi yang menerima input dua buah bilangan real positif yang menyatakan berat total masing-masing isi kantong terpal. Program akan terus meminta input bilangan tersebut hingga salah satu kantong terpal berisi 9 kg atau lebih.
**Setiap hari Pak Andi membawa banyak barang belanjaan dari pasar dengan mengendarai sepeda motor. Barang belanjaan tersebut dibawa dalam kantong terpal di kiri-kanan motor. Sepeda motor tidak akan oleng jika selisih berat barang di kedua kantong sisi tidak lebih dari 9 kg.
Buatlah program Pak Andi yang menerima input dua buah bilangan real positif yang menyatakan berat total masing-masing isi kantong terpal. Program akan terus meminta input bilangan tersebut hingga salah satu kantong terpal berisi 9 kg atau lebih.
Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):**

![Screenshot 2024-10-03 100719](https://github.com/user-attachments/assets/3d234268-4ab5-4bc1-bbf4-49a8452c4320)

**Pada modifikasi program tersebut, program akan menampilkan true jika selisih kedua isi kantong lebih dari atau sama dengan 9 kg. Program berhenti memproses apabila total berat isi kedua kantong melebihi 150 kg atau salah satu kantong beratnya negatif.
Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):**

![Screenshot 2024-10-03 100729](https://github.com/user-attachments/assets/736f0ca8-3a2c-4add-a3cc-b9835a6ee078)

### Source Code :
```go
package main

import (
	"bufio"   // Mengimpor paket bufio untuk membaca input dari pengguna
	"fmt"     // Mengimpor paket fmt untuk menampilkan output
	"os"      // Mengimpor paket os untuk menggunakan input dari sistem operasi
	"strconv" // Mengimpor paket strconv untuk konversi string ke tipe data numerik
	"strings" // Mengimpor paket strings untuk manipulasi string
)

func main() {
	reader := bufio.NewReader(os.Stdin) // Membuat reader untuk membaca input pengguna

	// Loop untuk meminta input pengguna secara berulang
	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		input, _ := reader.ReadString('\n') // Membaca input pengguna
		input = strings.TrimSpace(input)    // Menghapus spasi dan karakter newline di awal dan akhir

		// Memisahkan input menjadi dua bagian berdasarkan spasi
		berat := strings.Split(input, " ")

		// Mengecek apakah input terdiri dari dua angka
		if len(berat) != 2 {
			fmt.Println("Input tidak valid. Masukkan dua angka.") // Menampilkan pesan kesalahan jika input tidak valid
			continue                                              // Melanjutkan ke iterasi berikutnya untuk meminta input ulang
		}

		// Mengonversi input dari string ke float64
		berat1, _ := strconv.ParseFloat(berat[0], 64)
		berat2, _ := strconv.ParseFloat(berat[1], 64)

		// Jika salah satu berat negatif, sepeda tidak oleng dan keluar dari loop
		if berat1 < 0 || berat2 < 0 {
			fmt.Println("Sepeda motor pak Andi akan oleng: false") // Menampilkan bahwa sepeda tidak oleng
			break                                                  // Keluar dari loop
		}

		// Jika total berat lebih dari 150 kg, sepeda tidak oleng dan keluar dari loop
		if berat1+berat2 > 150 {
			fmt.Println("Sepeda motor pak Andi akan oleng: false") // Sepeda tidak oleng karena berat berlebihan
			break                                                  // Keluar dari loop
		}

		// Menghitung selisih berat antara kedua kantong
		selisih := berat1 - berat2
		if selisih < 0 {
			selisih = -selisih // Mengubah selisih menjadi nilai absolut (selisih selalu positif)
		}

		// Jika selisih berat lebih dari atau sama dengan 9 kg, sepeda akan oleng
		if selisih >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}

	// Pesan yang ditampilkan setelah proses selesai
	fmt.Println("Proses selesai.")
}
```
### Output:
![image](https://github.com/user-attachments/assets/3b1431ef-b348-43e8-9ac0-5b3dfac1d171)

### Full code Screenshot:
![carbon (6)](https://github.com/user-attachments/assets/b65524c3-271f-4f70-93e1-d04f99167a77)

### Deskripsi Program :
Program ini berfungsi untuk menentukan apakah sepeda motor milik Pak Andi akan oleng (tidak stabil) berdasarkan berat belanjaan yang dimasukkan ke dalam dua kantong. 
Pengguna dapat memasukkan berat belanjaan secara berulang, dan program akan memberikan informasi apakah sepeda motor akan oleng atau tidak berdasarkan beberapa kondisi yang ditentukan.

### Algoritma Program

- Inisialisasi:

Gunakan bufio untuk membaca input pengguna dengan efisien.
Mulai loop untuk meminta input dari pengguna secara berulang.

- Input:

Tampilkan prompt untuk meminta pengguna memasukkan berat dua kantong.
Baca input pengguna dan hapus karakter whitespace yang tidak perlu.

- Validasi Input:

Pisahkan input berdasarkan spasi dan cek apakah terdapat dua angka.
Jika tidak, tampilkan pesan kesalahan dan lanjutkan ke iterasi berikutnya.

- Konversi dan Pengecekan Berat:

Konversi input dari string ke float64 menggunakan strconv.ParseFloat.
Jika salah satu berat negatif, tampilkan bahwa sepeda tidak oleng dan keluar dari loop.
Jika total berat lebih dari 150 kg, tampilkan bahwa sepeda tidak oleng dan keluar dari loop.

- Hitung Selisih:

Hitung selisih berat antara kedua kantong dan ubah menjadi nilai absolut jika negatif.
Jika selisih lebih dari atau sama dengan 9 kg, tampilkan bahwa sepeda akan oleng. Jika tidak, tampilkan bahwa sepeda tidak oleng.

- Output Akhir:

Tampilkan pesan bahwa proses telah selesai setelah loop keluar.

### Cara Kerja Program
- Program dimulai dengan mendeklarasikan reader untuk membaca input.
- Dalam loop, program meminta pengguna untuk memasukkan berat dua kantong.
- Program memvalidasi input dan memisahkan angka-angka menjadi dua variabel.
- Setelah konversi, program melakukan pengecekan untuk menentukan apakah sepeda motor akan oleng berdasarkan kriteria yang telah ditentukan.
- Hasilnya dicetak ke layar untuk setiap iterasi input, dan jika kondisi tertentu terpenuhi (seperti berat negatif atau total berat lebih dari 150 kg), loop akan berhenti.


### 3. Diberikan sebuah persamaan sebagai berikut ini.

![image](https://github.com/user-attachments/assets/d7e30bf0-2cca-46fa-b2a4-2c3e2baae551)

**Buatlah sebuah program yang menerima input sebuah bilangan sebagai K, kemudian menghitung dan menampilkan nilai f(K) sesuai persamaan di atas.
Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):**

![image](https://github.com/user-attachments/assets/d870a04f-d59a-4723-8ee5-8e1dc219ec11)

### Source Code :
```go
package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel k dengan tipe float64
	var k float64
	// Menampilkan pesan untuk memasukkan nilai K
	fmt.Print("Nilai K = ")
	// Membaca input dari pengguna dan menyimpannya ke variabel k
	fmt.Scanln(&k)

	// Hitung f(K) dengan rumus: f(k) = ((4*k + 2)^2) / ((4*k + 1) * (4*k + 3))
	f_k := (4*k + 2) * (4*k + 2) / ((4*k + 1) * (4*k + 3))

	// Menampilkan hasil perhitungan f(K) dengan presisi 10 angka di belakang koma
	fmt.Printf("Nilai f(K) = %.10f\n", f_k)
}
```
### Output:
![Screenshot 2024-10-03 074925](https://github.com/user-attachments/assets/75c40d98-64ff-4b73-b747-e83cd6f7bdde)

### Full code Screenshot:
![carbon (7)](https://github.com/user-attachments/assets/02f2077d-4fde-4908-bb11-617699a2dd94)

### Deskripsi Program :
Program ini berfungsi untuk menghitung nilai dari fungsi matematis 𝑓(𝐾) berdasarkan input yang diberikan oleh pengguna. 
Pengguna diminta untuk memasukkan nilai 𝐾 dan program akan menghitung serta menampilkan nilai 𝑓(𝐾) dengan presisi tinggi.

### Algoritma Program

- Inisialisasi:

Deklarasikan variabel k dengan tipe float64 untuk menyimpan input dari pengguna.

- Input:

Tampilkan pesan untuk meminta pengguna memasukkan nilai 𝐾
Gunakan fmt.Scanln untuk membaca nilai 𝐾 dari input pengguna.

- Hitung Fungsi:

Hitung nilai 
𝑓(𝐾) menggunakan rumus yang telah ditentukan.

- Output:

Tampilkan hasil perhitungan dengan format yang menunjukkan 10 angka di belakang koma.

### Cara Kerja Program
- Program mulai dengan mendeklarasikan variabel k untuk menyimpan nilai yang diinputkan oleh pengguna.
- Program kemudian meminta pengguna untuk memasukkan nilai 𝐾
- Setelah menerima input, program menghitung 𝑓(𝐾) menggunakan rumus yang telah ditetapkan.
- Program mencetak hasilnya ke layar dengan presisi 10 angka di belakang koma.

### 4. PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut!
**Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.
Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):**

![image](https://github.com/user-attachments/assets/1008c544-23ad-459f-b53c-8ecfd629352c)

### Source Code :
```go
package main

import "fmt"

// Fungsi untuk menghitung biaya pengiriman
func hitungBiaya(beratGram int) (int, int, int) {
	// Konversi berat dari gram ke kg
	beratKg := beratGram / 1000
	sisaGram := beratGram % 1000

	// Hitung biaya pengiriman per kg
	biayaKg := beratKg * 10000

	// Hitung biaya pengiriman untuk sisa gram
	var biayaSisa int
	if sisaGram >= 500 {
		biayaSisa = sisaGram * 5
	} else if sisaGram > 0 && beratKg <= 10 {
		biayaSisa = sisaGram * 15
	} else {
		biayaSisa = 0
	}

	return biayaKg, biayaSisa, beratKg
}

func main() {
	var beratGram int
	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scanln(&beratGram)

	// Panggil fungsi hitungBiaya untuk mendapatkan biaya pengiriman
	biayaKg, biayaSisa, beratKg := hitungBiaya(beratGram)

	// Tampilkan hasil perhitungan
	fmt.Printf("Total berat: %d kg %d gram\n", beratKg, beratGram%1000)
	fmt.Printf("Biaya pengiriman: Rp. %d\n", biayaKg+biayaSisa)
	fmt.Printf("  - Biaya per kg: Rp. %d\n", biayaKg)
	fmt.Printf("  - Biaya sisa gram: Rp. %d\n", biayaSisa)
}
```
### Output:
![image](https://github.com/user-attachments/assets/9c08af01-7ca3-451b-ac57-b2bee28d9556)

### Full code Screenshot:
![carbon (8)](https://github.com/user-attachments/assets/c4542d27-c4bc-45c2-9ee2-6274f7d3519d)

### Deskripsi Program :
Program ini berfungsi untuk menghitung biaya pengiriman parsel berdasarkan beratnya dalam gram. 
Program ini meminta pengguna untuk memasukkan berat parsel, kemudian menghitung biaya pengiriman berdasarkan berat tersebut dan menampilkan hasilnya dengan jelas.

### Algoritma Program

- Inisialisasi:

Deklarasikan fungsi hitungBiaya yang menerima berat parsel dalam gram sebagai parameter dan mengembalikan biaya pengiriman per kilogram, biaya pengiriman untuk sisa gram, dan berat dalam kilogram.

- Hitung Berat dan Biaya:

Dalam fungsi hitungBiaya:
Hitung berat dalam kilogram dan sisa gram.
Hitung biaya pengiriman untuk berat dalam kilogram (setiap kilogram dikenakan biaya Rp 10.000).
Hitung biaya pengiriman untuk sisa gram dengan ketentuan tertentu:
Jika sisa lebih dari atau sama dengan 500 gram, dikenakan biaya Rp 5 per gram.
Jika sisa lebih dari 0 dan berat total (kg) tidak lebih dari 10, dikenakan biaya Rp 15 per gram.
Jika sisa tidak memenuhi kondisi di atas, biaya untuk sisa gram adalah 0.

- Input Pengguna:

Minta pengguna untuk memasukkan berat parsel dalam gram.

- Output:

Tampilkan berat total (dalam kg dan gram), total biaya pengiriman, dan rincian biaya per kilogram dan biaya sisa gram.

### Cara Kerja Program
- Program dimulai dengan mendeklarasikan fungsi hitungBiaya.
- Fungsi ini menghitung biaya berdasarkan berat parsel yang dimasukkan oleh pengguna.
- Di dalam fungsi main, pengguna diminta untuk memasukkan berat parsel.
- Setelah itu, fungsi hitungBiaya dipanggil untuk mendapatkan biaya pengiriman, yang kemudian ditampilkan ke layar.

### 5. Buatlah program yang menerima input sebuah bilangan bulat b dan b > 1. Program harus dapat mencari dan menampilkan semua faktor dari bilangan tersebut!
**Sebuah bilangan bulat b memiliki faktor bilangan f > 0 jika f habis membagi b. Contoh: 2 merupakan faktor dari bilangan 6 karena 6 habis dibagi 2.
Buatlah program yang menerima input sebuah bilangan bulat b dan b > 1. Program harus dapat mencari dan menampilkan semua faktor dari bilangan tersebut!
Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):**

![image](https://github.com/user-attachments/assets/d174fa11-3d6b-48a3-a353-1c14b4c3e323)

**Bilangan bulat b > 0 merupakan bilangan prima p jika dan hanya jika memiliki persis dua faktor bilangan saja, yaitu 1 dan dirinya sendiri.
Lanjutkan program sebelumnya. Setelah menerima masukan sebuah bilangan bulat b > 0. Program tersebut mencari dan menampilkan semua faktor bilangan tersebut. Kemudian, program menentukan apakah b merupakan bilangan prima.
Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):**

![image](https://github.com/user-attachments/assets/262cc7ad-da31-45f3-8c86-bed9ce51b475)

### Source Code :
```go
package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel untuk menyimpan input bilangan
	var b int

	// Meminta pengguna untuk memasukkan bilangan
	fmt.Print("Bilangan: ")
	fmt.Scanln(&b)

	// Mencetak faktor-faktor dari bilangan
	fmt.Print("Faktor: ")
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
		}
	}

	// Mengecek apakah bilangan adalah bilangan prima
	prima := true
	if b <= 1 {
		prima = false
	} else {
		// Loop untuk mencari pembagi selain 1 dan b
		for i := 2; i*i <= b; i++ {
			if b%i == 0 {
				prima = false
				break
			}
		}
	}

	// Menampilkan apakah bilangan prima atau bukan
	fmt.Println("\nPrima:", prima)
}
```
### Output:
![image](https://github.com/user-attachments/assets/9c3eeecd-21e0-4fd4-80b5-193212dfe35a)

### Full code Screenshot:
![carbon (9)](https://github.com/user-attachments/assets/37db835c-4ca0-434f-92fa-018813f8501a)

### Deskripsi Program :
Program ini berfungsi untuk menghitung faktor dari bilangan bulat yang dimasukkan oleh pengguna serta menentukan apakah bilangan tersebut adalah bilangan prima atau bukan.

### Algoritma Program

- Input:

Program meminta pengguna untuk memasukkan sebuah bilangan bulat.

- Mencetak Faktor:

Menggunakan loop dari 1 hingga bilangan yang dimasukkan untuk menemukan dan mencetak semua faktor.
Jika bilangan dapat dibagi tanpa sisa oleh angka dalam loop, maka angka tersebut adalah faktor.

- Pengecekan Bilangan Prima:

Jika bilangan kurang dari atau sama dengan 1, maka bilangan tersebut bukan bilangan prima.
Untuk bilangan lebih dari 1, program akan melakukan loop dari 2 hingga akar bilangan tersebut untuk mencari pembagi.
Jika ditemukan pembagi lain selain 1 dan bilangan itu sendiri, maka bilangan tersebut bukan bilangan prima.

- Output:

Menampilkan semua faktor dari bilangan dan menyatakan apakah bilangan tersebut adalah bilangan prima.

### Cara Kerja Program
- Program dimulai dengan mendeklarasikan variabel b untuk menyimpan input bilangan bulat dari pengguna.
- Setelah pengguna memasukkan bilangan, program akan mencetak faktor-faktor bilangan tersebut.
- Program kemudian memeriksa apakah bilangan tersebut adalah bilangan prima.
- Hasil akhir dicetak ke layar.

## Kesimpulan 
Berdasarkan hasil praktikum Review Struktur Kontrol pada bahasa pemrograman Go, dapat disimpulkan bahwa pemahaman tentang struktur kontrol sangat penting dalam membangun program yang efisien. 
Dari praktikum ini, dapat disimpulkan sebagai berikut:

1. Memahami berbagai jenis struktur kontrol seperti perulangan (`for`), pengkondisian (`if, else, switch`), serta kontrol alur lainnya (`break, continue, goto`), yang berfungsi untuk mengarahkan jalannya program sesuai kondisi.
2. Mengimplementasikan struktur kontrol dengan benar dalam program untuk menyusun logika yang lebih dinamis dan mampu menangani berbagai skenario program.
3. Mengoptimalkan penggunaan struktur kontrol sehingga dapat menulis kode yang lebih ringkas, efisien, dan mudah dipelihara sesuai dengan kebutuhan logika program.
   
## Daftar Pustaka
[1] A. A. A. Donovan and B. W. Kernighan, *The Go Programming Language*. Boston, MA: Addison-Wesley, 2015.
