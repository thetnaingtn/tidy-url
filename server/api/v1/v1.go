package v1

import (
	"context"
	"fmt"
	"math"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/thetnaingtn/tidy-url/internal/config"
	v1pb "github.com/thetnaingtn/tidy-url/proto/gen/api/v1"
	"github.com/thetnaingtn/tidy-url/server/frontend"
	"github.com/thetnaingtn/tidy-url/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type APIV1Service struct {
	v1pb.UnimplementedUrlShortenerServer
	store      store.Store
	grpcServer *grpc.Server
	config     *config.Config
}

func NewAPIV1Service(store store.Store, config *config.Config, server *grpc.Server) *APIV1Service {
	apiService := &APIV1Service{
		store:      store,
		grpcServer: server,
		config:     config,
	}

	v1pb.RegisterUrlShortenerServer(server, apiService)

	return apiService
}

func (s *APIV1Service) RegisterGateway(ctx context.Context, mux *http.ServeMux) error {
	address := fmt.Sprintf("%s:%s", s.config.Addr, s.config.Port)
	conn, err := grpc.NewClient(address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt32)),
	)

	if err != nil {
		return err
	}

	gwmux := runtime.NewServeMux()

	if err := v1pb.RegisterUrlShortenerHandler(ctx, gwmux, conn); err != nil {
		return err
	}

	grpcWebOptions := []grpcweb.Option{
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true
		}),
	}
	frontendService := frontend.NewFrontendService(s.store, s.config)
	grpcWebProxy := grpcweb.WrapServer(s.grpcServer, grpcWebOptions...)

	handler := func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Path) >= 8 && r.URL.Path[:8] == "/api.v1." {
			fmt.Println("Handling gRPC-Web request")
			grpcWebProxy.ServeHTTP(w, r)
			return
		}

		if len(r.URL.Path) >= 4 && r.URL.Path[:4] == "/v1/" {
			fmt.Println("Handling gRPC-Gateway request")
			gwmux.ServeHTTP(w, r)
			return
		}

		frontendService.ServeHTTP(w, r)
	}

	mux.Handle("/", http.HandlerFunc(handler))

	return nil
}
