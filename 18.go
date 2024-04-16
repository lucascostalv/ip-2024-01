package main

import "fmt"

func main() {
	// Declaração das variáveis para armazenar o valor inicial da progressão, a razão e o número de elementos
	var a1, r, n int

	// Solicita ao usuário para inserir os valores
	fmt.Println("Digite o valor inicial da progressão, a razão e o número de elementos, separados por espaço:")
	fmt.Scanln(&a1, &r, &n)

	// Calcula a soma dos n primeiros elementos da progressão aritmética
	soma := calcularSomaProgressao(a1, r, n)

	// Imprime o resultado
	fmt.Println(soma)
}

// Função para calcular a soma dos n primeiros elementos de uma progressão aritmética
func calcularSomaProgressao(a1, r, n int) int {
	// Inicializa a soma
	soma := 0

	// Itera de 0 até n-1, somando cada elemento da progressão
	for i := 0; i < n; i++ {
		termo := a1 + i*r
		soma += termo
	}

	return soma
}
