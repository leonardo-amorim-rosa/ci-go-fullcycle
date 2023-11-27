package main

import "testing"

func TestSoma(t *testing.T) {

	total := Soma(15, 1)

	if total != 30 {
		t.Errorf("Operação soma inválida. Resultado: %d Esperado: %d", total, 30)
	}
}
