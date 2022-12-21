package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"topro/pkg/utils"
)

func (h *Handler) ReadTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	byt, err := ioutil.ReadAll(r.Body)

	// var task utils.Task
	var user utils.User
	error := json.Unmarshal(byt, &user)

	if err != nil || error != nil {
		response, err := utils.BadRequest()
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, string(response))

		return
	}
	

	err = h.IReadTask.IReadTask(user)
	if err != nil {
		fmt.Println(4)
		log.Println(err)
		fmt.Fprintf(w, string(err.Error()))
		return
	}

	response, err := utils.Success()
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, string(response))
	fmt.Println("Success")

}
