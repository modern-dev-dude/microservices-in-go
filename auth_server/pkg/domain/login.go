package domain

import (
	"database/sql"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

type Login struct {
	Username   string         `db:"username"`
	CustomerId sql.NullString `db:"customer_id"`
	Accounts   sql.NullString `db:"accounts_numbers"`
	Role       string         `db:"role"`
}

func (l Login) ClaimsForAccessToken() AccessTokenClaims {
	if l.Accounts.Valid && l.CustomerId.Valid {
		return l.claimsForUser()
	} else {
		return l.claimsForAdmin()
	}
}

func (l Login) claimsForUser() AccessTokenClaims {
	accounts := strings.Split(l.Accounts.String, ",")
	return AccessTokenClaims{

		CustomerId: l.CustomerId.String,
		Role:       l.Role,
		Username:   l.Username,
		Accounts:   accounts,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TOKEN_DURATION)),
		},
	}
}

func (l Login) claimsForAdmin() AccessTokenClaims {
	return AccessTokenClaims{
		Username: l.Username,
		Role:     l.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TOKEN_DURATION)),
		},
	}
}
