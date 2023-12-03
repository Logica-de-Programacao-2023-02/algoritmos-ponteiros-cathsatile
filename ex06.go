package main

import "fmt"

type Book struct {
	Titulo string
	Autor  string
}

func alteraTituloSeAnonimo(livro *Book) {
	// Verifica se o ponteiro é válido
	if livro == nil {
		fmt.Println("Erro: Ponteiro nulo.")
		return
	}

	if livro.Autor == "Anônimo" {
		livro.Titulo = "Desconhecido"
	}
}

func main() {
	meulivro := Book{
		Titulo: "O Livro Desconhecido",
		Autor:  "Anônimo",
	}

	fmt.Printf("Antes da alteração: Título: %s, Autor: %s\n", meulivro.Titulo, meulivro.Autor)
	alteraTituloSeAnonimo(&meulivro)
	fmt.Printf("Após a alteração: Título: %s, Autor: %s\n", meulivro.Titulo, meulivro.Autor)
}
