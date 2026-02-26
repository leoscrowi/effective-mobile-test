package v1

import (
	"github.com/leoscrowi/effective-mobile-test/internal/app/subscription"
	usecase2 "github.com/leoscrowi/effective-mobile-test/internal/app/subscription/usecase"
	"gorm.io/gorm"
)

type SubscriptionController struct {
	usecase subscription.Usecase
}

func NewSubscriptionController(db *gorm.DB) *SubscriptionController {
	return &SubscriptionController{usecase: usecase2.NewSubscriptionUsecase(db)}
}
