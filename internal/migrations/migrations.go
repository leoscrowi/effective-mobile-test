package migrations

import (
	"fmt"
	"github.com/leoscrowi/effective-mobile-test/domain/subscription"
	"gorm.io/gorm"
)

var models = []interface{}{
	&subscription.Subscription{},
}

func Migrate(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	for _, model := range models {
		if err := tx.AutoMigrate(model); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to migrate model %T: %w", model, err)
		}
	}

	return tx.Commit().Error
}
