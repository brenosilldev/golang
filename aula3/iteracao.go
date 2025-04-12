package main

import "strings"

func Repetir(caractere string) string {
	var repeticoes string
	for i := 0; i < 5; i++ {
		repeticoes += (strings.ToUpper(caractere) + "\n") //  adiciona o caractere
	}
	return repeticoes
}

func main() {
	repeticoes := Repetir("nrepet")
	println(repeticoes)
}
