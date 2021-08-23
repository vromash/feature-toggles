package models

import (
	"github.com/jinzhu/gorm"
)

type FeatureCustomer struct {
	gorm.Model

	FeatureId  int  `json:"feature_id"`
	CustomerId int  `json:"customer_id"`
	Active     bool `json:"active"`
}

func AddCustomerToFeature(data map[string]interface{}) error {
	featureCustomer := FeatureCustomer{
		FeatureId:  data["feature_id"].(int),
		CustomerId: data["customer_id"].(int),
		Active:     data["active"].(bool),
	}

	if err := db.Create(&featureCustomer).Error; err != nil {
		return err
	}

	return nil
}

func GetCustomersFeatures(customerId int) ([]*FeatureCustomer, error) {
	var featureCustomers []*FeatureCustomer
	err := db.Where("customer_id = ?", customerId).Find(&featureCustomers).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return featureCustomers, nil
}

func RemoveCustomerFromFeature(customerId int, featureId int) error {
	if err := db.Where("customer_id = ? AND feature_id = ?", customerId, featureId).Delete(FeatureCustomer{}).Error; err != nil {
		return err
	}

	return nil
}
