package authDelete

import (
	"context"
	"errors"
	"fmt"
	"log"
	"topro/pkg/utils"
)

func (t IAuthDeleteProvider) IAuthDelete(account utils.User) (err error) {
	var a int
	correctId := fmt.Sprintf("select id from users where id = %v", account.ID)
	err = t.Postgres.QueryRow(context.Background(), correctId).Scan(&a)
	if err != nil {
		_, err = utils.NotFound()
		if err != nil {
			log.Println(err)
		}
		return errors.New("Not Found")
	}
	t.RAuthDelete.RDeleteTask(account)
	return

}
