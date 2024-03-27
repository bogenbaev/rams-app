package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"rams/internal/repository/auth"
	"rams/internal/repository/real_property"
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
	GetUserByLogin(ctx context.Context, login string) (user models.User, err error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		RealProperty:  real_property.NewRealPropertyRepository(db),
		Authorization: auth.NewAuthRepository(db),
	}
}

type Repository struct {
	RealProperty
	Authorization
}
