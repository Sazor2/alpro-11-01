package main

import "fmt"

func main() {
	pi := 3.14
	var r, luas float64

	fmt.Scan(&r)

	luas = pi * r * r

	fmt.Println(luas)
}
