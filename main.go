package main

import (
	"log"

	"github.com/danto9/moneybag/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
