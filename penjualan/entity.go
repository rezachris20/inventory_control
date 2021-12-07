package penjualan

import (
	"backend-simple-pos/product"
	"time"
)

type InvPemesanan struct {
	ID                 int
	Noref              string
	TanggalPemesanan   time.Time
	UserID             int
	CreatedAt          time.Time
	UpdatedAt          time.Time
	InvPemesananDetail []InvPemesananDetail
}

type InvPemesananDetail struct {
	ID             int
	InvPemesananID int
	ProductID      int
	Product        product.Product
	Qty            int
	HargaSatuan    int
	Total          float32
	Status         int
}

type PemesananWithDetail struct {
	ID               int
	Noref            string
	TanggalPemesanan time.Time
	Qty              int
	HargaSatuan      int
	Total            float32
	Status           int
	ProductID        int
	ProductName      string
}
