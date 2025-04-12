package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("mundo")
	esperado := "Olá, mundo mundo"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
