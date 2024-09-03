package main

import "fmt"

type Shape interface {
	getArea() float64
	getParimeter() float64
}

type Rectangele struct {
	width  float64
	length float64
}

func (r Rectangele) getArea() float64 {
	return r.width * r.length
}

func (r Rectangele) getParimeter() float64 {
	return 2 * (r.width + r.length)
}

type Square struct {
	side float64
}

func (s Square) getArea() float64 {
	return s.side * s.side
}

func (s Square) getParimeter() float64 {
	return 4 * s.side
}

// func getAreaOfRectangele(r Rectangele)  {
// 	fmt.Println("Luas:", r.getArea())
// }

// func getAreaOfSquare(s Square)  {
// 	fmt.Println("Luas:", s.getArea())
// }

func getArea(s Shape)  {
	fmt.Println("Luas:", s.getArea())
}

func main() {
	var shape1 Shape = Square{side: 5}
	var shape2 Shape = Rectangele{width: 10, length: 5}
	var shape3 Shape = Rectangele{width: 10, length: 5}

	fmt.Printf("area of shape1: %#v \n", shape1)
	fmt.Printf("area of shape2: %#v \n", shape2)
	fmt.Printf("area of shape3: %#v \n", shape3)

	fmt.Println("=======================================")

	shapes := []Shape{shape1, shape2, shape3}

	for _, shape := range shapes {
		fmt.Printf("%#v memiliki luas %.1f dan keliling %.1f \n", shape, shape.getArea(), shape.getParimeter())
	}

	fmt.Println("=======================================")

	// getAreaOfRectangele(Rectangele{width: 8, length: 5})
	// getAreaOfSquare(Square{side: 9})

	getArea(Rectangele{width: 8, length: 5})
	getArea(Square{side: 9})

	fmt.Println("=======================================")

	var square1 Shape = Square{side: 5}
	fmt.Println("getArea", square1.getArea())
	fmt.Println("getParimeter", square1.getParimeter())

	fmt.Println("=======================================")

	var originalSquare Square = square1.(Square)
	fmt.Println("getArea:", originalSquare.getArea())
	fmt.Println("getParimeter:", originalSquare.getParimeter())
	fmt.Println("side:", originalSquare.side)

	fmt.Println("=======================================")

	var anything any

	anything = "hello"
	fmt.Printf("anything %T: %v \n", anything, anything)

	anything = 18
	fmt.Printf("anything %T: %v \n", anything, anything)

	anything = false
	fmt.Printf("anything %T: %v \n", anything, anything)

	anything = []string{"java", "php", "python"}
	fmt.Printf("anything %T: %v \n", anything, anything)
}