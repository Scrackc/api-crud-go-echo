package main

import (
	"log"

	"github.com/Scrackc/api-echo-crud/auth"
	"github.com/Scrackc/api-echo-crud/handler"
	"github.com/Scrackc/api-echo-crud/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	err := auth.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("no se pudieron cargar los certificados: %v", err)
	}

	store := storage.NewMemory()

	e := echo.New()
	e.Use(middleware.Recover(), middleware.Logger())
	handler.RoutePerson(e, &store)
	handler.RouteAuth(e, &store)

	e.Start(":8080")
}
