package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rams/pkg/models"
	"strconv"
)

// @Summary		Create real property
// @Tags			real_property(недвижимость)
// @Description	Создание недвижимости.
// @ID				create-real-property
// @Accept			json
// @Produce		json
// @Param			input	body		models.RealProperty		true	"информация о недвижимости"
// @Success		200		{object}	map[string]interface{}	"При успешном выполнении возврашает created!"
// @Failure		403		{object}	map[string]interface{}
// @Failure		400,404	{object}	map[string]interface{}
// @Failure		500		{object}	map[string]interface{}
// @Failure		default	{object}	map[string]interface{}
// @Router			/real_property [post]
func (h *Handler) Create(ctx *gin.Context) {
	var input models.RealProperty
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := h.validate.Struct(input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.RealProperty.Create(ctx, input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"reason": "created!"})
}

// @Summary		GetList real property
// @Tags			real_property(недвижимость)
// @Description	Получение список недвижимости.
// @ID				get-list-real-property
// @Accept			json
// @Produce		json
// @Success		200		{array}		models.RealProperty
// @Failure		403		{object}	map[string]interface{}
// @Failure		400,404	{object}	map[string]interface{}
// @Failure		500		{object}	map[string]interface{}
// @Failure		default	{object}	map[string]interface{}
// @Router			/real_property [get]
func (h *Handler) GetList(ctx *gin.Context) {
	contracts, err := h.services.RealProperty.GetList(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, contracts)
}

// @Summary		GetByID real property
// @Tags			real_property(недвижимость)
// @Description	Получение недвижимость по идентификатору.
// @ID				get-by-id-real-property
// @Accept			json
// @Produce		json
// @Param			id		path		integer	true	"идентификатор недвижимости"
// @Success		200		{object}	models.RealProperty
// @Failure		403		{object}	map[string]interface{}
// @Failure		400,404	{object}	map[string]interface{}
// @Failure		500		{object}	map[string]interface{}
// @Failure		default	{object}	map[string]interface{}
// @Router			/real_property/:id [get]
func (h *Handler) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	contract, err := h.services.RealProperty.GetByID(ctx, models.RealProperty{ID: id})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, contract)
}
