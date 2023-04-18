package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int

	fmt.Println("Escreva 3 numeros: ")
	fmt.Scan(&a, &b, &c)

	d = a + b + c

	fmt.Print("Soma = ", d)
}
