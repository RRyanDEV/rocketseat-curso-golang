package switchcase

import (
	"fmt"
	"time"
)

func TrocaECasos() {
	fmt.Println("Quando é Domingo?")
	today := time.Now().Weekday()
	switch time.Sunday /*(Expressão)*/ {
	case today + 0:
		fmt.Println("É hoje")
	case today + 1:
		fmt.Println("É amanhã")
	case today + 2:
		fmt.Println("É em dois dias")
	default:
		fmt.Println("Ta longe ainda!")
	}
}
