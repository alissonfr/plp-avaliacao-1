package escrita

import (
	"fmt"
	"strings"
)

func criarListaDeCompras() map[string]int {
	lista := make(map[string]int)
	lista["manga"] = 4
	lista["banana"] = 12
	lista["laranja"] = 12
	lista["melancia"] = 1

	return lista
}

func imprimirLista(lista map[string]int) {
	for fruta, quantidade := range lista {
		fmt.Printf("vou comprar %d unidade(s) de %s\n", quantidade, fruta)
	}
}

func somarQuantidade(lista map[string]int) int {
	quantidade := 0
	for _, q := range lista {
		quantidade += q
	}

	return quantidade
}

func filtrar(lista map[string]int, filtro string) map[string]int {
	filtrado := make(map[string]int)
	for fruta, quantidade := range lista {
		if strings.Contains(fruta, filtro) {
			filtrado[fruta] = quantidade
		}
	}

	return filtrado
}

func main() {
	lista := criarListaDeCompras()
	imprimirLista(lista)

	fmt.Printf("quantidade total de itens = %d\n", somarQuantidade(lista))

	lista = filtrar(lista, "ana")
	imprimirLista(lista)
}
