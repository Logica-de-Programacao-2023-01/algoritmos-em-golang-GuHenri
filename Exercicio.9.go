package main

import "fmt"

func main() {

	var a float32

	fmt.Print("Escreva o valor do produto:")
	fmt.Scan(&a)

	desconto := a / 10
	preco := a - desconto
	fmt.Print("Preco do produto com o desconto de %10 aplicado: R$ ", preco)

}
