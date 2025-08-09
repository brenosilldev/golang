package main

import "fmt"

// type Gato struct{}
// type Cachorro struct{}

// // Função pode retorna varios valores
// func (c Cachorro) Falar() string {
// 	return "au au"
// }

// func (c Gato) Falar() string {
// 	return "miau"
// }

// // Pode passar apenas metodos
// type Animal interface {
// 	Falar() string
// }

// // Função pode retorna varios valores
// func FazerAnimalFalar(a Animal) {
// 	fmt.Println(a.Falar())
// }

// func main() {

// 	var a Cachorro
// 	var b Gato

// 	FazerAnimalFalar(a)
// 	FazerAnimalFalar(b)

// }

// type Pessoa struct {
// 	nome      string
// 	sobrenome string
// 	idade     int
// }

// func (p Pessoa) Falar() string {
// 	return fmt.Sprintf("%s %s tem %d anos", p.nome, p.sobrenome, p.idade)
// }

// type Pessoas interface {
// 	Falar() string
// }

// func FazerPessoaFalar(p Pessoas) {
// 	fmt.Println(p.Falar())
// }

/*************  ✨ Windsurf Command ⭐  *************/
// // main exemplo de como usar interfaces
// //
// // Mostra como uma interface pode ser usada com diferentes estruturas
// // e como pode ser usada para polimorfismo
// /*******  faf5c6e4-1adf-4d83-bb7e-fdbaaab16369  *******/func main() {
// 	p := Pessoa{"Breno", "Silva", 21}
// 	FazerPessoaFalar(p)

// 	b := Pessoa{"Breno", "Silva", 22}
// 	FazerPessoaFalar(b)

// }

func main() {
	//Interface vazia
	var a interface{} = 10 //implementa qualquer tipo de dado
	var b interface{} = "Hello"

	showType(a)
	showType(b)

}

// Função que mostra o tipo e o valor de uma interface
func showType(t interface{}) {
	fmt.Printf("O tipo de t é %T e o valor é %v\n", t, t)
}
