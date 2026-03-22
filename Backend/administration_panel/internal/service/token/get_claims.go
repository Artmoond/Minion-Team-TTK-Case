package token

import (
	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/custom_err"
	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	jwt.RegisteredClaims
	Roles []string `json:"roles"`
}

func (t *tokenService) GetClaims(tokenString string) (*MyClaims, error) {
	claims := &MyClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.secret), nil
	})

	if err != nil {
		return nil, custom_err.ErrGetClaims
	}

	if !token.Valid {
		return nil, custom_err.ErrTokenInvalid
	}

	return claims, nil
}
