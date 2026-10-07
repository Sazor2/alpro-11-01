Markdown# <h1 align="center">Laporan Praktikum Modul 3 - Tipe Data dan Instruksi Dasar</h1>
<p align="center">Nabil Aqbar Kurnia Wijaya Putra - [Isi NIM kamu di sini]</p>

# A. Tipe Data Dasar pada Go
Bahasa Go menyediakan beberapa tipe data dasar untuk menyimpan nilai. Tipe data int digunakan untuk menyimpan bilangan bulat (seperti jumlah barang, hari, atau nominal uang), sedangkan float64 digunakan untuk menampung bilangan real/desimal dengan tingkat presisi tinggi (seperti hasil konversi suhu atau jarak). Pemilihan tipe data yang tepat sangat krusial agar tidak terjadi kesalahan operasi aritmatika.

# B. Operator Aritmetika dan Pembagian Integer
Operator aritmetika standar pada Go meliputi penjumlahan (+), pengurangan (-), perkalian (*), pembagian (/), dan sisa hasil bagi atau modulus (%). Ketika menggunakan pembagian (/) dengan dua operand bertipe integer, Go secara otomatis melakukan pembagian bilangan bulat yang membuang sisa desimalnya. Sebaliknya, jika pembagian melibatkan tipe data float64 atau konstanta desimal (seperti 4.0 / 5.0), Go akan memprosesnya sebagai pembagian real.

# C. Input dan Output Standard
Proses membaca masukan dilakukan menggunakan fungsi fmt.Scan atau fmt.Scanln dari package fmt, yang menerima alamat pointer variabel (&variabel). Untuk mencetak keluaran, fmt.Println digunakan untuk mencetak dengan baris baru, sedangkan fmt.Printf digunakan saat format keluaran memerlukan aturan khusus, seperti pembatasan angka di belakang koma (%.1f).

# Guided
## 1. kasir.go
```Go
package main

import "fmt"

func main() {
	var totalBelanja, diskon, totalBayar float64

	fmt.Print("Masukkan total belanja: ")
	fmt.Scan(&totalBelanja)

	fmt.Print("Masukkan diskon (%): ")
	fmt.Scan(&diskon)

	totalBayar = totalBelanja - (totalBelanja * (diskon / 100))

	fmt.Println("Total yang harus dibayar:", totalBayar)
}
```
### Deskripsi
Program kasir.go berfungsi untuk mengkalkulasikan total pembayaran setelah dipotong diskon. Program menerima masukan berupa total belanja awal dan persentase diskon bertipe float64, kemudian menghitung nilai potongan dan menampilkan jumlah akhir yang harus dibayar.

Contoh masukan: 100000 10

Contoh keluaran:

Plaintext
Total yang harus dibayar: 90000
# 2. konversi.go
``` Go
package main

import "fmt"

func main() {
	var celcius float64

	fmt.Print("Masukkan suhu Celcius: ")
	fmt.Scan(&celcius)

	reamur := celcius * 4 / 5
	fahrenheit := (celcius * 9 / 5) + 32
	kelvin := celcius + 273.15

	fmt.Println("Reamur:", reamur)
	fmt.Println("Fahrenheit:", fahrenheit)
	fmt.Println("Kelvin:", kelvin)
}
```
# Deskripsi
Program konversi.go menerima masukan suhu dalam satuan Celsius bertipe float64, lalu mengonversinya ke dalam tiga skala suhu lain yaitu Reamur, Fahrenheit, dan Kelvin menggunakan rumus konversi aritmatika dasar.

Contoh masukan: 100

Contoh keluaran:

Plaintext
Reamur: 80
Fahrenheit: 212
Kelvin: 373.15
## 3. tukar.go
```Go
package main

import "fmt"

func main() {
	var a, b, temp int

	fmt.Print("Masukkan nilai A dan B: ")
	fmt.Scan(&a, &b)

	temp = a
	a = b
	b = temp

	fmt.Println("Setelah ditukar -> A:", a, ", B:", b)
}
```
# Deskripsi
Program tukar.go bertujuan untuk menukarkan nilai dari dua variabel bulat (a dan b) dengan memanfaatkan satu variabel penampung sementara (temp). Hasil penukaran kemudian ditampilkan kembali ke terminal.

Contoh masukan: 3 7

Contoh keluaran:

Plaintext
Setelah ditukar -> A: 7 , B: 3
### Unguided
# 1. konversi_suhu_derajat.go
```Go
package main

import "fmt"

func main() {
	var celsius float64
	fmt.Scan(&celsius)

	reamur := (4.0 / 5.0) * celsius

	fmt.Println(reamur)
}
```
# Deskripsi
Program konversi_suhu_derajat.go membaca nilai suhu Celsius bertipe float64 dan menghitung konversinya ke skala Reamur menggunakan rumus R = (4.0 / 5.0) * C[cite: 8]. Penulisan operand 4.0 / 5.0 dengan presisi desimal memastikan bahwa sistem melakukan pembagian real, bukan pembagian integer yang dapat menghasilkan nilai 0[cite: 8].

Contoh masukan: 100

Contoh keluaran:

Plaintext
80
# 2. konversi_hari.go
```Go
package main

import "fmt"

func main() {
	var totalHari int
	fmt.Scan(&totalHari)

	// 1 tahun = 360 hari (12 bulan x 30 hari)
	tahun := totalHari / 360
	sisa := totalHari % 360

	// 1 bulan = 30 hari
	bulan := sisa / 30
	sisa = sisa % 30

	// 1 minggu = 7 hari
	minggu := sisa / 7
	sisaHari := sisa % 7

	fmt.Println(tahun)
	fmt.Println(bulan)
	fmt.Println(minggu)
	fmt.Println(sisaHari)
}
```
![Output program konversi2](unguided/konversi2/output.png)
# Deskripsi
Program konversi_hari.go mengonversi masukan berupa jumlah hari (integer) menjadi rincian jumlah tahun, bulan, minggu, dan sisa hari[cite: 9]. Perhitungan dilakukan berurutan menggunakan operator pembagian integer (/) untuk mendapatkan kuantitas tiap satuan waktu dan operator modulus (%) untuk memperoleh sisa hari yang akan diproses ke tingkat berikutnya[cite: 9].

Contoh masukan: 400

Contoh keluaran:

Plaintext
1
1
1
3
Kesimpulan
Melalui praktikum Modul 03 ini, konsep dasar deklarasi variabel, pemilihan tipe data (int dan float64), serta penggunaan operator aritmatika pada bahasa Go berhasil dipahami dan diimplementasikan[cite: 8, 9]. Penggunaan tipe data yang tepat sangat mempengaruhi hasil perhitungan, terutama pada pembagian integer vs pembagian real[cite: 8]. Selain itu, penggunaan fungsi fmt.Scan dan fmt.Println mempermudah penanganan masukan dan keluaran standar dalam pembuatan program konversi maupun manipulasi data[cite: 9].

Referensi
The Go Authors. (n.d.). A Tour of Go: Basics. Diakses pada 30 September 2026 melalui https://go.dev/tour/basics

The Go Authors. (n.d.). Package fmt. Diakses pada 30 September 2026 melalui https://pkg.go.dev/fmt