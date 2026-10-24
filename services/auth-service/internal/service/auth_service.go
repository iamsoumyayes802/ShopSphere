package service

import (
	"auth-service/internal/model"
	"auth-service/internal/repository"
	"auth-service/internal/security"
	"context"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
	passwordRepo *repository.PasswordRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Signup(
	ctx context.Context,
	email string,
	password string,
) (*model.User, error) {
	email = strings.ToLower(strings.TrimSpace(email))

	if email == "" || password == "" {
		return nil, errors.New("Email and Password are required")
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID: uuid.New().String(),
		Email: email,
		PasswordHash: string(hash),
		Role: "USER",
	}

	if err := s.userRepo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func generateJWT(user *model.User) (string, error) {
	secret := os.Getenv("JWT_SECRET_PRIVATE_KEY")

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role": user.Role,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		claims,
	)

	return token.SignedString([]byte(secret))
}

func (s *AuthService) Signin(
	ctx context.Context,
	email string,
	password string,
) (string, error) {
	email = strings.ToLower(strings.TrimSpace(email))

	user, err := s.userRepo.FindByEmail(
		ctx,
		email,
	)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	return generateJWT(user)
}

func (s *AuthService) ForgotPassword(
	ctx context.Context,
	email string,
) error {
	user, err := s.userRepo.FindByEmail(ctx, email)

	if err != nil {
		return err
	}

	token, err := security.GeneratePasswordResetToken()

	if err != nil {
		return err
	}

	hashedToken := security.HashResetToken(token)
	expiresAt := time.Now().UTC().Add(15 * time.Minute)

	if err := s.passwordRepo.CreateResetToken(ctx, user.ID, hashedToken, expiresAt,); err != nil {
		return err
	}

	// resetLink := "https://your-app.example/reset-password?token=" + token

}
