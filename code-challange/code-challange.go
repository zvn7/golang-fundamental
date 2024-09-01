package main

import (
	"fmt"
)

// code challange chapter 2

// func main() {
//  	var a,b string
//  	fmt.Scanln(&a)
//  	fmt.Scanln(&b)

// 	fmt.Println(a, b)
// }

// code challange chapter 3

// func main() {
//     var x int
//     var y float64

// 	x = 2
// 	y = 3.12

// 	result1 := x + int(y)
// 	result2 := float64(x) - y
// 	result3 := x * int(y)

// 	fmt.Printf("%d\n", result1)
// 	fmt.Printf("%.2f\n", result2)
// 	fmt.Printf("%d\n", result3)
// }

// func main() {
// 	var name string
// 	var address string
// 	var email string

// 	scanner := bufio.NewScanner(os.Stdin)
// 	fmt.Print("Enter your name: ")
// 	scanner.Scan()
// 	name = scanner.Text()
// 	fmt.Print("Enter your address: ")
// 	scanner.Scan()
// 	address = scanner.Text()
// 	fmt.Print("Enter your email: ")
// 	scanner.Scan()
// 	email = scanner.Text()

// 	fmt.Printf("Halo, saya %s. Saya tinggal di %s. Alamat email saya adalah %s", name, address, email)
// }

// func main() {
// 	var jam = 5

// 	if jam >= 4 && jam <= 5 {
// 		fmt.Println("Bangun pagi")
// 	} else if jam >= 6 && jam <= 7 {
// 		fmt.Println("Mandi dan Sarapan")
// 	} else if jam >=8 && jam <= 11 {
// 		fmt.Println("Berangkat Sekolah")
// 	} else if jam == 12 {
// 		fmt.Println("Pulang Sekolah")
// 	} else if jam >= 13 && jam <= 15 {
// 		fmt.Println("Tidur Siang")
// 	} else if jam >= 16 && jam <= 17 {
// 		fmt.Println("Waktu Main")
// 	} else {
// 		fmt.Println("Waktu Belajar dan Istirahat")
// 	}
// }

// code challange chapter 7

func main() {
	var input int

	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&input)

	for i := input; i > 0; i-- {
		fmt.Printf("%d I will become a go developer\n", i)
	}
}