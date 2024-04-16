package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Digite o número de temperaturas em Fahrenheit:")
	fmt.Scanln(&n)

	// Loop para cada temperatura
	for i := 0; i < n; i++ {
		var fahrenheit float64
		fmt.Println("Digite a temperatura em Fahrenheit:")
		fmt.Scanln(&fahrenheit)

		// Calcula a temperatura em Celsius
		celsius := (fahrenheit - 32) * 5 / 9

		// Imprime o resultado formatado
		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", fahrenheit, celsius)
	}
}
