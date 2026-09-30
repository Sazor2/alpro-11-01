package main

import "fmt"

func main() {
	var y, x int
	fmt.Scan(&y, &x)

	sisaKue := y % x

	fmt.Println(sisaKue)
}
