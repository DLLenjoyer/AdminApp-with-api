package main

import (
	"log"
	"github.com/DLLenjoyer/AdminApp-with-api/config"
	"github.com/DLLenjoyer/AdminApp-with-api/routes"
	
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.SetupDatabase()
	if err != nil {
		log.Fatalf("Не удалось подключится к database: %v", err)
	}

	r := gin.Default()

	routes.RegisterRoutes(r, db)

	if err := r.Run(":8024"); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
