// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gdex/liquidity/v1beta1/liquidity.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// SwapDirection enumerates swap directions.
type SwapDirection int32

const (
	// SWAP_DIRECTION_UNSPECIFIED specifies unknown swap direction.
	SwapDirectionUnspecified SwapDirection = 0
	// SWAP_DIRECTION_X_TO_Y specifies x to y swap direction.
	SwapDirectionXToY SwapDirection = 1
	// SWAP_DIRECTION_Y_TO_X specifies y to x swap direction.
	SwapDirectionYToX SwapDirection = 2
)

var SwapDirection_name = map[int32]string{
	0: "SWAP_DIRECTION_UNSPECIFIED",
	1: "SWAP_DIRECTION_X_TO_Y",
	2: "SWAP_DIRECTION_Y_TO_X",
}

var SwapDirection_value = map[string]int32{
	"SWAP_DIRECTION_UNSPECIFIED": 0,
	"SWAP_DIRECTION_X_TO_Y":      1,
	"SWAP_DIRECTION_Y_TO_X":      2,
}

func (x SwapDirection) String() string {
	return proto.EnumName(SwapDirection_name, int32(x))
}

func (SwapDirection) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_91aa8d5b251b52ba, []int{0}
}

// Params defines the parameters for the liquidity module.
type Params struct {
	InitialPoolCoinSupply github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=initial_pool_coin_supply,json=initialPoolCoinSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"initial_pool_coin_supply"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_91aa8d5b251b52ba, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

// Pool defines a liquidity pool.
type Pool struct {
	Id                uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ReserveCoinDenoms []string `protobuf:"bytes,2,rep,name=reserve_coin_denoms,json=reserveCoinDenoms,proto3" json:"reserve_coin_denoms,omitempty"`
	ReserveAddress    string   `protobuf:"bytes,3,opt,name=reserve_address,json=reserveAddress,proto3" json:"reserve_address,omitempty"`
	PoolCoinDenom     string   `protobuf:"bytes,4,opt,name=pool_coin_denom,json=poolCoinDenom,proto3" json:"pool_coin_denom,omitempty"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_91aa8d5b251b52ba, []int{1}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

// SwapRequest defines a swap request.
type SwapRequest struct {
	Requester       string                                 `protobuf:"bytes,1,opt,name=requester,proto3" json:"requester,omitempty"`
	Direction       SwapDirection                          `protobuf:"varint,2,opt,name=direction,proto3,enum=gdex.liquidity.v1beta1.SwapDirection" json:"direction,omitempty"`
	Price           github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price"`
	RemainingAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=remaining_amount,json=remainingAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"remaining_amount"`
}

func (m *SwapRequest) Reset()         { *m = SwapRequest{} }
func (m *SwapRequest) String() string { return proto.CompactTextString(m) }
func (*SwapRequest) ProtoMessage()    {}
func (*SwapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_91aa8d5b251b52ba, []int{2}
}
func (m *SwapRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SwapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SwapRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SwapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwapRequest.Merge(m, src)
}
func (m *SwapRequest) XXX_Size() int {
	return m.Size()
}
func (m *SwapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SwapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SwapRequest proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("gdex.liquidity.v1beta1.SwapDirection", SwapDirection_name, SwapDirection_value)
	proto.RegisterType((*Params)(nil), "gdex.liquidity.v1beta1.Params")
	proto.RegisterType((*Pool)(nil), "gdex.liquidity.v1beta1.Pool")
	proto.RegisterType((*SwapRequest)(nil), "gdex.liquidity.v1beta1.SwapRequest")
}

func init() {
	proto.RegisterFile("gdex/liquidity/v1beta1/liquidity.proto", fileDescriptor_91aa8d5b251b52ba)
}

var fileDescriptor_91aa8d5b251b52ba = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4d, 0x6f, 0xd3, 0x30,
	0x1c, 0xc6, 0x93, 0xac, 0x4c, 0xaa, 0xd1, 0xd6, 0xce, 0x50, 0x14, 0x55, 0x53, 0x56, 0x55, 0xa2,
	0x54, 0x48, 0x24, 0x0c, 0xae, 0x5c, 0xba, 0xa6, 0x48, 0x3d, 0xb0, 0x55, 0x69, 0x27, 0x5a, 0x2e,
	0x51, 0x1a, 0x7b, 0xc1, 0xa2, 0xb1, 0x53, 0xdb, 0x1d, 0xf4, 0x1b, 0xa0, 0x9d, 0xe0, 0xc4, 0x69,
	0x27, 0x3e, 0x04, 0x5f, 0xa1, 0xc7, 0x1d, 0x11, 0x87, 0x09, 0xda, 0x2f, 0x82, 0xf2, 0x32, 0xb2,
	0xa2, 0x5e, 0x76, 0xb2, 0xf5, 0xf8, 0x79, 0x7e, 0xf9, 0xfb, 0x89, 0x0c, 0x1a, 0x01, 0xc2, 0x9f,
	0xac, 0x09, 0x99, 0xce, 0x08, 0x22, 0x72, 0x6e, 0x9d, 0x1f, 0x8e, 0xb1, 0xf4, 0x0e, 0x73, 0xc5,
	0x8c, 0x38, 0x93, 0x0c, 0x3e, 0x8a, 0x7d, 0x66, 0xae, 0x66, 0xbe, 0xea, 0xc3, 0x80, 0x05, 0x2c,
	0xb1, 0x58, 0xf1, 0x2e, 0x75, 0xd7, 0xa7, 0x60, 0xbb, 0xe7, 0x71, 0x2f, 0x14, 0x30, 0x00, 0x3a,
	0xa1, 0x44, 0x12, 0x6f, 0xe2, 0x46, 0x8c, 0x4d, 0x5c, 0x9f, 0x11, 0xea, 0x8a, 0x59, 0x14, 0x4d,
	0xe6, 0xba, 0x5a, 0x53, 0x9b, 0xc5, 0x23, 0x73, 0x71, 0x7d, 0xa0, 0xfc, 0xba, 0x3e, 0x68, 0x04,
	0x44, 0xbe, 0x9f, 0x8d, 0x4d, 0x9f, 0x85, 0x96, 0xcf, 0x44, 0xc8, 0x44, 0xb6, 0x3c, 0x13, 0xe8,
	0x83, 0x25, 0xe7, 0x11, 0x16, 0x66, 0x97, 0x4a, 0xa7, 0x92, 0xf1, 0x7a, 0x8c, 0x4d, 0xda, 0x8c,
	0xd0, 0x7e, 0x02, 0xab, 0x7f, 0x53, 0x41, 0x21, 0x96, 0xe0, 0x2e, 0xd0, 0x08, 0x4a, 0xd8, 0x05,
	0x47, 0x23, 0x08, 0x9a, 0xe0, 0x01, 0xc7, 0x02, 0xf3, 0x73, 0x9c, 0x7e, 0x1c, 0x61, 0xca, 0x42,
	0xa1, 0x6b, 0xb5, 0xad, 0x66, 0xd1, 0xd9, 0xcb, 0x8e, 0x62, 0x90, 0x9d, 0x1c, 0xc0, 0x27, 0xa0,
	0x74, 0xe3, 0xf7, 0x10, 0xe2, 0x58, 0x08, 0x7d, 0x2b, 0x1e, 0xd4, 0xd9, 0xcd, 0xe4, 0x56, 0xaa,
	0xc2, 0x06, 0x28, 0xe5, 0x57, 0x4a, 0xa8, 0x7a, 0x21, 0x31, 0xee, 0x44, 0xd9, 0x68, 0x09, 0xb1,
	0xfe, 0x55, 0x03, 0xf7, 0xfb, 0x1f, 0xbd, 0xc8, 0xc1, 0xd3, 0x19, 0x16, 0x12, 0xee, 0x83, 0x22,
	0x4f, 0xb7, 0x98, 0xa7, 0x1d, 0x38, 0xb9, 0x00, 0xdb, 0xa0, 0x88, 0x08, 0xc7, 0xbe, 0x24, 0x8c,
	0xea, 0x5a, 0x4d, 0x6d, 0xee, 0xbe, 0x78, 0x6c, 0x6e, 0x2e, 0xdf, 0x8c, 0xa9, 0xf6, 0x8d, 0xd9,
	0xc9, 0x73, 0xd0, 0x06, 0xf7, 0x22, 0x4e, 0x7c, 0x9c, 0x4e, 0x7e, 0xa7, 0x8a, 0x6d, 0xec, 0x3b,
	0x69, 0x18, 0x8e, 0x40, 0x99, 0xe3, 0xd0, 0x23, 0x94, 0xd0, 0xc0, 0xf5, 0x42, 0x36, 0xa3, 0x32,
	0xbd, 0xe1, 0x9d, 0xff, 0x59, 0xe9, 0x1f, 0xa7, 0x95, 0x60, 0x9e, 0xfe, 0x50, 0xc1, 0xce, 0xda,
	0xf4, 0xf0, 0x15, 0xa8, 0xf6, 0xdf, 0xb6, 0x7a, 0xae, 0xdd, 0x75, 0x3a, 0xed, 0x41, 0xf7, 0xe4,
	0xd8, 0x3d, 0x3d, 0xee, 0xf7, 0x3a, 0xed, 0xee, 0xeb, 0x6e, 0xc7, 0x2e, 0x2b, 0xd5, 0xfd, 0x8b,
	0xcb, 0x9a, 0xbe, 0x16, 0x39, 0xa5, 0x22, 0xc2, 0x3e, 0x39, 0x23, 0x18, 0xc1, 0xe7, 0xa0, 0xf2,
	0x5f, 0x7a, 0xe8, 0x0e, 0x4e, 0xdc, 0x51, 0x59, 0xad, 0x56, 0x2e, 0x2e, 0x6b, 0x7b, 0x6b, 0xc1,
	0xe1, 0x80, 0x8d, 0x36, 0x24, 0x46, 0x71, 0x62, 0x58, 0xd6, 0x36, 0x24, 0x46, 0x03, 0x36, 0xac,
	0x16, 0x3e, 0x7f, 0x37, 0x94, 0xa3, 0x37, 0x8b, 0x3f, 0x86, 0xb2, 0x58, 0x1a, 0xea, 0xd5, 0xd2,
	0x50, 0x7f, 0x2f, 0x0d, 0xf5, 0xcb, 0xca, 0x50, 0xae, 0x56, 0x86, 0xf2, 0x73, 0x65, 0x28, 0xef,
	0xac, 0x5b, 0x85, 0x48, 0x4c, 0x11, 0xe6, 0x21, 0xa1, 0xd2, 0x3a, 0xf3, 0xe2, 0x35, 0xb0, 0x6e,
	0xbf, 0xb4, 0xa4, 0x9d, 0xf1, 0x76, 0xf2, 0x60, 0x5e, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xed,
	0x34, 0xf4, 0x3e, 0x88, 0x03, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.InitialPoolCoinSupply.Size()
		i -= size
		if _, err := m.InitialPoolCoinSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLiquidity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PoolCoinDenom) > 0 {
		i -= len(m.PoolCoinDenom)
		copy(dAtA[i:], m.PoolCoinDenom)
		i = encodeVarintLiquidity(dAtA, i, uint64(len(m.PoolCoinDenom)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ReserveAddress) > 0 {
		i -= len(m.ReserveAddress)
		copy(dAtA[i:], m.ReserveAddress)
		i = encodeVarintLiquidity(dAtA, i, uint64(len(m.ReserveAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ReserveCoinDenoms) > 0 {
		for iNdEx := len(m.ReserveCoinDenoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ReserveCoinDenoms[iNdEx])
			copy(dAtA[i:], m.ReserveCoinDenoms[iNdEx])
			i = encodeVarintLiquidity(dAtA, i, uint64(len(m.ReserveCoinDenoms[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Id != 0 {
		i = encodeVarintLiquidity(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SwapRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SwapRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SwapRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.RemainingAmount.Size()
		i -= size
		if _, err := m.RemainingAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLiquidity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLiquidity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Direction != 0 {
		i = encodeVarintLiquidity(dAtA, i, uint64(m.Direction))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Requester) > 0 {
		i -= len(m.Requester)
		copy(dAtA[i:], m.Requester)
		i = encodeVarintLiquidity(dAtA, i, uint64(len(m.Requester)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLiquidity(dAtA []byte, offset int, v uint64) int {
	offset -= sovLiquidity(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.InitialPoolCoinSupply.Size()
	n += 1 + l + sovLiquidity(uint64(l))
	return n
}

func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovLiquidity(uint64(m.Id))
	}
	if len(m.ReserveCoinDenoms) > 0 {
		for _, s := range m.ReserveCoinDenoms {
			l = len(s)
			n += 1 + l + sovLiquidity(uint64(l))
		}
	}
	l = len(m.ReserveAddress)
	if l > 0 {
		n += 1 + l + sovLiquidity(uint64(l))
	}
	l = len(m.PoolCoinDenom)
	if l > 0 {
		n += 1 + l + sovLiquidity(uint64(l))
	}
	return n
}

func (m *SwapRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Requester)
	if l > 0 {
		n += 1 + l + sovLiquidity(uint64(l))
	}
	if m.Direction != 0 {
		n += 1 + sovLiquidity(uint64(m.Direction))
	}
	l = m.Price.Size()
	n += 1 + l + sovLiquidity(uint64(l))
	l = m.RemainingAmount.Size()
	n += 1 + l + sovLiquidity(uint64(l))
	return n
}

func sovLiquidity(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLiquidity(x uint64) (n int) {
	return sovLiquidity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidity
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialPoolCoinSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
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
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitialPoolCoinSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidity
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
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidity
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReserveCoinDenoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
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
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReserveCoinDenoms = append(m.ReserveCoinDenoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReserveAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
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
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReserveAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolCoinDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
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
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolCoinDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidity
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
func (m *SwapRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidity
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
			return fmt.Errorf("proto: SwapRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SwapRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requester", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
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
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requester = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Direction", wireType)
			}
			m.Direction = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Direction |= SwapDirection(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
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
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainingAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
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
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RemainingAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidity
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
func skipLiquidity(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLiquidity
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
					return 0, ErrIntOverflowLiquidity
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
					return 0, ErrIntOverflowLiquidity
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
				return 0, ErrInvalidLengthLiquidity
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLiquidity
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLiquidity
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLiquidity        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLiquidity          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLiquidity = fmt.Errorf("proto: unexpected end of group")
)