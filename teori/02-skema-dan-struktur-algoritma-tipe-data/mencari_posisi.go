package main

import "fmt"

func main() {
	var posisi, posisi0, kecepatan, waktu int

	fmt.Scan(&posisi0)
	fmt.Scan(&kecepatan)
	fmt.Scan(&waktu)

	posisi = posisi0 + kecepatan*waktu
	fmt.Println("Posisi atau jarak adalah:", posisi)
}
