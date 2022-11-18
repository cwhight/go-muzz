package model

import 	(
	"github.com/google/uuid"
	_ "github.com/go-playground/validator"
)

type SwipeRequest struct {
	UserId uuid.UUID `json:"userId" validate:"required"`
	ProfileId uuid.UUID `json:"profileId" validate:"required"`
	Preference string `json:"preference" validate:"required"`
}
