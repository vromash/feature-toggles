package feature_service

import (
	"time"

	"github.com/vromash/toggle-feature-switcher/internal/models"
)

type Feature struct {
	ID            int
	DisplayName   string
	TechnicalName string
	ExpiresOn     time.Time
	Description   string
	Inverted      bool
	CustomerIds   []string
	Active        bool
}

func (f *Feature) Add() error {
	feature := map[string]interface{}{
		"display_name":   f.DisplayName,
		"technical_name": f.TechnicalName,
		"expires_on":     f.ExpiresOn,
		"description":    f.Description,
		"inverted":       f.Inverted,
	}

	if err := models.CreateFeature(feature); err != nil {
		return err
	}

	return nil
}

func (f *Feature) Update() error {
	feature := map[string]interface{}{
		"display_name":   f.DisplayName,
		"technical_name": f.TechnicalName,
		"expires_on":     f.ExpiresOn,
		"description":    f.Description,
		"inverted":       f.Inverted,
	}

	if err := models.EditFeature(f.ID, feature); err != nil {
		return err
	}

	return nil
}

func (f *Feature) ExistByID() (bool, error) {
	return models.ExistsFeatureByID(f.ID)
}

func ExistByID(id int) (bool, error) {
	return models.ExistsFeatureByID(id)
}

func GetFeatureByID(id int) (*models.Feature, error) {
	feature, err := models.GetFeature(id)
	if err != nil {
		return nil, err
	}

	return feature, nil
}
