package token

import "time"

type TokenService interface {
	GenerateToken(userId int64, role string) (string, error)
}
type tokenService struct {
	secret string
	ttl    time.Duration
}

func NewTokenService(secret string, ttl time.Duration) TokenService {
	return &tokenService{
		secret: secret,
		ttl:    ttl,
	}
}
