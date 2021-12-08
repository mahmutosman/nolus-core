// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: suspend/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QuerySuspendRequest struct {
}

func (m *QuerySuspendRequest) Reset()         { *m = QuerySuspendRequest{} }
func (m *QuerySuspendRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySuspendRequest) ProtoMessage()    {}
func (*QuerySuspendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b2d8646a9f10f23, []int{0}
}
func (m *QuerySuspendRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySuspendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySuspendRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySuspendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySuspendRequest.Merge(m, src)
}
func (m *QuerySuspendRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySuspendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySuspendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySuspendRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QuerySuspendResponse struct {
	// params defines the parameters of the module.
	QuerySuspend *QuerySuspend `protobuf:"bytes,1,opt,name=querySuspend,proto3" json:"querySuspend,omitempty"`
}

func (m *QuerySuspendResponse) Reset()         { *m = QuerySuspendResponse{} }
func (m *QuerySuspendResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySuspendResponse) ProtoMessage()    {}
func (*QuerySuspendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b2d8646a9f10f23, []int{1}
}
func (m *QuerySuspendResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySuspendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySuspendResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySuspendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySuspendResponse.Merge(m, src)
}
func (m *QuerySuspendResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySuspendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySuspendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySuspendResponse proto.InternalMessageInfo

func (m *QuerySuspendResponse) GetQuerySuspend() *QuerySuspend {
	if m != nil {
		return m.QuerySuspend
	}
	return nil
}

type QuerySuspend struct {
	Creator  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Suspend  bool   `protobuf:"varint,2,opt,name=suspend,proto3" json:"suspend,omitempty"`
	AdminKey string `protobuf:"bytes,3,opt,name=admin_key,json=adminKey,proto3" json:"admin_key,omitempty"`
}

func (m *QuerySuspend) Reset()         { *m = QuerySuspend{} }
func (m *QuerySuspend) String() string { return proto.CompactTextString(m) }
func (*QuerySuspend) ProtoMessage()    {}
func (*QuerySuspend) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b2d8646a9f10f23, []int{2}
}
func (m *QuerySuspend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySuspend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySuspend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySuspend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySuspend.Merge(m, src)
}
func (m *QuerySuspend) XXX_Size() int {
	return m.Size()
}
func (m *QuerySuspend) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySuspend.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySuspend proto.InternalMessageInfo

func (m *QuerySuspend) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *QuerySuspend) GetSuspend() bool {
	if m != nil {
		return m.Suspend
	}
	return false
}

func (m *QuerySuspend) GetAdminKey() string {
	if m != nil {
		return m.AdminKey
	}
	return ""
}

func init() {
	proto.RegisterType((*QuerySuspendRequest)(nil), "nomo.cosmzone.suspend.QuerySuspendRequest")
	proto.RegisterType((*QuerySuspendResponse)(nil), "nomo.cosmzone.suspend.QuerySuspendResponse")
	proto.RegisterType((*QuerySuspend)(nil), "nomo.cosmzone.suspend.QuerySuspend")
}

func init() { proto.RegisterFile("suspend/query.proto", fileDescriptor_5b2d8646a9f10f23) }

var fileDescriptor_5b2d8646a9f10f23 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x2e, 0x2d, 0x2e,
	0x48, 0xcd, 0x4b, 0xd1, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0xcd, 0xcb, 0xcf, 0xcd, 0xd7, 0x4b, 0xce, 0x2f, 0xce, 0xad, 0xca, 0xcf, 0x4b, 0xd5, 0x83,
	0x2a, 0x51, 0x12, 0xe5, 0x12, 0x0e, 0x04, 0xa9, 0x0a, 0x86, 0xf0, 0x83, 0x52, 0x0b, 0x4b, 0x53,
	0x8b, 0x4b, 0x94, 0xe2, 0xb9, 0x44, 0x50, 0x85, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xdc,
	0xb9, 0x78, 0x0a, 0x91, 0xc4, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x94, 0xf5, 0xb0, 0x1a,
	0xae, 0x87, 0x62, 0x04, 0x8a, 0x46, 0xa5, 0x78, 0x2e, 0x1e, 0x64, 0x59, 0x21, 0x09, 0x2e, 0xf6,
	0xe4, 0xa2, 0xd4, 0xc4, 0x92, 0xfc, 0x22, 0xb0, 0x99, 0x9c, 0x41, 0x30, 0x2e, 0x48, 0x06, 0x6a,
	0x9e, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x8c, 0x2b, 0x24, 0xcd, 0xc5, 0x99, 0x98, 0x92,
	0x9b, 0x99, 0x17, 0x9f, 0x9d, 0x5a, 0x29, 0xc1, 0x0c, 0xd6, 0xc5, 0x01, 0x16, 0xf0, 0x4e, 0xad,
	0x34, 0x62, 0xe7, 0x62, 0x05, 0x5b, 0xe0, 0x14, 0x74, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72,
	0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7,
	0x72, 0x0c, 0x51, 0x16, 0xe9, 0x99, 0x25, 0x39, 0x89, 0x49, 0xba, 0x10, 0xc7, 0x17, 0xa5, 0xa6,
	0x64, 0x16, 0x17, 0x67, 0xe6, 0xe6, 0xeb, 0xe5, 0xa5, 0x96, 0xe8, 0x83, 0xc4, 0xf4, 0x61, 0x1e,
	0xd2, 0xaf, 0xd0, 0x87, 0x05, 0x69, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0x4c, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x81, 0x2e, 0x14, 0xdc, 0x6a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	Suspend(ctx context.Context, in *QuerySuspendRequest, opts ...grpc.CallOption) (*QuerySuspendResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Suspend(ctx context.Context, in *QuerySuspendRequest, opts ...grpc.CallOption) (*QuerySuspendResponse, error) {
	out := new(QuerySuspendResponse)
	err := c.cc.Invoke(ctx, "/nomo.cosmzone.suspend.Query/Suspend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Suspend(context.Context, *QuerySuspendRequest) (*QuerySuspendResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nomo.cosmzone.suspend.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Suspend",
			Handler:    _Query_Suspend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "suspend/query.proto",
}

func _Query_Suspend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySuspendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Suspend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nomo.cosmzone.suspend.Query/Suspend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Suspend(ctx, req.(*QuerySuspendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func (m *QuerySuspendRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySuspendRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySuspendRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QuerySuspendResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySuspendResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySuspendResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.QuerySuspend != nil {
		{
			size, err := m.QuerySuspend.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySuspend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySuspend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySuspend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AdminKey) > 0 {
		i -= len(m.AdminKey)
		copy(dAtA[i:], m.AdminKey)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.AdminKey)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Suspend {
		i--
		if m.Suspend {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QuerySuspendRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QuerySuspendResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.QuerySuspend != nil {
		l = m.QuerySuspend.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QuerySuspend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Suspend {
		n += 2
	}
	l = len(m.AdminKey)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QuerySuspendRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QuerySuspendRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySuspendRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QuerySuspendResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QuerySuspendResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySuspendResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuerySuspend", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.QuerySuspend == nil {
				m.QuerySuspend = &QuerySuspend{}
			}
			if err := m.QuerySuspend.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QuerySuspend) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QuerySuspend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySuspend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Suspend", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Suspend = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdminKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
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
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
