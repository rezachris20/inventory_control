package handler

import (
	"backend-simple-pos/auth"
	"backend-simple-pos/helper"
	"backend-simple-pos/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *UserHandler {
	return &UserHandler{userService, authService}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Register account gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account gagal", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("Register account gagal", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formater := user.FormatUser(newUser, token)
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formater)

	c.JSON(http.StatusOK, response)

}

func (h *UserHandler) Login(c *gin.Context) {
	// User memasukkan input (email dan password)
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnauthorized, "error", errorMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedInUser.ID)
	if err != nil {
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedInUser, token)

	response := helper.APIResponse("Successfully loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	formatter := user.FormatUser(currentUser, "")

	response := helper.APIResponse("Successfully fetch user data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) Users(c *gin.Context) {
	users, err := h.userService.GetAllUsers()

	if err != nil {
		response := helper.APIResponse("Load Data Users Failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Successfully Load Data Users", http.StatusOK, "success", user.FormatUsers(users))

	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetUserDetail(c *gin.Context) {
	var inputID *user.GetUserDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to fetch data user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userByID, err := h.userService.GetUserByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed to fetch data user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success load data users", http.StatusOK, "success", user.FormatUser(userByID, ""))
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) UpdatedUser(c *gin.Context) {
	var inputID user.GetUserDetailInput
	var inputData user.UpdateUserInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to updated user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
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

	updatedUser, err := h.userService.UpdateUser(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to updated user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success updated user", http.StatusOK, "success", user.FormatUser(updatedUser, ""))
	c.JSON(http.StatusOK, response)

}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	var inputID user.GetUserDeleteInput

	err := c.ShouldBindJSON(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.userService.DeleteUser(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete user", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// simpan gambar di folder "images/"
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	path := fmt.Sprintf("images/user/%d-%s", userID, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}

		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// service : panggil repo
	// JWT (sementara hardcore, seakan2 user yang memiliki ID 1)
	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}

		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar successfully uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}
