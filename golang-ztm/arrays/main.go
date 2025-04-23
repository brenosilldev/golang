package main

import (
	"fmt"
)

// type Room struct {
// 	name    string
// 	cleaned bool
// }

// func checkCleanlniess(rommms [4]Room) {

// 	for i := 0; i < len(rommms); i++ {
// 		j := rommms[i]

// 		if j.cleaned {
// 			fmt.Println("Room", j.name, "is cleaned")
// 		} else {
// 			fmt.Println("Room", j.name, "is not cleaned")
// 		}
// 	}
// }

// func main() {
// 	/*

// var myArray [5]int
// fmt.Println(myArray)
// myArra1 := [5]int{1, 2, 3, 4, 5}
// fmt.Println(myArra1)

// myArra2 := [...]int{1, 2, 3, 5}
// fmt.Println(myArra2)

// myArra3 := [5]int{1, 2, 3, 4}
// fmt.Println(myArra3)

// for i := 0; i < len(myArra3); i++ {
// 	myArra3[i] = i + 3
// 	fmt.Println("i: ", myArra3[i])
// }
// 	*/

// 	room := [...]Room{
// 		{name: "Room 1"},
// 		{name: "Room 2"},
// 		{name: "Room 3"},
// 		{name: "Room 4"},
// 	}

// 	room[2].cleaned = true
// 	room[3].cleaned = true

// 	checkCleanlniess(room)

// }

type Producst struct {
	name  string
	price float64
}

func ProducsList(products [4]Producst) {
	var cost, totalItems int

	for i := 0; i < len(products); i++ {
		item := products[i]
		cost += int(item.price)

		if item.name != "" {
			totalItems += 1
		}

	}
	fmt.Println("Last item on th list:", products[totalItems-1])
	fmt.Println("total items", totalItems)
	fmt.Println("Total cost:", cost)

}

func main() {

	produto := [4]Producst{
		{name: "Product 1", price: 10.0},
		{name: "Product 2", price: 20.0},
		{name: "Product 3", price: 30.0},
	}

	ProducsList(produto)

	produto[3] = Producst{name: "Product 4", price: 40.0}

	ProducsList(produto)

}
