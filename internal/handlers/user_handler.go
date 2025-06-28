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

// Signup godoc
// @Summary Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param request body models.UserCreateRequest true "User creation request"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Failure 409 {object} apperror.AppError
// @Router /signup [post]
func (h *UserHandler) Signup(c *gin.Context) {
	var req models.UserCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		HandleError(c, apperror.ErrBadRequest)
		return
	}

	user, err := h.userService.CreateUser(req.Username, req.Email, req.Password)
	if err != nil {
		HandleError(c, err)
		return
	}

	response := models.UserResponse{
		Message: "User created successfully",
		User: models.UserInner{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			Role:      user.Role,
		},
	}

	c.JSON(http.StatusCreated, response)
}

// Login godoc
// @Summary Login as an existing user
// @Tags users
// @Accept json
// @Produce json
// @Param request body models.UserLoginRequest true "User login request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} apperror.AppError
// @Failure 404 {object} apperror.AppError
// @Failure 401 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var userlogin models.UserLoginRequest

	if err := c.ShouldBindJSON(&userlogin); err != nil {
		HandleError(c, apperror.ErrBadRequest)
		return
	}

	token, err := h.userService.LoginUser(userlogin.Username, userlogin.Password)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// GetUsers godoc
// @Summary Retrieve all existing users
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} apperror.AppError
// @Router /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetUsers()
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}
