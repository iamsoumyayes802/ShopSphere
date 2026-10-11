package server

import (
	"net/http"
	"log"
	"os"
	"product-service/internal/handler"
	"product-service/internal/repository"
	"product-service/internal/service"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := repository.NewPostgresPool(databaseURL)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	productRepo := repository.NewProductRepository(db)

	productService := service.NewProductService(productRepo)

	productHandler := handler.NewProductHandler(productService)

	http.HandleFunc("/products", productHandler.Create)

	log.Println("Auth service is running on: 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}