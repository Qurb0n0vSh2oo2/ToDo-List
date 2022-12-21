package handler

import (
	"fmt"
	"log"
	"net/http"
)

func (h *Handler) PreCheck(hh http.HandlerFunc) http.HandlerFunc {
	log.Println(1)
	return func(w http.ResponseWriter, r *http.Request) {
		password := r.Header.Get("Auth")

		
		err := h.IMW.IMiddleware(password)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, string(err.Error()))
			return
		}

		hh(w, r)
	}
}
