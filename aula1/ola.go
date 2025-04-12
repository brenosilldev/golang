package main

import "fmt"

func Ola(nome string) string {
	return "Olá, mundo" + " " + nome
}

func main() {
	fmt.Println(Ola("Mundo"))
	fmt.Println(Ola("Golang"))
}
