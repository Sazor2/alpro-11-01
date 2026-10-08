package main

import "fmt"

func main() {
	var tahun int
	var bulan string

	fmt.Scan(&tahun, &bulan)

	// Mengecek apakah tahun merupakan tahun kabisat
	isKabisat := (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)

	switch bulan {
	case "Jan", "Mar", "Mei", "Jul", "Agu", "Okt", "Des":
		fmt.Println(31)
	case "Apr", "Jun", "Sep", "Nov":
		fmt.Println(30)
	case "Feb":
		if isKabisat {
			fmt.Println(29)
		} else {
			fmt.Println(28)
		}
	default:
		fmt.Println("-")
	}
}
