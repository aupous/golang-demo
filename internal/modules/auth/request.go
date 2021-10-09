package auth

type SignInRequset struct {
	Email    string `json:"email" bind:"required,email"`
	Password string `json:"password" bind:"required"`
}
