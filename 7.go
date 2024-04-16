package main

import (
	"fmt"
)

func main() {
	// Declaração das variáveis para armazenar as entradas
	var fahrenheit, polegadas float64

	// Solicita ao usuário para inserir a temperatura em Fahrenheit
	fmt.Println("Digite a temperatura em Fahrenheit:")
	fmt.Scanln(&fahrenheit)

	// Solicita ao usuário para inserir a quantidade de chuva em polegadas
	fmt.Println("Digite a quantidade de chuva em polegadas:")
	fmt.Scanln(&polegadas)

	// Converte a temperatura de Fahrenheit para Celsius
	celsius := (5*fahrenheit - 160) / 9

	// Converte a quantidade de chuva de polegadas para milímetros
	milimetros := polegadas * 25.4

	// Imprime os resultados formatados
	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", celsius)
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f\n", milimetros)
}
