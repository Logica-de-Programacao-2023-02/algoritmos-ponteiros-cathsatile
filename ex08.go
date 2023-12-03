package main

import (
	"fmt"
)

func modificarValor(ponteiro *int) {
	// Verifica se o ponteiro é válido
	if ponteiro == nil {
		fmt.Println("Erro: Ponteiro nulo.")
		return
	}
	*ponteiro = 42
}

func main() {
	var minhaVariavel int
	ponteiroParaVariavel := &minhaVariavel

	fmt.Printf("Valor antes da modificação: %d\n", minhaVariavel)
	modificarValor(ponteiroParaVariavel)
	fmt.Printf("Valor após a modificação: %d\n", minhaVariavel)
}
