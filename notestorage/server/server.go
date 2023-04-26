package server

import (
	"net/http"
	"notestorage/handlers"

	"github.com/gorilla/mux"
)

func handlersRouter(notesHandler *handlers.NotesHandler) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/note", notesHandler.Create).Methods("POST")
	r.HandleFunc("/note/{id}", notesHandler.Read).Methods("GET")
	r.HandleFunc("/note/{id}", notesHandler.Update).Methods("PUT")
	r.HandleFunc("/note/{id}", notesHandler.Delete).Methods("DELETE")

	r.HandleFunc("/note", notesHandler.List).Methods("GET")

	return r
}

func Run(notesHandler *handlers.NotesHandler, addr string) error {
	mux := handlersRouter(notesHandler)

	ErrHTTP := http.ListenAndServe(addr, mux)

	return ErrHTTP
}
