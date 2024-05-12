package util

import (
	"crypto/sha256"
	"encoding/hex"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// Fungsi untuk membuat access token
func GenerateAccessToken(data interface{}, typeToken string) (string, error) {
	// Buat claims JWT
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["data"] = data
	claims["type"] = typeToken

	// Buat token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("ACCESS_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func HashString(str string) string {
	data := []byte(str)
	hash := sha256.New()

	// Tulis data ke hash
	hash.Write(data)

	// Dapatkan hasil hash dalam bentuk byte array
	hashBytes := hash.Sum(nil)

	// Konversi byte array hash ke string heksadesimal
	hashString := hex.EncodeToString(hashBytes)

	return hashString
}
