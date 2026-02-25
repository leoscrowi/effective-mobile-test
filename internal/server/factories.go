package server

import (
	v1 "github.com/leoscrowi/effective-mobile-test/internal/app/subscription/delivery/http/v1"
	"gorm.io/gorm"
)

func GetControllers(db *gorm.DB) []RouteSetup {
	return []RouteSetup{
		v1.NewSubscriptionController(db),
	}
}
