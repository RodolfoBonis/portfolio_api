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

func OpenConnection() error {
	Db, err = gorm.Open("postgres", connectorURL(buildConnectorConfig()))

	if err != nil {
		return err
	}

	Db.LogMode(true)

	return nil
}

func RunMigration() {
	Db.AutoMigrate(&entities.User{}, &entities.Post{}, &entities.Comment{}, &entities.SocialMedia{}, &entities.Tag{}, &entities.TagsPosts{})

	Db.Model(&entities.Post{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	Db.Model(&entities.SocialMedia{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	Db.Model(&entities.Comment{}).AddForeignKey("post_id", "posts(id)", "RESTRICT", "RESTRICT")
	Db.Model(&entities.TagsPosts{}).AddForeignKey("post_id", "posts(id)", "RESTRICT", "RESTRICT")
	Db.Model(&entities.TagsPosts{}).AddForeignKey("tag_id", "tags(id)", "RESTRICT", "RESTRICT")
}
