package utils

import "os"

func IsProduction() bool {
	var isProduction bool
	appEnv, exists := os.LookupEnv("APP_ENV")

	if !exists || (exists && (appEnv == "development" || appEnv == "debug")) {
		isProduction = false
	} else if appEnv == "production" {
		isProduction = true
	}

	return isProduction
}
