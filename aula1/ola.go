package main

// import "fmt"?:

// var prefixoOlaPortugues string = ""

func Ola(nome string, idioma string) string {

	// if idioma == "espanhol" {
	// 	prefixoOlaPortugues = "Hola, "
	// } else if idioma == "ingles" {
	// 	prefixoOlaPortugues = "Hello, "
	// } else if idioma == "frances" {
	// 	prefixoOlaPortugues = "Bonjour, "
	// } else if idioma == "italiano" {
	// 	prefixoOlaPortugues = "Ciao, "
	// } else if idioma == "portugues" {
	// 	prefixoOlaPortugues = "Olá, "
	// } else if nome == "" {
	// 	prefixoOlaPortugues = "Olá, Mundo"
	// }

	// switch idioma {
	// case "espanhol":
	// 	prefixoOlaPortugues = "Hola, "
	// case "ingles":
	// 	prefixoOlaPortugues = "Hello, "
	// case "frances":
	// 	prefixoOlaPortugues = "Bonjour, "
	// case "italiano":
	// 	prefixoOlaPortugues = "Ciao, "
	// case "portugues":
	// 	prefixoOlaPortugues = "Olá, "
	// default:
	// 	prefixoOlaPortugues = "Olá, "
	// }

	return prefixSaudade(idioma) + nome
}

// criado função prefixSaudade
// func prefixSaudade(idioma string) (prefixo string) {
func prefixSaudade(idioma string) (prefixo string) {

	switch idioma {
		case "espanhol":
			prefixo = "Hola, "
		case "ingles":
			prefixo = "Hello, "
		case "frances":
			prefixo = "Bonjour, "
		case "italiano":
			prefixo = "Ciao, "
		case "portugues":
			prefixo = "Olá, "
		default:
			prefixo = "Olá, "
	}
	return prefixo
}

// func main() {
// 	fmt.Println(Ola("Mundo"))
// 	fmt.Println(Ola("Golang"))
// }
