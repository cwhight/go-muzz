package handler

import (
	"net/http"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/cwhight/go-muzz/internal/db"
)

type ProfileHandler struct {
	userDb db.UserDb
}

func  NewProfileHandler(userDb db.UserDb) ProfileHandler {
	return ProfileHandler{userDb: userDb}
}

func (h *ProfileHandler) GetProfiles(c echo.Context) error {
	userId := c.QueryParam("userId")
	if userId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "user id must be provided")
	}

	idAsUuid, err := uuid.Parse(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "user id must be valid uuid")
	}
	
	profiles, err :=  h.userDb.GetProfileMatches(idAsUuid)
	if err != nil {
		return echo.NewHTTPError(500, "an internal error occurred")
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusCreated)
	return json.NewEncoder(c.Response()).Encode(profiles)
}