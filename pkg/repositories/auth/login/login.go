package rlogin

import (
	"context"
	"errors"
	"fmt"
	"log"
	"topro/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

func (t *RAuthLoginProvider) RAuthLogin(account utils.User) (err error) {
	fmt.Println(3)
	var tok string
	query := fmt.Sprintf(`select password from users where username = '%v'`, account.Username)
	_ = t.Postgres.QueryRow(context.Background(), query).Scan(&tok)
	log.Println(tok)
	log.Println(account.Password)
	err = bcrypt.CompareHashAndPassword([]byte(tok), []byte(account.Password))
	if err != nil {
		_, err := utils.NotFound()
		if err != nil {
			log.Println(err)
		}
		return errors.New("Not Found!")
	}

	return errors.New("Success!")
}
