package db

import (
	"github.com/cwhight/go-muzz/internal/model"
	"github.com/google/uuid"
)

type UserDb struct {

}

func NewUserDb() UserDb {
	return UserDb{}
}

func (d *UserDb) CreateUser() (model.User, error) {
	return model.User{
		Profile: model.Profile {
			Name: "test",
			Id: uuid.New(),
			Age: 29,
			Gender: "male",
		},
		Password: "password",
		Email: "email@test.com",
		
	}, nil
}

func (d *UserDb) GetProfileMatches(userId uuid.UUID) ([]model.Profile, error) {
	return []model.Profile{
		model.Profile {
			Name: "test",
			Id: uuid.New(),
			Age: 29,
			Gender: "male",
		},
		model.Profile {
			Name: "test",
			Id: uuid.New(),
			Age: 29,
			Gender: "male",
		},
	}, nil
}