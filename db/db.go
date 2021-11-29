package db

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/0xPiyush/my-parking/db/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var DB *gorm.DB

func Connect() {
	conn, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the Database")
		panic(err)
	}

	DB = conn

	conn.AutoMigrate(&models.User{}, &models.Location{}, &models.Booking{})

	applyConfig(conn)
}

func applyConfig(db *gorm.DB) {
	config, err := loadConfig("./dbconfig.json")
	if err != nil {
		fmt.Println("Failed to load config file")
		panic(err)
	}

	for _, location := range config.Locations {
		db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&location)
	}
}

type Config struct {
	Locations []models.Location `json:"locations"`
}

func loadConfig(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return config, err
	}
	return config, nil
}
