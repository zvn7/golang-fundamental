package main

import "fmt"

func main() {
	var fruits = [4]string {"Apple", "Orange", "Banana", "Grape"}

	fmt.Println(fruits[2])
	fmt.Println(fruits)
	fruits[2] = "Cherry"
	fmt.Println(fruits[2])
	fmt.Println(fruits)

	fmt.Println("==============================================")

	var scores [3]int
	scores[0] = 2
	scores[1] = 67
	scores[2] = 99

	fmt.Println(scores)

	fmt.Println("==============================================")

	var days = [...]string {"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(days)
	fmt.Println("Length:", len(days))

	fmt.Println("==============================================")

	for i := 0; i < len(days); i++ {
		fmt.Printf("Elemen ke-%d: %s\n", i, days[i])
	}

	fmt.Println("==============================================")

	for i, day := range days {
		fmt.Printf("Elemen ke-%d: %s\n", i, day)
	}

	fmt.Println("==============================================")

	for _, day := range days {
		fmt.Printf("Nama hari: %s\n", day)
	}

	fmt.Println("==============================================")

	for i := range days{
		fmt.Printf("Index ke-%d\n", i)
	}

	fmt.Println("==============================================")

	var numbers = [6]int {7, 3, 10, 19, 44, 9}

	for _, number := range numbers {
		if number % 2 == 0 {
			fmt.Println(number)
		}
	}

	fmt.Println("==============================================")

	fmt.Println("sebelum :", numbers)
	for i := 0; i < len(numbers); i++ {
		numbers[i] *= 2
	}
	fmt.Println("sesudah :", numbers)

	fmt.Println("==============================================")

	var matrix = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(matrix)
	fmt.Println(matrix[1][2])
}