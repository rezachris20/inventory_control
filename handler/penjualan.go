package handler

import (
	"backend-simple-pos/auth"
	"backend-simple-pos/helper"
	"backend-simple-pos/penjualan"
	"backend-simple-pos/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type penjualanHandler struct {
	penjualanService penjualan.Service
	authService      auth.Service
}

func NewPenjualanHandler(penjualanService penjualan.Service, authService auth.Service) *penjualanHandler {
	return &penjualanHandler{penjualanService, authService}
}

func (h *penjualanHandler) GetAllPemesanan(c *gin.Context) {
	pemesanan, err := h.penjualanService.GetAllPemesanan()
	if err != nil {
		response := helper.APIResponse("Gagal load pemesanan", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Sukses load pemesanan", http.StatusOK, "success", penjualan.FormatPemesanansNew(pemesanan))
	c.JSON(http.StatusOK, response)
}

func (h *penjualanHandler) CreatePemesanan(c *gin.Context) {
	var input penjualan.InputMultiplePemesanan

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Gagal membuat record pemesanan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	pemesanan, err := h.penjualanService.CreateMultiplePemesanan(input)
	if err != nil {
		response := helper.APIResponse("Gagal membuat record pemesanan", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Sukses membuat record pemesanan", http.StatusOK, "success", penjualan.FormatPemesananNew(pemesanan))
	c.JSON(http.StatusOK, response)
}

func (h *penjualanHandler) UpdatePemesananDetail(c *gin.Context) {
	var inputID penjualan.InputIDPemesananDetail
	if err := c.ShouldBindUri(&inputID); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Gagal update record pemesanan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var inputData penjualan.InputPemesanaDetail
	if err := c.ShouldBindJSON(&inputData); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Gagal update record pemesanan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	update, err := h.penjualanService.UpdatePemesananDetail(&inputID, &inputData)
	if err != nil {
		response := helper.APIResponse("Gagal update record pemesanan", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Sukses update record pemesanan", http.StatusOK, "success", penjualan.FormatPemesananNew(update))
	c.JSON(http.StatusOK, response)
}

func (h *penjualanHandler) DeletePemesananDetail(c *gin.Context) {
	var inputID penjualan.InputIDPemesananDetail

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Gagal delete record pemesanan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	res, err := h.penjualanService.DeletePemesananDetail(&inputID)
	if err != nil {
		response := helper.APIResponse("Gagal delete record pemesanan", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Sukses delete record pemesanan", http.StatusOK, "success", res)
	c.JSON(http.StatusOK, response)

}
