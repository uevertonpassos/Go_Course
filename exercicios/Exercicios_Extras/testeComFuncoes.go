package main

import (
	"fmt"
)

func main() {
	resultadoSoma := soma(3, 6)
	fmt.Printf("A soma é %v\n", resultadoSoma)

	resultadoDuplicada, resultadoMultiplicacao := duplicada(10, 3)
	fmt.Printf("A soma é %v e a multiplicação é %v\n", resultadoDuplicada, resultadoMultiplicacao)

	_, resultadoMultiplicacao = duplicada(3, 2)
	fmt.Printf("A multiplicação é %v\n", resultadoMultiplicacao)
}

func soma(n1, n2 int8) int8 {
	return n1 + n2
}

func duplicada(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	multiplicacao := n1 * n2
	return soma, multiplicacao
}

//tentando retornar duas funções ao msm tempo no go, ainda não entendi direito, reforçar
