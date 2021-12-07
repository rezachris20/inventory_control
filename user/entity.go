package user

import (
	"backend-simple-pos/role"
	"time"
)

type User struct {
	ID             int
	Nama           string
	Username       string
	Email          string
	Password       string
	AvatarFileName string
	RoleID         int
	Role           role.Role
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

