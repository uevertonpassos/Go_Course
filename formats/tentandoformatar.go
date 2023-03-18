package main

import (
	"fmt"
	"modulo/importando"
)

func main() {

	var x int8 = 30
	var y int8 = 20

	fmt.Println(importando.Numeros(3, 4))
	fmt.Printf("%v + %v = %v e é um %T", x, y, x+y, x+y)
	// ISTO É UM PRINT FORMATADO, FAVOR LEMBRAR DE USAR O PRINTF AO INVES DO
	//PRINTLN

}
