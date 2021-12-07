package penjualan

import (
	"backend-simple-pos/constant"
	"fmt"
)

type PemesananFormatter struct {
	ID               int     `json:"id"`
	Noref            string  `json:"noref"`
	TanggalPemesanan string  `json:"tanggal_pemesanan"`
	Qty              int     `json:"qty"`
	HargaSatuan      int     `json:"harga_satuan"`
	Total            float32 `json:"total"`
	Satuan           int     `json:"satuan"`
	ProductID        int     `json:"product_id"`
	ProductName      string  `json:"product_name"`
}

type NewFormatPemesanan struct {
	ID               int                        `json:"id"`
	Noref            string                     `json:"noref"`
	TanggalPemesanan string                     `json:"tanggal_pemesanan"`
	DetailPemesanan  []NewFormatPemesananDetail `json:"detail_pemesanan"`
}

type NewFormatPemesananDetail struct {
	ID          int     `json:"id"`
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Qty         int     `json:"qty"`
	HargaSatuan int     `json:"harga_satuan"`
	Total       float32 `json:"total"`
	Status      int     `json:"status"`
}

func FormatPemesananNew(pemesanan InvPemesanan) NewFormatPemesanan {

	formatter := NewFormatPemesanan{}
	formatter.ID = pemesanan.ID
	formatter.Noref = constant.Pemesanan + "." + pemesanan.Noref
	formatter.TanggalPemesanan = pemesanan.TanggalPemesanan.Format("2006-01-02")
	formatter.DetailPemesanan = FormatPemesanansDetailNew(pemesanan.InvPemesananDetail)
	fmt.Println(pemesanan.InvPemesananDetail)
	return formatter
}

func FormatPemesanansNew(pemesanan []InvPemesanan) []NewFormatPemesanan {

	var newFormatPemesanans []NewFormatPemesanan

	for _, detail := range pemesanan {
		newFormatPemesanan := FormatPemesananNew(detail)
		newFormatPemesanans = append(newFormatPemesanans, newFormatPemesanan)
	}
	return newFormatPemesanans
}

func FormatPemesananDetailNew(detail InvPemesananDetail) NewFormatPemesananDetail {
	formatter := NewFormatPemesananDetail{}
	formatter.ID = detail.ID
	formatter.ProductID = detail.ProductID
	formatter.Qty = detail.Qty
	formatter.HargaSatuan = detail.HargaSatuan
	formatter.Total = detail.Total
	formatter.Status = detail.Status
	formatter.ProductName = detail.Product.Nama

	return formatter
}

func FormatPemesanansDetailNew(details []InvPemesananDetail) []NewFormatPemesananDetail {
	var formatDetails []NewFormatPemesananDetail

	for _, detail := range details {
		formatDetail := FormatPemesananDetailNew(detail)
		formatDetails = append(formatDetails, formatDetail)
	}
	return formatDetails
}

func FormatPemesanan(detail PemesananWithDetail) PemesananFormatter {

	formatter := PemesananFormatter{
		ID:               detail.ID,
		Noref:            detail.Noref,
		TanggalPemesanan: detail.TanggalPemesanan.Format("2006-01-02"),
		Qty:              detail.Qty,
		HargaSatuan:      detail.HargaSatuan,
		Total:            detail.Total,
		Satuan:           detail.Status,
		ProductID:        detail.ProductID,
		ProductName:      detail.ProductName,
	}

	return formatter
}
