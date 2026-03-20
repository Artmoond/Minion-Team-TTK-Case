package token

import (
	"log"

	"github.com/Artmoond/Minion-Team-TTK-Case/internal/entity/custom_err"
	"github.com/golang-jwt/jwt/v4"
)

func (t *tokenService) GenerateToken(userId int64, role string) (string, error) {
	claims := jwt.MapClaims{"user_id": userId, "role": role, "exp": t.ttl}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(t.secret)
	if err != nil {
		log.Println("Error while signing token: " + err.Error())
		return "", custom_err.ErrParseTokenToString
	}

	return tokenString, nil
}
