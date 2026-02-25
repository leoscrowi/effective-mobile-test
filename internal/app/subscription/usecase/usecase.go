package usecase

import (
	"context"
	"fmt"
	subscriptiondomain "github.com/leoscrowi/effective-mobile-test/domain/subscription"
	"github.com/leoscrowi/effective-mobile-test/internal/app/subscription"
	"github.com/leoscrowi/effective-mobile-test/internal/app/subscription/dto"
	repository "github.com/leoscrowi/effective-mobile-test/internal/app/subscription/repository/postgresql"
	"gorm.io/gorm"
)

type SubscriptionUsecase struct {
	repository subscription.Repository
}

func NewSubscriptionUsecase(db *gorm.DB) *SubscriptionUsecase {
	return &SubscriptionUsecase{repository: repository.NewSubscriptionRepository(db)}
}

func (s *SubscriptionUsecase) CreateSubscription(ctx context.Context, d dto.CreateSubscriptionRequest) (dto.CreateSubscriptionResponse, error) {
	op := `SubscriptionUsecase.CreateSubscription`
	id, err := s.repository.CreateSubscription(ctx, subscriptiondomain.Subscription{
		ServiceName: d.ServiceName,
		Price:       d.Price,
		UserID:      d.UserID,
		StartDate:   d.StartDate,
		EndDate:     d.EndDate,
	})

	if err != nil {
		return dto.CreateSubscriptionResponse{ID: id}, fmt.Errorf("%s: %v;", op, err)
	}

	return dto.CreateSubscriptionResponse{ID: id}, nil
}

func (s *SubscriptionUsecase) ReadSubscription(ctx context.Context, d dto.ReadSubscriptionRequest) (dto.ReadSubscriptionResponse, error) {
	op := `SubscriptionUsecase.ReadSubscription`
	sub, err := s.repository.ReadSubscription(ctx, d.ID)
	if err != nil {
		return dto.ReadSubscriptionResponse{}, fmt.Errorf("%s: %v;", op, err)
	}

	return dto.ReadSubscriptionResponse{
		ID:          sub.ID,
		ServiceName: sub.ServiceName,
		Price:       sub.Price,
		UserID:      sub.UserID,
		StartDate:   sub.StartDate,
		EndDate:     sub.EndDate,
	}, nil
}

func (s *SubscriptionUsecase) EditSubscription(ctx context.Context, d dto.EditSubscriptionRequset) error {
	op := `SubscriptionUsecase.EditSubscription`
	err := s.repository.EditSubscription(ctx, subscriptiondomain.Subscription{
		ID:          d.ID,
		ServiceName: d.ServiceName,
		Price:       d.Price,
		UserID:      d.UserID,
		StartDate:   d.StartDate,
		EndDate:     d.EndDate,
	})

	if err != nil {
		return fmt.Errorf("%s: %v;", op, err)
	}

	return nil
}

func (s *SubscriptionUsecase) DeleteSubscription(ctx context.Context, dto dto.DeleteSubscriptionRequest) error {
	op := `SubscriptionUsecase.DeleteSubscription`
	err := s.repository.DeleteSubscription(ctx, dto.ID)
	if err != nil {
		return fmt.Errorf("%s: %v;", op, err)
	}

	return nil
}

func (s *SubscriptionUsecase) ReadSubscriptionsList(ctx context.Context) (dto.ReadSubscriptionsListResponse, error) {
	op := `SubscriptionUsecase.ReadSubscriptionsList`
	subs, err := s.repository.ReadSubscriptionsList(ctx)
	if err != nil {
		return dto.ReadSubscriptionsListResponse{}, fmt.Errorf("%s: %v;", op, err)
	}

	var subsResponse []dto.ReadSubscriptionResponse
	for _, sub := range subs {
		subsResponse = append(subsResponse, dto.ReadSubscriptionResponse{
			ID:          sub.ID,
			ServiceName: sub.ServiceName,
			Price:       sub.Price,
			UserID:      sub.UserID,
			StartDate:   sub.StartDate,
			EndDate:     sub.EndDate,
		})
	}

	return dto.ReadSubscriptionsListResponse{Subscriptions: subsResponse}, nil
}
