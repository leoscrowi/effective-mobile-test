package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	subscriptiondomain "github.com/leoscrowi/effective-mobile-test/domain/subscription"
	"gorm.io/gorm"
)

type SubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (s *SubscriptionRepository) CreateSubscription(ctx context.Context, subscription subscriptiondomain.Subscription) (uuid.UUID, error) {
	op := `SubscriptionRepository.CreateSubscription`
	result := s.db.WithContext(ctx).Create(&subscription)

	if result.Error != nil {
		return uuid.Nil, fmt.Errorf("%s: %v", op, result.Error)
	}

	return subscription.ID, nil
}

func (s *SubscriptionRepository) ReadSubscription(ctx context.Context, uuid uuid.UUID) (subscriptiondomain.Subscription, error) {
	op := `SubscriptionRepository.ReadSubscription`
	var sub subscriptiondomain.Subscription

	err := s.db.WithContext(ctx).Where("id = ?", uuid).First(&sub).Error
	if err != nil {
		return subscriptiondomain.Subscription{}, fmt.Errorf("%s: %v", op, err)
	}

	return sub, nil
}

func (s *SubscriptionRepository) EditSubscription(ctx context.Context, Subscription subscriptiondomain.Subscription) error {
	op := `SubscriptionRepository.EditSubscription`
	err := s.db.WithContext(ctx).Model(&subscriptiondomain.Subscription{}).Where("id = ?", Subscription.ID).Updates(Subscription).Error
	if err != nil {
		return fmt.Errorf("%s: %v", op, err)
	}

	return nil
}

func (s *SubscriptionRepository) DeleteSubscription(ctx context.Context, uuid uuid.UUID) error {
	op := `SubscriptionRepository.DeleteSubscription`
	err := s.db.WithContext(ctx).Where("id = ?", uuid).Delete(&subscriptiondomain.Subscription{}).Error
	if err != nil {
		return fmt.Errorf("%s: %v", op, err)
	}

	return nil
}

func (s *SubscriptionRepository) ReadSubscriptionsList(ctx context.Context) ([]subscriptiondomain.Subscription, error) {
	op := `SubscriptionRepository.ReadSubscriptionsList`
	var subs []subscriptiondomain.Subscription
	err := s.db.WithContext(ctx).Find(&subs).Error
	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	return subs, nil
}
