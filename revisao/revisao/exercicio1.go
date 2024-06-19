package revisao

import "fmt"

func Funcao1() {

	listaDeDespesas := []string{"mercado", "aluguel", "eletricidade", "faculdade"}

	fmt.Println("Lista de despesas:")
	for _, despesa := range listaDeDespesas {
		fmt.Println(despesa)
	}

	fmt.Println("\nVerificação de despesa específica:")
	var despesaBuscada string =  "mercado"

	fmt.Scanln(&despesaBuscada)

	encontrouDespesa := false
	for _, despesa := range listaDeDespesas {
		if despesa == despesaBuscada {
			encontrouDespesa = true
			break
		}
	}
	if encontrouDespesa {
		fmt.Printf("%s está na lista de despesas.\n", despesaBuscada)
	} else {
		fmt.Printf("%s não está na lista de despesas.\n", despesaBuscada)
	}

	fmt.Printf("\nTotal de despesas na lista: %d\n", len(listaDeDespesas))
}