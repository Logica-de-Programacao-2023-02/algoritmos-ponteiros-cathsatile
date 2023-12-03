package main

import (
	"fmt"
	"math"
)

func isPrimo(numero int) bool {
	if numero < 2 {
		return false
	}

	limite := int(math.Sqrt(float64(numero)))

	for i := 2; i <= limite; i++ {
		if numero%i == 0 {
			return false
		}
	}

	return true
}

func preencherComPrimos(slice *[]int, tamanho int) {
	if slice == nil || tamanho <= 0 {
		fmt.Println("Erro: Ponteiro nulo ou tamanho inválido.")
		return
	}

	numero := 2
	for len(*slice) < tamanho {
		if isPrimo(numero) {
			*slice = append(*slice, numero)
		}
		numero++
	}
}

func main() {
	var primos []int
	tamanho := 5
	preencherComPrimos(&primos, tamanho)
	fmt.Printf("Os %d primeiros números primos são: %v\n", tamanho, primos)
}
