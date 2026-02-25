package dto

import (
	"github.com/google/uuid"
	"github.com/govalues/decimal"
	"time"
)

type EditSubscriptionRequset struct {
	ID          uuid.UUID       `json:"id"`
	ServiceName string          `json:"service_name"`
	Price       decimal.Decimal `json:"price"`
	UserID      uuid.UUID       `json:"user_id"`
	StartDate   time.Time       `json:"start_date"`
	EndDate     time.Time       `json:"end_date"`
}
