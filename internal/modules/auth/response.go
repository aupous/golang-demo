package auth

import "awesomeProject/internal/model"

type SignInResponse struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
}
