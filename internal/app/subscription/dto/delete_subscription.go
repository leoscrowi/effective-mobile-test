package dto

import "github.com/google/uuid"

// DeleteSubscriptionRequest описывает запрос на удаление подписки
type DeleteSubscriptionRequest struct {
	ID uuid.UUID `json:"id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
}
