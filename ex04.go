package main

import "fmt"

func somaUltimosDigitos(ponteiroInteiro *int) {
	if ponteiroInteiro == nil {
		fmt.Println("Erro: Ponteiro nulo.")
		return
	}
	valor := *ponteiroInteiro

	ultimoDigito := valor % 10
	valor /= 10
	penultimoDigito := valor % 10
	soma := ultimoDigito + penultimoDigito
	*ponteiroInteiro = soma
}

func main() {
	var minhaVariavel int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&minhaVariavel)

	somaUltimosDigitos(&minhaVariavel)
	fmt.Printf("O novo valor é: %d\n", minhaVariavel)
}
