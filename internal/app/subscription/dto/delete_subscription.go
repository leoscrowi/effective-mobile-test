package dto

import "github.com/google/uuid"

type DeleteSubscriptionRequest struct {
	ID uuid.UUID `json:"id"`
}
