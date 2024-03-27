package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"rams/internal/repository/real_property"
	"rams/pkg/models"
)

type RealProperty interface {
	Create(ctx context.Context, realPro models.RealProperty) error
	GetList(ctx context.Context) ([]models.RealProperty, error)
	GetByID(ctx context.Context, realPro models.RealProperty) (models.RealProperty, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		RealProperty: real_property.NewRealPropertyRepository(db),
	}
}

type Repository struct {
	RealProperty
}
