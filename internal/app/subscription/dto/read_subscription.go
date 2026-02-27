package dto

import (
	"github.com/google/uuid"
)

// ReadSubscriptionRequest описывает запрос на получение подписки
type ReadSubscriptionRequest struct {
	ID uuid.UUID `json:"id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
}

// ReadSubscriptionResponse описывает ответ с информацией о подписке
type ReadSubscriptionResponse struct {
	ID          uuid.UUID `json:"id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	ServiceName string    `json:"service_name" example:"Yandex Music"`
	Price       int64     `json:"price" example:"400"`
	UserID      uuid.UUID `json:"user_id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	StartDate   string    `json:"start_date" example:"07-2026"`
	EndDate     string    `json:"end_time" example:"08-2026"`
}

// ReadSubscriptionsListResponse описывает ответ со списком подписок
type ReadSubscriptionsListResponse struct {
	Subscriptions []ReadSubscriptionResponse `json:"subscriptions"`
}
