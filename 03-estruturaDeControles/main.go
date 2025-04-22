package main

import (
	"fmt"

	"github.com/rryandev/rocketseat-curso-go/03-estruturaDeControles/forr"
	ifElse "github.com/rryandev/rocketseat-curso-go/03-estruturaDeControles/if-else"
	"github.com/rryandev/rocketseat-curso-go/03-estruturaDeControles/rangee"
	switchCase "github.com/rryandev/rocketseat-curso-go/03-estruturaDeControles/switch-case"
)

func main() {
	fmt.Println("============")
	ifElse.SeSenão()
	ifElse.AcessoMapUsingIF()
	fmt.Println("============")
	switchCase.TrocaECasos()
	fmt.Println("============")
	forr.LacoFor()
	forr.While()
	forr.ForInSlice()
	fmt.Println("============")
	rangee.Range()
}
