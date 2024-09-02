package main

import "fmt"

func main() {
	num := []int {7, 9, 1, 8, 3}
	kecil, _ := minMax(num)

	fmt.Println("kecil:", kecil)
	// fmt.Println("besar:", besar)
}

func minMax(numbers []int) (min int, max int)  {
	min = numbers[0]
	max = numbers[0]

	for _, n := range numbers{
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	return
}