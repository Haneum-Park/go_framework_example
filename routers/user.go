package routers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func GetUsers(c echo.Context) error {
	users := []User{{1, "ryna"}, {2, "joker"}}
	return c.JSON(http.StatusBadRequest, users)
}

func PostUsers(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	fmt.Println(user)

	return c.JSON(http.StatusCreated, nil)
}
