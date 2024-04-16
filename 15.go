package main

import (
	"fmt"
)

func main() {
	// Declaração da variável para armazenar o valor de N
	var N int

	// Solicita ao usuário para inserir o valor de N
	fmt.Println("Digite um valor inteiro N (5 < N < 2000):")
	fmt.Scanln(&N)

	// Itera de 2 até N, incrementando de 2 em 2 (apenas números pares)
	for i := 2; i <= N; i += 2 {
		// Calcula o quadrado de cada número par e imprime o resultado formatado
		fmt.Printf("%dˆ2 = %d\n", i, i*i)
	}
}
