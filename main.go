package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	_ "rams/docs"
	"rams/internal/app"
	"rams/internal/handler"
	"rams/internal/repository"
	"rams/internal/server"
	"rams/internal/service"
	"syscall"
)

// @title						RAMS-app
// @version					1.0
// @description				API Server for TodoList Application
// @host						localhost:8080
// @BasePath					/
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	cfg := app.InitConfigs()

	db := repository.NewPostgresDB(repository.Config{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		Username: cfg.Database.Username,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.DBName,
		SSLMode:  cfg.Database.SSLMode,
	})

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	newHandler := handler.NewHandler(services)

	srv := new(server.Server)
	go func() {
		if err := srv.Run("8080", newHandler.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("RAMSApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("RAMSApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
