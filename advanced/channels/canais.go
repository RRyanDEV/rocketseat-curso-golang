package channels

import (
	"fmt"
	"time"
)

/*
O Channel é como se fosse um tubo ele faz conexão entre as goRoutines
Assim sendo a forma mais prática de enviar valores entre goRoutines
*/

func Canais() {
	ch := make(chan int, 3) // Variavel que só recebe inteiros, e difinido o tamanho dedo canal(3)

	/*
		ch <- 9 Forma de atribuir valores em um canal
		<-ch    Forma para ler valores de um canal
	*/
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // Sempre encerrar um canal, para evitar deadlock
		fmt.Println("Finalizamos a Escrita!")
	}()
	time.Sleep(time.Second * 2)
	// <-ch (Tira o valor que foi inserido, funcionando como uma fila de prioridade)
	for valor := range ch {
		fmt.Println("Valor do canal:", valor)
	}
	fmt.Println("========================")
}

func Concorrentes() {
	chann := make(chan int)
	go producer(chann)
	go consumer(chann)
	/*
	Não da para garantir a ordem que as funções serão executadas
	Da pra utilizar mais de uma função que faça a mesma coisa
	*/
	go consumer(chann)
	go consumer(chann)
	time.Sleep(time.Second * 1)
}

func producer(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
}
func consumer(c chan int) {
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("Consumer Finalizado!")
}
