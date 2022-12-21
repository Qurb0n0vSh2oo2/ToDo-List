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

func DeletingUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Println("Deleting")

	var account utils.User
	byt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response, err := utils.BadRequest()
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, string(response))

		return
	}
	
	if err := json.Unmarshal(byt, &account); err != nil {
		response, err := utils.BadRequest()
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, string(response))
		return
	}

	var a int64
	correctId := fmt.Sprintf("select id from users where id = %v", account.ID)
	err = dbconn.QueryRow(context.Background(), correctId).Scan(&a)
	if err != nil {
		response, err := utils.NotFound()
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, string(response))
		fmt.Println(">>>NOT FOUND<<<")

		return
	}
	query := fmt.Sprintf("delete from users where id='%v'", account.ID)
	_, err = dbconn.Exec(context.Background(), query)
	utils.Check(err)
	w.WriteHeader(http.StatusOK)
	response, err := utils.Success()
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, string(response))
	fmt.Println("DONE!")

}
