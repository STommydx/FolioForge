package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string `env:"HOST, required"`
	Port     string `env:"PORT, default=5432"`
	Username string `env:"USER, required"`
	Password string `env:"PASSWORD, required"`
	Database string `env:"DB, default=folioforge"`
}

func NewPostgresDB(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Username,
		config.Password,
		config.Database,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil

}
