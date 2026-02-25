package subscription

import (
	"context"
	"github.com/google/uuid"
	subscriptiondomain "github.com/leoscrowi/effective-mobile-test/domain/subscription"
)

type Repository interface {
	CreateSubscription(ctx context.Context, Subscription subscriptiondomain.Subscription) (uuid.UUID, error)
	ReadSubscription(ctx context.Context, uuid uuid.UUID) (subscriptiondomain.Subscription, error)
	EditSubscription(ctx context.Context, Subscription subscriptiondomain.Subscription) error
	DeleteSubscription(ctx context.Context, uuid uuid.UUID) error
	ReadSubscriptionsList(ctx context.Context) ([]subscriptiondomain.Subscription, error)
}
