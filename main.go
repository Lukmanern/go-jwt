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

// createToken generates a JWT token with the given
// signing key, username, and expiration time.
func createToken(signingKey []byte, username string, expiration time.Time) (string, error) {
	// Define the claims to be included in the token.
	claims := jwt.MapClaims{
		usernameKey:   username,          // The username of the token's owner.
		expirationKey: expiration.Unix(), // The Unix timestamp when the token will expire.
	}

	// Create a new token with the specified
	// claims and signing algorithm.
	token := jwt.NewWithClaims(signingAlgorithm, claims)

	// Sign the token with the given key
	// and return the resulting string.
	return token.SignedString(signingKey)
}

// verifyToken verifies the authenticity and
// validity of a JWT token using the given signing key.
func verifyToken(tokenString string, signingKey []byte) (jwt.MapClaims, error) {
	// Parse the token string and verify
	// its signature using the provided signing key.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != signingAlgorithm {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	// Check if the token is valid and
	// contains the expected claims.
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	// Return an error if the token
	// is invalid or missing expected claims.
	return nil, fmt.Errorf("invalid token")
}
