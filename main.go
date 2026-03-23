package main

import (
	"ip-finder/app"
	"log"
	"os"
)

func main() {
	applicacao := app.Generate()
	if erro := applicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
