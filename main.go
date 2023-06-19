package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"go_framework/db/mysql"
	logger "go_framework/middlewares/logger"
	router "go_framework/routers"
	util "go_framework/utils"
)

func init() {
	var envFilePath string
	if util.IsProduction() {
		envFilePath = ".env"
	} else {
		envFilePath = ".dev.env"
	}

	err := godotenv.Load(envFilePath)
	if err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}

	e := router.Router()

	e.Logger = logger.GetEchoLogger()

	e.Logger.Info("Echo Framework Server is Starting...")

	mysql.TestConnect()

	e.Logger.Info(fmt.Sprintf("Server started at port : %s", PORT))
	if e.Debug {
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", PORT)))
	} else {
		e.Logger.Fatal(e.StartTLS(fmt.Sprintf(":%s", PORT), "cert.pem", "privkey.pem"))
	}
}
