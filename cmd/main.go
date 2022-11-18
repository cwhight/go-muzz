package main

import (
	"fmt"
	"github.com/cwhight/go-muzz/internal/handler"
	"github.com/cwhight/go-muzz/internal/db"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Validator = &handler.CustomValidator{Validator: validator.New()}

	userDb := db.NewUserDb()

	swipeHandler := handler.NewSwipeHandler(userDb)
	userHandler := handler.NewUserHandler(userDb)
	profilesHandler := handler.NewProfileHandler(userDb)

	e.GET("/profiles", profilesHandler.GetProfiles)
	e.POST("/user/create", userHandler.CreateUser)
	e.POST("/swipe", swipeHandler.Swipe)

	e.Logger.Fatal(e.Start(":3000"))
	fmt.Println("hello")
}