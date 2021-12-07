package handler

import (
	"backend-simple-pos/auth"
	"backend-simple-pos/customer"
	"backend-simple-pos/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type customerHandler struct {
	customerService customer.Service
	authService     auth.Service
}

func NewCustomerHandler(customerService customer.Service, authService auth.Service) *customerHandler {
	return &customerHandler{customerService, authService}
}

func (h *customerHandler) RegisterCustomer(c *gin.Context) {
	var input customer.RegisterCustomerInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Register customer failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newCustomer, err := h.customerService.RegisterCustomer(input)
	if err != nil {
		response := helper.APIResponse("Register customer failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Register customer success", http.StatusOK, "success", newCustomer)
	c.JSON(http.StatusOK, response)
}

func (h *customerHandler) UpdateCustomer(c *gin.Context) {
	var inputID customer.InputDetailCustomer
	var inputData customer.RegisterCustomerInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Update Customer Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to updated user", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedCustomer, err := h.customerService.UpdateCustomer(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Update Customer Failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success Update Customer", http.StatusOK, "success", updatedCustomer)
	c.JSON(http.StatusOK, response)
}

func (h *customerHandler) DeleteCustomer(c *gin.Context) {
	var input customer.InputDeleteCustomer

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to delete customer", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.customerService.DeleteCustomer(input)
	if err != nil {
		response := helper.APIResponse("Failed to delete customer", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete customer", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *customerHandler) GetAllCustomer(c *gin.Context) {
	customers, err := h.customerService.GetAllCustomer()
	if err != nil {
		response := helper.APIResponse("Failed to load customers", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to load customer", http.StatusOK, "success", customers)
	c.JSON(http.StatusOK, response)

}
