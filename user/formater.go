package user

import "backend-simple-pos/role"

type UserFormater struct {
	ID       int       `json:"id"`
	Nama     string    `json:"nama"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Token    string    `json:"token"`
	RoleID   int       `json:"role_id"`
	ImageURL string    `json:"image_url"`
	Role     role.Role `json:"role"`
}

func FormatUser(user User, token string) UserFormater {
	formater := UserFormater{
		ID:       user.ID,
		Nama:     user.Nama,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
		RoleID:   user.RoleID,
		ImageURL: user.AvatarFileName,
		Role:     user.Role,
	}
	return formater
}

func FormatUsers(users []User) []UserFormater {
	usersFormatter := []UserFormater{}

	for _, user := range users {
		userFormatter := FormatUser(user, "")
		usersFormatter = append(usersFormatter, userFormatter)
	}

	return usersFormatter
}
