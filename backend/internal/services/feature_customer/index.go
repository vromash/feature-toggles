package feature_customer_service

import (
	"github.com/vromash/toggle-feature-switcher/internal/models"
)

type FeatureCustomer struct {
	ID         int
	CustomerID int
	FeatureID  int
	Active     bool
}

func (f *FeatureCustomer) Add() error {
	featureCustomer := map[string]interface{}{
		"customer_id": f.CustomerID,
		"feature_id":  f.FeatureID,
		"active":      f.Active,
	}

	if err := models.AddCustomerToFeature(featureCustomer); err != nil {
		return err
	}

	return nil
}

func (f *FeatureCustomer) ExistByID() (bool, error) {
	return models.ExistsFeatureByID(f.ID)
}

func GetByCustomerId(id int) ([]*models.FeatureCustomer, error) {
	featureCustomers, err := models.GetCustomersFeatures(id)
	if err != nil {
		return nil, err
	}

	return featureCustomers, nil
}
