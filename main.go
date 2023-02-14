package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	// Set the signing key
	signingKey := []byte("my-secret-key")

	// Create a new token with a claims payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "johndoe",
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		fmt.Println("Error signing token:", err)
		return
	}

	// Print the token string
	fmt.Println("Token:", tokenString)

	// Verify the token
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		fmt.Println("Error parsing token:", err)
		return
	}

	// Print the token claims
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		fmt.Println("Username:", claims["username"])
		fmt.Println("Expiration Time:", time.Unix(int64(claims["exp"].(float64)), 0))
	} else {
		fmt.Println("Invalid token")
	}
}
