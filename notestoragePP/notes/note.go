package notes

type Note struct {
	ID   uint64 `json:"-"`
	Text string `json:"text"`

	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
}

func NewNote() *Note {
	return &Note{}
}
