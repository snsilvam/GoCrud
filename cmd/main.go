package main

import (
	"log"
	"net/http"

	"github.com/snsilvam/GoCrud/handler"
	"github.com/snsilvam/GoCrud/storage"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	store := storage.NewMemory()
	mux := http.NewServeMux()
	handler.RoutePerson(mux, &store)
	log.Println("Servidor iniciado en el puerto 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error en el servidor %v\n", err)
	}
}
