package main

import (
	"fmt"
)

func main() {
	// Declaração das variáveis para armazenar os coeficientes da equação quadrática
	var a, b, c float64

	// Solicita ao usuário para inserir os coeficientes A, B e C
	fmt.Println("Digite os coeficientes da equação quadrática (A, B e C), separados por espaço:")
	fmt.Scanln(&a, &b, &c)

	// Calcula o discriminante
	delta := b*b - 4*a*c

	// Imprime o resultado formatado
	fmt.Printf("O VALOR DE DELTA E = %.2f\n", delta)
}
