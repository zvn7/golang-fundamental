package main

import "fmt"

type kendaraan struct{
	merek string
	tahun int
	model string
	harga int
	mesin
}

type mesin struct{
	tipe string
	cc int
}

func main() {
	/*
	var a kendaraan

	a.merek = "Toyota"
	a.tahun = 2022
	a.model = "Avanza"
	a.harga = 10000000

	fmt.Println("a:", a)

	fmt.Println("merek:", a.merek)
	fmt.Println("tahun:", a.tahun)
	fmt.Println("model:", a.model)
	fmt.Println("harga:", a.harga)

	fmt.Println("=======================================")

	var b = kendaraan{}

	b.merek = "Wuling"
	b.tahun = 2021
	b.model = "X1"
	b.harga = 20000000

	fmt.Println("b:", b)

	var c = kendaraan{"Honda", 2020, "CRV", 30000000}

	fmt.Println("c:", c)

	fmt.Println("=======================================")

	var d = kendaraan{tahun: 2023, harga: 50000000, merek: "Suzuki", model: "SX4"}

	fmt.Println("d:", d)

	fmt.Println("=======================================")

	var x = kendaraan{merek: "Mercedes", model: "Sedan", tahun: 2022, harga: 50000000}
	var y kendaraan = x

	fmt.Println("x:", x)
	fmt.Println("y:", y)

	fmt.Printf("alamat x: %p, alamat y: %p\n", &x, &y)

	y.model = "SUV"

	fmt.Println("x:", x)
	fmt.Println("y:", y)

	fmt.Println("=======================================")

	// var z = kendaraan{merek: "Honda", model: "Civic", tahun: 2021, harga: 30000000}

	// updateKendaraan(z)
	// fmt.Println("z:", z)

	// fmt.Println("=======================================")

	// var k *kendaraan = &z
	// fmt.Printf("alamat z: %p, alamat k: %p\n", &z, k)

	// k.model = "Accord"
	// fmt.Println("z:", z)
	// fmt.Println("k:", k)

	fmt.Println("=======================================")

	var z = kendaraan{merek: "Honda", model: "Civic", tahun: 2021, harga: 30000000}

	updateKendaraan(&z)
	fmt.Println("z:", z)

	fmt.Println("=======================================")

	var k = new(kendaraan)
	fmt.Printf("alamat k: %p\n", k)
	*/

	var a = kendaraan{
		merek: "Toyota",
		tahun: 2022,
		model: "Avanza",
		harga: 10000000,
		mesin: mesin{
			tipe: "matic",
			cc: 1.0,
		},
	}

	fmt.Println("a:", a)
	fmt.Println("tipe mesin a:", a.mesin.tipe)
}

func updateKendaraan(newKendaraan *kendaraan)  {
	newKendaraan.merek = "Toyota"
	newKendaraan.tahun = 2022
	newKendaraan.model = "Avanza"
	newKendaraan.harga = 10000000

	fmt.Println("update kendaraan:", newKendaraan)
}