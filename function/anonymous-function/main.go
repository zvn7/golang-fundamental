package main

import "fmt"

func main() {

	// anonymous function
	func()  {
		fmt.Println("Hello World!")
	}()

	// anonymous function dalam variable
	halo := func()  {
		fmt.Println("Hello Dunia!")
	}

	halo()

	// passing argument ke anonymous function
	func(word string)  {
		fmt.Println(word)
	}("Hello EnigmaCamp!")

	// passing argument dalam variable function
	hello := func(name string)  {
		fmt.Println("Hello", name)
	}

	hello("Ilham")

	//passing anonymous function sebagai argument
	greetEnglish := func(name string) string {
		return "Hello " + name
	}

	greetRussian := func(name string) string {
		return "Привет " + name
	}

	greetKorean := func(name string) string {
		return "Annyeonghaseyo " + name
	}

	greet("Ilham", greetEnglish)
	greet("Ilham", greetRussian)
	greet("Ilham", greetKorean)

	add := func (num1, num2 int) int {
		return num1 + num2
	}

	multiply := func (num1, num2 int) int {
		return num1 * num2
	}

	calculate(2,6, add)
	calculate(2,6, multiply)
}

func greet(name string, f func(name string) string){
	fmt.Println(f(name))
}

func calculate(a, b int, operator func(z, y int) int){
	fmt.Println(operator(a, b))
}