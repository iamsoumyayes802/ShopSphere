package security

import (
	"auth-service/internal/model"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user *model.User) (string, error) {
	keyBytes, err := os.ReadFile(os.Getenv("JWT_PRIVATE_KEY_PATH"))

	if err != nil {
		return "", err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyBytes)

	if err != nil {
		return "", err
	}

	now := time.Now()

	claims := jwt.RegisteredClaims{
		Subject:   user.ID,
		Issuer:    "shopsphere-auth",
		Audience:  jwt.ClaimStrings{"shopsphere"},
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(15 * time.Minute)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	token.Claims = jwt.MapClaims{
		"sub": claims.Subject,
		"iss": claims.Issuer,
		"aud": claims.Audience,
		"iat": claims.IssuedAt.Unix(),
		"exp": claims.ExpiresAt.Unix(),
		"role": user.Role,
	}

	return token.SignedString(privateKey)
}

func GeneratePasswordResetToken() (string, error) {
	b := make([]byte, 32)

	if 	_, err := rand.Read(b); err != nil {
		return "", nil
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

func HashResetToken(token string) string {
	sum := sha256.Sum256([]byte(token))
    return hex.EncodeToString(sum[:])
}

