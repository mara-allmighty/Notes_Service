package main

import (
	"log"
	"notes_service/internal/service"

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

	// создаем группу api
	api := router.Group("api")

	// Note routes
	api.GET("/note/:id", svc.GetNoteById)
	api.POST("/note", svc.CreateNote)
	api.PUT("/note/:id", svc.UpdateNoteById)
	api.DELETE("/note/:id", svc.DeleteNoteById)

	// запускаем сервер, чтобы слушал 8000 порт
	router.Logger.Fatal(router.Start(":8000"))
}
