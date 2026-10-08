package main

import (
	"auth-service/internal/handler"
	"auth-service/internal/middleware"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"log"
	"net/http"
	"os"
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

	userRepo := repository.NewUserRepository(db)

	authService := service.NewAuthService(userRepo)

	authHandler := handler.NewAuthHandler(authService)

	http.HandleFunc("/signup", authHandler.Signup)
	http.HandleFunc("/signin", authHandler.Signin)

	meHandler := http.HandlerFunc(handler.GetMe)

	http.Handle(
		"/users/me",
		middleware.Auth(meHandler),
	)

	// admin 
	// http.Handle(
	// 	"/admin/test",
	// 	middleware.Auth(
	// 		middleware.RequireRole("ADMIN")(
	// 			http.HandlerFunc(adminHandler),
	// 		),
	// 	),
	// )	

	log.Println("Auth service is running on: 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
