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

func NewService(repos *repository.Repository) *Service {
	return &Service{
		RealProperty: real_property.NewRealPropertyService(repos),
	}
}

type Service struct {
	RealProperty
}
