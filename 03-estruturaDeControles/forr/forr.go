package forr

import "fmt"

func LacoFor() {
	// **FOR** inicialização; expressão; fim iteração;
	sum := 0
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		sum += i
	}
}

func While() {
	//While "não existe" em Go
	// Para estanciar um While, utilizar o FOR
	num := 0
	for num < 20 {
		fmt.Println(`Loop do "While"`)
		num += 1
	}
	fmt.Println(num)
}

func ForInSlice() {
	nums := []int{1, 2, 3, 4, 5}
	
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}
