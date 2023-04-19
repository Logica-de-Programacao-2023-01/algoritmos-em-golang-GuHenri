package main

import "fmt"

func main() {

	var a int

	fmt.Print(" Digite um numero: ")
	fmt.Scan(&a)
	for i := 1; i < 11; i++ {
		b := a * i
		fmt.Printf("%d x %d = %d\n", a, i, b)
	}
}
