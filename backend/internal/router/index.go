package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vromash/toggle-feature-switcher/internal/models"
	feature_service "github.com/vromash/toggle-feature-switcher/internal/services/feature"
	feature_customer_service "github.com/vromash/toggle-feature-switcher/internal/services/feature_customer"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong1",
			})
		})

		apiv1.POST("/features", getFeaturesByCustomerId)

		apiv1.POST("/feature", addFeature)
		apiv1.PUT("/feature", updateFeature)
		apiv1.DELETE("/feature", archiveFeature)

		apiv1.POST("/feature/customer", addCustomerToFeature)
	}

	return router
}

type GetFeaturesForm struct {
	CustomerID int `json:"customerID" binding:"required"`
}

type ReturnFeature struct {
	ID            int
	Active        bool
	DisplayName   string
	TechnicalName string
	ExpiresOn     time.Time
	Description   string
	Inverted      bool
}

func getFeaturesByCustomerId(c *gin.Context) {
	var form GetFeaturesForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	featureCustomers, err := feature_customer_service.GetByCustomerId(form.CustomerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var features []ReturnFeature
	for _, featureCustomer := range featureCustomers {
		feature, errDb := models.GetFeature((featureCustomer.FeatureId))
		if errDb != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errDb.Error()})
			return
		}

		features = append(
			features,
			ReturnFeature{
				ID:            int(feature.ID),
				Active:        featureCustomer.Active,
				DisplayName:   feature.DisplayName,
				TechnicalName: feature.TechnicalName,
				ExpiresOn:     feature.ExpiresOn,
				Description:   feature.Description,
				Inverted:      feature.Inverted,
			})
	}

	c.JSON(http.StatusOK, features)
}

type NewFeatureForm struct {
	DisplayName   string    `json:"displayName" binding:"required"`
	TechnicalName string    `json:"technicalName" binding:"required"`
	ExpiresOn     time.Time `json:"expiresOn" binding:"required"`
	Description   string    `json:"description" binding:"required"`
	Inverted      *bool     `json:"inverted" binding:"required"`
}

func addFeature(c *gin.Context) {
	var form NewFeatureForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	featureService := feature_service.Feature{
		DisplayName:   form.DisplayName,
		TechnicalName: form.TechnicalName,
		ExpiresOn:     form.ExpiresOn,
		Description:   form.Description,
		Inverted:      *form.Inverted,
	}

	if err := featureService.Add(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, "Feature added: %s", form.DisplayName)
}

type AddCustomerFeatureForm struct {
	FeatureID  int   `json:"featureId" binding:"required"`
	CustomerID int   `json:"customerId" binding:"required"`
	Active     *bool `json:"active" binding:"required"`
}

func addCustomerToFeature(c *gin.Context) {
	var form AddCustomerFeatureForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exists, errCheck := feature_service.ExistByID(form.FeatureID)
	if errCheck != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errCheck.Error()})
		return
	}

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Feature Id doesn't exist"})
		return
	}

	featureCustomer := feature_customer_service.FeatureCustomer{
		FeatureID:  form.FeatureID,
		CustomerID: form.CustomerID,
		Active:     *form.Active,
	}

	if err := featureCustomer.Add(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, "Customer added to feature")
}

type UpdateFeatureForm struct {
	ID            int       `json:"id" binding:"required"`
	DisplayName   string    `json:"displayName" binding:"required"`
	TechnicalName string    `json:"technicalName" binding:"required"`
	ExpiresOn     time.Time `json:"expiresOn" binding:"required"`
	Description   string    `json:"description" binding:"required"`
	Inverted      *bool     `json:"inverted" binding:"required"`
}

func updateFeature(c *gin.Context) {
	var form UpdateFeatureForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	featureService := feature_service.Feature{
		ID:            form.ID,
		DisplayName:   form.DisplayName,
		TechnicalName: form.TechnicalName,
		ExpiresOn:     form.ExpiresOn,
		Description:   form.Description,
		Inverted:      *form.Inverted,
	}

	exists, errCheck := featureService.ExistByID()
	if errCheck != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errCheck.Error()})
		return
	}

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id doesn't exist"})
		return
	}

	if err := featureService.Update(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, "Feature edited: %s", form.DisplayName)
}

func archiveFeature(c *gin.Context) {

}
