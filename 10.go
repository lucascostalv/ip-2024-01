package main

import (
	"fmt"
)

func main() {
	// Declaração das variáveis para armazenar os elementos da matriz 2x2
	var a, b, c, d float64

	// Solicita ao usuário para inserir os elementos da matriz
	fmt.Println("Digite os elementos da matriz 2x2 (a, b, c e d), separados por espaço:")
	fmt.Scanln(&a, &b, &c, &d)

	// Calcula o determinante da matriz
	determinante := a*d - b*c

	// Imprime o resultado formatado
	fmt.Printf("O VALOR DO DETERMINANTE E = %.2f\n", determinante)
}
