package array

import "fmt"

func Arr() {

	// Estrutura de tamanho fixo.
	// Armaneza valores do mesmo tipo.
	// Como fazer leitura e como adicionar valores no array.

	var gavetas [2]string
	gavetas[0] = "Faca"
	gavetas[1] = "Garfo"
	// gavetas[2] = "Colher"

	fmt.Println(gavetas)
	fmt.Println(gavetas[0]) // Faca (Resultado com base no índice no Array)
	fmt.Println(gavetas[1]) // Garfo (Resultado com base no índice no Array)
}
