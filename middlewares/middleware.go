package middlewares

import (
	"net/http"

	logger "go_framework/middlewares/logger"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

var fieldSeq = map[string]int{
	"TIME":    0,
	"LEVEL":   1,
	"METHOD":  2,
	"HOST":    3,
	"URI":     4,
	"status":  5,
	"MESSAGE": 6,
}

func CORSWithConfig() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		// ! AllowOrigins는 Production의 경우 엄격해야 한다.
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
	})
}

// Hook is a function to process middleware.
func Hook() echo.MiddlewareFunc {
	return logger.HookLogger
}

// func Hook() echo.MiddlewareFunc {
// 	Log := logger.Logger()

// 	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
// 		LogURI:    true,
// 		LogStatus: true,
// 		LogMethod: true,
// 		LogHost:   true,
// 		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
// 			commonFields := logrus.Fields{
// 				"METHOD": values.Method,
// 				"HOST":   values.Host,
// 				"URI":    values.URI,
// 				"status": values.Status,
// 			}

// 			// ? log level condition
// 			if values.Status >= 400 {
// 				Log.SetLevel(logrus.ErrorLevel)
// 				Log.WithFields(commonFields).Error()
// 				if values.Error != nil {
// 					Log.WithFields(logrus.Fields{
// 						"error": values.Error,
// 					}).Error()
// 				}
// 			} else {
// 				Log.SetLevel(logrus.InfoLevel)
// 				Log.WithFields(commonFields).Info()
// 			}
// 			return nil
// 		},
// 	})
// }
