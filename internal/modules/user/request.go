package user

type CreateUserRequest struct {
	Name string `json:"name"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}