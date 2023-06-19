package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	"go_framework/db/mysql"
)

func SignUp(c echo.Context) error {
	role := c.Param("role")
	var result FindAllResult
	if role == "admin" {
		result := AdminFindAll(c)

		if result.Err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": result.Msg,
			})
		}
	}

	db := mysql.Connect()
	if err := db.Create(&result.Admin); err.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed SignUp",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "SUCCESS SignUp",
	})
}

// func SignIn(c echo.Context) {
// 	admin
// }
