package ifelse

import "fmt"

func SeSenão() {
	nota := 90
	// Não utiliza os () para informar a condição
	if nota >= 90 {
		fmt.Println("Aprovado com Nota Máxima")
	} else if nota >= 65 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
	// O else, SEMPRE deve vir seguido ao braket final do IF ou ELSE-IF.
}

func AcessoMapUsingIF() {
	players := map[string]int{
		"Ryan":    90,
		"Dex":     36,
		"Rodrigo": 10,
	}

	if value, ok := players["Ryan"]; ok {
		fmt.Println("Possui: ", value, "Pontos")
	}
}
