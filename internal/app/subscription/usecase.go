package subscription

import (
	"context"
	_ "github.com/google/uuid"
	"github.com/leoscrowi/effective-mobile-test/internal/app/subscription/dto"
)

type Usecase interface {
	CreateSubscription(ctx context.Context, dto dto.CreateSubscriptionRequest) (dto.CreateSubscriptionResponse, error)
	ReadSubscription(ctx context.Context, dto dto.ReadSubscriptionRequest) (dto.ReadSubscriptionResponse, error)
	EditSubscription(ctx context.Context, dto dto.EditSubscriptionRequset) error
	DeleteSubscription(ctx context.Context, dto dto.DeleteSubscriptionRequest) error
	ReadSubscriptionsList(ctx context.Context) (dto.ReadSubscriptionsListResponse, error)
	GetSubscriptionsAmount(ctx context.Context, request dto.GetSubscriptionsAmountRequest) (dto.GetSubscriptionsAmountResponse, error)
}
