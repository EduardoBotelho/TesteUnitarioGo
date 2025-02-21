package main

import "testing"

//Padrao Triple AAA
// Arrange(Preparar)
//Act (atua)
//Assert(verificar as assercoes)
func ShouldSumCorrect(t *testing.T) { //convenção de nomes ShouldSum/correct(assinatura do método)
	//arrange
	teste := soma(3, 2, 1)
	//act
	resultado := 6

	//assert
	if teste != resultado {
		t.Error("Valor Esperado: ", resultado, "Valor retornado: ", teste)
	}
}
func ShouldSumIncorrect(t *testing.T) { //ShouldSumIncorrect
	//arrange
	teste := soma(3, 2, 1)
	//act
	resultado := 7

	//assert
	if teste != resultado {
		t.Error("Valor Esperado: ", resultado, "Valor retornado: ", teste)
	}

}
func ShouldMultCorrect(t *testing.T) {
	//arrange
	teste := multiplica(10, 10)
	//act
	resultado := 100
	//assert
	if teste != resultado {
		t.Error("Valor Esperado:", resultado, "Valor retornado:", teste)
	}

}
func ShouldMultInCorrect(t *testing.T) {
	//arrange
	teste := multiplica(10, 10)
	//act
	resultado := 2560
	//assert
	if teste != resultado {
		t.Error("Valor Esperado:", resultado, "Valor retornado:", teste)
	}

}
