package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	userv1 "github.com/Vaziria/pkl/backend/gen/catur/user/v1"
	"github.com/Vaziria/pkl/backend/gen/catur/user/v1/userv1connect"
)

// 1. validasi ?

type AuthService struct{}

// Login implements [userv1connect.AuthServiceHandler].
func (a *AuthService) Login(context.Context, *connect.Request[userv1.LoginRequest]) (*connect.Response[userv1.LoginResponse], error) {
	return nil, errors.New("Programmer males, belum ditulis")
}

// Register implements [userv1connect.AuthServiceHandler].
func (a *AuthService) Register(ctx context.Context, req *connect.Request[userv1.RegisterRequest]) (*connect.Response[userv1.RegisterResponse], error) {
	var err error
	resp := &userv1.RegisterResponse{
		Name: req.Msg.Username,
	}

	return connect.NewResponse(resp), err

}

func main() {

	// server
	mux := http.NewServeMux()

	// validasi
	validator := validate.NewInterceptor()

	// option rpc
	opts := connect.WithInterceptors(
		validator,
	)

	// service
	authServiceImpl := &AuthService{}
	path, authService := userv1connect.NewAuthServiceHandler(authServiceImpl, opts)
	mux.Handle(path, authService)

	// running webserver & config protocol
	server := NewServer(mux)

	slog.Info("Running Webserver")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func NewServer(mux *http.ServeMux) *http.Server {
	// The modern replacement for the deprecated x/net/http2/h2c: net/http has spoken
	// unencrypted HTTP/2 natively since Go 1.24.
	protocols := new(http.Protocols)
	protocols.SetHTTP1(true)
	protocols.SetUnencryptedHTTP2(true)

	return &http.Server{
		Addr:      "localhost:8080",
		Handler:   mux,
		Protocols: protocols,
	}
}
