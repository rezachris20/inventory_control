package supplier

import "backend-simple-pos/category_product"

type SupplierFormatter struct {
	ID                int                              `json:"id"`
	Nama              string                           `json:"nama"`
	NoHp              string                           `json:"no_hp"`
	Alamat            string                           `json:"alamat"`
	Email             string                           `json:"email"`
	CategoriProductID int                              `json:"categori_product_id"`
	CategoriProduct   category_product.CategoriProduct `json:"categori_product"`
}

func FormatSupplier(supplier Supplier) SupplierFormatter {
	formatter := SupplierFormatter{
		ID:                supplier.ID,
		Nama:              supplier.Nama,
		NoHp:              supplier.NoHp,
		Alamat:            supplier.Alamat,
		Email:             supplier.Email,
		CategoriProductID: supplier.CategoriProductID,
		CategoriProduct:   supplier.CategoriProduct,
	}

	return formatter
}

func FormatSuppliers(suppliers []Supplier) []SupplierFormatter {
	var suppliersFormatter []SupplierFormatter

	for _, supplier := range suppliers {
		supplierFormatter := FormatSupplier(supplier)
		suppliersFormatter = append(suppliersFormatter, supplierFormatter)
	}

	return suppliersFormatter
}
