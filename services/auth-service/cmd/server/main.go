package main

import (
	"auth-service/internal/email"
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
	passwordRepo := repository.NewPasswordRepository(db)
	emailSender := email.NewSMTPSender("localhost", "6543", "soumya", "soumya@123", "soumyadipdey802@gmail.com")

	authService := service.NewAuthService(userRepo, passwordRepo, emailSender, "http://localhost:8080/")

	authHandler := handler.NewAuthHandler(authService)

	http.HandleFunc("/signup", authHandler.Signup)
	http.HandleFunc("/signin", authHandler.Signin)

	meHandler := http.HandlerFunc(handler.GetMe)

	http.Handle(
		"/users/me",
		middleware.Auth(meHandler),
	)	

	log.Println("Auth service is running on: 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
