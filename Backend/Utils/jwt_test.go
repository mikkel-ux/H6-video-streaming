package utils

import (
	"testing"
)

var testSecretKey = []byte("test_secret_key")

func TestCreateAndValidateToken(t *testing.T) {
	secretKey = testSecretKey

	tokenString, err := CreateToken(123)
	if err != nil {
		t.Fatalf("Failed to create token: %v", err)
	}
	if tokenString == "" {
		t.Fatal("Expected non-empty token string")
	}

	token, err := ValidateToken(tokenString)
	if err != nil {
		t.Fatalf("Token should be valid, got error: %v", err)
	}
	if !token.Valid {
		t.Fatal("Expected token to be valid")
	}

	invalidToken := tokenString + "corrupted"
	_, err = ValidateToken(invalidToken)
	if err == nil {
		t.Fatal("Expected error for invalid token, got nil")
	}
}
