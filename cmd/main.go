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

	// -- Accessible routes --
	a := router.Group("/api")
	a.GET("/note/:id", svc.GetNoteById)

	// -- Restricted routes --
	r := router.Group("/api")
	r.Use(middlewares.AuthMiddleware()) // attach AuthMiddleware to routes

	r.POST("/note", svc.CreateNote)
	r.PUT("/note/:id", svc.UpdateNoteById)
	r.DELETE("/note/:id", svc.DeleteNoteById)

	// запускаем сервер, чтобы слушал 8000 порт
	router.Logger.Fatal(router.Start(":8000"))
}
