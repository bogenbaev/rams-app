package models

type AppConfigs struct {
	Port     string
	Database *Postgres
}

type Postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}
