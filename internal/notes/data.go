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
func (nr *NotesRepo) GetNoteById(id int) (*Note, error) {
	var note Note

	err := nr.db.QueryRow(`SELECT user_id, title, body, created_at FROM notes WHERE id = $1`, id).Scan(&note.User_id, &note.Title, &note.Body, &note.Created_at)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

// Создать новую заметку
func (nr *NotesRepo) CreateNote(user_id int, title, body string) error {
	_, err := nr.db.Exec(`INSERT INTO notes (user_id, title, body) VALUES ($1, $2, $3)`, user_id, title, body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// Обновить заметку по Id
func (nr *NotesRepo) UpdateNoteById(id int, newTitle, newBody string) error {
	_, err := nr.db.Exec(`UPDATE notes SET title = $1, body = $2 WHERE id = $3`, newTitle, newBody, id)
	if err != nil {
		return err
	}

	return nil
}

// Удалить заметку по Id
func (nr *NotesRepo) DeleteNoteById(id int) error {
	_, err := nr.db.Exec(`DELETE FROM notes WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
