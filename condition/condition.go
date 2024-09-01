package main

import "fmt"

func main() {
	if true{
		fmt.Println("Kode ini dijalankan")
	}
	fmt.Println("Done")

	fmt.Println("==============================")

	if result := "lulus"; result == "lulus" {
		fmt.Println("Selamat anda", result)
	} else {
		fmt.Println("Maaf anda", result)
	}

	fmt.Println("==============================")

	hour := 6

	if hour > 8 && hour < 17 {
		fmt.Println("Masih dalam rentan waktu yang diperbolehkan")
	} else {
		fmt.Println("Diluar rentan waktu yang diperbolehkan")
	}

	fmt.Println("==============================")

	GPA := 2.2
	var graduationStatus string

	if GPA >= 3.5 && GPA <= 4.0 {
		graduationStatus = "CUMLAUDE"
	} else if GPA >= 3.0 && GPA <= 3.5 {
		graduationStatus = "MEMUASKAN"
	} else if GPA > 2.75 && GPA < 3.0 {
		graduationStatus = "CUKUP MEMUASKAN"
	} else {
		graduationStatus = "KURANG MEMUASKAN"
	}

	fmt.Printf("Selamat kamu lulus dengan predikat %s dengan IPK %.2f\n", graduationStatus, GPA)

	fmt.Println("==============================")

	x := 3
	y := -1

	if x > 0 {
		if y > 0 {
			fmt.Println("x dan y lebih besar dari 0")
		} else {
			fmt.Println("x lebih besar dari 0 dan y kurang dari atau sama dengan 0")
		}
	}

	fmt.Println("==============================")

	var poin = 9
	// switch poin {
	// 	case 8: {
	// 		fmt.Println("bagus")
	// 	}
	// 	case 7, 6, 5: {
	// 		fmt.Println("cukup")
	// 	}
	// 	default: {
	// 		fmt.Println("kurang")
	// 	}
	// }

	switch {
		case poin >= 8: {
			fmt.Println("bagus")
		}
		fallthrough
		case poin >= 6 && poin < 8: {
			fmt.Println("cukup")
		}
		case poin >= 4 && poin < 6: {
			fmt.Println("kurang")
		}
		default: {
			fmt.Println("gagal")
			fmt.Println("belajar lagi")
		}
	}
}