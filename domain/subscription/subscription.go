package subscription

import (
	"github.com/google/uuid"
	"github.com/govalues/decimal"
	"time"
)

type Subscription struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	ServiceName string
	Price       decimal.Decimal `gorm:"type:decimal(10,2);"`
	UserID      uuid.UUID
	StartDate   time.Time
	EndDate     time.Time
}
