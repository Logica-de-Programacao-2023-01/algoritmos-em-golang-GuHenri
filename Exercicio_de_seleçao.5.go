package main

import "fmt"

func main() {

	var a int

	fmt.Print("Digite um numero:")
	fmt.Scan(&a)

	if a%3 == 0 && a%5 == 0 {
		fmt.Print("O numero digitado é múltiplo de 3 e de 5 ao mesmo tempo")
	} else {
		fmt.Print("O numero digitado não é múltiplo de 3 e 5 ao mesmo tempo")
	}

}
