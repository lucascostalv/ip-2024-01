package main

import (
    "fmt"
)

func main() {
    // Declaração das variáveis para armazenar as notas
    var nota1, nota2, nota3 float64

    // Solicita ao usuário para inserir as notas
    fmt.Println("Digite as três notas do aluno separadas por espaço:")

    // Lê as notas fornecidas pelo usuário
    fmt.Scanln(&nota1, &nota2, &nota3)

    // Calcula a média aritmética
    media := (nota1 + nota2 + nota3) / 3.0

    // Imprime a média com duas casas decimais
    fmt.Printf("MEDIA = %.2f\n", media)

    // Verifica se o aluno está aprovado ou reprovado
    if media >= 6.0 {
        fmt.Println("APROVADO")
    } else {
        fmt.Println("REPROVADO")
    }
}
