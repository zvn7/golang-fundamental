package main

import "fmt"

func main() {
	var x int = 4
	var y *int = &x

	fmt.Println("x:", x)
	fmt.Println("alamat memori x:", &x)
	fmt.Println("y:", y)
	fmt.Println("alamat memori y:", &y)

	fmt.Println("nilai reference dari pointer y:", *y)

	*y = *y + 1

	fmt.Println("x:", x)
	fmt.Println("*y:", *y)

	fmt.Println("=======================================")

	var a int = 8
	
	increase(&a)
	fmt.Println("a:", a)
}

func increase(n *int)  {
	*n = *n + 1
	fmt.Println("n di function increase:", *n)
}