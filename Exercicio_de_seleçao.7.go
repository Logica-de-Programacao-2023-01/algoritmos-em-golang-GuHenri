package main

import "fmt"

func main() {

	var final, a, aumento float32

	fmt.Print("Digite seu salario: ")
	fmt.Scan(&a)

	if a < 1000 {
		aumento = a / 100 * 10
	} else {
		aumento = a / 100 * 5
	}
	final = aumento + a
	fmt.Print("Seu novo salario é igual a: R$ ", final)

}
