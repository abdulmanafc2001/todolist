package main

import (
	"log"

	"github.com/abdulmanafc2001/todolist/pkg/database"
	"github.com/abdulmanafc2001/todolist/pkg/handlers"
	"github.com/abdulmanafc2001/todolist/pkg/repository"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := database.ConnectToDB()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to db", db.Name())
	repo := repository.NewRepository(db)

	handlers := handlers.NewHandler(repo)

	handlers.Run()
}
