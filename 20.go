package main

import "fmt"

func main() {
	// Declaração das variáveis para armazenar as horas, minutos e segundos
	var horas, minutos, segundos int

	// Solicita ao usuário para inserir os valores de horas, minutos e segundos
	fmt.Println("Digite o valor das horas:")
	fmt.Scanln(&horas)
	fmt.Println("Digite o valor dos minutos:")
	fmt.Scanln(&minutos)
	fmt.Println("Digite o valor dos segundos:")
	fmt.Scanln(&segundos)

	// Calcula o tempo total em segundos
	tempoEmSegundos := calcularTempoEmSegundos(horas, minutos, segundos)

	// Imprime o resultado
	fmt.Printf("O TEMPO EM SEGUNDOS E = %d\n", tempoEmSegundos)
}

// Função para calcular o tempo total em segundos
func calcularTempoEmSegundos(horas, minutos, segundos int) int {
	totalSegundos := horas*3600 + minutos*60 + segundos
	return totalSegundos
}
