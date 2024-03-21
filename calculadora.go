// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	var num1 int
	var num2 int
	var op string

	fmt.Println("Digite o primeiro numero")
	fmt.Scan(&num1)

	if num1 > 0 {
		fmt.Println("Digite o segundo numero")
		fmt.Scan(&num2)
	} else {
		fmt.Println("O numero inserido é negativo, digite apenas numeros postivos")
	}

	if num2 > 0 {
		fmt.Println("Qual operação deseja realizar?")
		fmt.Scan(&op)
	} else {
		fmt.Println("O numero inserido é negativo, digite apenas numeros postivos")
	}

	switch {
	case op == "soma" || op == "Soma" || op == "+":
		fmt.Println("A soma dos números é ", num1+num2)
	case op == "subtração" || op == "Subtração" || op == "-":
		fmt.Println("A subtração dos números é ", num1-num2)
	case op == "Divisão" || op == "divisão" || op == "/":
		fmt.Println("A divisão dos números é ", num1/num2)
	case op == "Multiplicação" || op == "Multiplicação" || op == "*":
		fmt.Println("A multiplicacao dos números é ", num1*num2)
	default:
		fmt.Println("Fim! ")
	}

}
