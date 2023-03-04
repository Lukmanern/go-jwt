package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var signingAlgorithm = jwt.SigningMethodHS256

const (
	usernameKey   = "username"
	expirationKey = "exp"
)

func main() {
	// Set the signing key
	signingKey := []byte("my-secret-key")

	tokenString, err := createToken(signingKey, "johndoe", time.Now().Add(time.Hour*24))
	if err != nil {
		fmt.Println("Error creating token:", err)
		return
	}

	claims, err := verifyToken(tokenString, signingKey)
	if err != nil {
		fmt.Println("Error verifying token:", err)
		return
	}
	fmt.Println("Username:", claims[usernameKey])
	fmt.Println("Expiration Time:", time.Unix(int64(claims[expirationKey].(float64)), 0))

	// Create a new token with a claims payload
	claims = jwt.MapClaims{
		"username": "johndoe",
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(signingAlgorithm, claims)

	// Sign the token with the secret key
	tokenString, err = token.SignedString(signingKey)
	if err != nil {
		fmt.Println("Error signing token:", err)
		return
	}

	// Print the token string
	fmt.Println("Token:", tokenString)

	// Verify the token
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != signingAlgorithm {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
	if err != nil {
		fmt.Println("Error parsing token:", err)
		return
	}

	// Print the token claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		fmt.Println("Invalid token")
		return
	}
	fmt.Println("Username:", claims[usernameKey])
	fmt.Println("Expiration Time:", time.Unix(int64(claims[expirationKey].(float64)), 0))
}

func createToken(signingKey []byte, username string, expiration time.Time) (string, error) {
	claims := jwt.MapClaims{
		usernameKey:   username,
		expirationKey: expiration.Unix(),
	}
	token := jwt.NewWithClaims(signingAlgorithm, claims)
	return token.SignedString(signingKey)
}

func verifyToken(tokenString string, signingKey []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != signingAlgorithm {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
