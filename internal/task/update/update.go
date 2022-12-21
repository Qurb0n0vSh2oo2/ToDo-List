package updateTask

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"
	"topro/pkg/utils"
)

func (t ITaskProvider) IUpdateTask(task utils.Task) (err error) {
	pattern := `^\w`
	MatchedName, err := regexp.Match(pattern, []byte(task.Name))
	utils.Check(err)
	MatchedDescr, err := regexp.Match(pattern, []byte(task.Description))
	utils.Check(err)
	MatchedStatus, err := regexp.Match(pattern, []byte(task.Status))
	utils.Check(err)

	var userID int
	correctUserId := fmt.Sprintf("select id from users where id = %v", task.UserID)
	err = t.Postgres.QueryRow(context.Background(), correctUserId).Scan(&userID)
	if err != nil {
		_, err = utils.NotFound()
		if err != nil {
			log.Println(err)
		}
		return errors.New("There is no user!")
	}

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

	if MatchedName && MatchedDescr && MatchedStatus {

		t.RTaskUpdate.RUpdateTask(task)

	}else {
		_, err := utils.FillBkanks()
		if err != nil {
			log.Println(err)
		}
		return errors.New("Fill all blanks, please!")
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
