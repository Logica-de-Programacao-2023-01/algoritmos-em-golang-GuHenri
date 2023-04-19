package main

import "fmt"

func main() {
	var a int
	soma := 0
	soma2 := 0

	for i := 0; i <= 100; i++ {
		fmt.Print("Escreva um numero: ")
		fmt.Scan(&a)
		soma = soma + a
		soma2 = soma2 + 1
		if a == 0 {
			soma2 = soma2 - 1
			break
		}
	}
	media := soma / soma2
	fmt.Print("A média aritmética é igual a: ", media)
}
