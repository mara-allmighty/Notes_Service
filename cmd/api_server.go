package main

import (
	"notes_service/internal/service"
	"notes_service/internal/service/middlewares"
	"notes_service/pkg/logs"

	"github.com/labstack/echo/v4"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() {
	logger := logs.NewLogger(false)

	db, err := PostgresConnection()
	if err != nil {
		logger.Fatal(err)
	}

	svc := service.NewService(db, logger)

	router := echo.New()

	// Auth routes
	router.GET("/login", svc.LogIn)
	router.POST("/signup", svc.SignUp)

	// -- Restricted routes --
	api := router.Group("/api")
	// attach auth middleware
	api.Use(middlewares.AuthMiddleware())

	api.GET("/note/:id", svc.GetNoteById)
	api.GET("/notes", svc.GetNotes)
	api.POST("/note", svc.CreateNote)
	api.PUT("/note/:id", svc.UpdateNoteById)
	api.DELETE("/note/:id", svc.DeleteNoteById)

	router.Logger.Fatal(router.Start(s.addr))
}
