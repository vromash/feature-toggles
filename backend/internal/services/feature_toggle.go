package services

import (
	"time"

	"github.com/vromash/toggle-feature-switcher/internal/models"
)

type FeatureToggle struct {
	DisplayName   string
	TechnicalName string
	ExpiresOn     time.Time
	Description   string
	Inverted      bool
	CustomerIds   []string
}

func AddFeatureToggle(f *FeatureToggle) error {
	featureToggle := map[string]interface{}{
		"display_name":   f.DisplayName,
		"technical_name": f.TechnicalName,
		"expires_on":     f.ExpiresOn,
		"description":    f.Description,
		"inverted":       f.Inverted,
	}

	if err := models.CreateFeatureToggle(featureToggle); err != nil {
		return err
	}

	return nil
}
