package main

import "fmt"

func main() {

	var a float32

	fmt.Print("Escreva seu peso em Kg:")
	fmt.Scan(&a)

	libras := a * 2.20462

	fmt.Printf("Seu peso em libras sera igual a: %.2f ", libras)

}
