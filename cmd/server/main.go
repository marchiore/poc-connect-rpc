package main

import (
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/validate"

	"poc/gen/proto/Motherboard/motherboardv1connect"
	motherboard "poc/internal/app/motherboard"
	motherboardrepo "poc/internal/repository/motherboard"
)

func main() {
	repo := &motherboardrepo.MotherboardRepository{}
	service := motherboard.NewMotherboardService(repo)
	server := motherboard.NewMotherboardServer(service)

	mux := http.NewServeMux()

	path, handler := motherboardv1connect.NewMotherboardServiceHandler(
		server,
		connect.WithInterceptors(validate.NewInterceptor()),
	)

	mux.Handle(path, handler)
	p := new(http.Protocols)
	p.SetHTTP1(true)
	// Use h2c so we can serve HTTP/2 without TLS.
	p.SetUnencryptedHTTP2(true)
	s := http.Server{
		Addr:      "localhost:8080",
		Handler:   mux,
		Protocols: p,
	}
	s.ListenAndServe()
}
