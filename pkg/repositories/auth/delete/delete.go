package rAuthDelete

import (
	"context"
	"errors"
	"fmt"
	"log"
	"topro/pkg/utils"
)

func (t *RAuthDeleteProvider) RDeleteTask(account utils.User) (err error) {
	query := fmt.Sprintf("delete from users where id='%v'", account.ID)
	_, err = t.Postgres.Exec(context.Background(), query)
	if err != nil {
		_, err = utils.NotFound()
		if err != nil {
			log.Println(err)
		}
		return errors.New("Not Found!")
	}
	return
}
