package main

import "fmt"

func main() {
	// fmt.Println("Hasil penjumlahan:", add(4, 7))

	total := add(8, 9)
	fmt.Println("Hasil penjumlahan:", total)

	total2 := multiply(7, 9)
	fmt.Println("Hasil perkalian:", total2)

	calculation := add(8, multiply(8, 6))
	fmt.Println("Hasil kalkulasi:", calculation)
}

func add(a int, b int) int {
	result := a + b
	return result
}

func multiply(a, b int) int  {
	result := a * b
	return result
}