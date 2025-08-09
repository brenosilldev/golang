package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//Criando um arquivo
	arquivo, err := os.Create("arquivo.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		panic(err)
	}
	//Escrevendo no arquivo
	_, err = arquivo.Write([]byte("Hello World ,Seja bem vindo ao meu mundo\n"))
	_, err = arquivo.Write([]byte("Hello World ,Seja bem vindo ao meu mundo 1.0 \n"))
	_, err = arquivo.Write([]byte("Hello World ,Seja bem vindo ao meu mund o 2.0\n"))

	//Lendo o arquivo
	arquivo2, err := os.ReadFile("arquivo.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		panic(err)
	}

	fmt.Println(string(arquivo2))

	arquivo4, err := os.ReadFile("teste1.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		panic(err)
	}

	fmt.Println(string(arquivo4))

	arquivo3, err := os.Open("arquivo.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		panic(err)
	}

	reader := bufio.NewReader(arquivo3)
	//Lendo o arquivo com buffer a cada 20 bytes
	buffer := make([]byte, 30)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	arquivo5, err := os.Open("teste1.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		panic(err)
	}

	reader2 := bufio.NewReader(arquivo5)
	//Lendo o arquivo com buffer a cada 20 bytes
	buffer2 := make([]byte, 3)

	for {
		n, err := reader2.Read(buffer2)
		if err != nil {
			break
		}
		fmt.Println(string(buffer2[:n]))
	}

}
