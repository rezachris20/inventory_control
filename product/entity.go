package product

import (
	"backend-simple-pos/supplier"
	"time"
)

type Product struct {
	ID         int               `json:"id"`
	Nama       string            `json:"nama"`
	SupplierID int               `json:"supplier_id"`
	Supplier   supplier.Supplier `json:"supplier"`
	Satuan     string            `json:"satuan"`
	Hargabeli  int               `json:"hargabeli"`
	Hargajual  int               `json:"hargajual"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	Image      string            `json:"image"`
}
