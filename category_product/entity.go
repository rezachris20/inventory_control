package category_product

import "time"

type CategoriProduct struct {
	ID        int       `json:"id"`
	Nama      string    `json:"nama"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
