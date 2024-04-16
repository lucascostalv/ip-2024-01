package main

import (
	"fmt"
	"math"
)

func main() {
	// Declaração da variável para armazenar o valor de n
	var n int

	// Solicita ao usuário para inserir o valor de n
	fmt.Println("Digite um número inteiro positivo e maior que 1:")
	_, err := fmt.Scanln(&n)

	// Verifica se houve algum erro na leitura
	if err != nil || n <= 1 {
		fmt.Println("Número inválido!")
		return
	}

	// Calcula o valor do somatório
	soma := calcularSomatorio(n)

	// Imprime o resultado com 6 casas decimais
	fmt.Printf("%.6f\n", soma)
}

// Função para calcular o somatório S = 1/1 + 1/2 + 1/3 + ... + 1/n
func calcularSomatorio(n int) float64 {
	var soma float64
	for k := 1; k <= n; k++ {
		soma += 1.0 / float64(k)
	}
	return soma
}
