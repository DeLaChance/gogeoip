package main

import (
	"fmt"
	"go.uber.org/fx"
	
	"config"
)

func main() {
	fx.New(
		fx.Invoke(initializeApp),
		fx.Invoke(config.StartRestController),
	).Run()
}

func initializeApp() {
	fmt.Println("Initializing app")
}