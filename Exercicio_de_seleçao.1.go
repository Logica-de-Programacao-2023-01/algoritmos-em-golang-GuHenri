package main

import "fmt"

func main() {

	var a, b, maior int

	fmt.Println("Escreva dois numeros:")
	fmt.Scan(&a, &b)

	if a > b {
		maior = a
	} else {
		maior = b
	}

	fmt.Print("O maior é: ", maior)

}
