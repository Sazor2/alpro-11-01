package main

import "fmt"

func main() {
	var nama string
	var skorMatematika, skorBahasaIggris int

	//Membaca Input
	fmt.Scan(&nama)
	fmt.Scan(&skorMatematika)
	fmt.Scan(&skorBahasaIggris)

	//Menghitung total & rata-rata (pembagian bilangan bulat)
	total := skorMatematika + skorBahasaIggris
	rataRata := total / 2

	//Menampilkan output
	fmt.Println(nama)
	fmt.Println(total)
	fmt.Println(rataRata)
}
