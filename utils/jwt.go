package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}

func GenerateJWT(userId uint) (string, error) {
    LoadEnv()
    
    token := jwt.New(jwt.SigningMethodHS256)
    
    claims := token.Claims.(jwt.MapClaims)
    claims["userId"] = userId
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours
    
    jwtSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }
    
    return tokenString, nil
}
