package main

import "fmt"

func main() {
	var firstName string = "Ilham"
	var lastName string = "Maulana"

	fmt.Println(firstName, lastName)
	fmt.Printf("Halo, Ilham Maulana\n")
	fmt.Printf("Halo, %s %s \n", firstName, lastName)

	var (
		age int
		address string
	)

	age = 18
	address = "Jakarta"

	fmt.Printf("Halo, saya %s %s berumur %d, dan saya tinggal di %s\n", firstName, lastName, age, address)

	var (bootcampName, bootcampAddress = "EnigmaCamp", "Jakarta Selatan")
	fmt.Println(bootcampName, bootcampAddress)

	day := "Monday"
	date := "24"
	month := "August"

	fmt.Println(day, date, month+" 2024")

	var number = 22
	number = 18

	const phi = 3.14

	fmt.Println(number, phi)

	bootcamp, _ := "EnigmaCamp", "Aktif7"
	fmt.Println(bootcamp)
}
