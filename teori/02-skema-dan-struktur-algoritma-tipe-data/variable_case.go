package main

import "fmt"

func main() {
	var name string
	var age int
	var address string

	name = "Nabil Aqbar Kurniawijaya Putra"
	fmt.Println("Nama :", name)

	age = 21
	fmt.Println("Usia :", age)

	address = "Jl. Raya Cibaduyut No. 123, Bandung"
	fmt.Println("Alamat :", address)

	var middlename = "Aqbar"
	fmt.Println("Nama Tengah :", middlename)

	lastname := "Putra"
	fmt.Println("Nama Belakang :", lastname)

	var (
		fullname  = "Nabil Aqbar Kurniawijaya Putra"
		firstname = "Nabil"
	)

	fmt.Println(fullname)
	fmt.Println(firstname)

}
