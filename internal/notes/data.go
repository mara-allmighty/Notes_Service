package notes

import (
	"database/sql"
	"fmt"
)

type NotesRepo struct {
	db *sql.DB
}

func NewNotesRepo(db *sql.DB) *NotesRepo {
	return &NotesRepo{db: db}
}

// Получить заметку по Id
func (nDB *NotesRepo) GetNoteById(id int) (*Note, error) {
	var note Note

	err := nDB.db.QueryRow(`SELECT title, body, created_at FROM notes WHERE id = $1`, id).Scan(&note.Title, &note.Body, &note.Created_at)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

// Создать новую заметку
func (nDB *NotesRepo) CreateNote(title, body string) error {
	_, err := nDB.db.Exec(`INSERT INTO notes (title, body) VALUES ($1, $2)`, title, body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// Обновить заметку по Id
func (nDB *NotesRepo) UpdateNoteById(id int, newTitle, newBody string) error {
	_, err := nDB.db.Exec(`UPDATE notes SET title = $1, body = $2 WHERE id = $3`, newTitle, newBody, id)
	if err != nil {
		return err
	}

	return nil
}

// Удалить заметку по Id
func (nDB *NotesRepo) DeleteNoteById(id int) error {
	_, err := nDB.db.Exec(`DELETE FROM notes WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
