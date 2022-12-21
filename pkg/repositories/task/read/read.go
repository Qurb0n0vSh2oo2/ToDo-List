package rreadtasks

import (
	"context"
	"errors"
	"fmt"
	"log"
	"topro/pkg/utils"
)

func (t *RTaskReadProvider) RReadTask(user utils.User) (err error) {
	var allrows []utils.Task
	query := fmt.Sprintf("select id, name, description, status, user_id from tasks where user_id=%v", user.ID)
	// _, err := dbconn.Exec(context.Background(), query)
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

	firsttask := utils.Task{}
	err = rows.Scan(&firsttask.ID, &firsttask.Name, &firsttask.Description, &firsttask.Status, &firsttask.UserID)
	if err != nil {
		log.Println(err)
		return
	}

	allrows = append(allrows, firsttask)

	for rows.Next() {
		saverows := utils.Task{}
		err = rows.Scan(&saverows.ID, &saverows.Name, &saverows.Description, &saverows.Status, &saverows.UserID)
		if err != nil {
			log.Println(err)
			return
		}

		allrows = append(allrows, saverows)
	}

	fmt.Println("allrows:", allrows)

	// w.WriteHeader(http.StatusOK)
	_, err = utils.Success()
	if err != nil {
		log.Println(err)
	}
	return errors.New("Success!")
}
