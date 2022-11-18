package handler

import  (
	"github.com/go-playground/validator"
	"net/http"
	"github.com/labstack/echo/v4"

)

type CustomValidator struct {
    Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
