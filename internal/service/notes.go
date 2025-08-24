package service

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GET - localhost:8000/api/note/:id
func (s *Service) GetNoteById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid params"))
	}

	note, err := s.notesDB.GetNoteById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("internal server error"))
	}

	return c.JSON(http.StatusOK, note)
}

// POST - localhost:8000/api/note
func (s *Service) CreateNote(c echo.Context) error {
	var note Note

	err := c.Bind(&note)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid params"))
	}

	err = s.notesDB.CreateNote(note.Title, note.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("internal server error"))
	}

	return c.JSON(http.StatusOK, "OK")
}

// UPDATE - localhost:8000/api/note/:id
func (s *Service) UpdateNoteById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid params"))
	}

	var note Note
	err = c.Bind(&note)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid params"))
	}

	err = s.notesDB.UpdateNoteById(id, note.Title, note.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("internal server error"))
	}

	return c.JSON(http.StatusOK, "OK")
}

// DELETE - localhost:8000/api/note/:id
func (s *Service) DeleteNoteById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid params"))
	}

	err = s.notesDB.DeleteNoteById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("internal server error"))
	}

	return c.JSON(http.StatusOK, "OK")
}
