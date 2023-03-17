package main

import (
	"bufio"
	"fmt"
	"modulo/auxiliar"
	"os"
)

func main() {
	fmt.Println("Tentando buid")
	auxiliar.Escrevendo()
	auxiliar.Fixando()
	bufio.NewReader(os.Stdin).ReadBytes('\n') //pausando o código compilado

}
