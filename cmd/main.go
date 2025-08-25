package main

import (
	"log"
	"notes_service/internal/service"
	"notes_service/middlewares"

	"github.com/labstack/echo/v4"
)

func main() {
	// подключаемся к базе
	db, err := PostgresConnection()
	if err != nil {
		log.Fatal(err)
	}

	svc := service.NewService(db)

	router := echo.New()

	// Auth routes
	router.GET("/login", svc.LogIn)
	router.POST("/signup", svc.SignUp)

	// -- Restricted routes --
	api := router.Group("/api")
	api.Use(middlewares.AuthMiddleware()) // attach AuthMiddleware to routes

	api.GET("/note/:id", svc.GetNoteById)
	api.POST("/note", svc.CreateNote)
	api.PUT("/note/:id", svc.UpdateNoteById)
	api.DELETE("/note/:id", svc.DeleteNoteById)

	api.GET("/profile", svc.GetUserProfile)

	// запускаем сервер, чтобы слушал 8000 порт
	router.Logger.Fatal(router.Start(":8000"))
}
