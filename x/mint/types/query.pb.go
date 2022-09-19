// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: squad/mint/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8bfd93def1f9142, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params defines the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8bfd93def1f9142, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QueryLastBlockTimeRequest is the request type for the Query/LastBlockTime RPC method.
type QueryLastBlockTimeRequest struct {
}

func (m *QueryLastBlockTimeRequest) Reset()         { *m = QueryLastBlockTimeRequest{} }
func (m *QueryLastBlockTimeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryLastBlockTimeRequest) ProtoMessage()    {}
func (*QueryLastBlockTimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8bfd93def1f9142, []int{2}
}
func (m *QueryLastBlockTimeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLastBlockTimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLastBlockTimeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLastBlockTimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLastBlockTimeRequest.Merge(m, src)
}
func (m *QueryLastBlockTimeRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryLastBlockTimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLastBlockTimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLastBlockTimeRequest proto.InternalMessageInfo

// QueryLastBlockTimeResponse is the response type for the Query/LastBlockTime RPC method.
type QueryLastBlockTimeResponse struct {
	LastBlockTime *time.Time `protobuf:"bytes,1,opt,name=last_block_time,json=lastBlockTime,proto3,stdtime" json:"last_block_time,omitempty" yaml:"last_block_time"`
}

func (m *QueryLastBlockTimeResponse) Reset()         { *m = QueryLastBlockTimeResponse{} }
func (m *QueryLastBlockTimeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryLastBlockTimeResponse) ProtoMessage()    {}
func (*QueryLastBlockTimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8bfd93def1f9142, []int{3}
}
func (m *QueryLastBlockTimeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLastBlockTimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLastBlockTimeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLastBlockTimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLastBlockTimeResponse.Merge(m, src)
}
func (m *QueryLastBlockTimeResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryLastBlockTimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLastBlockTimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLastBlockTimeResponse proto.InternalMessageInfo

func (m *QueryLastBlockTimeResponse) GetLastBlockTime() *time.Time {
	if m != nil {
		return m.LastBlockTime
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "squad.mint.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "squad.mint.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QueryLastBlockTimeRequest)(nil), "squad.mint.v1beta1.QueryLastBlockTimeRequest")
	proto.RegisterType((*QueryLastBlockTimeResponse)(nil), "squad.mint.v1beta1.QueryLastBlockTimeResponse")
}

func init() { proto.RegisterFile("squad/mint/v1beta1/query.proto", fileDescriptor_a8bfd93def1f9142) }

var fileDescriptor_a8bfd93def1f9142 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcd, 0x6a, 0xdb, 0x30,
	0x00, 0xb6, 0xc2, 0x96, 0x83, 0x46, 0x18, 0x68, 0x61, 0x6c, 0x5e, 0x26, 0x0f, 0x8f, 0xfd, 0xc0,
	0x88, 0x44, 0xb2, 0xcb, 0xd8, 0xd1, 0x97, 0x5d, 0x06, 0xdb, 0x42, 0x4e, 0xbb, 0x04, 0x39, 0xd3,
	0x3c, 0x33, 0xcb, 0x72, 0x22, 0x79, 0x2c, 0x87, 0x42, 0xe9, 0x13, 0x04, 0x7a, 0xec, 0xb5, 0x0f,
	0x93, 0x63, 0xa0, 0x97, 0x9e, 0xd2, 0x92, 0xf4, 0x09, 0xfa, 0x04, 0xc5, 0xb2, 0x52, 0x9a, 0xc4,
	0xa5, 0xbd, 0xd9, 0xfa, 0x3e, 0x7d, 0x7f, 0x82, 0x58, 0x8d, 0x72, 0xf6, 0x8b, 0x8a, 0x38, 0xd5,
	0xf4, 0x5f, 0x27, 0xe4, 0x9a, 0x75, 0xe8, 0x28, 0xe7, 0xe3, 0x09, 0xc9, 0xc6, 0x52, 0x4b, 0x84,
	0x0c, 0x4e, 0x0a, 0x9c, 0x58, 0xdc, 0x6d, 0x46, 0x32, 0x92, 0x06, 0xa6, 0xc5, 0x57, 0xc9, 0x74,
	0x5b, 0x91, 0x94, 0x51, 0xc2, 0x29, 0xcb, 0x62, 0xca, 0xd2, 0x54, 0x6a, 0xa6, 0x63, 0x99, 0x2a,
	0x8b, 0x7a, 0x16, 0x35, 0x7f, 0x61, 0xfe, 0x9b, 0xea, 0x58, 0x70, 0xa5, 0x99, 0xc8, 0x2c, 0xe1,
	0x65, 0x45, 0x10, 0xe3, 0x6a, 0x60, 0xbf, 0x09, 0xd1, 0x8f, 0x22, 0xd6, 0x77, 0x36, 0x66, 0x42,
	0xf5, 0xf8, 0x28, 0xe7, 0x4a, 0xfb, 0xdf, 0xe0, 0x93, 0x8d, 0x53, 0x95, 0xc9, 0x54, 0x71, 0xf4,
	0x09, 0xd6, 0x33, 0x73, 0xf2, 0x0c, 0xbc, 0x02, 0xef, 0x1f, 0x75, 0x5d, 0xb2, 0xdb, 0x82, 0x94,
	0x77, 0x82, 0x07, 0xb3, 0x85, 0xe7, 0xf4, 0x2c, 0xdf, 0x7f, 0x01, 0x9f, 0x1b, 0xc1, 0xaf, 0x4c,
	0xe9, 0x20, 0x91, 0xc3, 0xbf, 0xfd, 0x58, 0xf0, 0xb5, 0xdb, 0x3e, 0x80, 0x6e, 0x15, 0x6a, 0x5d,
	0x43, 0xf8, 0x38, 0x61, 0x4a, 0x0f, 0xc2, 0x02, 0x19, 0x14, 0xfd, 0xae, 0xed, 0xcb, 0xf2, 0x64,
	0x5d, 0x9e, 0xf4, 0xd7, 0xe5, 0x03, 0x7c, 0xb9, 0xf0, 0x9e, 0x4e, 0x98, 0x48, 0x3e, 0xfb, 0x5b,
	0x97, 0xfd, 0xe9, 0x99, 0x07, 0x7a, 0x8d, 0xe4, 0xa6, 0x57, 0xf7, 0xb8, 0x06, 0x1f, 0x9a, 0x08,
	0x68, 0x0f, 0xd6, 0xcb, 0x06, 0xe8, 0x6d, 0x55, 0xbb, 0xdd, 0xb1, 0xdc, 0x77, 0x77, 0xf2, 0xca,
	0x22, 0xbe, 0x7f, 0x70, 0x72, 0x71, 0x58, 0x6b, 0x21, 0x97, 0x56, 0xbc, 0x49, 0x39, 0x14, 0x3a,
	0x02, 0xb0, 0xb1, 0x31, 0x03, 0x6a, 0xdf, 0x2a, 0x5f, 0x35, 0xa6, 0x4b, 0xee, 0x4b, 0xb7, 0xa1,
	0x3e, 0x98, 0x50, 0x6f, 0xd0, 0xeb, 0xaa, 0x50, 0x5b, 0xd3, 0x05, 0x5f, 0x66, 0x4b, 0x0c, 0xe6,
	0x4b, 0x0c, 0xce, 0x97, 0x18, 0x4c, 0x57, 0xd8, 0x99, 0xaf, 0xb0, 0x73, 0xba, 0xc2, 0xce, 0xcf,
	0x76, 0x14, 0xeb, 0x3f, 0x79, 0x48, 0x86, 0x52, 0xd0, 0xa1, 0x54, 0x42, 0x1a, 0xb5, 0x76, 0xc2,
	0x42, 0x65, 0x85, 0xff, 0x97, 0xd2, 0x7a, 0x92, 0x71, 0x15, 0xd6, 0xcd, 0x93, 0x7d, 0xbc, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0x60, 0x2d, 0xa8, 0xb2, 0x27, 0x03, 0x00, 0x00,
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
	// Params returns the total set of minting parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// LastBlockTime returns the last block time.
	LastBlockTime(ctx context.Context, in *QueryLastBlockTimeRequest, opts ...grpc.CallOption) (*QueryLastBlockTimeResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/squad.mint.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) LastBlockTime(ctx context.Context, in *QueryLastBlockTimeRequest, opts ...grpc.CallOption) (*QueryLastBlockTimeResponse, error) {
	out := new(QueryLastBlockTimeResponse)
	err := c.cc.Invoke(ctx, "/squad.mint.v1beta1.Query/LastBlockTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params returns the total set of minting parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// LastBlockTime returns the last block time.
	LastBlockTime(context.Context, *QueryLastBlockTimeRequest) (*QueryLastBlockTimeResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) LastBlockTime(ctx context.Context, req *QueryLastBlockTimeRequest) (*QueryLastBlockTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LastBlockTime not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/squad.mint.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_LastBlockTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLastBlockTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LastBlockTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/squad.mint.v1beta1.Query/LastBlockTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LastBlockTime(ctx, req.(*QueryLastBlockTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "squad.mint.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "LastBlockTime",
			Handler:    _Query_LastBlockTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "squad/mint/v1beta1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryLastBlockTimeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLastBlockTimeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLastBlockTimeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryLastBlockTimeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLastBlockTimeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLastBlockTimeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastBlockTime != nil {
		n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.LastBlockTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.LastBlockTime):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintQuery(dAtA, i, uint64(n2))
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
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryLastBlockTimeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryLastBlockTimeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LastBlockTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.LastBlockTime)
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
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryLastBlockTimeRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryLastBlockTimeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLastBlockTimeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryLastBlockTimeResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryLastBlockTimeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLastBlockTimeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockTime", wireType)
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
			if m.LastBlockTime == nil {
				m.LastBlockTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.LastBlockTime, dAtA[iNdEx:postIndex]); err != nil {
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
