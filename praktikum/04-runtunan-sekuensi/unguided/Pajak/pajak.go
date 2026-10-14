package main

import "fmt"

func main() {
	var penghasilan float64
	var pajak float64

	fmt.Print("Penghasilan dalam juta: ")
	fmt.Scan(&penghasilan)

	if penghasilan <= 50 && penghasilan >= 0 {
		pajak = 0.05 * float64(penghasilan)
	} else if penghasilan > 50 && penghasilan < 100 {
		pajak = (0.05 * 50) + (0.10 * float64(penghasilan-50))
	} else if penghasilan >= 100 && penghasilan < 200 {
		pajak = (0.05 * 50) + (0.10 * 50) + (0.15 * float64(penghasilan-100))
	} else if penghasilan >= 200 {
		pajak = (0.05 * 50) + (0.10 * 50) + (0.15 * 100) + (0.20 * float64(penghasilan-200))
	}

	fmt.Printf("Total pajak yang harus dibayarkan: %.2f juta", pajak)
}
