package product

import "mime/multipart"

type InputProduct struct {
	Nama       string                `form:"nama" binding:"required"`
	SupplierID int                   `form:"supplier_id" binding:"required"`
	Satuan     string                `form:"satuan" binding:"required"`
	Hargabeli  int                   `form:"hargabeli" binding:"required"`
	Hargajual  int                   `form:"hargajual" binding:"required"`
	Image      *multipart.FileHeader `form:"image"`
	ImagePath  string
}

type InputDetailProduct struct {
	ID int `uri:"id" binding:"required"`
}

type InputDeleteProduct struct {
	ID int `json:"id" binding:"required"`
}
