// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	var mensagem string = "Hello"

	slice := []string{mensagem, "Nome1", "Nome2", "Nome3"}

	segundoValor := slice[2]

	fmt.Println(slice[0], segundoValor)
}
