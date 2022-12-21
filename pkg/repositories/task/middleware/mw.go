package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"topro/pkg/utils"
)

func (t *RMwProvider) RMW(password string) (err error) {
	log.Println(3)
	var userPass utils.User
	var user utils.User
	query := fmt.Sprintf(`select password from users where password = '%v'`, password)

	rows, err := t.Postgres.Query(context.Background(), query)
	if err != nil {
		log.Fatalln(err.Error())
	} else if !rows.Next() {
		// w.WriteHeader(http.StatusNotFound)
		_, err = utils.NotFound()
		if err != nil {
			log.Println(err)
		}
		return errors.New("Not Found!")
	}

	err = t.Postgres.QueryRow(context.Background(), query).Scan(&userPass.Password)

	// _ = rows.Scan(&tok)
	// err = bcrypt.CompareHashAndPassword([]byte(password), []byte(userPass.Password))
	// fmt.Println(tok)
	fmt.Println(password)
	fmt.Println(user.Password)
	if password != userPass.Password {
		// var response []byte
		// response, err = utils.NotFound()
		// if err != nil {
		// 	log.Println(err)
		// }

		return errors.New(fmt.Sprintf("Error while comparing: %v", err.Error()))
	}
	return
}
