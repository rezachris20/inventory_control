package handler

import (
	"backend-simple-pos/auth"
	"backend-simple-pos/category_product"
	"backend-simple-pos/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type categoryProductHandler struct {
	categoryProductService category_product.Service
	authService            auth.Service
}

func NewCategoryProductHandler(categoryProductService category_product.Service, authService auth.Service) *categoryProductHandler {
	return &categoryProductHandler{categoryProductService, authService}
}

func (h *categoryProductHandler) GetAllProductCategory(c *gin.Context) {
	categoryProducts, err := h.categoryProductService.GetAllCategoryProduct()
	if err != nil {
		response := helper.APIResponse("Failed to load category product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to load category product", http.StatusOK, "success", categoryProducts)
	c.JSON(http.StatusOK, response)
}

func (h *categoryProductHandler) RegisterProductCategory(c *gin.Context) {
	var input category_product.InputCategoryProduct

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors":errors}

		response := helper.APIResponse("Failed to register category product",http.StatusUnprocessableEntity,"error",errorMessage)
		c.JSON(http.StatusUnprocessableEntity,response)
		return
	}

	newProductCategory, err := h.categoryProductService.RegisterCategoryProduct(input)
	if err != nil {
		response := helper.APIResponse("Failed to register category product",http.StatusBadRequest,"error",err)
		c.JSON(http.StatusBadRequest,response)
		return
	}

	response := helper.APIResponse("Success to register category product",http.StatusOK,"success",newProductCategory)
	c.JSON(http.StatusOK,response)
}

func (h *categoryProductHandler) UpdateProductCategory(c *gin.Context) {
	var inputID category_product.InputDetailCategoryProduct
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors":errors}

		response := helper.APIResponse("Failed to update category product",http.StatusUnprocessableEntity,"error",errorMessage)
		c.JSON(http.StatusUnprocessableEntity,response)
		return
	}

	var inputData category_product.InputCategoryProduct
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors":errors}

		response := helper.APIResponse("Failed to update category product",http.StatusUnprocessableEntity,"error",errorMessage)
		c.JSON(http.StatusUnprocessableEntity,response)
		return
	}

	updateProductCategory,err := h.categoryProductService.UpdateCategoryProduct(inputID,inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update category product",http.StatusBadRequest,"error",err)
		c.JSON(http.StatusBadRequest,response)
		return
	}
	response := helper.APIResponse("Success to update category product",http.StatusOK,"error",updateProductCategory)
	c.JSON(http.StatusOK,response)
}

func (h *categoryProductHandler) DeleteProductCategory(c *gin.Context){
	var input category_product.InputDeleteCategoryProduct

	err := c.ShouldBindJSON(&input)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors" : error}

		response := helper.APIResponse("Failed delete category product",http.StatusUnprocessableEntity,"error",errorMessage)
		c.JSON(http.StatusUnprocessableEntity,response)
		return
	}

	deleteCategoryProduct, err := h.categoryProductService.DeleteCategoryProduct(input)
	if err != nil {
		response := helper.APIResponse("Failed delete category product",http.StatusBadRequest,"error",err)
		c.JSON(http.StatusBadRequest,response)
		return
	}

	response := helper.APIResponse("Success delete category product",http.StatusOK,"success",deleteCategoryProduct)
	c.JSON(http.StatusOK,response)
}