package main

import "fmt"

func somar(n1, n2 int8) int8 {
	return n1 + n2
}

func nomear() string {
	return ("Ueverton Passos")
}

func main() {
	soma := somar(12, 12)
	fmt.Println(soma)

	fmt.Println("Esse código foi feito por")
	fmt.Println(nomear())
}
