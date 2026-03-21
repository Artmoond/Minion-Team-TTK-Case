package token

type TokenService interface {
	GetClaims(tokenString string) (*MyClaims, error)
}

type tokenService struct {
	secret string
}

func NewTokenService(secret string) TokenService {
	return &tokenService{secret}
}
