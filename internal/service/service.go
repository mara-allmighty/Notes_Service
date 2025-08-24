package service

import (
	"database/sql"
	"notes_service/internal/notes"
)

type Service struct {
	db      *sql.DB
	notesDB *notes.NotesDB
}

func NewService(db *sql.DB) *Service {
	svc := &Service{db: db}
	svc.initDBs(db)
	return svc
}

func (s *Service) initDBs(db *sql.DB) {
	s.notesDB = notes.NewNotesDB(db)
}
