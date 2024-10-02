package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/modern-dev-dude/microservices-in-go/auth_server/pkg/domain"
	"github.com/modern-dev-dude/microservices-in-go/auth_server/pkg/dto"
	"log"
)

type AuthService interface {
	Login(request dto.LoginRequest) (*dto.LoginResponse, error)
	Verify(urlParams map[string]string) error
	Refresh(request dto.RefreshTokenRequest) (*dto.LoginResponse, error)
}

type DefaultAuthService struct {
	repo            domain.AuthRepository
	rolePermissions domain.RolePermissions
}

func NewLoginService(repo domain.AuthRepository, permissions domain.RolePermissions) DefaultAuthService {
	return DefaultAuthService{repo, permissions}
}

func (s DefaultAuthService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	login, err := s.repo.FindBy(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	claims := login.ClaimsForAccessToken()
	authToken := domain.NewAuthToken(claims)

	var accessToken, refreshToken string
	if accessToken, err = authToken.NewAccessToken(); err != nil {
		return nil, err
	}

	if refreshToken, err = s.repo.GenerateAndSaveRefreshTokenToStore(authToken); err != nil {
		return nil, err
	}

	return &dto.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s DefaultAuthService) Verify(urlParams map[string]string) error {
	if jwtToken, err := jwtTokenFromString(urlParams["token"]); err != nil {
		return errors.New("error getting token from url")
	} else {
		if jwtToken.Valid {
			claims := jwtToken.Claims.(*domain.AccessTokenClaims)

			if claims.IsUserRole() {
				if !claims.IsRequestVerifiedWithTokenClaims(urlParams) {
					return errors.New("request not verified with the token claims")
				}
			}
			isAuthorized := s.rolePermissions.IsAuthorizedFor(claims.Role, urlParams["routeName"])
			if !isAuthorized {
				return errors.New(fmt.Sprintf("%s role is not authorized", claims.Role))
			}
			return nil
		} else {
			return errors.New("invalid token")
		}
	}

}

func (s DefaultAuthService) Refresh(request dto.RefreshTokenRequest) (*dto.LoginResponse, error) {
	if err := request.IsAccessTokenValid(); err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			// continue with the refresh token functionality
			if err = s.repo.RefreshTokenExists(request.RefreshToken); err != nil {
				return nil, errors.New("error refreshing token")
			}
			// generate a access token from refresh token.
			var accessToken string
			if accessToken, err = domain.NewAccessTokenFromRefreshToken(request.RefreshToken); err != nil {
				return nil, errors.New("error refreshing token")
			}
			return &dto.LoginResponse{AccessToken: accessToken}, nil
		}
		return nil, errors.New("invalid token")
	}
	return nil, errors.New("cannot generate a new access token until the current one expires")
}

func jwtTokenFromString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(domain.HMAC_SAMPLE_SECRET), nil
	})
	if err != nil {
		log.Println("Error while parsing token: " + err.Error())
		return nil, err
	}
	return token, nil
}
