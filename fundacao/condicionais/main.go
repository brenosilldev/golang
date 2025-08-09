package main

import (
	"fmt"
	
)

func main() {

	// && - AND - E - Todos os valores devem ser verdadeiros
	// || - OR - OU - Pelo menos um valor deve ser verdadeiro
	// ! - NOT - NÃO - Inverte o valor da expressão

	idade := 18
	idade2 := 17

	if idade >= 18 {
		fmt.Println("Você é maior de idade")
	} else {
		fmt.Println("Você é menor de idade")
	}

	if idade >= 18 && idade2 >= 18 {
		fmt.Println("Vocês são maiores de idade")
	} else {
		fmt.Println("Vocês são menores de idade")
	}

	switch {
	case idade >= 18:
		fmt.Println("Você é maior de idade")
	case idade < 18:
		fmt.Println("Você é menor de idade")
	default:
		fmt.Println("Você não tem idade")
	}

}
