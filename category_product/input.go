package category_product

type InputCategoryProduct struct {
	Nama string `json:"nama" binding:"required"`
}

type InputDetailCategoryProduct struct {
	ID int `uri:"id" binding:"required"`
}

type InputDeleteCategoryProduct struct {
	ID int `json:"id" binding:"required"`
}
