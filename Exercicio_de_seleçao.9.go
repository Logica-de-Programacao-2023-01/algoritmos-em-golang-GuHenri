package main

import "fmt"

func main() {

	var a, b, c int
	fmt.Println("Escreva 3 numeros: ")
	fmt.Scan(&a, &b, &c)

	if a < b && b < c {
		fmt.Printf("Numeros em ordem crescente: \n%d \n%d \n%d", a, b, c)
	} else if a < b && c < b {
		fmt.Printf("Numeros em ordem crescente: \n%d \n%d \n%d", a, c, b)
	} else if b > c && c > a {
		fmt.Printf("Numeros em ordem crescente: \n%d \n%d \n%d", b, c, a)
	} else if b < a && a < c {
		fmt.Printf("Numeros em ordem crescente: \n%d \n%d \n%d", b, a, c)
	} else if c < a && a < b {
		fmt.Printf("Numeros em ordem crescente: \n%d \n%d \n%d", c, a, b)
	} else {
		fmt.Printf("Numeros em ordem crescente: \n%d \n%d \n%d", c, b, a)
	}
}
