package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rams/pkg/models"
	"strconv"
)

type signInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary	CreateNewUser
// @Description	`Регистрация пользователя в систему.`
// @Tags		Auth
// @VendorCode	create account
// @ID			create-account
// @Accept		json
// @Produce	json
// @Param		input	body		models.User	true	"account info"
// @Success	200		{integer}	integer
// @Router		/auth/sign_up [post]
func (h *Handler) SignUp(ctx *gin.Context) {
	var input models.User
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"reason": "invalid input body"})
		return
	}

	if err := h.validate.Struct(input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := h.services.Authorization.CreateUser(ctx, input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"reason": "created!"})
}

// @Summary	login user
// @Description	`Аутентификация пользователя в систему.`
// @Tags		Auth
// @VendorCode	login user
// @ID			login
// @Accept		json
// @Produce	json
// @Param		input	body		signInInput	true	"account info"
// @Success	200		{object}	models.AuthLoginResponse
// @Failure	403		{object}	map[string]interface{}
// @Failure	400,404	{object}	map[string]interface{}
// @Failure	500		{object}	map[string]interface{}
// @Failure	default	{object}	map[string]interface{}
// @Router		/auth/login [post]
func (h *Handler) SignIn(ctx *gin.Context) {
	var input signInInput

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	token, err := h.services.Authorization.GenerateToken(ctx, input.Login, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	account, err := h.services.Authorization.GetUserByLogin(ctx, input.Login)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	output := models.AuthLoginResponse{
		Token:    token,
		FullName: account.FullName,
	}

	ctx.JSON(http.StatusOK, output)
}

// @Summary	GetUser получение пользователя по идентификатору
// @VendorCode	GetUserById
// @Accept		json
// @Produce	json
// @Tags		Auth
// @Success	200
// @Param		id		path		integer	true	"account info"
// @Success	200		{integer}	models.User
// @Failure	403		{object}	map[string]interface{}
// @Failure	400,404	{object}	map[string]interface{}
// @Failure	500		{object}	map[string]interface{}
// @Failure	default	{object}	map[string]interface{}
// @Router		/auth/users/:id [get]
func (h *Handler) GetUserByID(ctx *gin.Context) {
	userIDParam := ctx.Param("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.Authorization.GetUserByID(ctx, models.User{ID: userID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

// @Summary	GetAllUser получение список пользователей
// @VendorCode	GetAllUser
// @Accept		json
// @Produce	json
// @Tags		Auth
// @Success	200		{array}		models.User
// @Failure	403		{object}	map[string]interface{}
// @Failure	400,404	{object}	map[string]interface{}
// @Failure	500		{object}	map[string]interface{}
// @Failure	default	{object}	map[string]interface{}
// @Router		/auth/users [get]
func (h *Handler) GetListUser(ctx *gin.Context) {
	users, err := h.services.Authorization.GetListUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}
