package main

import "fmt"

func main() {
	var resultado, diferenca = sum(1, 2)
	fmt.Println("1 + 2 = ", resultado)
	fmt.Println("1 - 2 = ", diferenca)
}

// Função pode retorna varios valores
func sum(a, b int) (int, int) {
	return a + b, a - b
}
