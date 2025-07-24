package main

// Struct - Estrutura de dados em Go

import "fmt"

type endereco struct {
	logradouro string
	numero     int
}

type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
	ativo     bool
	endereco  endereco
}

// Alterar o valor da struct direto
func (p *Pessoa) Desativar() {
	p.ativo = false
	fmt.Println("Pessoa desativada")
}

func main() {
	// breno := pessoa{"Breno", "Silva", 21}
	// fmt.Println(breno)

	// nome, sobrenome, idade := breno.nome, breno.sobrenome, breno.idade
	// fmt.Println(nome, sobrenome, idade)

	// breno.idade = 22
	// fmt.Println(breno)

	endereco := endereco{
		logradouro: "Rua 1",
		numero:     1,
	}

	bruno := Pessoa{
		nome:      "Bruno",
		sobrenome: "Silva",
		idade:     21,
		endereco:  endereco,
	}
	bruno.ativo = true

	fmt.Println(bruno)
	bruno.Desativar()
	fmt.Println(bruno)

	// fmt.Println(bruno)

}
