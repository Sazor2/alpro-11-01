package main

import "fmt"

func main() {
	var t int
	fmt.Print("Input t (jumlah transaksi): ")
	fmt.Scan(&t)
	totalPendapatan := 0

	for i := 1; i <= t; i++ {
		var jenisFilm, hari string
		var jumlahTiket int

		fmt.Printf("Input transaksi ke-%d (jenis, hari, tiket): ", i)
		fmt.Scan(&jenisFilm, &hari, &jumlahTiket)

		hargaSatuan := 0

		if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
			if jenisFilm == "Reguler" {
				hargaSatuan = 40000
			} else if jenisFilm == "IMAX" {
				hargaSatuan = 80000
			} else if jenisFilm == "3D" {
				hargaSatuan = 60000
			}
		} else if hari == "Jumat" || hari == "Sabtu" || hari == "Minggu" {
			if jenisFilm == "Reguler" {
				hargaSatuan = 50000
			} else if jenisFilm == "IMAX" {
				hargaSatuan = 90000
			} else if jenisFilm == "3D" {
				hargaSatuan = 70000
			}
		}
		totalHarga := hargaSatuan * jumlahTiket

		if jumlahTiket >= 5 {
			diskon := (totalHarga * 10) / 100
			totalHarga = totalHarga - diskon
		}
		fmt.Printf("Output total harga transaksi ke-%d: %d\n", i, totalHarga)

		totalPendapatan = totalPendapatan + totalHarga
	}
	fmt.Println("Output total pendapatan: ", totalPendapatan)
}
