package main

// /*
// 	Ponteiros são variáveis que armazenam o endereço de outra variável na memória.
// 	Isso é útil quando queremos manipular o valor de uma variável sem fazer uma cópia dela.

// 	var x int = 10       Uma variável normal, tipo uma casa com o número 10 dentro
// 	var p *int = &x      Ponteiro pra um int, guarda o endereço da variável x

// 	*p += 5 Aumenta o valor de x em 5, usando o ponteiro p

// 	fmt.Println("x =", x)         Mostra o valor de x
// 	fmt.Println("p =", p)         Mostra o endereço onde x está guardado
// 	fmt.Println("*p =", *p)       Mostra o valor guardado no endereço (ou seja, o valor de x)

// */

// // type Counter struct {
// // 	value int
// // }

// // func increment(n *Counter) {
// // 	n.value += 1 // Incrementa o valor do número usando o ponteiro

// // 	fmt.Println("n.value =", n)
// // }

// // func replace(old *string, new string, n *Counter) {

// // 	*old = new   // Substitui o valor do número usando o ponteiro
// // 	increment(n) // Chama a função increment passando o ponteiro como argumento
// // 	fmt.Println("old =", old)
// // 	fmt.Println("new =", new)

// // }
// // func main() {
// // 	valor := Counter{}

// // 	hello := "Hello"
// // 	world := "World"
// // 	fmt.Println("hello =", hello)
// // 	fmt.Println("world =", world)

// // 	replace(&hello, "hi", &valor) // Passa o endereço da variável hello e world para a função replace
// // 	fmt.Println("hello =", hello)
// // 	fmt.Println("world =", world)

// // }

// const (
// 	Active   = true
// 	Inactive = false
// )

// type Security bool

// type item struct {
// 	nome string
// 	tag  Security
// }

// func activeItem(i *item) {
// 	i.tag = Active // Altera o valor do item para ativo
// }

// func inactiveItem(i *item) {
// 	fmt.Println(i)
// 	// fmt.Println(&*i)
// 	i.tag = Inactive // Altera o valor do item para inativo
// }

// func checkout(items []item) {
// 	for i := 0; i < len(items); i++ {
// 		inactiveItem(&items[i]) // Passa o endereço do item para a função inactiveItem

// 	}
// }

// func main() {
// 	chinelo := item{"Chinelo", Active}
// 	tenis := item{"Tenis", Active}
// 	camisa := item{"Camisa", Active}
// 	sapato := item{"Sapato", Active}

// 	items := []item{chinelo, tenis, camisa, sapato}
// 	fmt.Println(items)
// 	fmt.Println("Antes do checkout:")
// 	fmt.Println("-----")

// 	inactiveItem(&items[0]) // Passa o endereço do item para a função inactiveItem
// 	fmt.Println("Desativando item")
// 	fmt.Println(items)
// 	fmt.Println("-----")

// 	fmt.Println("Depois do checkout:")
// 	checkout(items)
// 	fmt.Println(items)

// }

import (
	"fmt"
	"time"
)

type Titulo string
type Nome string

type Livro struct {
	checkout time.Time
	checkIn  time.Time
}

type Membro struct {
	nome   Nome
	titulo map[Titulo]Livro
}

type LivrosEntry struct {
	total  int
	lended int
}

type Biblioteca struct {
	livros  map[Titulo]LivrosEntry
	membros map[Nome]Membro
}

func printMembro(m *Membro) {
	for titulo, livro := range m.titulo {
		var retornotime string
		if livro.checkIn.IsZero() {
			retornotime = "não retornado"
		} else {
			retornotime = livro.checkIn.String()
		}

		fmt.Println(m.nome, ":", titulo, " - ", livro.checkout.String(), " - ", retornotime)
	}
}

func printBibliotecaMembros(b *Biblioteca) {
	for _, item := range b.membros {
		printMembro(&item)
	}
}

func printBibliotecaLivros(b *Biblioteca) {
	fmt.Println("Livros na biblioteca:")
	fmt.Println("-----")
	for titulo, item := range b.livros {
		fmt.Println(titulo, "/ total", item.total, "/ lended", item.lended)
	}

	fmt.Println("-----")

}

func checkoutLivros(b *Biblioteca, m *Membro, titulo Titulo) bool {
	books, found := b.livros[titulo]
	if !found {
		fmt.Println("Livro não encontrado na biblioteca")
		return false
	}

	if books.lended == books.total {
		fmt.Println("Todos os livros estão emprestados")
		return false
	}

	books.lended += 1
	b.livros[titulo] = books

	m.titulo[titulo] = Livro{time.Now(), time.Time{}}
	return true

}

func retornoLivros(b *Biblioteca, m *Membro, titulo Titulo) bool {
	livro, found := b.livros[titulo]
	if !found {
		fmt.Println("Livro não encontrado na biblioteca")
		return false
	}

	audit, found := m.titulo[titulo]
	if !found {
		fmt.Println("Livro não encontrado no membro")
		return false
	}

	livro.lended -= 1
	b.livros[titulo] = livro

	audit.checkIn = time.Now()
	m.titulo[titulo] = audit

	return true
}

func main() {

	bliblioteca := Biblioteca{
		livros: map[Titulo]LivrosEntry{
			"O Senhor dos Aneis": {total: 3, lended: 0},
			"Harry Potter":       {total: 2, lended: 0},
			"O Hobbit":           {total: 1, lended: 0},
		},
		membros: map[Nome]Membro{
			"J.R.R. Tolkien": {nome: "J.R.R. Tolkien", titulo: make(map[Titulo]Livro)},
			"J.K. Rowling":   {nome: "J.K. Rowling", titulo: make(map[Titulo]Livro)},
		},
	}

	fmt.Println("-----")
	printBibliotecaLivros(&bliblioteca)
	printBibliotecaMembros(&bliblioteca)

	membro := bliblioteca.membros["J.R.R. Tolkien"]
	check := checkoutLivros(&bliblioteca, &membro, "O Senhor dos Aneis")
	if check {
		printBibliotecaLivros(&bliblioteca)
		printBibliotecaMembros(&bliblioteca)
	}

	retorno := retornoLivros(&bliblioteca, &membro, "O Senhor dos Aneis 2")
	if retorno {
		printBibliotecaLivros(&bliblioteca)
		printBibliotecaMembros(&bliblioteca)
	}

	membro2 := bliblioteca.membros["J.K. Rowling"]
	check2 := checkoutLivros(&bliblioteca, &membro2, "Harry Potter")
	if check2 {
		printBibliotecaLivros(&bliblioteca)
		printBibliotecaMembros(&bliblioteca)
	}

	membro3 := bliblioteca.membros["J.K. Rowling"]
	check3 := checkoutLivros(&bliblioteca, &membro3, "Harry Potter")
	if check3 {
		printBibliotecaLivros(&bliblioteca)
		printBibliotecaMembros(&bliblioteca)
	}

	membro4 := bliblioteca.membros["J.K. Rowling"]
	check4 := checkoutLivros(&bliblioteca, &membro4, "Harry Potter")
	if check4 {
		printBibliotecaLivros(&bliblioteca)
		printBibliotecaMembros(&bliblioteca)
	}

}
