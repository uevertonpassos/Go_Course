package main

import (
	"fmt"
)

type batata int

var x batata

func main() {

	fmt.Printf("%v, é do tipo %T\n", x, x)
	x = 42
	fmt.Println(x)

}
