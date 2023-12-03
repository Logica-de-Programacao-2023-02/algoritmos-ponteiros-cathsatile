package main

import "fmt"

func inverteString(ponteiroString *string) {
	if ponteiroString == nil {
		fmt.Println("Erro: Ponteiro nulo.")
		return
	}
	bytesDaString := []byte(*ponteiroString)

	for i, j := 0, len(bytesDaString)-1; i < j; i, j = i+1, j-1 {
		bytesDaString[i], bytesDaString[j] = bytesDaString[j], bytesDaString[i]
	}
	*ponteiroString = string(bytesDaString)
}

func main() {
	var minhaString string

	fmt.Print("Digite uma string: ")
	fmt.Scan(&minhaString)
	inverteString(&minhaString)
	fmt.Printf("A string invertida é: %s\n", minhaString)
}
