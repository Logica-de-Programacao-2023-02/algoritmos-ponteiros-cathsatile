package main

import "fmt"

type Conta struct {
	Saldo   float64
	Titular string
}

func depositarNaConta(conta *Conta, valor float64) {
	// Verifica se o ponteiro é válido
	if conta == nil {
		fmt.Println("Erro: Ponteiro nulo.")
		return
	}

	conta.Saldo += valor
}

func main() {
	minhaConta := Conta{
		Saldo:   1000.0,
		Titular: "João",
	}

	fmt.Printf("Saldo antes do depósito: %.2f\n", minhaConta.Saldo)
	depositarNaConta(&minhaConta, 500.0)
	fmt.Printf("Saldo após o depósito: %.2f\n", minhaConta.Saldo)
}
