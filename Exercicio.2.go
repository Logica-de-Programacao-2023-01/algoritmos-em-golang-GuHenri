package main

import "fmt"

func main() {

	var a, b, c, d int

	fmt.Print("Escreva um numero: ")
	fmt.Scan(&a)
	b = a * 2
	c = a * 3
	d = a * 4

	fmt.Println("Dobro =  ", b)
	fmt.Println("Triplo =  ", c)
	fmt.Println("Quadruplo =  ", d)
}
