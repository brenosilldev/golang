package main

import "fmt"

func main() {

	// //Arrays
	// var arrays [3]int

	// arrays[0] = 1
	// arrays[1] = 2
	// arrays[2] = 3

	// for i, v := range arrays {
	// 	fmt.Println("index: ", i, "value: ", v)
	// }

	//Slices
	// slices := []int{10, 20, 30, 40, 50, 60}

	// fmt.Printf("len=%d cap=%d %v\n", len(slices), cap(slices), slices)
	// fmt.Printf("len=%d cap=%d %v\n", len(slices[:2]), cap(slices[:2]), slices[:2])
	// fmt.Printf("len=%d cap=%d %v\n", len(slices[2:]), cap(slices[2:]), slices[2:])

	// // append no slice
	// slices = append(slices, 70, 80, 90)
	// fmt.Printf("len=%d cap=%d %v\n", len(slices), cap(slices), slices)

	//maps

	salarios := map[string]int{
		"jose":  1000,
		"maria": 2000,
		"joao":  3000,
	}

	salarios["pedro"] = 4000

	// for i, v := range salarios {
	// 	fmt.Println("o salario de", i, "e", v)
	// }

	// Quando se coloca _ ele vai ignora o valor
	for _, v := range salarios {
		fmt.Println("o salario de", v, "e", v)
	}

	// // make cria um map vazio
	// salarios2 := make(map[string]int)
	// fmt.Println(salarios2)
	// salarios2["jose"] = 1000
	// fmt.Println(salarios2)

	// for i, v := range salarios2 {
	// 	fmt.Println("index: ", i, "value: ", v)
	// }

}
