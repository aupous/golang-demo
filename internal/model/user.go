package model

import "github.com/google/uuid"

type User struct {
	tableName struct{} `pg:"users,alias:users"`
	ID uuid.UUID `json:"id" pg:"type:uuid"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	Create(*User) error
	Update(*User) error
	Find() ([]*User, error)
	Delete(*User) error
}