package token

import (
	"log"
	"time"

	"github.com/Artmoond/Minion-Team-TTK-Case/internal/entity/custom_err"
	"github.com/golang-jwt/jwt/v5"
)

func (t *tokenService) GenerateToken(userId int64, roles []string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"roles":   roles,
		"exp":     time.Now().Add(t.ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(t.secret))
	if err != nil {
		log.Println("Error while signing token: " + err.Error())
		return "", custom_err.ErrParseTokenToString
	}

	return tokenString, nil
}
