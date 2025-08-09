package main

import (
	"fmt"
)

func main() {
	// var resultado, error = sum(1, 2)

	// if error != nil {
	// 	fmt.Println(error)
	// 	return
	// }

	// fmt.Println("Resultado da soma: ", resultado)

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

}

// Função pode retorna varios valores
// func sum(a, b int) (int, int) {
// 	return a + b, a - b
// }

// func sum(a, b int) (int, error) {
// 	if a+b >= 10 {
// 		return 0, errors.New("a não pode ser maior que b")
// 	}
// 	return a + b, nil
// }



func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
		fmt.Println("numero: ", "-", numero)
	}
	return total

}
