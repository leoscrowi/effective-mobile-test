package subscription

import (
	"github.com/google/uuid"
)

type Subscription struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ServiceName string
	Price       int64
	UserID      uuid.UUID
	StartDate   string
	EndDate     *string
}
