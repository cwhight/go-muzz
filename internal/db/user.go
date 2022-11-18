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
	// should randomly gen different values
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
	// should search all other users in DB, filtering based on criteria such as age, gender location etc.
	// not implemented for sake of time
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

func (d *UserDb) GetUser(email string) (*model.User, error) {
	// should return user from DB with given email
	return &model.User{
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
