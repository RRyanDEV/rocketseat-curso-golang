package methods

import "fmt"

// **Métodos em Go, são funções associadas a structs
/* Diferença entre métodos que recebem uma cópia da
struct e aqeueles que utilizam ponteiro (*)
*/
// Como criar métodos e permitindo interação e modificação.
type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) Apresentar() {
	/* Aqui é feito uma cópia da struct, sem alteração do valor
	definido pelo usuário
	*/
	fmt.Printf("Olá, meu nome é %s e tenho %d anos. \n", p.Nome, p.Idade)
}
func (p *Pessoa) Apresentarr() {
	/* Ele recebe a referência da struct original, sem os valores
	definidos.
	*/
	p.Nome = "Ricardo"
	fmt.Printf("Olá, meu nome é %s e tenho %d anos. \n", p.Nome, p.Idade)
}

func MostrarPessoa() {
	p1 := Pessoa{Nome: "Ryan", Idade: 22}
	p1.Apresentar()
	p2 := Pessoa{Nome: "Luis Miguel", Idade: 20} /* A struct original não é modificada*/
	p2.Apresentarr()
}
