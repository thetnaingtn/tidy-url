package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"

	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/soheilhy/cmux"

	"github.com/thetnaingtn/tidy-url/internal/config"
	v1 "github.com/thetnaingtn/tidy-url/server/api/v1"
	"github.com/thetnaingtn/tidy-url/store"
	"google.golang.org/grpc"
)

type Server struct {
	Store  *store.Store
	Config *config.Config

	grpcServer *grpc.Server
	server     http.Server
}

func NewServer(ctx context.Context, store *store.Store, config *config.Config) (*Server, error) {
	mux := http.NewServeMux()

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcrecovery.UnaryServerInterceptor(),
		),
	)

	service := v1.NewAPIV1Service(store, config, grpcServer)

	if err := service.RegisterGateway(ctx, mux); err != nil {
		return nil, err
	}

	return &Server{
		Store:      store,
		Config:     config,
		grpcServer: grpcServer,
		server:     http.Server{Handler: mux},
	}, nil
}

func (s *Server) Start() error {
	address := fmt.Sprintf("%s:%s", s.Config.Addr, s.Config.Port)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	muxServer := cmux.New(listener)

	go func() {
		grpcListener := muxServer.MatchWithWriters(
			cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"),
		)

		if err := s.grpcServer.Serve(grpcListener); err != nil {
			slog.Error("Failed to start gRPC server", "error", err)
		}
	}()

	go func() {
		httpListener := muxServer.Match(cmux.HTTP1Fast())
		if err := s.server.Serve(httpListener); err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to start HTTP server", "error", err)
		}
	}()

	go func() {
		if err := muxServer.Serve(); err != nil {
			slog.Error("Failed to start cmux server", "error", err)
		}
	}()

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*20)
	defer cancel()
	slog.Info("Shutting down tidy-url server...")
	if err := s.server.Shutdown(ctx); err != nil {
		if err != http.ErrServerClosed {
			slog.Error("Failed to shutdown HTTP server", "error", err)
			return err
		}
	}

	s.grpcServer.GracefulStop()
	s.Store.Close()
	slog.Info("tidy-url server shutdown gracefully")

	return nil
}
