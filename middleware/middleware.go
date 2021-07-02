package middleware

import (
	"log"
	"net/http"
)

func Log(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("peticion %q, método: %q", r.URL.Path, r.Method)
		f(w, r)
		log.Printf("peticion %q, método: %q", r.URL.Path, r.Method)
	}
}
func Authentication(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("password")
		if token != "7752984" {
			forbidden(w, r)
			return
		}
		f(w, r)
	}
}
func forbidden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("No tiene autorizacion =("))
}
