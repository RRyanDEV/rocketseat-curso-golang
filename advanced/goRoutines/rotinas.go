package goroutines

import (
	"fmt"
	"time"
)

func ExibirMsg() {
	fmt.Println("Eu sou uma goRoutine!")
}

func Routines() {
	go ExibirMsg()              // Go: usado para indicar que é a função é uma goRoutine
	time.Sleep(1 * time.Second) // Função que faz com que o programe aguarde para rodar a outra função
	/*
		goRoutine principal é a **Main. Caso ela se encerre antes de acontecer o outro processo
		ela se encerra.
	*/
	fmt.Println("Olá")
	fmt.Println("===================")
}

func sayHello() {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond) /* Bloqueia, não deixando ir para o proximo for, enquanto não acontecer
		o processo completo
		*/
	}
}

func sayWorld() {
	for i := 0; i < 3; i++ {
		fmt.Println("World")
		time.Sleep(150 * time.Millisecond) /* Bloqueia, não deixando ir para o proximo for, enquanto não acontecer
		o processo completo
		*/
	}
}

func NoneOrder() {
	/*
		função principal pode encerrar antes das Goroutines, resultando em uma execução concorrente
	*/
	go sayHello()
	go sayWorld()
	time.Sleep(1 * time.Second)
}
