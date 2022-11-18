package model 

import 	"github.com/google/uuid"

type Profile struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	Age int `json:"age"`
}

type User struct {
	Profile
	Email string `json:"email"`
	Password string `json:"age"`
}
