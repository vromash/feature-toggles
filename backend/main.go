package main

import (
	"github.com/vromash/toggle-feature-switcher/internal/models"
	"github.com/vromash/toggle-feature-switcher/internal/router"
)

func init() {
	models.CreateConnection()
}

func main() {
	router := router.CreateRouter()
	router.Run()
}
