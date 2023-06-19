package routers

import (
	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"

	handler "go_framework/handlers"
	customMiddleware "go_framework/middlewares"
	util "go_framework/utils"
)

func Router() *echo.Echo {
	var isProduction = util.IsProduction()

	// ? echo.New()를 사용하여 *Echo를 리턴 받는다.
	e := echo.New()

	// ? Debug 모드. isProduction이 true라면 production 모드이므로 !isProduction이여야 한다.
	e.Debug = !isProduction

	// ? echo middleware.
	// NOTE Custom Middleware는 middlewares에서 정의한다.
	e.Use(middleware.Recover())              // NOTE from panics anywhere in the chain
	e.Use(customMiddleware.Hook())           // NOTE Setting logger
	e.Use(customMiddleware.CORSWithConfig()) // NOTE Setting CORS
	// NOTE K, M, G, T, P
	e.Use(middleware.BodyLimit("10M"))

	e.GET("/", handler.Home)

	// ? Group User
	user := e.Group("/user")
	user.GET("", GetUsers)
	user.POST("", PostUsers)

	auth := e.Group("/auth")
	auth.POST("/register/:role", handler.SignUp)

	return e
}
