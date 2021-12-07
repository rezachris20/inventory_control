package product

import (
	"backend-simple-pos/supplier"
)

type ProductFormatter struct {
	ID         int               `json:"id"`
	Nama       string            `json:"nama"`
	SupplierID int               `json:"supplier_id"`
	Supplier   supplier.Supplier `json:"supplier"`
	Satuan     string            `json:"satuan"`
	Hargabeli  int               `json:"hargabeli"`
	Hargajual  int               `json:"hargajual"`
	Image      string            `json:"image"`
}

func FormatProduct(product Product) ProductFormatter {
	formatter := ProductFormatter{
		ID:         product.ID,
		Nama:       product.Nama,
		SupplierID: product.SupplierID,
		Supplier:   product.Supplier,
		Satuan:     product.Satuan,
		Hargabeli:  product.Hargabeli,
		Hargajual:  product.Hargajual,
		Image:      product.Image,
	}
	return formatter
}

func FormatProducts(products []Product) []ProductFormatter {
	var productsFormatter []ProductFormatter

	for _, product := range products {
		productFormatter := FormatProduct(product)
		productsFormatter = append(productsFormatter, productFormatter)
	}

	return productsFormatter
}
