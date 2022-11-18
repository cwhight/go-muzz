package model

import 	"github.com/google/uuid"

type Match struct {
	Matched bool `json:"matched"`
	MatchId uuid.UUID `json:"matchId"`
}
