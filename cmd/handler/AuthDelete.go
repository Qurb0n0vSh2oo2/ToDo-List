package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"topro/pkg/utils"
)

func (h *Handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	byt, err := ioutil.ReadAll(r.Body)

	var account utils.User
	error := json.Unmarshal(byt, &account)

	if err != nil || error != nil {
		response, err := utils.BadRequest()
		if err != nil {
			fmt.Println(1)
			log.Println(err)
		}
		fmt.Fprintf(w, string(response))

		return
	}

	err = h.IAuthDelete.IAuthDelete(account)

	if err != nil {
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
