package main

import "fmt"

func main() {

	var a int

	fmt.Print("Digite um numero: ")
	fmt.Scan(&a)

	fmt.Printf("Os multiplos de %d são: ", a)
	for i := 1; i <= a; i++ {
		if a%i == 0 {
			fmt.Println(i)
		}
	}
}
