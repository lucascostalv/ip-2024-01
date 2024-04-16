package main

import "fmt"

func main() {
	// Declaração das variáveis para armazenar os números inteiros
	var x, y int

	// Solicita ao usuário para inserir os números inteiros x e y
	fmt.Println("Digite dois números inteiros separados por um espaço:")
	fmt.Scanln(&x, &y)

	// Verifica se x é par
	if x%2 == 0 {
		// Se x for par, imprime a sequência de números pares iniciando com x
		imprimirSequenciaPares(x, y)
	} else {
		// Se x não for par, imprime a mensagem indicando que o primeiro número não é par
		fmt.Println("O PRIMEIRO NUMERO NAO E PAR")
	}
}

// Função para imprimir a sequência de números pares iniciando com o número dado
func imprimirSequenciaPares(x, y int) {
	// Inicializa o contador para o próximo número par
	numeroPar := x

	// Itera y vezes para imprimir a sequência de números pares
	for i := 0; i < y; i++ {
		fmt.Printf("%d ", numeroPar)
		numeroPar += 2 // Incrementa para o próximo número par
	}

	fmt.Println() // Imprime o caractere de quebra de linha
}
