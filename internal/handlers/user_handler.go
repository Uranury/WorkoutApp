package handlers

import (
	"net/http"

	"github.com/Uranury/WorkoutApp/internal/apperror"
	"github.com/Uranury/WorkoutApp/internal/models"
	"github.com/Uranury/WorkoutApp/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func HandleError(c *gin.Context, err error) {
	if appErr, ok := err.(*apperror.AppError); ok {
		c.JSON(appErr.StatusCode, gin.H{"error": appErr.Message})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}

func (h *UserHandler) Signup(c *gin.Context) {
	var req models.UserCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Convert request to user model
	user := req.ToUser()

	if err := h.userService.CreateUser(user); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *UserHandler) Login(c *gin.Context) {
	var userlogin models.UserLoginRequest

	if err := c.ShouldBindJSON(&userlogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := h.userService.LoginUser(userlogin.Username, userlogin.Password)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetUsers()
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}
