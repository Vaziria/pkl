package main

import (
	"github.com/Vaziria/pkl/backend/pkgs/pkl_config"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewDatabase(
	cfg *pkl_config.Config,

) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.Database.Filename), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(1) // avoids "database is locked"

	return db, nil
}
