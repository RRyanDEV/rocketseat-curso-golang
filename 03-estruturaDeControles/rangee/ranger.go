package rangee

import "fmt"

func Range() {
	nums := []int{1, 2, 3}
	users := []string{"Ryan", "Miguel", "Dex"}

	for key, value := range nums {
		fmt.Println(key, value)
	}
	for key, value := range users {
		fmt.Println(key, value)
	}
}
