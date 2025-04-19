package variaveis

import "fmt"

func Booleanos() {
	// Por padrão ele já vem como "False".
	maior := 10 > 5
	menor := 10 < 5

	fmt.Println("10 é maior que 5? ", maior)
	fmt.Println("10 é menor que 5?", menor)
}
