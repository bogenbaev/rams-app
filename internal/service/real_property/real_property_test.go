package real_property

//import (
//	"context"
//	"rams/pkg/models"
//	"testing"
//	"time"
//
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//)
//
//// Mock репозитория
//type MockRepo struct {
//	mock.Mock
//}
//
//func (m *MockRepo) Create(ctx context.Context, realPro models.RealProperty) error {
//	args := m.Called(ctx, realPro)
//	return args.Error(0)
//}
//
//func (m *MockRepo) GetList(ctx context.Context) ([]models.RealProperty, error) {
//	args := m.Called(ctx)
//	return args.Get(0).([]models.RealProperty), args.Error(1)
//}
//
//func (m *MockRepo) GetByID(ctx context.Context, realPro models.RealProperty) (models.RealProperty, error) {
//	args := m.Called(ctx, realPro)
//	return args.Get(0).(models.RealProperty), args.Error(1)
//}
//
//func TestRealPropertyService_Create(t *testing.T) {
//	mockRepo := new(MockRepo)
//	realProService := NewRealPropertyService(mockRepo)
//
//	realPro := models.RealProperty{
//		PropertyTypeID: 1,
//		PropertyType:   "House",
//		Address:        "123 Main St",
//		Price:          100000,
//		Rooms:          3,
//		Area:           1500,
//		Description:    "Spacious house with garden",
//		CreatedAt:      time.Now(),
//	}
//
//	// Указываем ожидаемый вызов метода Create у репозитория
//	mockRepo.On("Create", mock.Anything, realPro).Return(nil)
//
//	// Вызываем метод Create у сервиса
//	err := realProService.Create(context.Background(), realPro)
//
//	// Проверяем, что метод Create у репозитория был вызван с правильными аргументами
//	mockRepo.AssertCalled(t, "Create", mock.Anything, realPro)
//	// Проверяем, что нет ошибок
//	assert.NoError(t, err)
//}
//
//func TestRealPropertyService_GetList(t *testing.T) {
//	mockRepo := new(MockRepo)
//	realProService := NewRealPropertyService(mockRepo)
//
//	realPros := []models.RealProperty{
//		{
//			ID:             1,
//			PropertyTypeID: 1,
//			PropertyType:   "House",
//			Address:        "123 Main St",
//			Price:          100000,
//			Rooms:          3,
//			Area:           1500,
//			Description:    "Spacious house with garden",
//			CreatedAt:      time.Now(),
//		},
//		// Другие объекты недвижимости...
//	}
//
//	// Указываем ожидаемый вызов метода GetList у репозитория
//	mockRepo.On("GetList", mock.Anything).Return(realPros, nil)
//
//	// Вызываем метод GetList у сервиса
//	result, err := realProService.GetList(context.Background())
//
//	// Проверяем, что метод GetList у репозитория был вызван
//	mockRepo.AssertCalled(t, "GetList", mock.Anything)
//	// Проверяем, что результат соответствует ожиданиям
//	assert.NoError(t, err)
//	assert.Equal(t, realPros, result)
//}
