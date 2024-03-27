package repository

import (
	"embed"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var MigrationsFS embed.FS

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) *sqlx.DB {
	//migrationsDir := "migrations"

	//migrator := migrator2.MustGetNewMigrator(MigrationsFS, migrationsDir)

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)

	// Подключаемся к базе данных
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	//defer db.Close()

	// Проверяем, что подключение успешно установлено
	err = db.Ping()
	if err != nil {
		log.Fatalf("Ошибка проверки подключения: %v", err)
	}

	fmt.Println("Успешное подключение к базе данных!")

	//if err = migrator.ApplyMigrations(db); err != nil {
	//	log.Fatalf("Ошибка миграции: %v", err)
	//}
	//
	//fmt.Println("Применены миграции!!!!")

	return db
}
