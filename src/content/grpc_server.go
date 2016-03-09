package content

import (
	"errors"

	"github.com/otsimo/api/apipb"
	"golang.org/x/net/context"
)

type contentGrpcServer struct {
	server *Server
}

func (w *contentGrpcServer) List(ctx context.Context, query *apipb.ContentListRequest) (*apipb.ContentListResponse, error) {
	var contents []*apipb.Content

	for _, c := range w.server.Content.contents {
		if c.Language != query.Language {
			continue
		}
		if query.Status == apipb.ContentListRequest_ONLY_APPROVED && c.Draft {
			continue
		}
		if query.Status == apipb.ContentListRequest_ONLY_DRAFT && !c.Draft {
			continue
		}
		if query.Category != "" && query.Category != c.Category {
			continue
		}

		contents = append(contents, c)
	}
	return &apipb.ContentListResponse{
		AssetVersion: w.server.Content.assetVersion,
		Contents:     contents,
	}, nil
}

func (w *contentGrpcServer) Get(ctx context.Context, in *apipb.ContentGetRequest) (*apipb.Content, error) {
	for _, c := range w.server.Content.contents {
		if c.Slug == in.Slug {
			return c, nil
		}
	}
	return nil, errors.New("not found")
}
