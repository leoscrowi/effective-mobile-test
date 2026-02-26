package dto

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/govalues/decimal"
	"time"
)

type EditSubscriptionRequset struct {
	ID          uuid.UUID
	ServiceName string          `json:"service_name"`
	Price       decimal.Decimal `json:"price"`
	UserID      uuid.UUID       `json:"user_id"`
	StartDate   time.Time       `json:"start_date"`
	EndDate     time.Time       `json:"end_date"`
}

func ValidateDto(dto EditSubscriptionRequset) error {
	if dto.ID == uuid.Nil {
		return fmt.Errorf("id: cannot be empty")
	}

	if dto.ServiceName == "" {
		return fmt.Errorf("service_name: cannot be empty")
	}

	if dto.Price.IsZero() {
		return fmt.Errorf("price: cannot be zero")
	}

	if dto.UserID == uuid.Nil {
		return fmt.Errorf("user_id: cannot be empty")
	}

	if dto.StartDate.IsZero() {
		return fmt.Errorf("start_date: cannot be empty")
	}

	if dto.EndDate.IsZero() {
		return fmt.Errorf("end_date: cannot be empty")
	}

	if dto.EndDate.Before(dto.StartDate) {
		return fmt.Errorf("end_date: cannot be before start_date")
	}

	return nil
}
