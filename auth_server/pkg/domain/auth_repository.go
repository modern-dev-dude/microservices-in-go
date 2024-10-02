package domain

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
)

type AuthRepository interface {
	FindBy(username, password string) (*Login, error)
	GenerateAndSaveRefreshTokenToStore(authToken AuthToken) (string, error)
	RefreshTokenExists(authToken string) error
}

type AuthRepositoryDb struct {
	client *sqlx.DB
}

func NewAuthRepository(client *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{client: client}
}

func (d AuthRepositoryDb) FindBy(username, password string) (*Login, error) {
	var login Login
	sqlVerify := `SELECT username, u.customer_ID, role, group_concat(a.account_id) as account
	LEFT JOIN accounts a ON a.customer_id = u.customer_id
	WHERE username = ? and password = ?
	GROUP BY a.customer_id
	`

	err := d.client.Get(&login, sqlVerify, username, password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			log.Println("Error while verifying login request: " + err.Error())
			return nil, errors.New("unexpected error")
		}
	}

	return &login, nil
}

func (d AuthRepositoryDb) GenerateAndSaveRefreshTokenToStore(authToken AuthToken) (string, error) {
	refreshToken, err := authToken.newRefreshToken()

	if err != nil {
		return "", err
	}

	sqlInsert := "insert into refresh_token_store (refresh_token) values (?)"
	_, err = d.client.Exec(sqlInsert, refreshToken)
	if err != nil {
		log.Println("Error while saving refresh token")
		return "", err
	}

	return refreshToken, nil
}

func (d AuthRepositoryDb) RefreshTokenExists(refreshToken string) error {
	sqlSelect := "select refresh_token from refresh_token_store where refresh_token = ?"
	var token string
	err := d.client.Get(&token, sqlSelect, refreshToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("refresh token not registered in the store")
		} else {
			log.Println("Unexpected database error: " + err.Error())
			return errors.New("unexpected database error")
		}
	}
	return nil
}
