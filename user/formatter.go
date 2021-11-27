package user

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	// AvatarFileName string    `json:"avatarFileName"`
	Token string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {

	newUser := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}

	return newUser

}
