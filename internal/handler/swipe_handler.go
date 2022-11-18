package handler

import (
	"net/http"
	"encoding/json"
	"github.com/labstack/echo/v4"

	"github.com/cwhight/go-muzz/internal/auth"

	"github.com/cwhight/go-muzz/internal/db"
	"github.com/cwhight/go-muzz/internal/model"
)

type SwipeHandler struct {
	matchDb db.MatchDb
}

func NewSwipeHandler(matchDb db.MatchDb) SwipeHandler {
	return SwipeHandler{matchDb: matchDb}
}

func (h *SwipeHandler) Swipe(c echo.Context) error {
	cookie, err := c.Cookie("access-token")
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorised access")
	}

	user, err := auth.ParseCookie(cookie.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorised access")
	}

	swipe := new(model.Swipe)
    if err := c.Bind(swipe); err != nil {
      return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

	if err = c.Validate(swipe); err != nil {
		return err
	}

	if swipe.UserId != user.Id {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorised access")
	}

	err = h.matchDb.SaveSwipe(swipe)
	if err != nil {
		return echo.NewHTTPError(500, "an internal error occurred")
	}

	match, err := h.matchDb.CheckMatch(swipe)
	if err != nil {
		return echo.NewHTTPError(500, "an internal error occurred")
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusCreated)
	return json.NewEncoder(c.Response()).Encode(match)
}
