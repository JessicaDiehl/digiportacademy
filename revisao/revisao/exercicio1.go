package revisao
import "fmt"

func Funcao1() {

	listaDeDespesas := []string{"mercado", "aluguel", "eletricidade", "faculdade"}

	var i int

	for i=0; i< len(listaDeDespesas); i++{
		fmt.Printf("%s\n", listaDeDespesas[i])
	}
}