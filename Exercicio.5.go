package main

import "fmt"

func main() {

	var a int

	fmt.Println("Escreva quantos anos voce tem:")
	fmt.Scan(&a)

	dias := a * 365

	fmt.Print("Sua idade em dias: ", dias)

}
