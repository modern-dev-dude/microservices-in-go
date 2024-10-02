package dto

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/modern-dev-dude/microservices-in-go/auth_server/pkg/domain"
)

type RefreshTokenRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (r RefreshTokenRequest) IsAccessTokenValid() error {

	_, err := jwt.Parse(r.AccessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(domain.HMAC_SAMPLE_SECRET), nil
	})

	if err != nil {
		return err
	}
	return nil
}
