package controllers

import (
	"crud-go-echo-gorm/models"
	"crud-go-echo-gorm/services"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (uc *UserController) Create(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	createdUser, err := uc.UserService.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(http.StatusCreated, createdUser)
}

func (uc *UserController) FindAll(c echo.Context) error {
	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch users"})
	}

	return c.JSON(http.StatusOK, users)
}

func (uc *UserController) FindByID(c echo.Context) error {
	userID := c.Param("id")
	user, err := uc.UserService.GetUserByID(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}

func (uc *UserController) Update(c echo.Context) error {
	userID := c.Param("id")

	var updatedUser models.User
	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	user, err := uc.UserService.UpdateUser(userID, updatedUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
	}

	return c.JSON(http.StatusOK, user)
}

func (uc *UserController) Delete(c echo.Context) error {
	userID := c.Param("id")
	fmt.Println("-------------------------------------------")
	fmt.Println("-------------------------------------------")
	fmt.Println("Chegou no delete ID: ", userID)
	fmt.Println("-------------------------------------------")
	fmt.Println("-------------------------------------------")
	err := uc.UserService.DeleteUser(userID)
	if err != nil {
		// Verifica se o erro é relacionado à ausência do usuário
		if strings.Contains(err.Error(), "record not found") {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		// Outros erros internos
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
	}

	return c.NoContent(http.StatusNoContent)
}
