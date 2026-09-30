//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Vaziria/pkl/backend/pkgs/pkl_config"
	"github.com/google/wire"
)

func InitializeApp() (App, error) {

	wire.Build(
		pkl_config.NewConfig,
		NewMigrationFunc,
		NewDatabase,
		NewApp,
	)

	return nil, nil
}
