package main

import "fmt"

func main() {
	var celsius float64
	fmt.Scan(&celsius)

	reamur := celsius * 4.0 / 5.0
	fahrenheit := (celsius * 9.0 / 5.0) + 32.0
	kelvin := celsius + 273.15

	fmt.Println(reamur, fahrenheit, kelvin)
}
