package main

import "fmt"

func main() {

	var a int
	var b string

	fmt.Println("Escreva um numero:")
	fmt.Scan(&a)

	if a%2 == 0 {
		b = "par"
	} else {
		b = "impar"
	}

	fmt.Print("O numero escrito é ", b)

}
