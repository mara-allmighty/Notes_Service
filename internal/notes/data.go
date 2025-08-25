package notes

import (
	"database/sql"
	"errors"
	"fmt"
)

type NotesRepo struct {
	db *sql.DB
}

func NewNotesRepo(db *sql.DB) *NotesRepo {
	return &NotesRepo{db: db}
}

// Получить заметку по Id. Можно получить только свою заметку.
func (nr *NotesRepo) GetNoteById(user_id, id int) (*Note, error) {
	var note Note

	err := nr.db.QueryRow(`SELECT user_id FROM notes WHERE id = $1`, id).Scan(&note.User_id)
	if err != nil {
		return nil, err
	}
	if user_id != note.User_id {
		return nil, errors.New("access denied")
	}

	err = nr.db.QueryRow(`SELECT id, user_id, title, body, created_at FROM notes WHERE id = $1`, id).Scan(&note.Id, &note.User_id, &note.Title, &note.Body, &note.Created_at)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

// Получить свои заметки.
func (nr *NotesRepo) GetNotes(user_id int) ([]Note, error) {
	rows, err := nr.db.Query(`SELECT id, user_id, title, body, created_at FROM notes WHERE user_id = $1`, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userNotesSlice []Note

	for rows.Next() {
		var note Note
		err := rows.Scan(&note.Id, &note.User_id, &note.Title, &note.Body, &note.Created_at)
		if err != nil {
			return nil, err
		}
		userNotesSlice = append(userNotesSlice, note)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return userNotesSlice, nil
}

// Создать новую заметку.
func (nr *NotesRepo) CreateNote(user_id int, title, body string) error {
	_, err := nr.db.Exec(`INSERT INTO notes (user_id, title, body) VALUES ($1, $2, $3)`, user_id, title, body)
	if err != nil {
		return err
	}

	return nil
}

// Обновить заметку по Id. Можно обновлять только собственные заметки.
func (nr *NotesRepo) UpdateNoteById(user_id, id int, newTitle, newBody string) error {
	var note Note

	err := nr.db.QueryRow(`SELECT user_id FROM notes WHERE id = $1`, id).Scan(&note.User_id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if user_id != note.User_id {
		fmt.Println(err)
		return errors.New("access denied")
	}

	_, err = nr.db.Exec(`UPDATE notes SET title = $1, body = $2 WHERE id = $3`, newTitle, newBody, id)
	if err != nil {
		return err
	}

	return nil
}

// Удалить заметку по Id. Разрешается удалять только собственные заметки.
func (nr *NotesRepo) DeleteNoteById(user_id, id int) error {
	var note Note

	err := nr.db.QueryRow(`SELECT user_id FROM notes WHERE id = $1`, id).Scan(&note.User_id)
	if err != nil {
		return err
	}
	if user_id != note.User_id {
		return errors.New("access denied")
	}

	_, err = nr.db.Exec(`DELETE FROM notes WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
