package service

import (
	"database/sql"
	jwttoken "notes_service/internal/jwt_token"
	"notes_service/internal/notes"
	"notes_service/internal/users"

	"github.com/labstack/echo/v4"
)

type Service struct {
	db     *sql.DB
	logger echo.Logger
	token  jwttoken.Token

	notesRepo *notes.NotesRepo
	usersRepo *users.UsersRepo
}

func NewService(db *sql.DB, logger echo.Logger, token jwttoken.Token) *Service {
	svc := &Service{
		db:     db,
		logger: logger,
		token:  token,
	}
	svc.initRepos(db)
	return svc
}

func (s *Service) initRepos(db *sql.DB) {
	s.notesRepo = notes.NewNotesRepo(db)
	s.usersRepo = users.NewUsersRepo(db)
}
