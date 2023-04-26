package notes

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type NotesRepo struct {
	mu     *sync.RWMutex
	lastID uint64
	data   []*Note
}

func NewRepo() *NotesRepo {
	return &NotesRepo{
		mu:     &sync.RWMutex{},
		lastID: 0,
		data:   make([]*Note, 0, 10),
	}
}

func (repo *NotesRepo) GetAll(orderField string) ([]*Note, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	switch orderField {
	case "":
		fallthrough
	case "id":
		sort.Slice(repo.data, func(i, j int) bool {
			return repo.data[i].ID < repo.data[j].ID
		})
	case "text":
		sort.Slice(repo.data, func(i, j int) bool {
			return repo.data[i].Text < repo.data[j].Text
		})
	case "created_at":
		sort.Slice(repo.data, func(i, j int) bool {
			return repo.data[i].CreatedAt < repo.data[j].CreatedAt
		})
	case "updated_at":
		sort.Slice(repo.data, func(i, j int) bool {
			return repo.data[i].UpdatedAt < repo.data[j].UpdatedAt
		})
	}

	return repo.data, nil

}

func (repo *NotesRepo) GetByID(ID uint64) (*Note, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	for _, note := range repo.data {
		if note.ID == ID {
			return note, nil
		}
	}
	return nil, fmt.Errorf("no note found")
}

func (repo *NotesRepo) AddNote(note *Note) (uint64, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.lastID++
	note.ID = repo.lastID
	repo.data = append(repo.data, note)
	note.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	return repo.lastID, nil
}

func (repo *NotesRepo) EditNote(noteEdit *Note) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	for _, note := range repo.data {
		if note.ID == noteEdit.ID {
			note.Text = noteEdit.Text
			note.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			return nil
		}
	}

	return fmt.Errorf("no note found")
}

func (repo *NotesRepo) DelNote(ID uint64) (bool, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	i := -1
	for idx, note := range repo.data {
		if note.ID == ID {
			i = idx
			break
		}
	}
	if i < 0 {
		return false, fmt.Errorf("no note found")
	}

	if i < len(repo.data)-1 {
		copy(repo.data[i:], repo.data[i+1:])
	}
	repo.data[len(repo.data)-1] = nil
	repo.data = repo.data[:len(repo.data)-1]

	return true, nil
}
