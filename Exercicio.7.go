package main

import "fmt"

func main() {

	var a int

	fmt.Print("Escreva um numero:")
	fmt.Scan(&a)

	antecessor := a - 1
	sucessor := a + 1
	fmt.Printf("Numero anterior '%d' numero sucessor '%d'", antecessor, sucessor)

}
