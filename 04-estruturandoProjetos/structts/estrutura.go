package structts

import "fmt"

// Fundamental para criar e organizar tipos de dados.
// **Funciona como se fosse classe na P.O.O(Programação Orientada a Objetos).
// Criação de instâncias, como acessar e como modifcar seus valores.

type Cliente struct {
	Nome     string
	Idade    int
	Endereco Endereco
	Email    string
}

type Endereco struct {
	Rua    string
	Numero int
	Cep    string
	Estado string
}

func EstruturaDeClient() {
	cliente1 := Cliente{
		Nome:  "Ryan",
		Idade: 22,
		Endereco: Endereco{
			Rua:    "Rua Aspirante Maso",
			Numero: 394,
			Cep:    "26900-000",
			Estado: "RJ",
		},
	}
	cliente2 := Cliente{
		Nome:  "Luis Miguel",
		Idade: 20,
	}
	cliente2.Endereco.Numero = 144
	cliente2.Email = "google@google.com"

	fmt.Println(cliente1.Nome, cliente1.Idade, cliente1.Endereco.Rua)
	fmt.Println(cliente2.Nome, cliente2.Idade, cliente2.Email, cliente2.Endereco.Numero)
}
