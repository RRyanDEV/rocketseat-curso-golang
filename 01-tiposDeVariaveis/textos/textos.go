package textos

import (
	"fmt"
	"strings"
)

func Textos() {
	const nome string = "Ryan" //Segue o padrão de não poder ser mutável durante a execução do código.
	var text = "Anos"
	var idade = 21

	question := "Como vai?"

	var meet = "Olá" + " " + nome + " " + question

	fmt.Println(strings.ToUpper(meet))
	fmt.Println(nome, idade, text)
}
