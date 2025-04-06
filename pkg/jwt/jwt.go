package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewToken(claims map[string]any, secret string, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	for k, v := range claims {
		token.Claims.(jwt.MapClaims)[k] = v
	}

	token.Claims.(jwt.MapClaims)["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("can't signed token: %w", err)
	}

	return tokenString, nil
}

func ValidateToken(tokenString, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			expirationTime := time.Unix(int64(exp), 0)
			if expirationTime.Before(time.Now()) {
				return nil, fmt.Errorf("token has expired")
			}
		}

		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token or expired")
	}
}
