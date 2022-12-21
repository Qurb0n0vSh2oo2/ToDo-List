package deleteTask

import (
	"context"
	"errors"
	"fmt"
	"log"
	"topro/pkg/utils"
)

func (t ITaskProvider) IDeleteTask(task utils.Task) (err error) {
	var a int
	correctId := fmt.Sprintf("select id from tasks where id = %v", task.ID)
	err = t.Postgres.QueryRow(context.Background(), correctId).Scan(&a)
	if err != nil {
		_, err = utils.NotFound()
		if err != nil {
			log.Println(err)
		}
		return errors.New("Not Found!")
	}
	t.RTaskDelete.RDeleteTask(task)
	return

}

// func FillBlanks(w http.ResponseWriter, r *http.Request) (err error)  {
// 	response, err := utils.FillBkanks()
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	fmt.Sprintf(w, string(response))
// 	fmt.Println("Fill all blanks, please!")
// 	return
// }
