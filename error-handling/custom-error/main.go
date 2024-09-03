package main

import "fmt"

type Item struct {
	id    int
	name  string
	price int
}

type ItemNotFoundError struct {
	Id int
}

func (i *ItemNotFoundError) Error() string {
	return fmt.Sprintf("Item dengan Id: %d tidak ditemukan", i.Id)
}

var items = []Item{
	{id: 1, name: "Pencil", price: 10},
	{id: 2, name: "Book", price: 20},
	{id: 3, name: "Notebook", price: 30},
	{id: 4, name: "Paper", price: 40},
}

func getItemById(id int) (Item, error) {
	for _, item := range items{
		if item.id == id {
			return item, nil
		}
	}
	return Item{}, &ItemNotFoundError{Id: id}
}

func main() {
	result, err := getItemById(5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Hasil:", result)
}