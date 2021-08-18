package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vromash/toggle-feature-switcher/internal/services"
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

		apiv1.POST("/features", getFeatureToggles)

		apiv1.POST("/feature", addFeatureToggle)
		apiv1.PUT("/feature", updateFeatureToggle)
		apiv1.DELETE("/feature", archiveFeatureToggle)
	}

	return router
}

func getFeatureToggles(c *gin.Context) {

}

func addFeatureToggle(c *gin.Context) {
	featureToggle := services.FeatureToggle{
		DisplayName:   "test1_dn",
		TechnicalName: "test1_tn",
		ExpiresOn:     time.Now(),
		Description:   "test1_des",
		Inverted:      false,
		CustomerIds:   []string{""},
	}

	if err := services.AddFeatureToggle(&featureToggle); err != nil {
		return
	}
}

func updateFeatureToggle(c *gin.Context) {

}

func archiveFeatureToggle(c *gin.Context) {

}
