package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/Scrackc/api-echo-crud/auth"
	"github.com/labstack/echo/v4"
)

type FuncHandler func(w http.ResponseWriter, r *http.Request)

// Log .
func Log(f FuncHandler) FuncHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		log.Printf("petición: %q, método: %q, duration:%v", r.URL.Path, r.Method, time.Since(start).Milliseconds())
	}
}

// Authentication .
func Authentication(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		_, err := auth.ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "No autorizado"})
		}
		return f(c)
	}
}
