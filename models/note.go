/*
* en Note-struct i models/note.go

en hardkodet slice (datasett)

et GET /notes-endepunkt i handlers/note_handler.go

ruter lagt til i routes/routes.go

og du binder det sammen i main.go
* */

package models

type Note struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
