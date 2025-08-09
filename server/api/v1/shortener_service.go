package v1

import (
	"context"

	v1pb "github.com/thetnaingtn/tidy-url/proto/gen/api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *APIV1Service) MakeTidyUrl(ctx context.Context, req *v1pb.ShortenUrlRequest) (*v1pb.TidyUrl, error) {
	tidyUrl, err := s.store.Create(ctx, req.GetLongUrl())
	if err != nil {
		return nil, err
	}

	return &v1pb.TidyUrl{
		Id:         tidyUrl.Id,
		LongUrl:    tidyUrl.LongUrl,
		EncodedStr: tidyUrl.EncodedStr,
		CreatedAt:  timestamppb.New(tidyUrl.CreatedAt),
	}, nil
}
