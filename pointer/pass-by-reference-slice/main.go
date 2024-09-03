package main

import "fmt"

func main() {
	var numbers = []int{3, 4, 1, 5}
	var anotherNubers = numbers

	fmt.Println("numbers:", numbers)
	fmt.Println("anotherNubers:", anotherNubers)

	numbers[1] = 10

	fmt.Println("numbers:", numbers)
	fmt.Println("anotherNubers:", anotherNubers)

	fmt.Println("==============================================")

	var scores = []int{2, 67, 99, 30}

	multiplyBy10(scores)
	fmt.Println("scores:", scores)
}

func multiplyBy10(numbers []int)  {
	for i := range numbers {
		numbers[i] = numbers[i] * 10
	}
	fmt.Println("numbers di function:", numbers)
}