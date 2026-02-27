package dto

import (
	"github.com/google/uuid"
)

// GetSubscriptionsAmountRequest описывает запрос на подсчет суммарной стоимости подписок
type GetSubscriptionsAmountRequest struct {
	UserID      uuid.UUID `json:"user_id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	ServiceName string    `json:"service_name" example:"Yandex Music"`
}

// GetSubscriptionsAmountResponse описывает ответ с суммарной стоимостью подписок
type GetSubscriptionsAmountResponse struct {
	TotalAmount int64 `json:"total_amount" example:"400"`
}
