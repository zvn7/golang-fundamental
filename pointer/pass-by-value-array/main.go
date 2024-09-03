package main

import "fmt"

func main() {
	var numbers = [4]int{3, 4, 1, 5}
	var anotherNubers = numbers

	fmt.Println("sebelum========================================")

	fmt.Println("numbers:", numbers)
	fmt.Println("anotherNubers:", anotherNubers)

	fmt.Println("sesudah========================================")

	numbers[1] = 10

	fmt.Println("numbers:", numbers)
	fmt.Println("anotherNubers:", anotherNubers)

	fmt.Println("==============================================")

	var scores = [4]int{2, 67, 99, 30}
	multiplyBy10(scores)

	fmt.Println("scores:", scores)
}

func multiplyBy10(n [4]int)  {
	for i := range n {
		n[i] = n[i] * 10
	}
	fmt.Println("n:", n)
}