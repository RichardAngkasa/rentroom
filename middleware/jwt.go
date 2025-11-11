package middleware

import (
	"errors"
	"rentroom/utils"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(token string, allowedRoles ...string) (jwt.MapClaims, error) {
	if token == "" {
		return nil, errors.New("missing token")
	}
	claims, err := utils.ParseJWT(token)
	if err != nil {
		return nil, errors.New("invalid token")
	}
	role, ok := claims["role"].(string)
	if !ok {
		return nil, errors.New("invalid role")
	}
	for _, r := range allowedRoles {
		if role == r {
			return claims, nil
		}
	}
	return claims, errors.New("unauthorized")
}
