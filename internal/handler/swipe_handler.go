package handler

import (
	"net/http"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/cwhight/go-muzz/internal/db"
	"github.com/cwhight/go-muzz/internal/model"

)

type SwipeHandler struct {
	userDb db.UserDb
}

func NewSwipeHandler(userDb db.UserDb) SwipeHandler {
	return SwipeHandler{userDb: userDb}
}

func (h *SwipeHandler) Swipe(c echo.Context) error {
	swipeRequest := new(model.SwipeRequest)
    if err := c.Bind(swipeRequest); err != nil {
      return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

	if err := c.Validate(swipeRequest); err != nil {
		return err
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusCreated)
	return json.NewEncoder(c.Response()).Encode("")
}