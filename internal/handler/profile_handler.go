package handler

import (
	"net/http"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/cwhight/go-muzz/internal/db"
	"github.com/cwhight/go-muzz/internal/auth"

	"github.com/cwhight/go-muzz/internal/model"
)

type ProfileHandler struct {
	userDb db.UserDb
	matchDb db.MatchDb
}

func  NewProfileHandler(userDb db.UserDb, matchDb db.MatchDb) ProfileHandler {
	return ProfileHandler{userDb: userDb, matchDb: matchDb}
}

func (h *ProfileHandler) GetProfiles(c echo.Context) error {
	cookie, err := c.Cookie("access-token")
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorised access")
	}

	user, err := auth.ParseCookie(cookie.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorised access")
	}
	
	profiles, err :=  h.userDb.GetProfileMatches(user.Id)
	if err != nil {
		return echo.NewHTTPError(500, "an internal error occurred")
	}

	swipes, err := h.matchDb.GetSwipes(user.Id)
	if err != nil {
		return echo.NewHTTPError(500, "an internal error occurred")
	}

	filteredProfiles := []model.Profile{}
	for _, profile := range profiles {
		if !swipes[profile.Id] {
			filteredProfiles = append(filteredProfiles, profile)
		}
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusCreated)
	return json.NewEncoder(c.Response()).Encode(filteredProfiles)
}
