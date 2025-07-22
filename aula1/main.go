package main

import "fmt"

var x int

/*	Posso cria meu type  */
type ID int

func main() {

	var b ID

	b = 1

	fmt.Printf("b:  %T\n", b) //ID(b)
}
