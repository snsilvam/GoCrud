package handler

import (
	"encoding/json"
	"net/http"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}
func (p *person) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "mesagge": "Metodo no permitido"`))
		return
	}
	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "mesagge": "Metodo no permitido"`))
		return
	}
	err = p.storage.Create(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type": "error", "mesagge": "Problema al crear la persona"`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message_type": "messagge", "mesagge": "Persona creada"`))
	return
}
