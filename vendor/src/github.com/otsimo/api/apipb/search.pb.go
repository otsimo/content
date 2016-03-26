// Code generated by protoc-gen-gogo.
// source: search.proto
// DO NOT EDIT!

package apipb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for SearchService service

type SearchServiceClient interface {
	IndexDatabase(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*Response, error)
	ReindexAll(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*Response, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type searchServiceClient struct {
	cc *grpc.ClientConn
}

func NewSearchServiceClient(cc *grpc.ClientConn) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) IndexDatabase(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/apipb.SearchService/IndexDatabase", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) ReindexAll(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/apipb.SearchService/ReindexAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := grpc.Invoke(ctx, "/apipb.SearchService/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SearchService service

type SearchServiceServer interface {
	IndexDatabase(context.Context, *IndexRequest) (*Response, error)
	ReindexAll(context.Context, *IndexRequest) (*Response, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
}

func RegisterSearchServiceServer(s *grpc.Server, srv SearchServiceServer) {
	s.RegisterService(&_SearchService_serviceDesc, srv)
}

func _SearchService_IndexDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(IndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SearchServiceServer).IndexDatabase(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _SearchService_ReindexAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(IndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SearchServiceServer).ReindexAll(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _SearchService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SearchServiceServer).Search(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _SearchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apipb.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IndexDatabase",
			Handler:    _SearchService_IndexDatabase_Handler,
		},
		{
			MethodName: "ReindexAll",
			Handler:    _SearchService_ReindexAll_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _SearchService_Search_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptorSearch = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x4d, 0x2c,
	0x4a, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2c, 0xc8, 0x2c, 0x48, 0x92,
	0xe2, 0xcb, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0x86, 0x08, 0x1b, 0x6d, 0x67, 0xe4, 0xe2,
	0x0d, 0x06, 0xab, 0x0b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x32, 0xe7, 0xe2, 0xf5, 0xcc,
	0x4b, 0x49, 0xad, 0x70, 0x49, 0x2c, 0x49, 0x4c, 0x4a, 0x2c, 0x4e, 0x15, 0x12, 0xd6, 0x03, 0x6b,
	0xd5, 0x03, 0x8b, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x48, 0xf1, 0x43, 0x05, 0x83, 0x52,
	0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x95, 0x18, 0x84, 0x4c, 0xb8, 0xb8, 0x82, 0x52, 0x33, 0x41,
	0x8a, 0x1c, 0x73, 0x72, 0x88, 0xd6, 0x65, 0xce, 0xc5, 0x06, 0xb1, 0x5f, 0x48, 0x04, 0x2a, 0x09,
	0xe1, 0xc2, 0xb4, 0x88, 0xa2, 0x89, 0xc2, 0x34, 0x3a, 0x89, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0x01,
	0x88, 0x1f, 0x00, 0xf1, 0x84, 0xc7, 0x72, 0x0c, 0x8b, 0x98, 0x98, 0xfd, 0x43, 0x82, 0x93, 0xd8,
	0xc0, 0x1e, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x75, 0xcd, 0x4e, 0xf0, 0xff, 0x00, 0x00,
	0x00,
}
