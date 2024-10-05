package migrations

import (
	"gorm.io/gorm"
	"github.com/ngocxxu/grocery-store-svelte-be/internal/model"
)

type Migration struct {
	ID       string
	Migrate  func(*gorm.DB) error
	Rollback func(*gorm.DB) error
}

var migrations = []Migration{
	{
		ID: "202310061_create_users_table",
		Migrate: func(db *gorm.DB) error {
			return db.AutoMigrate(&model.User{})
		},
		Rollback: func(db *gorm.DB) error {
			return db.Migrator().DropTable("users")
		},
	},
	// Add more migrations
}

func RunMigrations(db *gorm.DB) error {
	for _, m := range migrations {
		if err := m.Migrate(db); err != nil {
			return err
		}
	}
	return nil
}

func RollbackMigrations(db *gorm.DB) error {
	for i := len(migrations) - 1; i >= 0; i-- {
		if err := migrations[i].Rollback(db); err != nil {
			return err
		}
	}
	return nil
}