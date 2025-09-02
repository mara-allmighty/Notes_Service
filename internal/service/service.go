package service

import (
	"database/sql"
	"notes_service/internal/notes"
	"notes_service/internal/users"

	"github.com/labstack/echo/v4"
)

type Service struct {
	db     *sql.DB
	logger echo.Logger

	notesRepo *notes.NotesRepo
	usersRepo *users.UsersRepo
}

func NewService(db *sql.DB, logger echo.Logger) *Service {
	svc := &Service{
		db:     db,
		logger: logger,
	}
	svc.initRepos(db)
	return svc
}

func (s *Service) initRepos(db *sql.DB) {
	s.notesRepo = notes.NewNotesRepo(db)
	s.usersRepo = users.NewUsersRepo(db)
}
