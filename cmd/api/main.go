package main

import (
	"log"

	"github.com/kamilrahmatullin/mama-ryadom-mvp/internal/db"
)

func main() {
	database, err := db.NewSQLiteDB()
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	log.Println("database connected")

	_ = database
}
