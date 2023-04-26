package main

import (
	"log"
	"notestorage/handlers"
	"notestorage/notes"
	"notestorage/server"
)

func main() {

	notesHandler := &handlers.NotesHandler{NotesRepo: notes.NewRepo()}

	log.Print("Server started!")
	ErrHTTP := server.Run(notesHandler, ":8080")
	if ErrHTTP != nil {
		log.Fatal(ErrHTTP)
	}
}
