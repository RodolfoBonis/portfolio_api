package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
	"portfolio_api/src/entities"
)

var Db *gorm.DB
var err error

func buildConnectorConfig() *entities.ConnectorConfig {
	_ = godotenv.Load(".env")
	connectorConfig := entities.ConnectorConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_SECRET"),
		DBName:   os.Getenv("DB_SCHEMA"),
	}
	return &connectorConfig
}

func connectorURL(connectorConfig *entities.ConnectorConfig) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		connectorConfig.Host,
		connectorConfig.Port,
		connectorConfig.User,
		connectorConfig.DBName,
		connectorConfig.Password,
	)
}

func OpenConnection() {
	Db, err = gorm.Open("postgres", connectorURL(buildConnectorConfig()))

	if err != nil {
		fmt.Println(err)
	}

	Db.LogMode(true)

}
