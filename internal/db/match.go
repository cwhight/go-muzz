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
	// store given swipe in swipe table
	return nil
}

func (d *MatchDb) GetSwipes(userId uuid.UUID) (map[uuid.UUID]bool, error) {
	// fetch all 'swipes' for given user id - return map[id]bool with all profile ids (i.e users that the userId has swiped for previously)
	return map[uuid.UUID]bool{uuid.New(): true}, nil
}

func (d *MatchDb) CheckMatch(swipe *model.Swipe) (*model.Match, error) {
	// returns false match if current user swiped 'NO'
	if swipe.Preference == "NO" {
		return &model.Match{Matched: false}, nil	
	}

	// checks for match
	matched, err :=  d.getMatch(swipe.UserId, swipe.ProfileId)
	if err != nil {
		return nil, err
	}

	// returns successful match and gens match ID if matching swipe exists
	if matched {
		return &model.Match{Matched: matched, MatchId: uuid.New()}, nil
	} 

	return &model.Match{Matched: matched}, nil
}

func (d *MatchDb) getMatch(userId uuid.UUID, profileId uuid.UUID) (bool, error) {
	// checks if 'swipe' exists for reverse of userId and profileId
	// i.e other user has swiped the current user
	// return true if has swiped and has macthed
	return true, nil
}
