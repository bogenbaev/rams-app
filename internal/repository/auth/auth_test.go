package auth

import (
	"context"
	"rams/pkg/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestCreateUser(t *testing.T) {
	// Создаем макет базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Создаем экземпляр sqlx.DB с макетом базы данных
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Создаем экземпляр вашего репозитория, использующего sqlx.DB
	userRepository := NewAuthRepository(sqlxDB)

	// Ожидаемый запрос вставки записи в базу данных
	mock.ExpectExec("INSERT INTO users").WithArgs("John Doe", "john@example.com", "john", "password", "2024-03-27").WillReturnResult(sqlmock.NewResult(1, 1))

	// Вызываем функцию, которая записывает данные в базу данных
	err = userRepository.CreateUser(context.Background(), models.User{
		FullName:  "John Doe",
		Email:     "john@example.com",
		Login:     "john",
		Password:  "password",
		CreatedAt: time.Date(2024, time.March, 27, 0, 0, 0, 0, time.UTC)})

	// Проверяем, что нет ошибок
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	// Проверяем, что все ожидаемые запросы были выполнены
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserByID(t *testing.T) {
	// Создаем макет базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Создаем экземпляр sqlx.DB с макетом базы данных
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Создаем экземпляр вашего репозитория, использующего sqlx.DB
	userRepository := NewAuthRepository(sqlxDB)

	// Ожидаемый запрос на выборку данных из базы данных
	rows := sqlmock.NewRows([]string{"id", "full_name", "email", "login", "password", "created_at"}).
		AddRow(1, "John Doe", "john@example.com", "john", "password", "2024-03-27")
	mock.ExpectQuery("SELECT (.+) FROM users WHERE id=?").WithArgs(1).WillReturnRows(rows)

	// Вызываем функцию, которая получает данные из базы данных
	user, err := userRepository.GetUserByID(context.Background(), 1)

	// Проверяем, что нет ошибок
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	// Проверяем, что все ожидаемые запросы были выполнены
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	_ = user
}

func TestGetAllUsers(t *testing.T) {
	// Создаем макет базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Создаем экземпляр sqlx.DB с макетом базы данных
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Создаем экземпляр вашего репозитория, использующего sqlx.DB
	userRepository := NewAuthRepository(sqlxDB)

	// Ожидаемый запрос на выборку всех пользователей из базы данных
	rows := sqlmock.NewRows([]string{"id", "full_name", "email", "login", "password", "created_at"}).
		AddRow(1, "John Doe", "john@example.com", "john", "password", "2024-03-27").
		AddRow(2, "Jane Smith", "jane@example.com", "jane", "password123", "2024-03-28")
	mock.ExpectQuery("SELECT (.+) FROM users").WillReturnRows(rows)

	// Вызываем функцию, которая получает всех пользователей из базы данных
	users, err := userRepository.GetListUser(context.Background())

	// Проверяем, что нет ошибок
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	// Проверяем, что все ожидаемые запросы были выполнены
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	_ = users
}
