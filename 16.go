package main

import "fmt"

func main() {
	// Declaração da variável para armazenar o salário do funcionário
	var salario float64

	// Solicita ao usuário para inserir o salário do funcionário
	fmt.Println("Digite o salário do funcionário:")
	fmt.Scanln(&salario)

	// Calcula o salário reajustado de acordo com as regras
	var salarioReajustado float64
	if salario <= 300.00 {
		salarioReajustado = salario * 1.5 // Reajuste de 50%
	} else {
		salarioReajustado = salario * 1.3 // Reajuste de 30%
	}

	// Imprime o resultado formatado
	fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", salarioReajustado)
}
