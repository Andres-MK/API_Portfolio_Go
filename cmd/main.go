package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Andres-MK/internal/api"
	"github.com/Andres-MK/internal/application/services"
	apiServices "github.com/Andres-MK/internal/infrastructure/api"
	dbActions "github.com/Andres-MK/internal/infrastructure/db"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	

	// Load Repositories
	emailRepo := dbActions.NewEmailRepository()
	emailApiRepo := apiServices.NewApiEmailRepository()

	// Load Services
	productService := services.NewEmailService(emailRepo, emailApiRepo)

	r := mux.NewRouter()

	// Load Controllers
	api.NewEmailController(r, productService)

	// Iniciar servidor
    fmt.Println("Iniciando el servidor en el puerto 8080...")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Fallo al iniciar servidor: %v", err)
    }
}