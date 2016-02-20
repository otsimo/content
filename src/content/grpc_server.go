package content

import (
	"github.com/otsimo/api/apipb"
	"golang.org/x/net/context"
)

type contentGrpcServer struct {
	server *Server
}

func (w *contentGrpcServer) List(ctx context.Context, query *apipb.ContentListRequest) (*apipb.ContentListResponse, error) {
	return &apipb.ContentListResponse{}, nil
}
