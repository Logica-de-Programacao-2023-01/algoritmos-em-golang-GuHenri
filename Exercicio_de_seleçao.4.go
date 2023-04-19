package main

import "fmt"

func main() {

	var a, c float32
	var b, d string

	fmt.Println("Digite seu sexo 'f' para feminino e 'm' para masculino:")
	fmt.Scan(&b)
	fmt.Println("Digite a sua altura:")
	fmt.Scan(&a)
	fmt.Println("Digite a seu peso:")
	fmt.Scan(&c)

	imc := c / (a * a)

	if imc < 18.5 {
		d = "Abaixo do peso ideal "
	} else if imc > 18.5 && imc < 29 {
		d = "Na pesagem ideal "
	} else {
		d = "Acima do peso ideal "
	}
	fmt.Printf("Taxa de IMC: %.2f\n", imc)
	fmt.Printf("Voce esta '%s' para seu sexo '%s' e altura '%.2f' ", d, b, a)
}
