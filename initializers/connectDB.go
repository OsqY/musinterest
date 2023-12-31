package initializers

import (
	"fmt"	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
    	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", config.DBHost, config.DBUsername, config.DBUserPassword, config.DBName ,config.DBPort)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} 

	fmt.Println("Successfully connected to database!")	
}
