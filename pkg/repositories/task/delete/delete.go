package rdeletetasks

import (
	"context"
	"errors"
	"fmt"
	"log"
	"topro/pkg/utils"
)

func (t *RTaskDeleteProvider) RDeleteTask(task utils.Task) (err error) {
	query := fmt.Sprintf("delete from tasks where id='%v'", task.ID)
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
