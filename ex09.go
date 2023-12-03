package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
	Preco  float64
}

func aplicarDesconto(livro *Livro) {
	if livro == nil {
		fmt.Println("Erro: Ponteiro nulo.")
		return
	}
	livro.Preco *= 0.90
}

func main() {
	meuLivro := Livro{
		Titulo: "Golang para Iniciantes",
		Autor:  "Programador Anônimo",
		Preco:  50.0,
	}
	fmt.Printf("Preço antes do desconto: %.2f\n", meuLivro.Preco)
	aplicarDesconto(&meuLivro)
	fmt.Printf("Preço após o desconto: %.2f\n", meuLivro.Preco)
}
