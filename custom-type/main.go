package main

import "fmt"

type Patient struct {
	Name string
	Age string
	Celsius
}

type Celsius float64

func (c Celsius) toFahrenheit() float64 {
	return float64(c * 9 / 5 + 32)
}

func (c *Celsius) add(value float64) {
	*c += Celsius(value)
}

func main() {
	var temperature Celsius = 36.5
	fmt.Println(temperature)

	fmt.Println("Suhu diruangan ini", temperature.toFahrenheit(), "derajat Fahrenheit")

	temperature.add(3)
	fmt.Println("suhu diruangan ini menjadi", temperature, "derajat Celsius")

	fmt.Println("=======================================")

	newPatient := Patient{Name: "Aam", Age: "20", Celsius: 36.5}

	fmt.Printf("New Patient : %+v\n", newPatient)
}