package main

import "fmt"

func main() {
	// Declaração da variável para armazenar a quantidade de horas de uso da charrete
	var horas int

	// Solicita ao usuário para inserir a quantidade de horas de uso da charrete
	fmt.Println("Digite a quantidade de horas de uso da charrete:")
	fmt.Scanln(&horas)

	// Calcula o valor do aluguel
	valor := calcularValorAluguel(horas)

	// Imprime o resultado formatado
	fmt.Printf("O VALOR A PAGAR E = %.2f\n", valor)
}

// Função para calcular o valor do aluguel da charrete
func calcularValorAluguel(horas int) float64 {
	// Taxa para cada 3 horas de uso
	taxa := 10.0

	// Calcula o número de blocos de 3 horas
	blocos := horas / 3

	// Calcula o número de horas restantes após completar os blocos de 3 horas
	horasRestantes := horas % 3

	// Calcula o valor do aluguel baseado nos blocos de 3 horas
	valor := float64(blocos) * taxa

	// Adiciona o valor correspondente às horas restantes
	if horasRestantes > 0 {
		valor += float64(horasRestantes) * 5.0
	}

	return valor
}
