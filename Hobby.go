// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	listaHobbies := make(map[string]string)

	listaHobbies["Jessica"] = "Assistir Serie"
	listaHobbies["Izadora"] = "Danca"
	listaHobbies["Natalia"] = "Volei"

	for name, hobby := range listaHobbies {
		fmt.Printf("O hooby da %s, é %s\n", name, hobby)

	}

}
