package main

import "fmt"

func main() {
	// Declaração da variável para armazenar a nota do aluno
	var nota float64

	// Solicita ao usuário para inserir a nota do aluno
	fmt.Println("Digite a nota do aluno:")
	fmt.Scanln(&nota)

	// Determina o conceito correspondente à nota
	conceito := determinarConceito(nota)

	// Imprime o resultado formatado
	fmt.Printf("NOTA = %.1f CONCEITO = %s\n", nota, conceito)
}

// Função para determinar o conceito correspondente à nota
func determinarConceito(nota float64) string {
	switch {
	case nota >= 9.0 && nota <= 10.0:
		return "A"
	case nota >= 7.5 && nota < 9.0:
		return "B"
	case nota >= 6.0 && nota < 7.5:
		return "C"
	default:
		return "D"
	}
}
