package pointers

import "fmt"

type Pessoa struct {
	Nome string
}

func Pointt() {
	p1 := Pessoa{Nome: "Ryan"}
	p2 := Pessoa{Nome: "Miguel"}

	var p3 *Pessoa = &p1
	/* Altero o valor de p1, pois estou passando o local na memória.
	Passando valor por referência
	Apontando o endereço na memória ai é possivel fazer alteração
	*/
	p3.Nome = "Vanessa"

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}
