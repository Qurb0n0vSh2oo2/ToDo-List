package readTask

import (
	"context"
	"errors"
	"fmt"
	"log"
	"topro/pkg/utils"
)

func (t ITaskProvider) IReadTask(user utils.User) (err error) {
	correctId := fmt.Sprintf("select id from users where id = %v", user.ID)
	_, err = t.Postgres.Exec(context.Background(), correctId)
	if err != nil {
		_, err := utils.NotFound()
		if err != nil {
			fmt.Println(2)
			log.Println(err)
		}
		return errors.New("Not Found!")
	} else {

		t.RTaskRead.RReadTask(user)
	}
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
