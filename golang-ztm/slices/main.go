package main

import "fmt"

// func printSlice(rota string, slice []string) {
// 	fmt.Println()
// 	fmt.Println("----", rota, "----")
// 	for i := 0; i < len(slice); i++ {
// 		fmt.Println(slice[i])
// 	}
// }

// func main() {
// arr := []int{1}
// for i := 1; i < 21; i++ {
// 	arr = append(arr, i)
// }

// fmt.Println(arr[:])
// fmt.Println(arr[1:])
// fmt.Println(arr[0:1])
// fmt.Println(arr[:0])
// fmt.Println(arr[:2])

// numbers := append(arr, 1, 2, 3, 4, 5)
// fmt.Println(numbers)

// numbers2 := []int{5, 6, 7, 8, 9}
// ... passa os o slice para o append
// combined := append(numbers, numbers2...)
// fmt.Println(combined)

// slice := make([]int, 5)
// slice[0] = 1
// slice[1] = 2
// slice[2] = 3
// slice[3] = 4
// slice[4] = 5

// fmt.Println(slice)

// 	route := []string{"São Paulo", "Rio de Janeiro", "Belo Horizonte", "Salvador", "Recife"}
// 	printSlice("Rota 1", route)

// 	printSlice("Rota 2", route[:3])

// 	route = append(route, "João Pessoa")
// 	printSlice("Rota 3", route[3:])

// }

//EXERCÍCIO 1
// Crie um slice de inteiros e imprima o dobro de cada elemento do slice.
// func main() {
// 	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	for i := 0; i < len(slice); i++ {
// 		slice[i] = slice[i] * 2
// 	}
// 	fmt.Println(slice)
// }

type Part string

func showline(line []Part) {
	for i := 0; i < len(line); i++ {
		fmt.Println(line[i])
	}

}

func main() {

	assembly := make([]Part, 3)
	assembly[0] = "Cabeça"
	assembly[1] = "Corpo"
	assembly[2] = "Perna"

	showline(assembly)
	fmt.Println(assembly)
	fmt.Println()

	assembly = append(assembly, "Braço", "Olho")
	fmt.Println(assembly)
	showline(assembly)
	fmt.Println()

	assembly = append(assembly, assembly[2:]...)
	fmt.Println(assembly)
}
