package main

import (
	"log"
	"testAPI/config"

	ad "testAPI/features/auth/data"
	ah "testAPI/features/auth/handler"
	as "testAPI/features/auth/service"

	pd "testAPI/features/presence/data"
	ph "testAPI/features/presence/handler"
	ps "testAPI/features/presence/service"

	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)

	authData := ad.New(db)
	authService := as.New(authData)
	authhandler := ah.New(authService)

	presenceData := pd.New(db)
	presenceService := ps.New(presenceData)
	presenceHandler := ph.New(presenceService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "- method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("/register", authhandler.Register())
	e.POST("/login", authhandler.Login())

	e.POST("/presence", presenceHandler.Insert(), echojwt.JWT([]byte(config.JWT_KEY)))
	e.POST("/presence/:id", presenceHandler.Approve(), echojwt.JWT([]byte(config.JWT_KEY)))
	e.GET("/presence", presenceHandler.GetData(), echojwt.JWT([]byte(config.JWT_KEY)))
	e.GET("/presence/user/:id", presenceHandler.GetDataByIdUser(), echojwt.JWT([]byte(config.JWT_KEY)))

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
