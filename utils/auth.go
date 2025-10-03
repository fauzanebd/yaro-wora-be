package utils

import (
	"errors"
	"time"
	"yaro-wora-be/config"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaim struct {
	Username string `json:"username"`
	UserID   uint   `json:"user_id"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a JWT token for the user
func GenerateJWT(username string, userID uint, role string) (string, error) {
	claims := JWTClaim{
		Username: username,
		UserID:   userID,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.AppConfig.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT validates a JWT token and returns the claims
func ValidateJWT(tokenString string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// SimpleAuthResponse represents a simple authentication response
type SimpleAuthResponse struct {
	Success bool      `json:"success"`
	Token   string    `json:"token,omitempty"`
	Message string    `json:"message"`
	User    *AuthUser `json:"user,omitempty"`
}

// AuthUser represents user data in auth response
type AuthUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
