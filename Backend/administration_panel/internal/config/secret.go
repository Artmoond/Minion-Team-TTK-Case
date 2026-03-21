package config

import (
	"log"
	"os"

	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/custom_err"
)

type TokenSecret interface {
	Secret() string
}

type tokenSecret struct {
	secret string
}

func NewTokenSecret() (*tokenSecret, error) {
	secret := os.Getenv("TOKEN_SECRET")
	if len(secret) == 0 {
		log.Println("TOKEN_SECRET environment variable not set")
		return nil, custom_err.ErrNilSecret
	}

	return &tokenSecret{secret: secret}, nil
}

func (t *tokenSecret) Secret() string {
	return t.secret
}
