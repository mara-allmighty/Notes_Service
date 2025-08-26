package service

import (
	"errors"
	"fmt"
	"net/http"
	externalapi "notes_service/internal/external_api"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GET - /api/note/:id
func (s *Service) GetNoteById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(http.StatusBadRequest, errors.New("invalid params"))
	}

	var user User
	user.Id = s.usersRepo.GetCurrentUser(c)

	note, err := s.notesRepo.GetNoteById(user.Id, id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.New("internal server error"))
	}

	return c.JSON(http.StatusOK, note)
}

// GET - /api/notes
func (s *Service) GetNotes(c echo.Context) error {
	var user User

	user.Id = s.usersRepo.GetCurrentUser(c)

	notes, err := s.notesRepo.GetNotes(user.Id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.New("notes not found"))
	}

	return c.JSON(http.StatusOK, notes)
}

// POST - /api/note
func (s *Service) CreateNote(c echo.Context) error {
	var user User
	var note Note

	err := c.Bind(&note)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(http.StatusBadRequest, errors.New("invalid params"))
	}

	quoteData, err := externalapi.GetQuote()
	if err != nil {
		s.logger.Error(err)
		return c.JSON(http.StatusFailedDependency, errors.New("external api error"))
	}
	note.Body += fmt.Sprintf("; ~quote: '%s' - %s", quoteData["quote"], quoteData["author"])

	user.Id = s.usersRepo.GetCurrentUser(c)

	err = s.notesRepo.CreateNote(user.Id, note.Title, note.Body)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.New("internal server error"))
	}

	// server response
	return c.JSON(http.StatusOK, echo.Map{
		"message": "note created successfully!",
		"note":    note,
	})
}

// PUT - /api/note/:id
func (s *Service) UpdateNoteById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(http.StatusBadRequest, errors.New("invalid params"))
	}

	var user User
	var note Note

	err = c.Bind(&note)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(http.StatusBadRequest, errors.New("invalid params"))
	}

	user.Id = s.usersRepo.GetCurrentUser(c)

	err = s.notesRepo.UpdateNoteById(user.Id, id, note.Title, note.Body)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.New("internal server error"))
	}

	// server response
	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("note with ID %d updated successfully!", id),
		"note":    note,
	})
}

// DELETE - /api/note/:id
func (s *Service) DeleteNoteById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(http.StatusBadRequest, errors.New("invalid params"))
	}

	var user User
	user.Id = s.usersRepo.GetCurrentUser(c)

	err = s.notesRepo.DeleteNoteById(user.Id, id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.New("internal server error"))
	}

	// server response
	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("note with ID %d deleted successfully!", id),
	})
}
