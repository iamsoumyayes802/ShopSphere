package password

import "golang.org/x/crypto/bcrypt"

func Hash(plainText string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)

	if err != nil {
		return "", nil
	}

	return string(hashedBytes), nil
}

func Verify(plainText, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(plainText),
	)

	return err == nil
} 