package main

// Pacotes

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {

	uuid := uuid.New()
	fmt.Println(uuid.Time())
	fmt.Println(uuid.Version())
	fmt.Println(uuid.String())
	fmt.Println(uuid.Variant())

}

// Go mod tidy - atualiza o go.mod e go.sum
// Go mod vendor - cria uma pasta vendor com todos os modulos do projeto
// Go mod verify - verifica se o go.mod e go.sum estão atualizados
// Go mod download - baixa todos os modulos do projeto
// Go mod graph - mostra o grafico de dependencias do projeto
// Go mod why - mostra porque um modulo foi escolhido
// Go mod why -r - mostra porque um modulo foi escolhido recursivamente
// Go mod why -r -m - mostra porque um modulo foi escolhido recursivamente e mostra o modulo que ele depende
// Go mod why -r -m -d - mostra porque um modulo foi escolhido recursivamente e mostra o modulo que ele depende e mostra o modulo que ele depende
// go get
