package handler

import (
	"backend-simple-pos/auth"
	"backend-simple-pos/helper"
	"backend-simple-pos/supplier"
	"github.com/gin-gonic/gin"
	"net/http"
)

type supplierHandler struct {
	supplierService supplier.Service
	authService     auth.Service
}

func NewSupplierHandler(supplierService supplier.Service, authService auth.Service) *supplierHandler {
	return &supplierHandler{supplierService, authService}
}

func (h *supplierHandler) GetAllSupliers(c *gin.Context) {
	suppliers, err := h.supplierService.GetAllSupplier()
	if err != nil {
		response := helper.APIResponse("Failed to load suppliers", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to load suppliers", http.StatusOK, "success", supplier.FormatSuppliers(suppliers))
	c.JSON(http.StatusOK, response)
}

func (h *supplierHandler) RegisterSupplier(c *gin.Context) {
	var input supplier.InputRegisterSupplier

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Register supplier failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newSupplier, err := h.supplierService.RegisterSupplier(input)
	if err != nil {
		response := helper.APIResponse("Register supplier failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Register supplier success", http.StatusOK, "success", supplier.FormatSupplier(newSupplier))
	c.JSON(http.StatusOK, response)
	return
}

func (h *supplierHandler) UpdateSupplier(c *gin.Context) {
	var inputID supplier.InputDetailSupplier
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Update supplier failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var inputData supplier.InputRegisterSupplier
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Update supplier failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusRequestEntityTooLarge, response)
		return
	}

	updateSupplier, err := h.supplierService.UpdateSupplier(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Update supplier failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Update supplier success", http.StatusOK, "error", supplier.FormatSupplier(updateSupplier))
	c.JSON(http.StatusOK, response)
	return
}

func (h *supplierHandler) DeleteSupplier(c *gin.Context) {
	var inputID supplier.InputDeleteSupplier

	err := c.ShouldBindJSON(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to delete supplier", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	deleteSupplier, err := h.supplierService.DeleteSupplier(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete supplier", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete supplier", http.StatusOK, "success", deleteSupplier)
	c.JSON(http.StatusOK, response)
}
