package penjualan

import (
	"backend-simple-pos/user"
)

type InputPemesanan struct {
	TanggalPemesanan string `json:"tanggal_pemesanan" binding:"required"`
	ProductID        int    `json:"product_id" binding:"required"`
	Qty              int    `json:"qty" binding:"required"`
	HargaSatuan      int    `json:"harga_satuan" binding:"required"`
	User             user.User
}

type InputMultiplePemesanan struct {
	TanggalPemesanan string `json:"tanggal_pemesanan" binding:"required"`
	Data             []Data `json:"data" binding:"required"`
	User             user.User
}

type Data struct {
	ProductID   int `json:"product_id" binding:"required"`
	Qty         int `json:"qty" binding:"required"`
	HargaSatuan int `json:"harga_satuan" binding:"required"`
}

type InputIDPemesananDetail struct {
	ID int `uri:"id" binding:"required"`
}

type InputPemesanaDetail struct {
	ProductID   int `json:"product_id" binding:"required"`
	Qty         int `json:"qty" binding:"required"`
	HargaSatuan int `json:"harga_satuan" binding:"required"`
}
