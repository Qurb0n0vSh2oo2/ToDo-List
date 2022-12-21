package users

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"topro/pkg/utils"
)

func ReadingUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var allrows []utils.User
	var checkbody utils.User
	byt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response, err := utils.BadRequest()
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, string(response))

		return
	}
	if err := json.Unmarshal(byt, &checkbody); err != nil {
		response, err := utils.BadRequest()
		if err != nil {
			log.Println(err)
		}

		fmt.Fprintf(w, string(response))
		return
	}

	query := fmt.Sprintf("select id, username, created_at from users where id=%v", checkbody.ID)
	// _, err := dbconn.Exec(context.Background(), query)
	rows, err := dbconn.Query(context.Background(), query)
	if err != nil {
		log.Fatalln(err.Error())
	} else if !rows.Next() {
		w.WriteHeader(http.StatusNotFound)
		response, err := utils.NotFound()
		if err != nil {
			log.Println(err)
		}

		fmt.Fprintf(w, string(response))
		return
	}

	firsttask := utils.User{}
	err = rows.Scan(&firsttask.ID, &firsttask.Username, &firsttask.Created_at)

	if err != nil {
		fmt.Println("1")
		log.Println(err)
		return
	}

	allrows = append(allrows, firsttask)

	for rows.Next() {
		saverows := utils.User{}
		err = rows.Scan(&saverows.ID, &saverows.Username, &saverows.Created_at)
		if err != nil {
			log.Println(err)
			return
		}

		allrows = append(allrows, saverows)
	}

	fmt.Println("allrows:", allrows)

	w.WriteHeader(http.StatusOK)
	response, err := utils.Success()
	if err != nil {
		log.Println(err)
	}

	fmt.Fprintf(w, string(response))
	fmt.Println("DONE!")
}
