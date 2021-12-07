package penjualan

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Service interface {
	GetAllPemesanan() ([]InvPemesanan, error)
	CreateMultiplePemesanan(inputMultiplePemesanan InputMultiplePemesanan) (InvPemesanan, error)
	UpdatePemesananDetail(inputDetailID *InputIDPemesananDetail, inputDetailData *InputPemesanaDetail) (InvPemesanan, error)
	DeletePemesananDetail(detail *InputIDPemesananDetail) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllPemesanan() ([]InvPemesanan, error) {
	pemesanan, err := s.repository.GetAll()
	if err != nil {
		return pemesanan, err
	}

	return pemesanan, nil
}

func (s *service) CreateMultiplePemesanan(inputMultiplePemesanan InputMultiplePemesanan) (InvPemesanan, error) {
	// Generate Noref Otomatis
	noref := s.repository.GetNoref()
	newNoref := ""

	if noref != "" {
		result, _ := strconv.Atoi(noref)
		newNoref = fmt.Sprintf("%06d", result+1)
	} else {
		newNoref = "000001"
	}

	date, _ := time.Parse("2006-01-02", inputMultiplePemesanan.TanggalPemesanan)

	pemesanan := InvPemesanan{
		Noref:            newNoref,
		TanggalPemesanan: date,
		UserID:           inputMultiplePemesanan.User.ID,
	}

	newPemesanan, err := s.repository.SaveMain(pemesanan)
	if err != nil {
		return newPemesanan, err
	}
	var invDetailPemesanans []InvPemesananDetail

	for _, detail := range inputMultiplePemesanan.Data {

		data := InvPemesananDetail{
			ProductID:      detail.ProductID,
			Qty:            detail.Qty,
			HargaSatuan:    detail.HargaSatuan,
			Total:          float32(detail.Qty * detail.HargaSatuan),
			InvPemesananID: newPemesanan.ID,
		}

		invDetailPemesanans = append(invDetailPemesanans, data)
	}

	_, err = s.repository.SaveDetail(invDetailPemesanans)

	if err != nil {
		return newPemesanan, err
	}

	result, _ := s.repository.GetByID(newPemesanan.ID)

	return result, nil
}

func (s *service) UpdatePemesananDetail(inputDetailID *InputIDPemesananDetail, inputDetailData *InputPemesanaDetail) (InvPemesanan, error) {
	var invPemesanan InvPemesanan

	pemesanan, err := s.repository.GetByIDDetail(inputDetailID.ID)
	if err != nil {
		return invPemesanan, err
	}

	pemesanan.ProductID = inputDetailData.ProductID
	pemesanan.Qty = inputDetailData.Qty
	pemesanan.HargaSatuan = inputDetailData.HargaSatuan
	pemesanan.Total = float32(inputDetailData.Qty * inputDetailData.HargaSatuan)

	detail, err := s.repository.UpdateDetail(pemesanan)
	if err != nil {
		return invPemesanan, err
	}

	result, _ := s.repository.GetByID(detail.InvPemesananID)

	return result, nil
}

func (s *service) DeletePemesananDetail(detail *InputIDPemesananDetail) (bool, error) {
	_, err := s.repository.GetByIDDetail(detail.ID)
	if err != nil {
		return false, errors.New("No supplier found that ID")
	}

	response, err := s.repository.DeleteDetail(detail.ID)
	if err != nil {
		return false, err
	}

	return response, nil
}
