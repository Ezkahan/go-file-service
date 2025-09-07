package http

import (
	"fmt"
	"net/http"

	"github.com/ezkahan/go-file-service/internal/delivery/http/validators"
	"github.com/ezkahan/go-file-service/internal/domain"
	"github.com/ezkahan/go-file-service/internal/usecase"
	"github.com/ezkahan/go-file-service/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Profile(ctx *gin.Context)
}

type authHandler struct {
	userService usecase.UserService
}

func NewAuthHandler(userService usecase.UserService) AuthHandler {
	return &authHandler{userService}
}

// Login handler
func (h *authHandler) Login(ctx *gin.Context) {
	var req validators.LoginRequest
	validate := validator.New()

	fmt.Println("login handler")

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": http.StatusBadRequest})
		return
	}

	if err := validate.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": utils.ParseValidationError(err), "code": http.StatusBadRequest})
		return
	}

	u, err := h.userService.VerifyCredential(ctx, req.Username, req.Password)
	if err != nil || u == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Username or password invalid", "code": http.StatusBadRequest})
		return
	}

	token, err := utils.GenerateToken(u.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token", "code": http.StatusInternalServerError})
		return
	}

	u.Token = token
	userRes := &domain.LoginResponse{
		ID:        u.ID,
		Username:  u.Username,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Role:      u.Role,
		Token:     u.Token,
	}

	ctx.JSON(http.StatusOK, gin.H{"data": userRes})
}

// Register handler
func (h *authHandler) Register(ctx *gin.Context) {
	var req validators.SaveUserRequest
	validate := validator.New()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": http.StatusBadRequest})
		return
	}

	if err := validate.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": utils.ParseValidationError(err), "code": http.StatusBadRequest})
		return
	}

	req.IP = ctx.ClientIP()
	req.Device = ctx.Request.UserAgent()

	newUser, err := h.userService.Save(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "code": http.StatusInternalServerError})
		return
	}

	token, err := utils.GenerateToken(newUser.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token", "code": http.StatusInternalServerError})
		return
	}
	newUser.Token = token

	ctx.JSON(http.StatusOK, gin.H{"data": newUser})
}

// Profile handler
func (h *authHandler) Profile(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "code": http.StatusUnauthorized})
		return
	}

	// Assuming userID is stored as string (UUID) in middleware
	u, err := h.userService.GetByID(userID.(uint)) // 🔧 change this if IDs are UUID strings
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": u})
}
