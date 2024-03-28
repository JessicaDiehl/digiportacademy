// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	type Funcionario struct {
		nome    string
		idade   int
		funcao  string
		salario float32
	}

	a := Funcionario{"Jessica", 27, "Engenheira", 10000}

	b := Funcionario{"Karol", 18, "Professora", 5000}

	c := Funcionario{"Pablo", 26, "Tecnico", 3000}

	// Ignore unused variables a and c
	_ = a
	_ = c

	fmt.Println("Funcionário b:", b)

}
