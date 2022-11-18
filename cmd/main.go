package main

import (
	"fmt"
	"github.com/cwhight/go-muzz/internal/handler"
	"github.com/cwhight/go-muzz/internal/db"
	"github.com/cwhight/go-muzz/internal/auth"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

)

func main() {
	e := echo.New()
	e.Validator = &handler.CustomValidator{Validator: validator.New()}

	userDb := db.NewUserDb()
	matchDb := db.NewMatchDb()

	swipeHandler := handler.NewSwipeHandler(matchDb)
	userHandler := handler.NewUserHandler(userDb)
	profilesHandler := handler.NewProfileHandler(userDb, matchDb)
	loginHandler := handler.NewLoginHandler(userDb)

	jwtConfig := middleware.JWTConfig{
		Claims:                  &auth.Claims{},
        SigningKey:              []byte(auth.GetJWTSecret()),
		TokenLookup:             "cookie:access-token",
		ErrorHandlerWithContext: auth.JWTErrorChecker,
    }

	e.GET("/profiles", profilesHandler.GetProfiles, middleware.JWTWithConfig(jwtConfig))
	e.POST("/user/create", userHandler.CreateUser)
	e.POST("/swipe", swipeHandler.Swipe, middleware.JWTWithConfig(jwtConfig))
	e.POST("/login", loginHandler.Login)

	e.Logger.Fatal(e.Start(":3000"))
}
