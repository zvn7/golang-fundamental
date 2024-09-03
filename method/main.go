package main

import "fmt"

type triangle struct{
	base float64
	height float64
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func (t *triangle) increaseSize() {
	t.base += 1
	t.height += 1
}

func main() {
	instanceTriangle := triangle{10, 12}
	area := instanceTriangle.area()

	fmt.Println(area)

	fmt.Println("=======================================")

	fmt.Println("instanceTriangle:", instanceTriangle)

	instanceTriangle.increaseSize()

	fmt.Println("instanceTriangle:", instanceTriangle)
}