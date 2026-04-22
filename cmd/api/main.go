package main

import (
	"log"

	"github.com/kamilrahmatullin/mama-ryadom-mvp/internal/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("server started on :8080")
	log.Fatal(a.Run(":8080"))
}
