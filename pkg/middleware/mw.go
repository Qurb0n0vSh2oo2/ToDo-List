package middleware

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func MiddlewareFunc(next http.Handler, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	ioutil.ReadAll(r.Body)
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}
