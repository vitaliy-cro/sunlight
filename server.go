package main

import (
	"sunlight/handlers"
	"sunlight/modules/database"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var db *gorm.DB

func main() {
	database.Prepare()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"response":"${latency_human}", time":"${time_rfc3339_nano}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status}}"` + "\n",
	}))

	// Unauthenticated routes
	e.POST("/auth/sign_up", handlers.SignUp)
	e.POST("/auth/sign_in", handlers.SignIn)

	// Restricted routes
	r := e.Group("/")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("auth/current_user", handlers.CurrentUser)

	e.Start(":1323")
}
