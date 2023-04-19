package main

import "fmt"

func main() {

	var a int
	var dia string

	fmt.Print("Digite um numero de 1 a 7: ")
	fmt.Scan(&a)
	if a < 1 || a > 7 {
		for i := 0; i < 100; i++ {
			fmt.Print("Digite novamente um numero de 1 a 7: ")
			fmt.Scan(&a)
			if a >= 1 && a <= 7 {
				break
			}
		}
	}
	switch a {
	case 1:
		dia = "domingo"
	case 2:
		dia = "segunda"
	case 3:
		dia = "terça"
	case 4:
		dia = "quarta"
	case 5:
		dia = "quinta"
	case 6:
		dia = "sexta"
	case 7:
		dia = "sabado"
	}
	fmt.Print("Dia da semana correspondente: ", dia)
}
