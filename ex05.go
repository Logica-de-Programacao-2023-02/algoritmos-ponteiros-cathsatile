package main

import (
	"fmt"
	"math"
)

func atualizaMediaComPi(ponteiroFloat *float64) {
	if ponteiroFloat == nil {
		fmt.Println("Erro: Ponteiro nulo.")
		return
	}
	valorAtual := *ponteiroFloat

	media := (valorAtual + math.Pi) / 2
	*ponteiroFloat = media
}

func main() {
	var minhaVariavel float64

	fmt.Print("Digite um número float64: ")
	fmt.Scan(&minhaVariavel)

	atualizaMediaComPi(&minhaVariavel)
	fmt.Printf("O novo valor é: %f\n", minhaVariavel)
}
