package main

import "fmt"

func main() {

	var maior, a, b int

	fmt.Print("Escreva um numero: ")
	fmt.Scan(&b)

	for i := 0; i <= 100; i++ {
		fmt.Print("Escreva outro numero: ")
		fmt.Scan(&a)
		if a > b {
			maior = a
		} else if a > maior {
			maior = a
		} else {
			maior = b
		}
		if a == 0 {
			break
		}
	}
	fmt.Print("O maior numero é: ", maior)
}
