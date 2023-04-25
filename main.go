package main

import (
	"log"
	"notestorage/handlers"
	"notestorage/notes"
	"notestorage/server"
)

func main() {

	notesHandler := &handlers.NotesHandler{NotesRepo: notes.NewRepo()}

	ErrHTTP := server.Run(notesHandler, ":8091")
	if ErrHTTP != nil {
		log.Fatal(ErrHTTP)
	}
}
