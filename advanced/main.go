package main

import (
	"fmt"

	"github.com/rryandev/rocketseat-curso-go/advanced/channels"
	"github.com/rryandev/rocketseat-curso-go/advanced/goRoutines"
	"github.com/rryandev/rocketseat-curso-go/advanced/pointers"
)

func main() {
	pointers.Pointt()
	fmt.Println("==============================================")
	goroutines.Routines()
	goroutines.NoneOrder()
	fmt.Println("==============================================")
	channels.Canais()
	channels.Concorrentes()
	fmt.Println("==============================================")
}
