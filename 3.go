package main

import (
	"fmt"
)

func main() {
	var n1, n2, n3 int

	// Solicita ao usuário para inserir os três números separados por espaço
	fmt.Println("Digite três números inteiros separados por espaço:")

	// Lê os números fornecidos pelo usuário
	fmt.Scanln(&n1, &n2, &n3)

	// Verifica se algum dos números tem mais de um dígito
	if n1 < 0 || n1 > 9 || n2 < 0 || n2 > 9 || n3 < 0 || n3 > 9 {
		fmt.Println("DIGITO INVALIDO")
		return
	}

	// Concatenação dos números
	numero := fmt.Sprintf("%d%d%d", n1, n2, n3)

	// Cálculo do quadrado do número
	quadrado := n1*n1*10000 + n2*n2*100 + n3*n3

	// Imprime o número e o quadrado
	fmt.Printf("%s, %d\n", numero, quadrado)
}
