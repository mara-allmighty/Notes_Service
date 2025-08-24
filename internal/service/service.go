package service

import (
	"database/sql"
	"notes_service/internal/notes"
	"notes_service/internal/users"
)

type Service struct {
	db *sql.DB

	notesRepo *notes.NotesRepo
	usersRepo *users.UsersRepo
}

func NewService(db *sql.DB) *Service {
	svc := &Service{
		db: db,
	}
	svc.initRepos(db)
	return svc
}

func (s *Service) initRepos(db *sql.DB) {
	s.notesRepo = notes.NewNotesRepo(db)
	s.usersRepo = users.NewUsersRepo(db)
}
