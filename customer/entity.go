package customer

import "time"

type Customer struct {
	ID        int       `json:"id"`
	Nama      string    `json:"nama"`
	NoHp      string    `json:"no_hp"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
