package real_property

import (
	"context"
	"github.com/jmoiron/sqlx"
	"rams/pkg/models"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRealPropertyRepository_Create(t *testing.T) {
	// Создаем макет базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Создаем экземпляр вашего репозитория
	repo := NewRealPropertyRepository(sqlx.NewDb(db, "sqlmock"))

	// Ожидаемый объект недвижимости для создания
	realPro := models.RealProperty{
		PropertyTypeID: 1,
		PropertyType:   "Apartment",
		Address:        "123 Main St",
		Price:          100000.00,
		Rooms:          3,
		Area:           100.00,
		Description:    "Spacious apartment",
		CreatedAt:      time.Now(),
	}

	// Ожидаемый SQL запрос
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Вызываем функцию создания объекта недвижимости в репозитории
	err = repo.Create(context.Background(), realPro)

	// Проверяем, что нет ошибок
	assert.NoError(t, err)

	// Проверяем, что все ожидаемые запросы были выполнены
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRealPropertyRepository_GetList(t *testing.T) {
	// Создаем макет базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Создаем экземпляр вашего репозитория
	repo := NewRealPropertyRepository(sqlx.NewDb(db, "sqlmock"))

	// Ожидаемые данные
	expectedRealPros := []models.RealProperty{
		{
			ID:             1,
			PropertyTypeID: 1,
			PropertyType:   "Apartment",
			Address:        "123 Main St",
			Price:          100000.00,
			Rooms:          3,
			Area:           100.00,
			Description:    "Spacious apartment",
			CreatedAt:      time.Now(),
		},
		// Добавьте другие объекты недвижимости при необходимости
	}

	// Ожидаемый SQL запрос
	rows := sqlmock.NewRows([]string{"id", "property_type_id", "property_type", "address", "price", "rooms", "area", "description", "created_at"})
	for _, rp := range expectedRealPros {
		rows.AddRow(rp.ID, rp.PropertyTypeID, rp.PropertyType, rp.Address, rp.Price, rp.Rooms, rp.Area, rp.Description, rp.CreatedAt)
	}
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	mock.ExpectCommit()

	// Вызываем функцию получения списка объектов недвижимости в репозитории
	realPros, err := repo.GetList(context.Background())

	// Проверяем, что нет ошибок
	assert.NoError(t, err)

	// Проверяем, что возвращенные данные соответствуют ожидаемым
	assert.Equal(t, expectedRealPros, realPros)

	// Проверяем, что все ожидаемые запросы были выполнены
	assert.NoError(t, mock.ExpectationsWereMet())
}
