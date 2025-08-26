package main

import (
	jwttoken "notes_service/internal/jwt_token"
	"notes_service/internal/middlewares"
	"notes_service/internal/service"
	"notes_service/logs"

	"github.com/labstack/echo/v4"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{addr: addr}
}

func (s *APIServer) Run() {
	token := jwttoken.Token{}
	logger := logs.NewLogger(false)

	db, err := PostgresConnection()
	if err != nil {
		logger.Fatal(err)
	}

	svc := service.NewService(db, logger, token)

	router := echo.New()

	// Auth routes
	router.GET("/login", svc.GetToken)
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
