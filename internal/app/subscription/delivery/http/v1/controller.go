package v1

import (
	"github.com/leoscrowi/effective-mobile-test/internal/app/subscription"
	usecase2 "github.com/leoscrowi/effective-mobile-test/internal/app/subscription/usecase"
	"gorm.io/gorm"
	"net/http"
)

type SubscriptionController struct {
	usecase subscription.Usecase
}

func (s SubscriptionController) ReadSubscriptionsList(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewSubscriptionController(db *gorm.DB) *subscription.Controller {
	return &SubscriptionController{usecase: usecase2.NewSubscriptionUsecase(db)}
}
