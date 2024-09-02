package main

import "fmt"

var name = "Ilham"

func main() {
	helloWorld()
	fmt.Println("main :", name)

	greeting("Aam", 18)
	add(7, 9)
}

func helloWorld()  {
	var name = "Maulana"
	fmt.Println("Hello World!")
	fmt.Println("Hello world :", name)
}

func greeting(name string, age int)  {
	fmt.Printf("Hello my name is %s and I'm %d years old\n", name, age)
}

func add(a int, b int)  {
	fmt.Println("Result add :", a + b)
}