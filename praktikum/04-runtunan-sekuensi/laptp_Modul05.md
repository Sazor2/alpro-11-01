# <h1 align="center">Tugas Pendahuluan Modul 05 - Perulangan For</h1>
<p align="center">Nabil Aqbar Kurniawijaya Putra - 109092600021</p>

### 1. Fungsi Utama Perulangan `for`
`for` merupakan _looping construct_ satu-satunya dalam Go, karena Go tidak memiliki `while` dan `do-while`. Fungsi utama dari `for` yaitu untuk mengeksekusi kode didalamnya secara berulang selama kondisi-kondisinya terpenuhi. 
Tiga komponen `for` diantaranya:
1. Inisialisasi, atau _init_, program didalamnya akan dieksekusi sekali diawal sebelum perulangan dimulai. Bagian ini biasanya berisi nilai awal dari variabel pembanding, seperti `i := 0`.
2. Kondisi, atau _condition_, program didalam kondisi akan dievaluasi sebelum setiap iterasi dijalankan. Jika bernilai `true`, maka kode didalam `for` akan dieksekusi, jika `false` maka perulangan akan berhenti.
3. _Post-statement_, program didalam _post-statement_ akan dieksekusi pada setiap iterasi berakhir setelah kode didalam kondisi selesai dieksekusi. Biasa digunakan untuk mengubah nilai variabel pembanding, seperti `i++`.

### 2. Perbedaan `for` Dengan Tiga Komponen dan `for` Yang Hanya Menggunakan Kondisi
- `for` yang menggunakan tiga komponen digunakan ketika jumlah iterasi sudah diketahui sejak awal, jika jumlah iterasi belum diketahui, maka `for` ini tidak bisa digunakan. 
- `for` yang menggunakan kondisi digunakan ketika jumlah iterasi belum diketahui atau belum pasti. `for` ini bergantung pada kondisi yang dievaluasi saat program berjalan.

### 3. ganjil.go

```go
package main
import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	total := 0

	for i := 1; i <= x; i++{
		if i % 2 != 0 {
			fmt.Println(i)
			total += i
		}
	}
	fmt.Println(total)
}
```

##### Output
![Screenshot Output Unguided](./tp/ganjil/pendahuluan_ganjil_output.png)


#### Deskripsi
Program untuk menampilkan semua bilangan ganjil dari 1 sampai ke-n. Program ini menggunakan pengulangan `for` untuk menampilkan semua bilangan ganjil.

### 4. perkalian.go

```go
package main
import "fmt"

func main() {
	var a, b int
	fmt.Print("Bilangan ")
	fmt.Scan(&a)
	fmt.Print("Dengan batas perkalian ")
	fmt.Scan(&b)

	fmt.Println("Hasil: ")
	for i := 1; i <= b; i++ {
		fmt.Printf("%d x %d = %d\n", a, i, a * i)
	}
}
```
#### Output
![Screenshot Output Unguided](./tp/perkalian/pendahuluan_perkalian_output.png)


#### Deskripsi
Program untuk menampilkan tabel perkalian dari sebuah bilangan $a$, mulai dari bilangan $a x 1$ hingga $a x b$, yang dimana $a$ adalah adalah bilangan yang dikalikan, dan $b$ adalah batas perkalian. Program ini menggunakan pengulangan `for` untuk menampilkan tabel perkalian.

## Kesimpulan
Kesimpulan dari tugas pendahuluan yaitu pengenalan tentang perulangan dengan menggunakan `for`. Fungsi utama dari `for` yaitu untuk mengeksekusi kode yang ada didalamnya selama kondisi-kondisinya terpenuhi.
`for` memiliki tiga komponen, yaitu inisialisasi, kondisi, dan _post-statement_. 
Selain itu juga cara pengaplikasian `for` dalam sebuah program.