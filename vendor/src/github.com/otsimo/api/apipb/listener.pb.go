// Code generated by protoc-gen-gogo.
// source: listener.proto
// DO NOT EDIT!

package apipb

import proto "github.com/gogo/protobuf/proto"
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

type DeviceInfo struct {
	VendorId           string `protobuf:"bytes,1,opt,name=vendorId,proto3" json:"vendorId,omitempty"`
	ClientSdk          string `protobuf:"bytes,2,opt,name=clientSdk,proto3" json:"clientSdk,omitempty"`
	BundleIdentifier   string `protobuf:"bytes,3,opt,name=bundleIdentifier,proto3" json:"bundleIdentifier,omitempty"`
	BundleVersion      string `protobuf:"bytes,4,opt,name=bundleVersion,proto3" json:"bundleVersion,omitempty"`
	BundleShortVersion string `protobuf:"bytes,5,opt,name=bundleShortVersion,proto3" json:"bundleShortVersion,omitempty"`
	DeviceType         string `protobuf:"bytes,6,opt,name=deviceType,proto3" json:"deviceType,omitempty"`
	DeviceName         string `protobuf:"bytes,7,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	OsName             string `protobuf:"bytes,8,opt,name=osName,proto3" json:"osName,omitempty"`
	SystemVersion      string `protobuf:"bytes,9,opt,name=systemVersion,proto3" json:"systemVersion,omitempty"`
	LanguageCode       string `protobuf:"bytes,10,opt,name=languageCode,proto3" json:"languageCode,omitempty"`
	CountryCode        string `protobuf:"bytes,11,opt,name=countryCode,proto3" json:"countryCode,omitempty"`
}

func (m *DeviceInfo) Reset()         { *m = DeviceInfo{} }
func (m *DeviceInfo) String() string { return proto.CompactTextString(m) }
func (*DeviceInfo) ProtoMessage()    {}

type Event struct {
	// UserId is profile id or child id
	UserId string `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id,omitempty"`
	// Child is secondary id
	ChildId string `protobuf:"bytes,2,opt,name=child_id,proto3" json:"child_id,omitempty"`
	// Event the event name
	Event string `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
	// Timestamp is millisecond unix time
	Timestamp int64 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// SubId will use for game id
	SubId string `protobuf:"bytes,5,opt,name=sub_id,proto3" json:"sub_id,omitempty"`
	// DeviceId is unique device identifier,
	DeviceId string `protobuf:"bytes,6,opt,name=device_id,proto3" json:"device_id,omitempty"`
	// AppId is the client app id
	AppId string `protobuf:"bytes,7,opt,name=app_id,proto3" json:"app_id,omitempty"`
	// Payload is a json data
	Payload []byte `protobuf:"bytes,11,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}

type AppEventData struct {
	Event     string      `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	AppId     string      `protobuf:"bytes,2,opt,name=app_id,proto3" json:"app_id,omitempty"`
	Device    *DeviceInfo `protobuf:"bytes,3,opt,name=device" json:"device,omitempty"`
	Timestamp int64       `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Payload   []byte      `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *AppEventData) Reset()         { *m = AppEventData{} }
func (m *AppEventData) String() string { return proto.CompactTextString(m) }
func (*AppEventData) ProtoMessage()    {}

func (m *AppEventData) GetDevice() *DeviceInfo {
	if m != nil {
		return m.Device
	}
	return nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for ListenerService service

type ListenerServiceClient interface {
	AppEvent(ctx context.Context, in *AppEventData, opts ...grpc.CallOption) (*Response, error)
	CustomEvent(ctx context.Context, opts ...grpc.CallOption) (ListenerService_CustomEventClient, error)
}

type listenerServiceClient struct {
	cc *grpc.ClientConn
}

func NewListenerServiceClient(cc *grpc.ClientConn) ListenerServiceClient {
	return &listenerServiceClient{cc}
}

func (c *listenerServiceClient) AppEvent(ctx context.Context, in *AppEventData, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/apipb.ListenerService/AppEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listenerServiceClient) CustomEvent(ctx context.Context, opts ...grpc.CallOption) (ListenerService_CustomEventClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ListenerService_serviceDesc.Streams[0], c.cc, "/apipb.ListenerService/CustomEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &listenerServiceCustomEventClient{stream}
	return x, nil
}

type ListenerService_CustomEventClient interface {
	Send(*Event) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type listenerServiceCustomEventClient struct {
	grpc.ClientStream
}

func (x *listenerServiceCustomEventClient) Send(m *Event) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenerServiceCustomEventClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ListenerService service

type ListenerServiceServer interface {
	AppEvent(context.Context, *AppEventData) (*Response, error)
	CustomEvent(ListenerService_CustomEventServer) error
}

func RegisterListenerServiceServer(s *grpc.Server, srv ListenerServiceServer) {
	s.RegisterService(&_ListenerService_serviceDesc, srv)
}

func _ListenerService_AppEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(AppEventData)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ListenerServiceServer).AppEvent(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ListenerService_CustomEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListenerServiceServer).CustomEvent(&listenerServiceCustomEventServer{stream})
}

type ListenerService_CustomEventServer interface {
	SendAndClose(*Response) error
	Recv() (*Event, error)
	grpc.ServerStream
}

type listenerServiceCustomEventServer struct {
	grpc.ServerStream
}

func (x *listenerServiceCustomEventServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *listenerServiceCustomEventServer) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ListenerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apipb.ListenerService",
	HandlerType: (*ListenerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppEvent",
			Handler:    _ListenerService_AppEvent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CustomEvent",
			Handler:       _ListenerService_CustomEvent_Handler,
			ClientStreams: true,
		},
	},
}

func (m *DeviceInfo) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DeviceInfo) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.VendorId) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintListener(data, i, uint64(len(m.VendorId)))
		i += copy(data[i:], m.VendorId)
	}
	if len(m.ClientSdk) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintListener(data, i, uint64(len(m.ClientSdk)))
		i += copy(data[i:], m.ClientSdk)
	}
	if len(m.BundleIdentifier) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintListener(data, i, uint64(len(m.BundleIdentifier)))
		i += copy(data[i:], m.BundleIdentifier)
	}
	if len(m.BundleVersion) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintListener(data, i, uint64(len(m.BundleVersion)))
		i += copy(data[i:], m.BundleVersion)
	}
	if len(m.BundleShortVersion) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintListener(data, i, uint64(len(m.BundleShortVersion)))
		i += copy(data[i:], m.BundleShortVersion)
	}
	if len(m.DeviceType) > 0 {
		data[i] = 0x32
		i++
		i = encodeVarintListener(data, i, uint64(len(m.DeviceType)))
		i += copy(data[i:], m.DeviceType)
	}
	if len(m.DeviceName) > 0 {
		data[i] = 0x3a
		i++
		i = encodeVarintListener(data, i, uint64(len(m.DeviceName)))
		i += copy(data[i:], m.DeviceName)
	}
	if len(m.OsName) > 0 {
		data[i] = 0x42
		i++
		i = encodeVarintListener(data, i, uint64(len(m.OsName)))
		i += copy(data[i:], m.OsName)
	}
	if len(m.SystemVersion) > 0 {
		data[i] = 0x4a
		i++
		i = encodeVarintListener(data, i, uint64(len(m.SystemVersion)))
		i += copy(data[i:], m.SystemVersion)
	}
	if len(m.LanguageCode) > 0 {
		data[i] = 0x52
		i++
		i = encodeVarintListener(data, i, uint64(len(m.LanguageCode)))
		i += copy(data[i:], m.LanguageCode)
	}
	if len(m.CountryCode) > 0 {
		data[i] = 0x5a
		i++
		i = encodeVarintListener(data, i, uint64(len(m.CountryCode)))
		i += copy(data[i:], m.CountryCode)
	}
	return i, nil
}

func (m *Event) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Event) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UserId) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintListener(data, i, uint64(len(m.UserId)))
		i += copy(data[i:], m.UserId)
	}
	if len(m.ChildId) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintListener(data, i, uint64(len(m.ChildId)))
		i += copy(data[i:], m.ChildId)
	}
	if len(m.Event) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintListener(data, i, uint64(len(m.Event)))
		i += copy(data[i:], m.Event)
	}
	if m.Timestamp != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintListener(data, i, uint64(m.Timestamp))
	}
	if len(m.SubId) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintListener(data, i, uint64(len(m.SubId)))
		i += copy(data[i:], m.SubId)
	}
	if len(m.DeviceId) > 0 {
		data[i] = 0x32
		i++
		i = encodeVarintListener(data, i, uint64(len(m.DeviceId)))
		i += copy(data[i:], m.DeviceId)
	}
	if len(m.AppId) > 0 {
		data[i] = 0x3a
		i++
		i = encodeVarintListener(data, i, uint64(len(m.AppId)))
		i += copy(data[i:], m.AppId)
	}
	if m.Payload != nil {
		if len(m.Payload) > 0 {
			data[i] = 0x5a
			i++
			i = encodeVarintListener(data, i, uint64(len(m.Payload)))
			i += copy(data[i:], m.Payload)
		}
	}
	return i, nil
}

func (m *AppEventData) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *AppEventData) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Event) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintListener(data, i, uint64(len(m.Event)))
		i += copy(data[i:], m.Event)
	}
	if len(m.AppId) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintListener(data, i, uint64(len(m.AppId)))
		i += copy(data[i:], m.AppId)
	}
	if m.Device != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintListener(data, i, uint64(m.Device.Size()))
		n1, err := m.Device.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Timestamp != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintListener(data, i, uint64(m.Timestamp))
	}
	if m.Payload != nil {
		if len(m.Payload) > 0 {
			data[i] = 0x2a
			i++
			i = encodeVarintListener(data, i, uint64(len(m.Payload)))
			i += copy(data[i:], m.Payload)
		}
	}
	return i, nil
}

func encodeFixed64Listener(data []byte, offset int, v uint64) int {
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
func encodeFixed32Listener(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintListener(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *DeviceInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.VendorId)
	if l > 0 {
		n += 1 + l + sovListener(uint64(l))
	}
	l = len(m.ClientSdk)
	if l > 0 {
		n += 1 + l + sovListener(uint64(l))
	}
	l = len(m.BundleIdentifier)
	if l > 0 {
		n += 1 + l + sovListener(uint64(l))
	}
	l = len(m.BundleVersion)
	if l > 0 {
		n += 1 + l + sovListener(uint64(l))
	}
	l = len(m.BundleShortVersion)
	if l > 0 {
		n += 1 + l + sovListener(uint64(l))
	}
	l = len(m.DeviceType)
	if l > 0 {
		n += 1 + l + sovListener(uint64(l))
	}
	l = len(m.DeviceName)
	if l > 0 {
		n += 1 + l + sovListener(uint64(l))
	}
	l = len(m.OsName)
	if l > 0 {
		n += 1 + l + sovListener(uint64(l))
	}
	l = len(m.SystemVersion)
	if l > 0 {
		n += 1 + l + sovListener(uint64(l))
	}
	l = len(m.LanguageCode)
	if l > 0 {
		n += 1 + l + sovListener(uint64(l))
	}
	l = len(m.CountryCode)
	if l > 0 {
		n += 1 + l + sovListener(uint64(l))
	}
	return n
}

func (m *Event) Size() (n int) {
	var l int
	_ = l
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovListener(uint64(l))
	}
	l = len(m.ChildId)
	if l > 0 {
		n += 1 + l + sovListener(uint64(l))
	}
	l = len(m.Event)
	if l > 0 {
		n += 1 + l + sovListener(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovListener(uint64(m.Timestamp))
	}
	l = len(m.SubId)
	if l > 0 {
		n += 1 + l + sovListener(uint64(l))
	}
	l = len(m.DeviceId)
	if l > 0 {
		n += 1 + l + sovListener(uint64(l))
	}
	l = len(m.AppId)
	if l > 0 {
		n += 1 + l + sovListener(uint64(l))
	}
	if m.Payload != nil {
		l = len(m.Payload)
		if l > 0 {
			n += 1 + l + sovListener(uint64(l))
		}
	}
	return n
}

func (m *AppEventData) Size() (n int) {
	var l int
	_ = l
	l = len(m.Event)
	if l > 0 {
		n += 1 + l + sovListener(uint64(l))
	}
	l = len(m.AppId)
	if l > 0 {
		n += 1 + l + sovListener(uint64(l))
	}
	if m.Device != nil {
		l = m.Device.Size()
		n += 1 + l + sovListener(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovListener(uint64(m.Timestamp))
	}
	if m.Payload != nil {
		l = len(m.Payload)
		if l > 0 {
			n += 1 + l + sovListener(uint64(l))
		}
	}
	return n
}

func sovListener(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozListener(x uint64) (n int) {
	return sovListener(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeviceInfo) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListener
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
			return fmt.Errorf("proto: DeviceInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VendorId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientSdk", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientSdk = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BundleIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BundleIdentifier = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BundleVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BundleVersion = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BundleShortVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BundleShortVersion = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceType = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceName = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OsName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OsName = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SystemVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SystemVersion = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LanguageCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LanguageCode = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CountryCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CountryCode = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListener(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthListener
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
func (m *Event) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListener
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
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChildId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChildId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Event = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListener(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthListener
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
func (m *AppEventData) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListener
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
			return fmt.Errorf("proto: AppEventData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppEventData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Event = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Device", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Device == nil {
				m.Device = &DeviceInfo{}
			}
			if err := m.Device.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListener
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
				return ErrInvalidLengthListener
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListener(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthListener
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
func skipListener(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowListener
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
					return 0, ErrIntOverflowListener
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
					return 0, ErrIntOverflowListener
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
				return 0, ErrInvalidLengthListener
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowListener
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
				next, err := skipListener(data[start:])
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
	ErrInvalidLengthListener = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowListener   = fmt.Errorf("proto: integer overflow")
)
