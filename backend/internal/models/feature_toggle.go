package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type FeatureToggle struct {
	gorm.Model

	DisplayName   string    `json:"display_name"`
	TechnicalName string    `json:"technical_name"`
	ExpiresOn     time.Time `json:"expires_on"`
	Description   string    `json:"description"`
	Inverted      bool      `json:"inverted"`
}

func CreateFeatureToggle(data map[string]interface{}) error {
	featureToggle := FeatureToggle{
		DisplayName:   data["display_name"].(string),
		TechnicalName: data["technical_name"].(string),
		ExpiresOn:     data["expires_on"].(time.Time),
		Description:   data["description"].(string),
		Inverted:      data["inverted"].(bool),
	}

	if err := db.Create(&featureToggle).Error; err != nil {
		return err
	}

	return nil
}
