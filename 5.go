package main

import (
	"fmt"
)

func main() {
	// Variáveis para armazenar os dados de entrada
	var conta int
	var tipoConsumidor byte
	var consumo float64
	consumo = 230
	conta = 123456

	// Solicita ao usuário para inserir os dados
	fmt.Println("Digite a conta do cliente, o consumo de água (metros cúbicos) e o tipo de consumidor (R, C ou I), separados por espaço:")

	// Lê os dados fornecidos pelo usuário
	// fmt.Scanln(&conta, &consumo, &tipoConsumidor)
	tipoConsumidor = 'R'
	// Variável para armazenar o valor da conta
	var valorConta float64

	// Calcula o valor da conta com base no tipo de consumidor
	// eu faço um switch case para verificar o tipo de consumidor
	switch tipoConsumidor {
		//se for residencial, eu calculo o valor da conta com base no consumo de água
		//e a tarifa de 0.05 por metro cúbico eu multiplico pelo consumo de água e somo com 5 reais
		case 'R':
			valorConta = 5.0 + 0.05*consumo
		case 'C':
			// irei verificar o consumo de agua
			// consumo:  230
			// entao valorConta ira ser deduzido por 80
			// (consumo - 80) 
			// 230-80
			// valor da conta:150
			// e cobrado ainda 0.25 por metros cubicos: 37.5
			// valorConta = 500.0 por causa dos primeiros 80 m³ 
			// valorConta = 537.5


			if consumo <= 80{
				valorConta=500.0
			}else{
				valorConta=500.0+0.25*(consumo-80)
			}

			
		case 'I':
			
			if consumo <= 100 {
				valorConta = 800.0
			} else {
				valorConta = 800.0 + 0.04*(consumo-100)
			}
	}

	// Imprime o resultado formatado
	fmt.Printf("CONTA = %d\n", conta)
	fmt.Printf("VALOR DA CONTA = %.2f\n", valorConta)
}
/*
Desenvolver um programa para calcular a conta de água para uma empresa de saneamento. O custo da água varia dependendo se o consumidor é residencial, comercial ou industrial. A regra para calcular a conta é:
• Residencial: R$ 5, 00 de taxa mais R$ 0, 05 por metros cúbicos gastos;
• Comercial: R$ 500, 00 para os primeiros 80 metros cúbicos gastos mais R$ 0, 25 por metros cúbicos gastos;
• Industrial: R$ 800, 00 para os primeiros 100 metros cúbicos gastos mais R$ 0, 04 por metros cúbicos gastos;
O programa deverá ler a conta do cliente, o consumo de água por metros cúbicos e o tipo de consumidor ( residencial, comercial e industrial ). Como resultado, o programa deve imprimir a conta do cliente e o valor em Reais a ser pago pelo mesmo.
Entrada
O programa deverá ler uma linha na entrada contendo: a conta do cliente (um número inteiro), o con- sumo de água por metros cúbicos (float) e o tipo do consumidor (um caractere: ‘C’ - COMERCIAL, ‘I’ - INDUSTRIAL ou ‘R’ - RESIDENCIAL). Há um espaço entre os valores na linha de entrada
Saída
O programa deve imprimir duas linhas, contendo o seguinte:
• CONTA = u, onde u é o código inteiro identificador da conta;
• VALOR DA CONTA = v, onde v é o valor da conta com duas casas decimais, a ser pago pelo consumidor;
Após o valor impresso em cada linha, o programa dev imprimir o caractere de quebra de linha; ‘\n’. Os valores de v,x e w devem conter duas casas decimais.
Exemplo
Abaixo são mostrados dois exemplos de entrada e saída, mas há apenas um caso de entrada (uma linha) para esse programa. Entrada 39393939 230 C Saída CONTA = 39393939 VALOR DA CONTA = 537.50
Entrada 888 3752 I Saída CONTA = 888 VALOR DA CONTA = 946.08
7

*/