package funcoes

import "fmt"

func FuncaoPublica() {
	fmt.Println("Função Pública")
	resultado := soma(2, 10)
	fmt.Println(resultado)

	Multiplica := func(x int) int { 
		return x * 2 // Função Anônima, visivel somente dentro desse escopo!
	}
	resultadoMulti := Multiplica(2)
	fmt.Println("O resultado da multiplicação é: ", resultadoMulti)
}

func soma(a, b int) int {
	return a + b
}
