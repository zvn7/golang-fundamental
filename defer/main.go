package main

import "fmt"

func main() {
	fmt.Println("start")
	defer loop()
	fmt.Println("done")
}

func add(num1, num2 int)  {
	result := num1 + num2
	fmt.Println(result)
}

func multiply(num1, num2 int) int {
	result := num1 * num2
	fmt.Println(result)
	return result
}

func loop()  {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done loop")
}