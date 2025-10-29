package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6}

	var resultado []int
	for _, n := range numeros {
		if n%2 == 0 {
			resultado = append(resultado, n*2)
		}
	}

	fmt.Println(resultado)
}
