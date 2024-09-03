package main

import (
	"fmt"
	"golang-introduction/access-modifier/library"
)

func main() {
	fmt.Println("Hour In A Day:", library.HourInADay)

	fmt.Println(library.DaystoMinutes(5))
}

