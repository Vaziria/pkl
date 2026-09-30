package main

import "gorm.io/gorm"

type MigrationFunc func() error

func NewMigrationFunc(db *gorm.DB) MigrationFunc {
	return func() error {
		err := db.AutoMigrate(
			User{},
		)

		return err

	}
}
