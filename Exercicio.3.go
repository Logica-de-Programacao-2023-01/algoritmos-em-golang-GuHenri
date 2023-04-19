package main

import "fmt"

func main() {

	var altura, peso float64

	fmt.Print("Escreva seu peso: ")
	fmt.Scan(&peso)
	fmt.Print("Escreva sua altura: ")
	fmt.Scan(&altura)

	imc := peso / (altura * altura)

	fmt.Printf("Seu IMC eh igual a: %.2f", imc)
}
