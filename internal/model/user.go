package model

import "github.com/google/uuid"

type User struct {
	tableName struct{}  `pg:"users,alias:users"`
	ID        uuid.UUID `json:"id" pg:"type:uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	FullName  string    `json:"fullName"`
	Age       int       `json:"age"`
	Job       string    `json:"job"`
}

type UserRepository interface {
	Create(*User) error
	Update(*User) error
	Search(search string) ([]*User, error)
	Find(request FindUserRequest) ([]*User, int64, error)
	FindByEmail(email string) (*User, error)
	Delete(*User) error
}

type FindUserRequest struct {
	Search  string `form:"search"`
	MaxAge  int    `form:"maxAge"`
	MinAge  int    `form:"minAge"`
	Job     string `form:"job"`
	Page    int    `form:"page"`
	PerPage int    `form:"perPage"`
}
