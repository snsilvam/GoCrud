package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/snsilvam/GoCrud/model"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}
func (p *person) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Metodo no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "Estructura de los campos invalida", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	err = p.storage.Create(&data)
	if err != nil {
		response := newResponse(Error, "Problema interno al intentar crear la persona", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "Persona creada con exito", nil)
	responseJSON(w, http.StatusCreated, response)
}
func (p *person) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Metodo no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "El id ingresado, tiene que ser un entero positivo ", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	data := model.Person{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "Estructura de los campos invalida", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	err = p.storage.Update(ID, &data)
	if err != nil {
		response := newResponse(Error, "Problema interno al intentar crear la persona", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	person := newResponse(Message, "Persona actualizada con exito", nil)
	responseJSON(w, http.StatusOK, person)

}
func (p *person) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newResponse(Error, "Metodo no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "Tiene que ingresar un id positivo", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	err = p.storage.Delete(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response := newResponse(Error, "El id ingresado, no existe", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "Error interno.. Intente más tarde", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "Persona eliminada de manera satisfactoria", nil)
	responseJSON(w, http.StatusOK, response)
}
func (p *person) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Metodo no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "Problema interno al intentar crear la persona", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response := newResponse(Message, "Personas encontradas con exito", data)
	responseJSON(w, http.StatusOK, response)
}
