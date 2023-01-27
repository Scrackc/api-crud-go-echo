package handler

import (
	"github.com/Scrackc/api-echo-crud/middleware"
	"github.com/labstack/echo/v4"
)

// RoutePerson .
func RoutePerson(e *echo.Echo, storage Storage) {
	h := newPerson(storage)
	persons := e.Group("/v1/persons")
	persons.Use(middleware.Authentication)
	persons.POST("/create", h.create)
	persons.GET("/get-all", h.getAll)
	persons.PUT("/update/:id", h.update)
	persons.DELETE("/delete/:id", h.delete)
	persons.GET("/get-by-id/:id", h.getByID)

}

func RouteAuth(e *echo.Echo, storage Storage) {
	h := newLogin(storage)
	e.POST("/v1/login", h.login)

}
