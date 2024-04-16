package main

import (
	"fmt"
)

func main() {
	var salarioMinimo, consumo float64

	// Solicita ao usuário para inserir o valor do salário mínimo e a quantidade de kW gasta pela residência
	fmt.Println("Digite o valor do salário mínimo:")
	fmt.Scanln(&salarioMinimo)
	fmt.Println("Digite a quantidade de kW gasta pela residência:")
	fmt.Scanln(&consumo)

	// Calcula o valor em reais de cada kW
	valorKw := (70.0 / 100.0) * salarioMinimo / 100.0

	// Calcula o valor em reais a ser pago pelo consumo da residência
	custoConsumo := valorKw * consumo

	// Calcula o novo valor a ser pago pela residência com um desconto de 10%
	custoDesconto := 0.9 * custoConsumo

	// Imprime os resultados formatados
	fmt.Printf("Custo por kW: R$ %.2f\n", valorKw)
	fmt.Printf("Custo do consumo: R$ %.2f\n", custoConsumo)
	fmt.Printf("Custo com desconto: R$ %.2f\n", custoDesconto)
}
