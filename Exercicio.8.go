package main

import "fmt"

func main() {

	var b, a float32

	fmt.Print("Escreva quantos dias voce trabalhou este mes:")
	fmt.Scan(&a)
	fmt.Print("Escreva o valor da sua diaria:")
	fmt.Scan(&b)

	salario := b * a

	fmt.Print("Seu salario este mes foi igual a: R$ ", salario)

}
