package user

type RegisterUserInput struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type CheckEmailInput struct {
	Email string `json:"email" validate:"required,email"`
}
