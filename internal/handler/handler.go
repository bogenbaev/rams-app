package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"rams/internal/service"
)

type Handler struct {
	services *service.Service
	validate *validator.Validate
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
		validate: validator.New(),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	realProperty := r.Group("/real_property")
	{
		realProperty.POST("", h.Create)
		realProperty.GET("", h.GetList)
		realProperty.GET("/:id", h.GetByID)
	}

	auth := r.Group("/auth")
	{
		auth.POST("/login", h.SignIn)
		auth.POST("/sign_up", h.SignUp)
		auth.GET("/users", h.GetListUser)
		auth.GET("users/:id", h.GetUserByID)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
