package main

import (
	"fmt"
	"time"

	"github.com/austintraver/loading"
	"github.com/cbluth/randomart"
	"github.com/lerichardson/isbensus/utils"
	"github.com/logrusorgru/aurora"
)

func main() {
	fmt.Println("Isbensus - designed by Levi Richardson in Seattle.")
	wheel := loading.New("%s Initializing isbensus")
	wheel.Start()
	time.Sleep(time.Second * 5)
	wheel.Stop()
	fmt.Println("Generating randomart image for neural network...")
	wheel2 := loading.New("%s Generating randomart image")
	wheel2.Start()
	time.Sleep(time.Second * 2)
	wheel2.Stop()
	fmt.Println("\n" + randomart.RandomArt(utils.RandomBytes(32), "./bensus.png"))
	fmt.Println("Inititalizing neural network based on past tweets...")
	utils.LoadBar()
	wheel3 := loading.New("%s Parsing result")
	wheel3.Start()
	time.Sleep(time.Second * 2)
	wheel3.Stop()
	fmt.Println("Result loading...")
	utils.LoadBar()
	fmt.Println("\nResult: ")
	fmt.Println(aurora.Bold(aurora.Red("Ben is always sus, never not sus")))
}
