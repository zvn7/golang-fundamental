package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var firstName string
	var lastName string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	birthYear := (2024 - age)

	fmt.Println("Your name is", firstName, lastName, "You were born in", birthYear)

	fmt.Println("==================================================================")

	var fullName string
	var address string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Data diri")
	fmt.Print("Enter your full name: ")
	scanner.Scan()
	fullName = scanner.Text()
	fmt.Print("Enter your age: ")
	scanner.Scan()
	age, _ = strconv.Atoi(scanner.Text())
	fmt.Print("Enter your address: ")
	scanner.Scan()
	address = scanner.Text()

	fmt.Printf("%s | %d | %s", fullName, age, address)

}