package v1

import (
	"context"
	"fmt"

	v1pb "github.com/thetnaingtn/tidy-url/proto/gen/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *APIV1Service) MakeTidyUrl(ctx context.Context, req *v1pb.MakeTidyUrlRequest) (*v1pb.MakeTidyUrlResponse, error) {
	longUrl := req.GetLongUrl()
	if longUrl == "" {
		return nil, status.Error(codes.InvalidArgument, "long_url cannot be empty")
	}

	tidyUrl, err := s.store.GetTidyUrlByOriginalUrl(ctx, longUrl)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve tidy URL")
	}

	if tidyUrl != nil {
		return &v1pb.MakeTidyUrlResponse{
			TidyUrl: s.createTidyUrl(s.config.BaseURL, tidyUrl.EncodedStr),
		}, nil
	}

	tidyUrl, err = s.store.Create(ctx, req.GetLongUrl())
	if err != nil {
		return nil, err
	}

	return &v1pb.MakeTidyUrlResponse{
		TidyUrl: s.createTidyUrl(s.config.BaseURL, tidyUrl.EncodedStr),
	}, nil
}

func (s *APIV1Service) ExpandTidyUrl(ctx context.Context, req *v1pb.ExpandTidyUrlRequest) (*v1pb.ExpandTidyUrlResponse, error) {
	tidyUrl, err := s.store.GetTidyUrlByEncodedStr(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	if tidyUrl == nil {
		return nil, status.Error(codes.NotFound, "tidy url not found")
	}

	return &v1pb.ExpandTidyUrlResponse{
		LongUrl: tidyUrl.LongUrl,
	}, nil
}

func (s *APIV1Service) createTidyUrl(baseUrl string, encodedStr string) string {
	return fmt.Sprintf("%s/v1/expand/%s", baseUrl, encodedStr)
}
