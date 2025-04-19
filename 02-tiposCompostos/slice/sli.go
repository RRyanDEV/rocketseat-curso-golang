package slice

import "fmt"

func Sli() {
	// Possui um tamanho fléxivel, podendo aumentar e diminuir. Dependendo da necessidade
	var gavetas []string
	gavetas = append(gavetas, "Colher", "Panos", "Pratos") //Acrensenta valores dentro do Slices 

	fmt.Println(gavetas)
	fmt.Println(len(gavetas)) // Verificar o tamanho do Slice.

	//slice[x:x-1]
	fmt.Println(gavetas[1:2]) // Ele verificar uma parte de dentro do slice.
	fmt.Println(gavetas[2])
	gavetas = gavetas[:2]
	fmt.Println("Resultado após remover um valor do slice: ", gavetas)	
}