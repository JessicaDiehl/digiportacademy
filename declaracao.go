// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	var idade int
	var permissaoDosPais bool

	fmt.Println("Quantos anos voce tem? ")
	fmt.Scan(&idade)

	fmt.Println("Voce tem permissao dos pais? ")
	fmt.Scan(&permissaoDosPais)

	podeAssistir := idade >= 18 || (idade > 13 || permissaoDosPais)

	fmt.Println("Voce pode assistir a esse filme: ", podeAssistir)

}
