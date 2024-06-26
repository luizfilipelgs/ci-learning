package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if (total != 30) {
		t.Errorf("Resultado Inválido: %d. Resultado esperado: %d",  total, 30)
	}
}