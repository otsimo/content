package content

import (
	"errors"
	"models"

	"github.com/Sirupsen/logrus"
	"github.com/otsimo/api/apipb"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
)

type contentGrpcServer struct {
	server *Server
}

func (w *contentGrpcServer) Push(ctx context.Context, in *apipb.Content) (*apipb.Response, error) {
	jwt, err := getJWTToken(ctx)
	if err != nil {
		logrus.Errorf("grpc_server.go: failed to get jwt %+v", err)
		return nil, errors.New("failed to get jwt")
	}
	id, email, err := w.authToken(jwt, true)
	if err != nil {
		logrus.Errorf("grpc_server.go: failed to authorize user %+v", err)
		return nil, errors.New("unauthorized user")
	}
	if !bson.IsObjectIdHex(id) {
		return nil, models.ErrorInvalidID
	}

	err = w.server.Insert(in, email, bson.ObjectIdHex(id))
	if err != nil {
		return nil, err
	}
	return &apipb.Response{Type: 0, Message: "success"}, nil
}

func (w *contentGrpcServer) Approve(ctx context.Context, in *apipb.ContentApproveRequest) (*apipb.Response, error) {
	jwt, err := getJWTToken(ctx)
	if err != nil {
		logrus.Errorf("grpc_server.go: failed to get jwt %+v", err)
		return nil, errors.New("failed to get jwt")
	}
	_, _, err = w.authToken(jwt, true)
	if err != nil {
		logrus.Errorf("grpc_server.go: failed to authorize user %+v", err)
		return nil, errors.New("unauthorized user")
	}
	err = w.server.Approve(in.Slug)
	if err != nil {
		return nil, err
	}
	return &apipb.Response{Type: 0, Message: "success"}, nil
}

func (w *contentGrpcServer) List(ctx context.Context, query *apipb.ContentListRequest) (*apipb.ContentListResponse, error) {
	res, err := w.server.Storage.List(*query)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, models.ErrorNotFound
	}
	return &apipb.ContentListResponse{
		Contents: res,
	}, nil
}
