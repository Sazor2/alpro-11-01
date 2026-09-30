package main

import "fmt"

func main() {
	var uang int
	fmt.Scan(&uang)

	sepuluhRibu := uang / 10000
	sisa := uang % 10000

	limaRibu := sisa / 5000
	sisa = sisa % 5000

	seribu := sisa / 1000

	fmt.Println(sepuluhRibu, limaRibu, seribu)
}
