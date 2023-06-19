package handlers

import (
	"go_framework/db/mysql"
	"go_framework/models"

	echo "github.com/labstack/echo/v4"

	util "go_framework/utils"
)

type FindAllResult struct {
	Admin *models.Admin
	Msg   string
	Err   error
}

func AdminFindAll(c echo.Context) FindAllResult {
	admin := new(models.Admin)
	var err error
	var msg string

	if err := c.Bind(admin); err != nil {
		msg = "bad request"
	}

	db := mysql.Connect()
	result := db.Find(&admin, "username=?", admin.Username)

	if result.RowsAffected != 0 {
		msg = "username already exists"
	}

	salt, hash := util.Encrypt(admin.Password)

	admin.Password = util.EncodingBytes(hash)
	admin.Salt = util.EncodingBytes(salt)

	return FindAllResult{
		Admin: admin,
		Msg:   msg,
		Err:   err,
	}
}
