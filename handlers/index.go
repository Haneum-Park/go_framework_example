package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

// ? Home 핸들러이며, 간단 헬스체크용
func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
