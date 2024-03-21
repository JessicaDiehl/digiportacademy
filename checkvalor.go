// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	var numero int

	fmt.Println("Digite um numero")
	fmt.Scan(&numero)

	if numero > 0 {
		fmt.Println("Positivo ", numero)
	} else if numero < 0 {
		fmt.Println("Negativo ", numero)
	} else {
		fmt.Println("O numero inserido é 0 ")
	}
}
