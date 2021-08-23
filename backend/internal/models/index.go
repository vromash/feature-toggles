package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/vromash/toggle-feature-switcher/pkg/utils"
)

var db *gorm.DB

func CreateConnection() {
	connectionUrl := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.ReadEnvVariable("DB_USER"),
		utils.ReadEnvVariable("DB_PASSWORD"),
		utils.ReadEnvVariable("DB_HOST"),
		utils.ReadEnvVariable("DB_DATABASE"))

	var err error
	db, err = gorm.Open("mysql", connectionUrl)
	if err != nil {
		log.Fatalf("models.createConnection err: %v", err)
	}

	// db.AutoMigrate(&Customer{})
	db.AutoMigrate(&Feature{})
	db.AutoMigrate(&FeatureCustomer{})
}
