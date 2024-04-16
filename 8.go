package main

import (
	"fmt"
	"math"
)

func main() {
	// Declaração das variáveis para armazenar o raio e a altura da lata
	var raio, altura float64

	// Solicita ao usuário para inserir o raio da lata
	fmt.Println("Digite o raio da lata (em metros):")
	fmt.Scanln(&raio)

	// Solicita ao usuário para inserir a altura da lata
	fmt.Println("Digite a altura da lata (em metros):")
	fmt.Scanln(&altura)

	// Calcula a área do círculo
	areaCirculo := math.Pi * raio * raio

	// Calcula a área lateral do cilindro
	areaLateral := 2 * math.Pi * raio * altura

	// Calcula o valor total da área da lata
	areaTotal := 2*areaCirculo + areaLateral

	// Calcula o custo da lata
	custo := areaTotal * 100.0

	// Imprime o resultado formatado
	fmt.Printf("O VALOR DO CUSTO E = %.2f\n", custo)
}
