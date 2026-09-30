package main

import "fmt"

func main() {
	var a, b int

	//Membaca Input
	fmt.Scan(&a)
	fmt.Scan(&b)

	//Menukar nilai
	a, b = b, a

	//Menampilkan output
	fmt.Println(a, b)
}
