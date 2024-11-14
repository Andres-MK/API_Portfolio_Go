package main

import (
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
	dsn := "server=localhost;database=portfolio;user=kraaash;password=123456;port=1433;encrypt=disable"
	gormDB, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Fallo al conectar a la base de datos: %v", err)
	}
	

	// Load Repositories
	emailRepo := dbActions.NewEmailRepository(gormDB)
	emailApiRepo := apiServices.NewApiEmailRepository()

	// Load Services
	productService := services.NewEmailService(emailRepo, emailApiRepo)

	r := mux.NewRouter()

	// Load Controllers
	api.NewEmailController(r, productService)

	// Iniciar servidor
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Fallo al iniciar servidor: %v", err)
    }
}