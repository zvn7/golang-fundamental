package main

import "fmt"

func main() {
	var sum func(int, int) int = add

	fmt.Println("Hasil penjumlahan:", sum(4, 7))
}

func add(a,b int) int {
	return a + b
}