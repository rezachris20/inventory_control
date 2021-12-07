package supplier

import (
	"backend-simple-pos/category_product"
	"time"
)

type Supplier struct {
	ID                int                              `json:"id"`
	Nama              string                           `json:"nama"`
	NoHp              string                           `json:"no_hp"`
	Alamat            string                           `json:"alamat"`
	Email             string                           `json:"email"`
	CategoriProductID int                              `json:"categori_product_id"`
	CategoriProduct   category_product.CategoriProduct `json:"categori_product" gorm:"foreignKey:CategoriProductID"`
	CreatedAt         time.Time                        `json:"created_at"`
	UpdatedAt         time.Time                        `json:"updated_at"`
}
