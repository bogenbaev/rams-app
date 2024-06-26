package repository

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"log"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) *sqlx.DB {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)

	//Подключаемся к базе данных
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	//defer db.Close()

	// Проверяем, что подключение успешно установлено
	if err = db.Ping(); err != nil {
		log.Fatalf("Ошибка проверки подключения: %v", err)
	}

	fmt.Println("Успешное подключение к базе данных!")

	// Читаем миграции из указанной директории и подключаемся к локальной базе данных PostgreSQL.
	m, err := migrate.New("file:///home/nugman/rams-app/internal/repository/migrations", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Применяем миграции
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	fmt.Println("Миграция успешно применена!")

	return db
}
