package main

import (
	"errors"
	"fmt"
)

func main() {
	var inteiro int32 = 123132
	fmt.Println(inteiro) //tipos inteiros são separados em 8, 16,32 e 64
	// de acordo com a quantidade de bytes que ocupa
	var flutuante float32 = 231.3 // o mesmo vale pros flutuantes
	fmt.Println(flutuante)

	var texto string = "Textos são strings"
	fmt.Println(texto)

	var booleanos bool = true //aceitam apenas true e false
	fmt.Println(booleanos)

	var erro error = errors.New("Isso é um erro interno") //ainda não entendi direito
	fmt.Println(erro)

	char := 'A' //retorna o numero na tabela ASCII
	fmt.Println(char)

}
