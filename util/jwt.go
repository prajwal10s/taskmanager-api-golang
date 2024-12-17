package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)




type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}



func GenerateJWT(userID string) (string, error) {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file, falling back to default secret")
	}

	// Fetch the secret
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set in .env")
	}

	// Create claims
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(72 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Println("Error signing token:", err)
		return "", err
	}

	log.Println("Generated token successfully")
	return signedToken, nil
}

func VerifyJWT(tokenString string) (*Claims, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}