// Type assertion

package main

import (
	"fmt"
	"golang/pacotes/matematica"
)

func main() {

	// var minhaVar interface{} = "Hello"
	// println(minhaVar.(string)) // Definindo o tipo de dado da interface

	// res, ok := minhaVar.(int) // convertendo o tipo de dado da interface para int
	// if !ok { // retornando um booleano se a conversão foi feita com sucesso
	// 	println("Não é um int")
	// } else {
	// 	println(res)
	// }

	fmt.Println(matematica.Soma(10, 20))

}
