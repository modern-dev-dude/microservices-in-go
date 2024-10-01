package domain

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
)

type AuthRepositoryDb struct {
	client *sqlx.DB
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
