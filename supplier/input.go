package supplier

type InputRegisterSupplier struct {
	Nama              string `json:"nama" binding:"required"`
	NoHp              string `json:"no_hp" binding:"required"`
	Alamat            string `json:"alamat" binding:"required"`
	Email             string `json:"email" binding:"required"`
	CategoriProductID int    `json:"categori_product_id" binding:"required"`
}

type InputDetailSupplier struct {
	ID int `uri:"id" binding:"required"`
}

type InputDeleteSupplier struct {
	ID int `json:"id" binding:"required"`
}
