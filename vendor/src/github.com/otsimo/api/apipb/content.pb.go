// Code generated by protoc-gen-gogo.
// source: content.proto
// DO NOT EDIT!

package apipb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ContentListRequest_ListStatus int32

const (
	ContentListRequest_BOTH          ContentListRequest_ListStatus = 0
	ContentListRequest_ONLY_DRAFT    ContentListRequest_ListStatus = 1
	ContentListRequest_ONLY_APPROVED ContentListRequest_ListStatus = 2
)

var ContentListRequest_ListStatus_name = map[int32]string{
	0: "BOTH",
	1: "ONLY_DRAFT",
	2: "ONLY_APPROVED",
}
var ContentListRequest_ListStatus_value = map[string]int32{
	"BOTH":          0,
	"ONLY_DRAFT":    1,
	"ONLY_APPROVED": 2,
}

func (x ContentListRequest_ListStatus) String() string {
	return proto.EnumName(ContentListRequest_ListStatus_name, int32(x))
}

type ContentListRequest_SortBy int32

const (
	ContentListRequest_WEIGHT ContentListRequest_SortBy = 0
	ContentListRequest_TIME   ContentListRequest_SortBy = 1
)

var ContentListRequest_SortBy_name = map[int32]string{
	0: "WEIGHT",
	1: "TIME",
}
var ContentListRequest_SortBy_value = map[string]int32{
	"WEIGHT": 0,
	"TIME":   1,
}

func (x ContentListRequest_SortBy) String() string {
	return proto.EnumName(ContentListRequest_SortBy_name, int32(x))
}

type Content struct {
	Slug      string   `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	Title     string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Language  string   `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`
	Date      int64    `protobuf:"varint,4,opt,name=date,proto3" json:"date,omitempty"`
	Draft     bool     `protobuf:"varint,5,opt,name=draft,proto3" json:"draft,omitempty"`
	WrittenAt string   `protobuf:"bytes,6,opt,name=written_at,proto3" json:"written_at,omitempty"`
	Author    string   `protobuf:"bytes,7,opt,name=author,proto3" json:"author,omitempty"`
	Category  string   `protobuf:"bytes,8,opt,name=category,proto3" json:"category,omitempty"`
	Url       string   `protobuf:"bytes,9,opt,name=url,proto3" json:"url,omitempty"`
	Weight    int32    `protobuf:"varint,10,opt,name=weight,proto3" json:"weight,omitempty"`
	Keywords  []string `protobuf:"bytes,11,rep,name=keywords" json:"keywords,omitempty"`
	Markdown  []byte   `protobuf:"bytes,13,opt,name=markdown,proto3" json:"markdown,omitempty"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}

type ContentListRequest struct {
	Status        ContentListRequest_ListStatus `protobuf:"varint,1,opt,name=status,proto3,enum=apipb.ContentListRequest_ListStatus" json:"status,omitempty"`
	Limit         int32                         `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Category      string                        `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Offset        int32                         `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Language      string                        `protobuf:"bytes,5,opt,name=language,proto3" json:"language,omitempty"`
	OnlyHtmlUrl   bool                          `protobuf:"varint,6,opt,name=only_html_url,proto3" json:"only_html_url,omitempty"`
	Sort          ContentListRequest_SortBy     `protobuf:"varint,7,opt,name=sort,proto3,enum=apipb.ContentListRequest_SortBy" json:"sort,omitempty"`
	ProfileId     string                        `protobuf:"bytes,10,opt,name=profile_id,proto3" json:"profile_id,omitempty"`
	ClientVersion string                        `protobuf:"bytes,11,opt,name=client_version,proto3" json:"client_version,omitempty"`
}

func (m *ContentListRequest) Reset()         { *m = ContentListRequest{} }
func (m *ContentListRequest) String() string { return proto.CompactTextString(m) }
func (*ContentListRequest) ProtoMessage()    {}

type ContentListResponse struct {
	Contents []*Content `protobuf:"bytes,1,rep,name=contents" json:"contents,omitempty"`
}

func (m *ContentListResponse) Reset()         { *m = ContentListResponse{} }
func (m *ContentListResponse) String() string { return proto.CompactTextString(m) }
func (*ContentListResponse) ProtoMessage()    {}

func (m *ContentListResponse) GetContents() []*Content {
	if m != nil {
		return m.Contents
	}
	return nil
}

type ContentGetRequest struct {
	Slug string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (m *ContentGetRequest) Reset()         { *m = ContentGetRequest{} }
func (m *ContentGetRequest) String() string { return proto.CompactTextString(m) }
func (*ContentGetRequest) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("apipb.ContentListRequest_ListStatus", ContentListRequest_ListStatus_name, ContentListRequest_ListStatus_value)
	proto.RegisterEnum("apipb.ContentListRequest_SortBy", ContentListRequest_SortBy_name, ContentListRequest_SortBy_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for ContentService service

type ContentServiceClient interface {
	List(ctx context.Context, in *ContentListRequest, opts ...grpc.CallOption) (*ContentListResponse, error)
	Get(ctx context.Context, in *ContentGetRequest, opts ...grpc.CallOption) (*Content, error)
}

type contentServiceClient struct {
	cc *grpc.ClientConn
}

func NewContentServiceClient(cc *grpc.ClientConn) ContentServiceClient {
	return &contentServiceClient{cc}
}

func (c *contentServiceClient) List(ctx context.Context, in *ContentListRequest, opts ...grpc.CallOption) (*ContentListResponse, error) {
	out := new(ContentListResponse)
	err := grpc.Invoke(ctx, "/apipb.ContentService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) Get(ctx context.Context, in *ContentGetRequest, opts ...grpc.CallOption) (*Content, error) {
	out := new(Content)
	err := grpc.Invoke(ctx, "/apipb.ContentService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ContentService service

type ContentServiceServer interface {
	List(context.Context, *ContentListRequest) (*ContentListResponse, error)
	Get(context.Context, *ContentGetRequest) (*Content, error)
}

func RegisterContentServiceServer(s *grpc.Server, srv ContentServiceServer) {
	s.RegisterService(&_ContentService_serviceDesc, srv)
}

func _ContentService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ContentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ContentServiceServer).List(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ContentService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ContentGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ContentServiceServer).Get(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _ContentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apipb.ContentService",
	HandlerType: (*ContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ContentService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ContentService_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func (m *Content) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Content) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Slug) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintContent(data, i, uint64(len(m.Slug)))
		i += copy(data[i:], m.Slug)
	}
	if len(m.Title) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintContent(data, i, uint64(len(m.Title)))
		i += copy(data[i:], m.Title)
	}
	if len(m.Language) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintContent(data, i, uint64(len(m.Language)))
		i += copy(data[i:], m.Language)
	}
	if m.Date != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintContent(data, i, uint64(m.Date))
	}
	if m.Draft {
		data[i] = 0x28
		i++
		if m.Draft {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if len(m.WrittenAt) > 0 {
		data[i] = 0x32
		i++
		i = encodeVarintContent(data, i, uint64(len(m.WrittenAt)))
		i += copy(data[i:], m.WrittenAt)
	}
	if len(m.Author) > 0 {
		data[i] = 0x3a
		i++
		i = encodeVarintContent(data, i, uint64(len(m.Author)))
		i += copy(data[i:], m.Author)
	}
	if len(m.Category) > 0 {
		data[i] = 0x42
		i++
		i = encodeVarintContent(data, i, uint64(len(m.Category)))
		i += copy(data[i:], m.Category)
	}
	if len(m.Url) > 0 {
		data[i] = 0x4a
		i++
		i = encodeVarintContent(data, i, uint64(len(m.Url)))
		i += copy(data[i:], m.Url)
	}
	if m.Weight != 0 {
		data[i] = 0x50
		i++
		i = encodeVarintContent(data, i, uint64(m.Weight))
	}
	if len(m.Keywords) > 0 {
		for _, s := range m.Keywords {
			data[i] = 0x5a
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if m.Markdown != nil {
		if len(m.Markdown) > 0 {
			data[i] = 0x6a
			i++
			i = encodeVarintContent(data, i, uint64(len(m.Markdown)))
			i += copy(data[i:], m.Markdown)
		}
	}
	return i, nil
}

func (m *ContentListRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ContentListRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintContent(data, i, uint64(m.Status))
	}
	if m.Limit != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintContent(data, i, uint64(m.Limit))
	}
	if len(m.Category) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintContent(data, i, uint64(len(m.Category)))
		i += copy(data[i:], m.Category)
	}
	if m.Offset != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintContent(data, i, uint64(m.Offset))
	}
	if len(m.Language) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintContent(data, i, uint64(len(m.Language)))
		i += copy(data[i:], m.Language)
	}
	if m.OnlyHtmlUrl {
		data[i] = 0x30
		i++
		if m.OnlyHtmlUrl {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.Sort != 0 {
		data[i] = 0x38
		i++
		i = encodeVarintContent(data, i, uint64(m.Sort))
	}
	if len(m.ProfileId) > 0 {
		data[i] = 0x52
		i++
		i = encodeVarintContent(data, i, uint64(len(m.ProfileId)))
		i += copy(data[i:], m.ProfileId)
	}
	if len(m.ClientVersion) > 0 {
		data[i] = 0x5a
		i++
		i = encodeVarintContent(data, i, uint64(len(m.ClientVersion)))
		i += copy(data[i:], m.ClientVersion)
	}
	return i, nil
}

func (m *ContentListResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ContentListResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Contents) > 0 {
		for _, msg := range m.Contents {
			data[i] = 0xa
			i++
			i = encodeVarintContent(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ContentGetRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ContentGetRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Slug) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintContent(data, i, uint64(len(m.Slug)))
		i += copy(data[i:], m.Slug)
	}
	return i, nil
}

func encodeFixed64Content(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Content(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintContent(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Content) Size() (n int) {
	var l int
	_ = l
	l = len(m.Slug)
	if l > 0 {
		n += 1 + l + sovContent(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovContent(uint64(l))
	}
	l = len(m.Language)
	if l > 0 {
		n += 1 + l + sovContent(uint64(l))
	}
	if m.Date != 0 {
		n += 1 + sovContent(uint64(m.Date))
	}
	if m.Draft {
		n += 2
	}
	l = len(m.WrittenAt)
	if l > 0 {
		n += 1 + l + sovContent(uint64(l))
	}
	l = len(m.Author)
	if l > 0 {
		n += 1 + l + sovContent(uint64(l))
	}
	l = len(m.Category)
	if l > 0 {
		n += 1 + l + sovContent(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovContent(uint64(l))
	}
	if m.Weight != 0 {
		n += 1 + sovContent(uint64(m.Weight))
	}
	if len(m.Keywords) > 0 {
		for _, s := range m.Keywords {
			l = len(s)
			n += 1 + l + sovContent(uint64(l))
		}
	}
	if m.Markdown != nil {
		l = len(m.Markdown)
		if l > 0 {
			n += 1 + l + sovContent(uint64(l))
		}
	}
	return n
}

func (m *ContentListRequest) Size() (n int) {
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovContent(uint64(m.Status))
	}
	if m.Limit != 0 {
		n += 1 + sovContent(uint64(m.Limit))
	}
	l = len(m.Category)
	if l > 0 {
		n += 1 + l + sovContent(uint64(l))
	}
	if m.Offset != 0 {
		n += 1 + sovContent(uint64(m.Offset))
	}
	l = len(m.Language)
	if l > 0 {
		n += 1 + l + sovContent(uint64(l))
	}
	if m.OnlyHtmlUrl {
		n += 2
	}
	if m.Sort != 0 {
		n += 1 + sovContent(uint64(m.Sort))
	}
	l = len(m.ProfileId)
	if l > 0 {
		n += 1 + l + sovContent(uint64(l))
	}
	l = len(m.ClientVersion)
	if l > 0 {
		n += 1 + l + sovContent(uint64(l))
	}
	return n
}

func (m *ContentListResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Contents) > 0 {
		for _, e := range m.Contents {
			l = e.Size()
			n += 1 + l + sovContent(uint64(l))
		}
	}
	return n
}

func (m *ContentGetRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Slug)
	if l > 0 {
		n += 1 + l + sovContent(uint64(l))
	}
	return n
}

func sovContent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozContent(x uint64) (n int) {
	return sovContent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Content) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Content: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Content: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slug", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Slug = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Language", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Language = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			m.Date = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Date |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Draft", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Draft = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WrittenAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WrittenAt = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Author", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Author = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Category", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Category = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			m.Weight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Weight |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keywords", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Keywords = append(m.Keywords, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Markdown", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthContent
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Markdown = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContent(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ContentListRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ContentListRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContentListRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Status |= (ContentListRequest_ListStatus(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Limit |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Category", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Category = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Offset |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Language", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Language = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnlyHtmlUrl", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OnlyHtmlUrl = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sort", wireType)
			}
			m.Sort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Sort |= (ContentListRequest_SortBy(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfileId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProfileId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientVersion = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContent(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ContentListResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ContentListResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContentListResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthContent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contents = append(m.Contents, &Content{})
			if err := m.Contents[len(m.Contents)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContent(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ContentGetRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ContentGetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContentGetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slug", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Slug = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContent(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipContent(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContent
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowContent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowContent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthContent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowContent
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipContent(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthContent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContent   = fmt.Errorf("proto: integer overflow")
)
