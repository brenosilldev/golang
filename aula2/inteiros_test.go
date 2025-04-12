package inteiros

import (
	"fmt"
	"testing"
)

// Soma recebe dois inteiros e retorna a soma del"

func TestAdicionador(t *testing.T) {
	soma := Soma(2, 2)
	// soma := Adicionador(2, 2)
	fmt.Println("Soma:", soma)
	esperado := 4

	if soma != esperado {
		t.Errorf("esperado '%d', resultado '%d'", esperado, soma)
	}
}