package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Feature struct {
	gorm.Model

	DisplayName   string    `json:"display_name"`
	TechnicalName string    `json:"technical_name"`
	ExpiresOn     time.Time `json:"expires_on"`
	Description   string    `json:"description"`
	Inverted      bool      `json:"inverted"`
}

func CreateFeature(data map[string]interface{}) error {
	feature := Feature{
		DisplayName:   data["display_name"].(string),
		TechnicalName: data["technical_name"].(string),
		ExpiresOn:     data["expires_on"].(time.Time),
		Description:   data["description"].(string),
		Inverted:      data["inverted"].(bool),
	}

	if err := db.Create(&feature).Error; err != nil {
		return err
	}

	return nil
}

func GetFeature(id int) (*Feature, error) {
	var feature Feature
	err := db.Where("id = ?", id).First(&feature).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &feature, nil
}

func EditFeature(id int, data map[string]interface{}) error {
	if err := db.Model(&Feature{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func ExistsFeatureByID(id int) (bool, error) {
	var article Feature
	err := db.Select("id").Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}
