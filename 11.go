package main

import "fmt"

func main() {
	// Declaração da variável para armazenar o número inteiro fornecido pelo usuário
	var numero int

	// Solicita ao usuário para inserir o número inteiro
	fmt.Println("Digite um número inteiro:")
	fmt.Scanln(&numero)

	// Verifica se o número é divisível por três e cinco
	if numero%3 == 0 && numero%5 == 0 {
		fmt.Println("O NUMERO E DIVISIVEL")
	} else {
		fmt.Println("O NUMERO NAO E DIVISIVEL")
	}
}
