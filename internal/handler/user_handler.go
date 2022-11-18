package handler

import	(
	"net/http"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/cwhight/go-muzz/internal/db"
)

type UserHandler struct {
	userDb db.UserDb
} 

func NewUserHandler(userDb db.UserDb) UserHandler {
	return UserHandler{userDb: userDb}
}


func(h *UserHandler) CreateUser(c echo.Context) error {
	user, err :=  h.userDb.CreateUser()
	if err != nil {
		return echo.NewHTTPError(500, "an internal error occurred")
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusCreated)
	return json.NewEncoder(c.Response()).Encode(user)

}