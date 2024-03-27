package service

import (
	"context"
	"rams/internal/repository"
	"rams/internal/service/real_property"
	"rams/pkg/models"
)

type RealProperty interface {
	Create(ctx context.Context, realPro models.RealProperty) error
	GetList(ctx context.Context) ([]models.RealProperty, error)
	GetByID(ctx context.Context, realPro models.RealProperty) (models.RealProperty, error)
}

type Authorization interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUser(ctx context.Context, username, password string) (user models.User, err error)
	GetUserByID(ctx context.Context, user models.User) (models.User, error)
	GetListUser(ctx context.Context) (users []models.User, err error)
	GenerateToken(ctx context.Context, username, password string) (string, error)
	GetUserByLogin(ctx context.Context, login string) (user models.User, err error)
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		RealProperty: real_property.NewRealPropertyService(repos),
	}
}

type Service struct {
	RealProperty
	Authorization
}
