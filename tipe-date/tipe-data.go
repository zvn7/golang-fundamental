package main

import "fmt"

func main() {
	fmt.Println("Tipe Data")
	fmt.Println("Numerik")

	var positveNumber uint8 = 90
	var negativeNumber int = -90
	fmt.Printf("Bilangan positif : %d\n", positveNumber)
	fmt.Printf("Bilangan negatif : %d\n", negativeNumber)

	var decimalNumber = 2.5
	fmt.Printf("Bilangan pecahan: %f\n", decimalNumber)
	fmt.Printf("Bilangan pecahan: %.2f\n", decimalNumber)

	fmt.Println("=====================")
	fmt.Println("Boolean")

	var exist = true
	fmt.Printf("exist ? %t\n", exist)

	fmt.Println("=====================")
	fmt.Println("String")

	var message string = "Hello World"
	fmt.Printf("Message : %s\n", message)

	var otherMessage = `Nama saya "Ilham",
Salam kenal,
Mari belajar di "EnigmaCamp"`
	fmt.Println(otherMessage)

	fmt.Println("Ilham" + " " + "Maulana")
	fmt.Println(len(message))
}
