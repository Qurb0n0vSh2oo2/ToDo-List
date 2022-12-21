package rupdatetasks

import (
	"context"
	"errors"
	"fmt"
	"log"
	"topro/pkg/utils"
)

func (t *RTaskUpdateProvider) RUpdateTask(task utils.Task) (err error) {
	query := fmt.Sprintf(`
	update tasks set name='%v', description='%v', status='%v', user_id='%v' where id='%v'`,
		task.Name, 
		task.Description, 
		task.Status, 
		task.UserID, 
		task.ID)

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
