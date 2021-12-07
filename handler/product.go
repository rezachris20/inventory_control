package handler

import (
	"backend-simple-pos/auth"
	"backend-simple-pos/helper"
	"backend-simple-pos/product"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type productHandler struct {
	productService product.Service
	authSevice     auth.Service
}

func NewProductHandler(productService product.Service, authSevice auth.Service) *productHandler {
	return &productHandler{productService, authSevice}

}

func (h *productHandler) GetAllProduct(c *gin.Context) {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		response := helper.APIResponse("Failed to load products", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to load products", http.StatusOK, "success", product.FormatProducts(products))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) RegisterProduct(c *gin.Context) {
	var inputData product.InputProduct

	err := c.ShouldBind(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to register products", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//Upload File
	path := fmt.Sprintf("images/product/%s", inputData.Image.Filename)
	err = c.SaveUploadedFile(inputData.Image, path)
	if err != nil {
		response := helper.APIResponse("Failed to Upload Image", http.StatusUnprocessableEntity, "error", err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	inputData.ImagePath = path

	newProduct, err := h.productService.Save(inputData)
	if err != nil {
		response := helper.APIResponse("Failed to register products", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to register products", http.StatusOK, "success", product.FormatProduct(newProduct))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
	var inputDetailProduct product.InputDetailProduct

	err := c.ShouldBindUri(&inputDetailProduct)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Update Customer Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var inputData product.InputProduct

	err = c.ShouldBind(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to register products", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//Setting path image
	path := fmt.Sprintf("images/product/%s", inputData.Image.Filename)
	inputData.ImagePath = path

	updatedProduct, err := h.productService.Update(inputDetailProduct, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update products", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Upload image
	err = c.SaveUploadedFile(inputData.Image, path)
	if err != nil {
		response := helper.APIResponse("Failed to upload image products", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to register products", http.StatusOK, "success", product.FormatProduct(updatedProduct))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) DeleteProduct(c *gin.Context) {
	var inputDeleteProduct product.InputDeleteProduct
	err := c.ShouldBindJSON(&inputDeleteProduct)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to delete products", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	deleteProduct, err := h.productService.Delete(inputDeleteProduct)
	if err != nil {
		response := helper.APIResponse("Failed to delete products", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete products", http.StatusOK, "success", deleteProduct)
	c.JSON(http.StatusOK, response)
	return
}
