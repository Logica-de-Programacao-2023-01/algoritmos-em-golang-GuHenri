package main

import "fmt"

func main() {

	var a int
	var b string
	fmt.Print("Escreva sua idade: ")
	fmt.Scan(&a)

	if a <= 9 {
		b = "mirim"
	} else if a > 9 && a <= 13 {
		b = "infantil"
	} else if a > 13 && a <= 17 {
		b = "juvenil"
	} else {
		b = "adulto"
	}
	fmt.Print("Sua classificação é: ", b)

}
