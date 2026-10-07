package main

import "fmt"

func main() {
	var n, i, bilangan, jumlah, d1, d4 int
	fmt.Scan(&n)
	jumlah = 0
	for i = 1; i <= n; i++ {
		fmt.Scan(&bilangan)
		d1 = bilangan / 1000
		d4 = bilangan % 10
		jumlah += d1 + d4
	}
	fmt.Println(jumlah)
}
