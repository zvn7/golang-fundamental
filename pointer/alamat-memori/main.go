package main

import "fmt"

func main() {
	var name string = "Ilham"
	var age int = 18

	fmt.Println("Alamat memori name:", &name)
	fmt.Println("Alamar memori age:", &age)
}