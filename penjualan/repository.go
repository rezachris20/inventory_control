package penjualan

import (
	"errors"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllPemesanan() ([]PemesananWithDetail, error)
	GetAllPemesananWithID(ID int) (PemesananWithDetail, error)
	SaveMain(pemesanan InvPemesanan) (InvPemesanan, error)
	SaveDetail(detail []InvPemesananDetail) ([]InvPemesananDetail, error)
	GetAll() ([]InvPemesanan, error)
	GetByID(ID int) (InvPemesanan, error)
	GetByIDDetail(ID int) (InvPemesananDetail, error)
	GetNoref() string
	UpdateDetail(detail InvPemesananDetail) (InvPemesananDetail, error)
	DeleteDetail(ID int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (s *repository) SaveMain(pemesanan InvPemesanan) (InvPemesanan, error) {
	err := s.db.Create(&pemesanan).Error
	if err != nil {
		return pemesanan, err
	}

	return pemesanan, nil
}

func (s *repository) SaveDetail(detail []InvPemesananDetail) ([]InvPemesananDetail, error) {
	err := s.db.Create(&detail).Error
	if err != nil {
		return detail, err
	}

	return detail, nil
}

func (s *repository) GetNoref() string {
	var result string
	row := s.db.Model(&InvPemesanan{}).Select("max(noref) AS noref").Row()
	row.Scan(&result)
	return result
}

func (s *repository) GetAllPemesanan() ([]PemesananWithDetail, error) {
	var invPemesanan []PemesananWithDetail

	err := s.db.Table("inv_pemesanans ip ").
		Select("ip.id,ip.noref,ip.tanggal_pemesanan,ipd.qty,ipd.harga_satuan,ipd.total,ipd.status,ipd.product_id,p.nama").
		Joins("left join inv_pemesanan_details ipd on ip.id = ipd.inv_pemesanan_id").
		Joins("left join products p on p.id = ipd.product_id").
		Scan(&invPemesanan).Error

	if err != nil {
		return invPemesanan, err
	}

	return invPemesanan, nil
}

func (s *repository) GetAllPemesananWithID(ID int) (PemesananWithDetail, error) {
	var result PemesananWithDetail

	err := s.db.Table("inv_pemesanans ip ").
		Select("ip.id,ip.noref,ip.tanggal_pemesanan,ipd.qty,ipd.harga_satuan,ipd.total,ipd.status,ipd.product_id,p.nama").
		Where("ip.id = ? ", ID).
		Joins("left join inv_pemesanan_details ipd on ip.id = ipd.inv_pemesanan_id").
		Joins("left join products p on p.id = ipd.product_id").
		Scan(&result).Error

	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *repository) GetAll() ([]InvPemesanan, error) {
	var invPemesanan []InvPemesanan
	err := s.db.Preload("InvPemesananDetail.Product.Supplier.CategoriProduct").Find(&invPemesanan).Error

	if err != nil {
		return invPemesanan, err
	}
	return invPemesanan, nil
}

func (s *repository) GetByID(ID int) (InvPemesanan, error) {
	var invPemesanan InvPemesanan
	err := s.db.Preload("InvPemesananDetail.Product.Supplier.CategoriProduct").Where("id = ? ", ID).First(&invPemesanan).Error

	if err != nil {
		return invPemesanan, err
	}
	return invPemesanan, nil
}

func (s *repository) GetByIDDetail(ID int) (InvPemesananDetail, error) {
	var invPemesananDetail InvPemesananDetail

	if err := s.db.Where("id = ? ", ID).First(&invPemesananDetail).Error; err != nil {
		return invPemesananDetail, errors.New("record not found")
	}
	return invPemesananDetail, nil
}

func (s *repository) UpdateDetail(detail InvPemesananDetail) (InvPemesananDetail, error) {
	if err := s.db.Save(&detail).Error; err != nil {
		return detail, err
	}
	return detail, nil
}

func (s *repository) DeleteDetail(ID int) (bool, error) {
	var detail InvPemesananDetail

	if err := s.db.Where("id = ?", ID).Delete(&detail).Error; err != nil {
		return false, err
	}

	return true, nil
}
