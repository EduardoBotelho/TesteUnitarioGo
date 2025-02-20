package main

import "testing"

func TestSoma(t *testing.T) { //convenção de nomes ShouldSum/correct(assinatura do método)
	teste := soma(3, 2, 1)
	resultado := 6

	if teste != resultado {
		t.Error("Valor Esperado: ", resultado, "Valor retornado: ", teste)
	}
}
func TestSoma2(t *testing.T) { //ShouldSumIncorrect
	teste := soma(3, 2, 1)
	resultado := 7

	if teste != resultado {
		t.Error("Valor Esperado: ", resultado, "Valor retornado: ", teste)
	}

}
