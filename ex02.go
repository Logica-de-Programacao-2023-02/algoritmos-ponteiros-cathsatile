package main

import "fmt"

func verificaParOuImpar(ponteiroInteiro *int) {
	if ponteiroInteiro == nil {
		fmt.Println("Erro: Ponteiro nulo.")
		return
	}

	if *ponteiroInteiro%2 == 0 {
		*ponteiroInteiro = 0
	} else {
		*ponteiroInteiro = 1
	}
}

func main() {
	var numero int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&numero)
	verificaParOuImpar(&numero)
	fmt.Printf("O valor atualizado é: %d\n", numero)
}
