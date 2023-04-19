package main

import "fmt"

func main() {

	var a float32

	fmt.Println("Escreva seu salario atual:")
	fmt.Scan(&a)

	porcentagem := a / 100 * 15
	novo := porcentagem + a
	fmt.Printf("Novo salario com um aumento de 15%%: R$ %.2f ", novo)

}
