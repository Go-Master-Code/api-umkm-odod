package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret")

// isi: func GenerateToken dan ValidateToken
func GenerateToken(userID, username, roleID, role, tenantID string) (string, error) {
	//create new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   userID,
		"username":  username,
		"role_id":   roleID,
		"role":      role,
		"tenant_id": tenantID,
		"iat":       time.Now().Unix(),                    // issued at
		"exp":       time.Now().Add(time.Hour * 1).Unix(), // harus exp bukan expired, token expired dalam 1 jam
	})
	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		return secretKey, nil
	})
}
