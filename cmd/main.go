package main

import (
	"notes_service/internal/service"
	"notes_service/logs"
	"notes_service/middlewares"

	"github.com/labstack/echo/v4"
)

func main() {
	// создаем логгер
	logger := logs.NewLogger(false)

	// подключаемся к базе
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
	api.Use(middlewares.AuthMiddleware()) // attach AuthMiddleware to routes

	api.GET("/note/:id", svc.GetNoteById)
	api.GET("/notes", svc.GetNotes)
	api.POST("/note", svc.CreateNote)
	api.PUT("/note/:id", svc.UpdateNoteById)
	api.DELETE("/note/:id", svc.DeleteNoteById)

	// http://localhost:8000
	router.Logger.Fatal(router.Start(":8000"))
}
