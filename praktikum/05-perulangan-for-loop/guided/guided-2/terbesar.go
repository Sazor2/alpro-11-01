package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var angka, max int
	fmt.Scan(&angka)

	for i := 1; i < n; i++ {
		fmt.Scan(&angka)

		if angka > max {
			max = angka
		}
	}
	fmt.Println(max)
}
