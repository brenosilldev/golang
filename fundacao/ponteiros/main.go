package main

import "fmt"

// Ponteiros são variáveis que armazenam endereços de memória

// func main() {
// 	a := 10
// 	fmt.Println(a)
// 	fmt.Println(&a) // endereço de memória
// 	fmt.Println("--------------------------------")
// 	// Ponteiro é uma variável que armazena o endereço de memória de outra variável
// 	var ponteiro *int = &a
// 	fmt.Println(ponteiro)
// 	fmt.Println(*ponteiro) // valor do endereço de memória
// 	fmt.Println("--------------------------------")
// 	fmt.Println("Alterando o valor de a")
// 	*ponteiro = 20
// 	fmt.Println(a)
// 	fmt.Println(*ponteiro)
// 	fmt.Println(ponteiro)

// }
//
// func somar(a, b *int) int {
// 	*a = 30
// 	*b = 40
// 	return *a + *b
// }

// func main() {
// 	a := 10
// 	fmt.Println(a)
// 	b := 20
// 	fmt.Println(b)
// 	fmt.Println(somar(&a, &b))
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// structs são tipos de dados que permitem agrupar valores de diferentes tipos
type Carro struct {
	Modelo     string
	Ano        int
	Cor        string
	Velocidade int
}

// métodos são funções que pertencem a um tipo de dado específico (structs)
// (c *Carro) é o receptor do método, ou seja, o objeto que recebe o método
// valor int é o parâmetro do método
// c.Modelo = "Ferrari" é uma alteração no objeto
// c.Velocidade += valor é uma alteração no objeto

func (c *Carro) andar(valor int) {
	fmt.Printf("O carro %s está andando\n", c.Modelo)
	c.Modelo = "Ferrari"
	c.Velocidade += valor
}

func main() {
	carro := Carro{
		Modelo: "Gol",
		Ano:    2020,
		Cor:    "Vermelho",
	}

	carro.andar(10)
	fmt.Println(carro.Modelo)
	fmt.Println(carro.Velocidade)
	carro.andar(20)
	fmt.Println(carro.Velocidade)
}
