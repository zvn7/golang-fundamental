package main

import "fmt"

func main() {
	numbersSlice := []int{3, 4, 1, 5, 6, 7, 8, 9}
	fmt.Println(numbersSlice)

	fmt.Println("Panjang slice:", len(numbersSlice))
	fmt.Println("Kapasitas slice:", cap(numbersSlice))

	fmt.Println("=============================================")

	scores := make([]int, 3, 7)
	scores[0] = 28
	scores[1] = 29
	scores[2] = 30
	fmt.Println(scores)

	fmt.Println("Panjang scores:", len(scores))
	fmt.Println("Kapasitas scores:", cap(scores))

	fmt.Println("=============================================")

	animals := make([]string, 3)
	animals[0] = "Elephant"
	animals[1] = "Lion"
	animals[2] = "Tiger"
	fmt.Println(animals)

	fmt.Println("Panjang animals:", len(animals))
	fmt.Println("Kapasitas animals:", cap(animals))

	fmt.Println("=============================================")

	heroes := [4]string{"Superman", "Batman", "Spiderman", "Flash"}
	fmt.Println("Heroes:", heroes)

	villain := make([]string, 3, 5)
	villain[0] = "Joker"
	villain[1] = "Harley"
	villain[2] = "Hulk"

	fmt.Println("Villain:", villain)
	fmt.Println("Panjang villain:", len(villain))
	fmt.Println("Kapasitas villain:", cap(villain))

	villain = append(villain, "Cyclops")
	villain = append(villain, "Venom", "Carnage")

	fmt.Println("Villain:", villain)
	fmt.Println("Panjang villain:", len(villain))
	fmt.Println("Kapasitas villain:", cap(villain))

	fmt.Println("=============================================")

	var numbers = [4]int{1, 2, 3, 4}
	var anotherNumbers = numbers

	fmt.Println("numbers:", numbers)
	fmt.Println("anotherNumbers:", anotherNumbers)

	anotherNumbers[1] = 6

	fmt.Println("numbers:", numbers)
	fmt.Println("anotherNumbers:", anotherNumbers)

	fmt.Println("=============================================")

	var prices = []int{2000, 3000, 4000, 5000}
	var anotherPrices = prices

	fmt.Println("prices:", prices)
	fmt.Println("anotehr prices:", anotherPrices)

	anotherPrices[1] = 6000

	fmt.Println("prices:", prices)
	fmt.Println("anotehr prices:", anotherPrices)

	fmt.Println("=============================================")

	ages := [4]int{10, 20, 30, 40}
	slicesAges := ages[0:3]

	fmt.Println("ages:", ages)
	fmt.Println("slices ages:", slicesAges)

	slicesAges[2] = 70

	fmt.Println("ages:", ages)
	fmt.Println("slices ages:", slicesAges)

	slicesAges = append(slicesAges, 88)

	fmt.Println("ages:", ages)
	fmt.Println("slices ages:", slicesAges)
}