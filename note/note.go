package note

import "time"

type Note struct {
	NoteAuthor    string
	NoteText      string
	ReferenceText string
	BookTitle     string
	BookAuthor    string
	CreatedAt     time.Time
}
