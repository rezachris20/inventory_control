package user

type RegisterUserInput struct {
	Nama     string `json:"nama" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleID   int    `json:"role_id" binding:"required"`
}

type UpdateUserInput struct {
	Nama     string `json:"nama" `
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleID   int    `json:"role_id"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetUserDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type GetUserDeleteInput struct {
	ID int `json:"id" binding:"required"`
}
