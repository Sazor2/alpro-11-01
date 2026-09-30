package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	var sepuluhribu int = x / 10000
	var sisa int = x % 10000

	var limaribu int = sisa / 5000
	sisa = sisa % 5000

	var duaribu int = sisa / 2000
	sisa = sisa % 2000

	fmt.Println(sepuluhribu, limaribu, duaribu)
}
