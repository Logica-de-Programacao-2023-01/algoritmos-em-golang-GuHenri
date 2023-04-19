package main

import "fmt"

func main() {

	var a, b, menor, c int

	fmt.Println("Escreva tres numeros:")
	fmt.Scan(&a, &b, &c)

	if c < a && c < b {
		menor = c
	} else if b < a {
		menor = b
	} else {
		menor = a
	}

	fmt.Print("O menor eh: ", menor)

}
