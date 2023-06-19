package mysql

import (
	"fmt"
	"os"

	logger "go_framework/middlewares/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASS")
	HOST := os.Getenv("MYSQL_HOST")
	PORT := os.Getenv("MYSQL_PORT")
	DBNAME := os.Getenv("MYSQL_DB")

	CONN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)

	db, err := gorm.Open(mysql.Open(CONN), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}

func TestConnect() {
	db := Connect()

	logger := logger.GetEchoLogger()

	if db == nil {
		logger.Fatal("DB is not connected")
	}

	logger.Info("DB is connected")

	// logrus.Info("DB is connected")
}
