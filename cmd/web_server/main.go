package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Vaziria/pkl/backend/pkgs/pkl_config"
)

func main() {
	appFunction, err := InitializeApp()

	if err != nil {
		panic(err)
	}

	err = appFunction(context.Background())

	if err != nil {
		panic(err)
	}

}

func NewServer(
	cfg *pkl_config.Config,
	mux *http.ServeMux,
) *http.Server {
	// The modern replacement for the deprecated x/net/http2/h2c: net/http has spoken
	// unencrypted HTTP/2 natively since Go 1.24.
	protocols := new(http.Protocols)
	protocols.SetHTTP1(true)
	protocols.SetUnencryptedHTTP2(true)

	return &http.Server{
		Addr:      fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler:   mux,
		Protocols: protocols,
	}
}
