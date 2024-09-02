package main

import "fmt"

func main() {
	user := map[string]string{
		"username":"Ilham",
		"email":"ilham@gmail.com",
	}

	fmt.Println(user)

	fmt.Println("===========================================")

	var scores = make(map[string]int)
	fmt.Println(scores)

	scores["java"] = 78
	scores["php"] = 90
	scores["python"] = 97
	fmt.Println("scores:", scores)

	fmt.Println("Nilai java:", scores["java"])
	fmt.Println("Nilai php:", scores["php"])

	fmt.Println("===========================================")

	scores["java"] = 97
	fmt.Println("scores:", scores)

	fmt.Println("===========================================")

	delete(scores, "java")
	fmt.Println("scores:", scores)

	fmt.Println("===========================================")

	names := map[int]string{
		1: "Ilham",
		2: "Maulana",
		3: "Fadhli",
		4: "Daffa",
		5: "Aldi",
	}

	for _, value := range names {
		fmt.Println("value :", value)
	}
}