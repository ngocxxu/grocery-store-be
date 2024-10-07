package migrations

import (
	"github.com/ngocxxu/grocery-store-svelte-be/internal/model"
	"gorm.io/gorm"
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
	{
		ID: "202310062_create_units_table", // Unit trước
		Migrate: func(db *gorm.DB) error {
						return db.AutoMigrate(&model.Unit{})
		},
		Rollback: func(db *gorm.DB) error {
						return db.Migrator().DropTable("units")
		},
},
{
		ID: "202310063_create_weight_options_table",
		Migrate: func(db *gorm.DB) error {
						return db.AutoMigrate(&model.WeightOption{})
		},
		Rollback: func(db *gorm.DB) error {
						return db.Migrator().DropTable("weight_options")
		},
},
{
		ID: "202310064_create_products_table",
		Migrate: func(db *gorm.DB) error {
						return db.AutoMigrate(&model.Product{})
		},
		Rollback: func(db *gorm.DB) error {
						return db.Migrator().DropTable("products")
		},
},
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