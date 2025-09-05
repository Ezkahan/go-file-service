package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key")

type JWTMapClaims struct {
	jwt.RegisteredClaims
	UserID    uint
	ExpiresAt int64
}

func GenerateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}

// GenerateToken generates a signed JWT for a given user ID
// func GenerateToken(userID string) (string, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"user_id": userID,
// 		"exp":     time.Now().Add(24 * time.Hour).Unix(),
// 	})
// 	return token.SignedString(jwtSecret)
// }

func GenerateToken(userId uint) (string, error) {
	var (
		err     error
		hour, _ = strconv.Atoi(os.Getenv("TOKEN_EXP_HOUR"))
	)

	claims := JWTMapClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add((time.Hour) * time.Duration(hour))),
		},
	}

	newClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := newClaim.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateToken(tokenString string) (JWTMapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTMapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return JWTMapClaims{}, err
	}

	if claims, ok := token.Claims.(*JWTMapClaims); ok && token.Valid {
		return *claims, err
	}

	return JWTMapClaims{}, fmt.Errorf("invalid token")
}

func ExtractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("authorization header required")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}
	return jwtToken[1], nil
}

func ExtractAuthID(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTMapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*JWTMapClaims); ok && token.Valid {
		if ok {
			return claims.UserID, nil
		}
	}

	return 0, errors.New("invalid token")
}
