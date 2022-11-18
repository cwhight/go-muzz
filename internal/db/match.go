package db

import (
	"github.com/cwhight/go-muzz/internal/model"
	"github.com/google/uuid"
)

type MatchDb struct {}

func NewMatchDb() MatchDb {
	return MatchDb{}
}

func (d *MatchDb) SaveSwipe(swipe *model.Swipe) error {
	return nil
}

func (d *MatchDb) GetSwipes(userId uuid.UUID) (map[uuid.UUID]bool, error) {
	return map[uuid.UUID]bool{uuid.New(): true}, nil
}

func (d *MatchDb) CheckMatch(swipe *model.Swipe) (*model.Match, error) {
	if swipe.Preference == "NO" {
		return &model.Match{Matched: false}, nil	
	}

	matched, err :=  d.getMatch(swipe.UserId, swipe.ProfileId)
	if err != nil {
		return nil, err
	}

	if matched {
		return &model.Match{Matched: matched, MatchId: uuid.New()}, nil
	} 

	return &model.Match{Matched: matched}, nil
}

func (d *MatchDb) getMatch(userId uuid.UUID, profileId uuid.UUID) (bool, error) {
	return true, nil
}
