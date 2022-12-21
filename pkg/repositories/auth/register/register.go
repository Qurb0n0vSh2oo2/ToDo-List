package rregister

import (
	"context"
	"errors"
	"fmt"
	"log"
	"topro/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

func (t *RAuthRegProvider) RAuthRegist(account utils.User) (err error) {
	fmt.Println(3)
	var correctUser utils.User
	token, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return err

	}
	
	queryy := fmt.Sprintf(`
		insert into users(username, password) values('%v', '%v')
		`, account.Username, string(token))
	_, err = t.Postgres.Exec(context.Background(), queryy)
	if err != nil {
		return errors.New("Exist!")
	}

	query := fmt.Sprintf(`
	select username from users where username = '%v'
	`, account.Username)
	err = t.Postgres.QueryRow(context.Background(), query).Scan(&correctUser.Username)

	// if err != nil {
	// 	_, err := utils.Exist()
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	return errors.New("Exist!")
	// }
	if account.Password != correctUser.Username {
		_, err = utils.Success()
		if err != nil {
			log.Println(err)
		}
		return errors.New("Success")

	}
	return errors.New("Exist!")

	// tokenUser, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
}
