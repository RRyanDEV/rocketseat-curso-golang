package variaveis

import "fmt"

// **Int8:** Inteiro de 8 Bits ( valores de -128 a 127)
// **Int16:**: Inteiro de 16 Bits ( valores de -32768 a 32767)
// **Int32:** Inteiro de 32 Bits ( valores de -2147483648 a 2147483647)
// **Int64:** Inteiro de 64 Bits ( valores de -2147483648 a 2147483647)

func ExibirNumeros() {
	Inteiros()
	Flutuantes()
}

func Inteiros() {
	var idade int = 21
	var contador int32 = 160
	var velocidadeDaLuz int64 = 299793458

	fmt.Println("Idade:", idade)
	fmt.Println("Contador:", contador)
	fmt.Println("Velocidade da Luz:", velocidadeDaLuz)
	fmt.Println("------")
}

func Flutuantes() {
	// var floatNumber float32 = 4.444
	var pi float64 = 3.14
	var raio float64 = 2.5
	var area = pi * raio * raio

	// fmt.Println("Número Flutuante:", floatNumber)
	fmt.Println("Área do Círculo:", area)
}
