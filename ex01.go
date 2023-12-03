package main

import (
	"fmt"
)

func somaNaturais(ponteiroInteiro *int, n int) {
	if ponteiroInteiro == nil {
		fmt.Println("Erro: Ponteiro nulo.")
		return
	}
	if n < 0 {
		fmt.Println("Erro: Número de termos deve ser não-negativo.")
		return
	}

	soma := 0
	for i := 1; i <= n; i++ {
		soma += i
	}
	*ponteiroInteiro = soma
}

func main() {
	var valor int
	n := 5
	somaNaturais(&valor, n)
	fmt.Printf("A soma dos %d primeiros números naturais é: %d\n", n, valor)
}
