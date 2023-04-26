package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"notestorage/notes"
	"strconv"

	"github.com/gorilla/mux"
)

type NotesHandler struct {
	NotesRepo *notes.NotesRepo
}

func (h *NotesHandler) Create(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, `{"error": "can't read req body"}`, http.StatusInternalServerError)
		return
	}

	note := notes.NewNote()
	err = json.Unmarshal(data, note)
	if err != nil {
		http.Error(w, `{"error": "Cant unmarshal JSON"}`, http.StatusBadRequest)
		return
	}

	_, err = h.NotesRepo.AddNote(note)
	if err != nil {
		http.Error(w, `{"error": "db error"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	resp, err := json.Marshal(note)
	if err != nil {
		http.Error(w, `{"error": "db error"}`, http.StatusInternalServerError)
		return
	}
	_, err = w.Write(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *NotesHandler) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString, found := vars["id"]
	if !found {
		http.Error(w, `{"error": "bad id"}`, http.StatusBadGateway)
		return
	}
	ID, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		http.Error(w, `{"error": "bad id"}`, http.StatusInternalServerError)
		return
	}
	note, err := h.NotesRepo.GetByID(ID)
	if err != nil || note == nil {
		http.Error(w, `{"error": "no note"}`, http.StatusNotFound)
		return
	}

	data, err := json.Marshal(note)
	if err != nil {
		http.Error(w, `{error: json error}`, http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *NotesHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString, found := vars["id"]
	if !found {
		http.Error(w, `{"error": "bad id"}`, http.StatusBadGateway)
		return
	}
	ID, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		http.Error(w, `{"error": "bad id"}`, http.StatusInternalServerError)
		return
	}
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, `{"error": "can't read req body"}`, http.StatusInternalServerError)
		return
	}
	note := notes.NewNote()
	err = json.Unmarshal(data, note)
	if err != nil {
		http.Error(w, `{"error": "Cant unmarshal JSON"}`, http.StatusBadRequest)
		return
	}
	note.ID = ID
	err = h.NotesRepo.EditNote(note)
	if err != nil {
		http.Error(w, `{"error": "no note"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	resp, err := json.Marshal(note)
	if err != nil {
		http.Error(w, `{"error": "db error"}`, http.StatusInternalServerError)
		return
	}
	_, err = w.Write(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *NotesHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString, found := vars["id"]
	if !found {
		http.Error(w, `{"error": "bad id"}`, http.StatusBadGateway)
		return
	}
	ID, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		http.Error(w, `{"error": "bad id"}`, http.StatusInternalServerError)
		return
	}
	_, err = h.NotesRepo.DelNote(ID)
	if err != nil {
		http.Error(w, `{"error": "no note"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *NotesHandler) List(w http.ResponseWriter, r *http.Request) {
	orderBy := r.URL.Query().Get("order_by")

	notes, err := h.NotesRepo.GetAll(orderBy)
	if err != nil || notes == nil {
		http.Error(w, `{"error": "no notes"}`, http.StatusNotFound)
		return
	}

	data, err := json.Marshal(notes)
	if err != nil {
		http.Error(w, `{error: json error}`, http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
