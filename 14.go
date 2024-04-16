package main

import (
	"fmt"
	"math"
)

func main() {
	// Declaração das variáveis para armazenar a altura e o comprimento da aresta do hexágono
	var altura, aresta float64

	// Solicita ao usuário para inserir a altura da pirâmide e o comprimento da aresta do hexágono
	fmt.Println("Digite a altura da pirâmide e o comprimento da aresta do hexágono, separados por espaço:")
	fmt.Scanln(&altura, &aresta)

	// Calcula a área da base do hexágono
	areaBase := calcularAreaBase(aresta)

	// Calcula o volume da pirâmide
	volume := (1.0 / 3.0) * areaBase * altura

	// Imprime o resultado formatado
	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS\n", volume)
}

// Função para calcular a área da base do hexágono
func calcularAreaBase(aresta float64) float64 {
	return (3.0 * math.Pow(aresta, 2) * math.Sqrt(3)) / 2.0
}
