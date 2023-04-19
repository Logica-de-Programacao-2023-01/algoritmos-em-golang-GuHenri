package main

import "fmt"

func main() {

	var a, b, resu float32

	fmt.Println("Digite dois numeros: ")
	fmt.Scan(&a, &b)

	if a < 0 || b < 0 {
		resu = a + b
		fmt.Print("O resultado da soma entre os dois é igual a: ", resu)
	} else {
		resu = a * b
		fmt.Print("O resultado da multiplicação entre os dois é igual a: ", resu)
	}

}
