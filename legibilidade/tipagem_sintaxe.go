package main

import "fmt"

const divisor float64 = 10

// TIPAGEM EM GO: FORTE, ESTÁTICA e INFERIDA/IMPLÍCITA
func somar(a int, b int) int {
	resultado := a + b

	return resultado
}

func dividir(a int) float64 {
	resultado := float64(a) / divisor

	return resultado
}

type Resultado struct {
	Soma    int
	Divisao float64
}

func somarDividir(a int, b int) Resultado {
	soma := somar(a, b)
	divisao := dividir(a)

	return Resultado{soma, divisao}
}

func main() {
	soma := somar(5, 4)
	fmt.Printf("resultado da soma: %d\n", soma)

	divisao := dividir(23)
	fmt.Printf("resultado da divisão: %.2f\n", divisao)

	resultado := somarDividir(5, 4)
	fmt.Printf("resultados (soma, divisão): %d, %.2f\n", resultado.Soma, resultado.Divisao)
}

