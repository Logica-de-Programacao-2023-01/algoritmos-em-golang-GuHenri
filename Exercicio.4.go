package main

import "fmt"

func main() {

	var a, b, c float64

	fmt.Println("Escreva 3 numeros que no final terao os respectivos pesos: '2,3 e 5', para o calculo de uma media ponderada")

	fmt.Scan(&a, &b, &c)

	pesos := a*2 + b*3 + c*5
	media := pesos / 10

	fmt.Print("A media ponderada sera igual a: ", media)

}
