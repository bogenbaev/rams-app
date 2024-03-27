package real_property

import (
	"context"
	"rams/internal/repository"
	"rams/pkg/models"
)

type RealPropertyService struct {
	repo *repository.Repository
}

func NewRealPropertyService(repo *repository.Repository) *RealPropertyService {
	return &RealPropertyService{
		repo: repo,
	}
}

func (s *RealPropertyService) Create(ctx context.Context, realPro models.RealProperty) error {
	return s.repo.Create(ctx, realPro)
}

func (s *RealPropertyService) GetList(ctx context.Context) ([]models.RealProperty, error) {
	return s.repo.GetList(ctx)
}

func (s *RealPropertyService) GetByID(ctx context.Context, realPro models.RealProperty) (models.RealProperty, error) {
	return s.repo.GetByID(ctx, realPro)
}
