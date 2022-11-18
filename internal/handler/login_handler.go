package handler

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/cwhight/go-muzz/internal/db"
	"github.com/cwhight/go-muzz/internal/auth"
	"github.com/cwhight/go-muzz/internal/model"
)

type LoginHandler struct {
	userDb db.UserDb
}

func NewLoginHandler(userDb db.UserDb) LoginHandler {
	return LoginHandler{userDb: userDb}
}

func (h *LoginHandler) Login(c echo.Context) error {
	request := new(model.LoginRequest)
    if err := c.Bind(request); err != nil {
      return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

	if err := c.Validate(request); err != nil {
		return err
	}

	user, err := h.userDb.GetUser(request.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "no user found with that email")
	}

	if user.Password != request.Password {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid password")
	}

	err = auth.GenerateTokensAndSetCookies(user, c)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token is incorrect")
	}

	return c.Redirect(http.StatusMovedPermanently, "/profiles")
}
