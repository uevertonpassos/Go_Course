package main

import "fmt"

var x int8 = 12
var y string = "Uev" // se eu não atribuir nenhum valor, o compilador retona ZERO
var z bool = true

func main() {
	s := fmt.Sprint(x, y, z)
	fmt.Println(s)
}
