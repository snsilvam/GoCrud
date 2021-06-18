package handler

import {
	"https://github.com/snsilvam/Apis-in-Go/tree/main/CRUD/Model"
	"net/http"
}
func RoutePerson(mux *http.ServerMux, storage Storage) {
	h := newPerson(storage)
	mux.HandleFunc("/v1/person/create", h.create)
}
