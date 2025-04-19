package mapp

import "fmt"

func Map() {
	var pessoas = map[string]int{}

	pessoas["Ryan"] = 21
	pessoas["Miguel"] = 20
	pessoas["Dex"] = 33

	if idade, ok := pessoas["Fabio"]; ok {
		println("Pessoa Existe No MAP", idade, ok)
	} else {
		println("Pessoa Não Existe No MAP")
	}

	fmt.Println(pessoas)
	fmt.Println(pessoas["Ryan"]) // Procura em um valor espeficifico dentro do Map
	delete(pessoas, "Dex") // Função que remove valores de dentro Map
	fmt.Println(pessoas) // Resultado apôs a remoção de pessoas
}
