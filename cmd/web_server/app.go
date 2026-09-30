package main

import (
	"context"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"connectrpc.com/validate"
	"github.com/Vaziria/pkl/backend/gen/catur/user/v1/userv1connect"
	"github.com/Vaziria/pkl/backend/pkgs/pkl_config"
	"gorm.io/gorm"
)

type App func(ctx context.Context) error

func NewApp(
	cfg *pkl_config.Config,
	db *gorm.DB,
	migrationFunc MigrationFunc,

) App {
	return func(ctx context.Context) error {
		slog.Info("running migrasi")
		err := migrationFunc()
		if err != nil {
			slog.Error(err.Error())
			return err
		}

		// server
		mux := http.NewServeMux()

		// validasi
		validator := validate.NewInterceptor()

		// option rpc
		opts := connect.WithInterceptors(
			validator,
		)

		// service
		authServiceImpl := NewAuthService(db)
		path, authService := userv1connect.NewAuthServiceHandler(authServiceImpl, opts)
		mux.Handle(path, authService)

		// cara register reflect
		names := []string{
			userv1connect.AuthServiceName,
		}
		reflector := grpcreflect.NewStaticReflector(names...)
		mux.Handle(grpcreflect.NewHandlerV1(reflector))
		mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

		// running webserver & config protocol
		server := NewServer(cfg, mux)

		slog.Info("Running Webserver")
		err = server.ListenAndServe()
		if err != nil {
			return err
		}
		return nil
	}
}
