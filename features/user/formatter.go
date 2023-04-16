package user

type RegisterUserFormatter struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// Token string `json:"token"`
}

type LoginUserFormatter struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	AvatarFileName string `json:"avatar"`
	Token          string `json:"token"`
}

func FormatRegisterUser(user User) RegisterUserFormatter {
	formatter := RegisterUserFormatter{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	return formatter
}
func FormatLoginUser(user User, token string) LoginUserFormatter {
	formatter := LoginUserFormatter{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		AvatarFileName: user.AvatarFileName,
		Token:          token,
	}
	return formatter
}
