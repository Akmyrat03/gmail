package main

import (
	"log"
	"mail-service/handler"
	"mail-service/repository"
	"mail-service/routes"
	"mail-service/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Veritabanı bağlantısı
	db, err := sqlx.Connect("postgres", "user=postgres dbname=gmail sslmode=disable password=postgres")
	if err != nil {
		log.Fatal("Veritabanına bağlanılamadı:", err)
	}

	// Repository, Service ve Handler katmanlarını oluştur
	messageRepo := repository.NewMessageRepository(db)
	contactService := service.NewContactService(messageRepo)
	contactHandler := handler.NewContactHandler(contactService)

	// Gin engine ve rotaları başlat
	r := gin.Default()
	routes.SetupRoutes(r, contactHandler)

	log.Println("Server çalışıyor: http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
