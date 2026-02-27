package main

import (
	"fmt"
	"github.com/leoscrowi/effective-mobile-test/internal/config"
	"github.com/leoscrowi/effective-mobile-test/internal/migrations"
	"github.com/leoscrowi/effective-mobile-test/internal/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"time"
)

// @title Subscription Management API
// @version 1.0
// @description API для управления подписками пользователей
// @host localhost:8080
// @basePath /
// @schemes http https
func main() {
	cfg := config.MustLoad()

	db, err := initDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}

	err = migrations.Migrate(db)
	if err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}

	s := server.NewServer(db)
	s.SetupRoutes()

	port := fmt.Sprintf(":%s", cfg.AppConfig.Port)
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(port, s.Router); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}

func initDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := getDsn(cfg)

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		PrepareStmt: true,
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("db connection established successfully")
	return db, nil
}

func getDsn(cfg *config.Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DatabaseConfig.Host,
		cfg.DatabaseConfig.User,
		cfg.DatabaseConfig.Password,
		cfg.DatabaseConfig.Name,
		cfg.DatabaseConfig.Port,
		cfg.DatabaseConfig.SslMode)
}
