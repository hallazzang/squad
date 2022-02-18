// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: squad/claim/v1beta1/claim.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

// ConditionType defines the type of condition that a recipient must execute in order to receive a claimable amount.
type ConditionType int32

const (
	// CONDITION_TYPE_UNSPECIFIED specifies an unknown condition type
	ConditionTypeUnspecified ConditionType = 0
	// CONDITION_TYPE_DEPOSIT specifies deposit condition type
	ConditionTypeDeposit ConditionType = 1
	// CONDITION_TYPE_SWAP specifies swap condition type
	ConditionTypeSwap ConditionType = 2
	// CONDITION_TYPE_FARMING specifies farming (stake) condition type
	ConditionTypeFarming ConditionType = 3
)

var ConditionType_name = map[int32]string{
	0: "CONDITION_TYPE_UNSPECIFIED",
	1: "CONDITION_TYPE_DEPOSIT",
	2: "CONDITION_TYPE_SWAP",
	3: "CONDITION_TYPE_FARMING",
}

var ConditionType_value = map[string]int32{
	"CONDITION_TYPE_UNSPECIFIED": 0,
	"CONDITION_TYPE_DEPOSIT":     1,
	"CONDITION_TYPE_SWAP":        2,
	"CONDITION_TYPE_FARMING":     3,
}

func (x ConditionType) String() string {
	return proto.EnumName(ConditionType_name, int32(x))
}

func (ConditionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_84886eaa62c7639a, []int{0}
}

// Airdrop defines airdrop information.
type Airdrop struct {
	// id specifies index of the airdrop
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// source_address defines the bech32-encoded source address
	// where the source of coins from
	SourceAddress string `protobuf:"bytes,2,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	// termination_address defines the bech32-encoded termination address
	// where the remaining source coins are sent to
	TerminationAddress string `protobuf:"bytes,3,opt,name=termination_address,json=terminationAddress,proto3" json:"termination_address,omitempty"`
	// conditions specifies a list of conditions
	Conditions []ConditionType `protobuf:"varint,4,rep,packed,name=conditions,proto3,enum=squad.claim.v1beta1.ConditionType" json:"conditions,omitempty"`
	// start_time specifies the start time of the airdrop
	StartTime time.Time `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time"`
	// end_time specifies the start time of the airdrop
	EndTime time.Time `protobuf:"bytes,6,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time"`
}

func (m *Airdrop) Reset()         { *m = Airdrop{} }
func (m *Airdrop) String() string { return proto.CompactTextString(m) }
func (*Airdrop) ProtoMessage()    {}
func (*Airdrop) Descriptor() ([]byte, []int) {
	return fileDescriptor_84886eaa62c7639a, []int{0}
}
func (m *Airdrop) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Airdrop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Airdrop.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Airdrop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Airdrop.Merge(m, src)
}
func (m *Airdrop) XXX_Size() int {
	return m.Size()
}
func (m *Airdrop) XXX_DiscardUnknown() {
	xxx_messageInfo_Airdrop.DiscardUnknown(m)
}

var xxx_messageInfo_Airdrop proto.InternalMessageInfo

// ClaimRecord defines claim record that corresponds to the airdrop.
type ClaimRecord struct {
	// airdrop_id specifies airdrop id
	AirdropId uint64 `protobuf:"varint,1,opt,name=airdrop_id,json=airdropId,proto3" json:"airdrop_id,omitempty"`
	// recipient specifies the bech32-encoded address that is eligible to claim airdrop
	Recipient string `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// initial_claimable_coins specifies the initial claimable coins
	InitialClaimableCoins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=initial_claimable_coins,json=initialClaimableCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"initial_claimable_coins"`
	// claimable_coins specifies the unclaimed claimable coins
	ClaimableCoins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=claimable_coins,json=claimableCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"claimable_coins"`
	// claimed_conditions specifies a list of condition types
	// initial values are empty and each condition type gets appended when claim is successfully executed
	ClaimedConditions []ConditionType `protobuf:"varint,5,rep,packed,name=claimed_conditions,json=claimedConditions,proto3,enum=squad.claim.v1beta1.ConditionType" json:"claimed_conditions,omitempty"`
}

func (m *ClaimRecord) Reset()         { *m = ClaimRecord{} }
func (m *ClaimRecord) String() string { return proto.CompactTextString(m) }
func (*ClaimRecord) ProtoMessage()    {}
func (*ClaimRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_84886eaa62c7639a, []int{1}
}
func (m *ClaimRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRecord.Merge(m, src)
}
func (m *ClaimRecord) XXX_Size() int {
	return m.Size()
}
func (m *ClaimRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRecord proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("squad.claim.v1beta1.ConditionType", ConditionType_name, ConditionType_value)
	proto.RegisterType((*Airdrop)(nil), "squad.claim.v1beta1.Airdrop")
	proto.RegisterType((*ClaimRecord)(nil), "squad.claim.v1beta1.ClaimRecord")
}

func init() { proto.RegisterFile("squad/claim/v1beta1/claim.proto", fileDescriptor_84886eaa62c7639a) }

var fileDescriptor_84886eaa62c7639a = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x6f, 0xd3, 0x30,
	0x1c, 0x4d, 0xda, 0xee, 0x4f, 0x3d, 0xad, 0x74, 0xde, 0x06, 0x25, 0x1a, 0x69, 0x34, 0x09, 0xa9,
	0x42, 0x5a, 0xc2, 0x06, 0x47, 0x24, 0xd4, 0x7f, 0x83, 0x48, 0xd0, 0x55, 0x69, 0xa7, 0x09, 0x2e,
	0x91, 0x13, 0x7b, 0xc5, 0xa2, 0x8d, 0x43, 0xec, 0x02, 0x3b, 0x73, 0x41, 0x3b, 0xed, 0x0b, 0x4c,
	0x1c, 0xb8, 0xf1, 0x2d, 0xb8, 0xed, 0xb8, 0x23, 0x27, 0x06, 0xdb, 0x07, 0xe0, 0x2b, 0x20, 0x27,
	0xe9, 0xd6, 0x8e, 0x5e, 0x26, 0x71, 0x6a, 0xf3, 0x7e, 0xef, 0xfd, 0x9e, 0xfd, 0x7e, 0xb6, 0x41,
	0x99, 0xbf, 0x1b, 0x22, 0x6c, 0xf9, 0x7d, 0x44, 0x07, 0xd6, 0xfb, 0x4d, 0x8f, 0x08, 0xb4, 0x99,
	0x7c, 0x99, 0x61, 0xc4, 0x04, 0x83, 0xcb, 0x31, 0xc1, 0x4c, 0xa0, 0x94, 0xa0, 0xad, 0xf4, 0x58,
	0x8f, 0xc5, 0x75, 0x4b, 0xfe, 0x4b, 0xa8, 0x5a, 0xb9, 0xc7, 0x58, 0xaf, 0x4f, 0xac, 0xf8, 0xcb,
	0x1b, 0xee, 0x5b, 0x82, 0x0e, 0x08, 0x17, 0x68, 0x10, 0xa6, 0x04, 0xdd, 0x67, 0x7c, 0xc0, 0xb8,
	0xe5, 0x21, 0x4e, 0xae, 0xcc, 0x18, 0x0d, 0x92, 0xfa, 0xfa, 0xf7, 0x0c, 0x98, 0xab, 0xd2, 0x08,
	0x47, 0x2c, 0x84, 0x05, 0x90, 0xa1, 0xb8, 0xa4, 0x1a, 0x6a, 0x25, 0xe7, 0x64, 0x28, 0x86, 0xf7,
	0x41, 0x81, 0xb3, 0x61, 0xe4, 0x13, 0x17, 0x61, 0x1c, 0x11, 0xce, 0x4b, 0x19, 0x43, 0xad, 0xe4,
	0x9d, 0xc5, 0x04, 0xad, 0x26, 0x20, 0xb4, 0xc0, 0xb2, 0x20, 0xd1, 0x80, 0x06, 0x48, 0x50, 0x16,
	0x5c, 0x72, 0xb3, 0x31, 0x17, 0x8e, 0x95, 0x46, 0x82, 0xe7, 0x00, 0xf8, 0x2c, 0xc0, 0x54, 0x62,
	0xbc, 0x94, 0x33, 0xb2, 0x95, 0xc2, 0xd6, 0xba, 0x39, 0x65, 0xd3, 0x66, 0x7d, 0x44, 0xeb, 0x1e,
	0x84, 0xa4, 0x96, 0x3b, 0xf9, 0x59, 0x56, 0x9c, 0x31, 0x2d, 0xac, 0x03, 0xc0, 0x05, 0x8a, 0x84,
	0x2b, 0xb7, 0x5d, 0x9a, 0x31, 0xd4, 0xca, 0xc2, 0x96, 0x66, 0x26, 0x99, 0x98, 0xa3, 0x4c, 0xcc,
	0xee, 0x28, 0x93, 0xda, 0xbc, 0xec, 0x70, 0x74, 0x56, 0x56, 0x9d, 0x7c, 0xac, 0x93, 0x15, 0xf8,
	0x14, 0xcc, 0x93, 0x00, 0x27, 0x2d, 0x66, 0x6f, 0xd0, 0x62, 0x8e, 0x04, 0x58, 0xe2, 0xeb, 0x5f,
	0xb2, 0x60, 0xa1, 0x2e, 0xd7, 0xed, 0x10, 0x9f, 0x45, 0x18, 0xde, 0x03, 0x00, 0x25, 0x91, 0xba,
	0x97, 0x79, 0xe6, 0x53, 0xc4, 0xc6, 0x70, 0x0d, 0xe4, 0x23, 0xe2, 0xd3, 0x90, 0x92, 0x40, 0xa4,
	0x89, 0x5e, 0x01, 0xf0, 0x93, 0x0a, 0xee, 0xd0, 0x80, 0x0a, 0x8a, 0xfa, 0x6e, 0x1c, 0x06, 0xf2,
	0xfa, 0xc4, 0x95, 0x13, 0x93, 0x91, 0x66, 0x2b, 0x0b, 0x5b, 0x77, 0xcd, 0x64, 0xa6, 0xa6, 0x9c,
	0xe9, 0x58, 0x54, 0x34, 0xa8, 0x3d, 0x94, 0x8b, 0xfb, 0x76, 0x56, 0xae, 0xf4, 0xa8, 0x78, 0x33,
	0xf4, 0x4c, 0x9f, 0x0d, 0xac, 0xf4, 0x00, 0x24, 0x3f, 0x1b, 0x1c, 0xbf, 0xb5, 0xc4, 0x41, 0x48,
	0x78, 0x2c, 0xe0, 0xce, 0x6a, 0xea, 0x55, 0x1f, 0x59, 0xc5, 0x30, 0x14, 0xe0, 0xd6, 0x75, 0xf3,
	0xdc, 0xff, 0x37, 0x2f, 0xf8, 0x93, 0xae, 0x7b, 0x00, 0xc6, 0x08, 0xc1, 0xee, 0xd8, 0x01, 0x99,
	0xb9, 0xe1, 0x01, 0x59, 0x4a, 0x7b, 0x5c, 0xd6, 0xf8, 0x83, 0x3f, 0x2a, 0x58, 0x9c, 0xa0, 0xc2,
	0x27, 0x40, 0xab, 0xef, 0xb4, 0x1a, 0x76, 0xd7, 0xde, 0x69, 0xb9, 0xdd, 0x57, 0xed, 0xa6, 0xbb,
	0xdb, 0xea, 0xb4, 0x9b, 0x75, 0x7b, 0xdb, 0x6e, 0x36, 0x8a, 0x8a, 0xb6, 0x76, 0x78, 0x6c, 0x94,
	0x26, 0x24, 0xbb, 0x01, 0x0f, 0x89, 0x4f, 0xf7, 0x29, 0xc1, 0xf0, 0x31, 0xb8, 0x7d, 0x4d, 0xdd,
	0x68, 0xb6, 0x77, 0x3a, 0x76, 0xb7, 0xa8, 0x6a, 0xa5, 0xc3, 0x63, 0x63, 0x65, 0x42, 0xd9, 0x20,
	0x21, 0xe3, 0x54, 0x40, 0x13, 0x2c, 0x5f, 0x53, 0x75, 0xf6, 0xaa, 0xed, 0x62, 0x46, 0x5b, 0x3d,
	0x3c, 0x36, 0x96, 0x26, 0x24, 0x9d, 0x0f, 0x28, 0x9c, 0xe2, 0xb2, 0x5d, 0x75, 0x5e, 0xda, 0xad,
	0x67, 0xc5, 0xec, 0x14, 0x97, 0x6d, 0x24, 0x2f, 0x5a, 0x4f, 0xcb, 0x7d, 0xfe, 0xaa, 0x2b, 0xb5,
	0x17, 0x27, 0xbf, 0x75, 0xe5, 0xe4, 0x5c, 0x57, 0x4f, 0xcf, 0x75, 0xf5, 0xd7, 0xb9, 0xae, 0x1e,
	0x5d, 0xe8, 0xca, 0xe9, 0x85, 0xae, 0xfc, 0xb8, 0xd0, 0x95, 0xd7, 0xe6, 0x3f, 0x23, 0x92, 0xd9,
	0x6e, 0xf4, 0x91, 0xc7, 0xad, 0xe4, 0x75, 0xfa, 0x98, 0xbe, 0x4f, 0xf1, 0xb8, 0xbc, 0xd9, 0xf8,
	0x22, 0x3c, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x2b, 0xde, 0xf7, 0xbb, 0x04, 0x00, 0x00,
}

func (m *Airdrop) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Airdrop) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Airdrop) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintClaim(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintClaim(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x2a
	if len(m.Conditions) > 0 {
		dAtA4 := make([]byte, len(m.Conditions)*10)
		var j3 int
		for _, num := range m.Conditions {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintClaim(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TerminationAddress) > 0 {
		i -= len(m.TerminationAddress)
		copy(dAtA[i:], m.TerminationAddress)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.TerminationAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SourceAddress) > 0 {
		i -= len(m.SourceAddress)
		copy(dAtA[i:], m.SourceAddress)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.SourceAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintClaim(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ClaimRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimedConditions) > 0 {
		dAtA6 := make([]byte, len(m.ClaimedConditions)*10)
		var j5 int
		for _, num := range m.ClaimedConditions {
			for num >= 1<<7 {
				dAtA6[j5] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j5++
			}
			dAtA6[j5] = uint8(num)
			j5++
		}
		i -= j5
		copy(dAtA[i:], dAtA6[:j5])
		i = encodeVarintClaim(dAtA, i, uint64(j5))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ClaimableCoins) > 0 {
		for iNdEx := len(m.ClaimableCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimableCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClaim(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.InitialClaimableCoins) > 0 {
		for iNdEx := len(m.InitialClaimableCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialClaimableCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClaim(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x12
	}
	if m.AirdropId != 0 {
		i = encodeVarintClaim(dAtA, i, uint64(m.AirdropId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaim(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaim(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Airdrop) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovClaim(uint64(m.Id))
	}
	l = len(m.SourceAddress)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	l = len(m.TerminationAddress)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	if len(m.Conditions) > 0 {
		l = 0
		for _, e := range m.Conditions {
			l += sovClaim(uint64(e))
		}
		n += 1 + sovClaim(uint64(l)) + l
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovClaim(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovClaim(uint64(l))
	return n
}

func (m *ClaimRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AirdropId != 0 {
		n += 1 + sovClaim(uint64(m.AirdropId))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	if len(m.InitialClaimableCoins) > 0 {
		for _, e := range m.InitialClaimableCoins {
			l = e.Size()
			n += 1 + l + sovClaim(uint64(l))
		}
	}
	if len(m.ClaimableCoins) > 0 {
		for _, e := range m.ClaimableCoins {
			l = e.Size()
			n += 1 + l + sovClaim(uint64(l))
		}
	}
	if len(m.ClaimedConditions) > 0 {
		l = 0
		for _, e := range m.ClaimedConditions {
			l += sovClaim(uint64(e))
		}
		n += 1 + sovClaim(uint64(l)) + l
	}
	return n
}

func sovClaim(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaim(x uint64) (n int) {
	return sovClaim(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Airdrop) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
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
			return fmt.Errorf("proto: Airdrop: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Airdrop: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return fmt.Errorf("proto: wrong wireType = %d for field SourceAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TerminationAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TerminationAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v ConditionType
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= ConditionType(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Conditions = append(m.Conditions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaim
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaim
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Conditions) == 0 {
					m.Conditions = make([]ConditionType, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v ConditionType
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaim
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= ConditionType(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Conditions = append(m.Conditions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Conditions", wireType)
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
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
func (m *ClaimRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
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
			return fmt.Errorf("proto: ClaimRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropId", wireType)
			}
			m.AirdropId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AirdropId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialClaimableCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialClaimableCoins = append(m.InitialClaimableCoins, types.Coin{})
			if err := m.InitialClaimableCoins[len(m.InitialClaimableCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimableCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimableCoins = append(m.ClaimableCoins, types.Coin{})
			if err := m.ClaimableCoins[len(m.ClaimableCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType == 0 {
				var v ConditionType
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= ConditionType(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ClaimedConditions = append(m.ClaimedConditions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaim
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaim
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.ClaimedConditions) == 0 {
					m.ClaimedConditions = make([]ConditionType, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v ConditionType
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaim
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= ConditionType(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ClaimedConditions = append(m.ClaimedConditions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedConditions", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
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
func skipClaim(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaim
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
					return 0, ErrIntOverflowClaim
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
					return 0, ErrIntOverflowClaim
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
				return 0, ErrInvalidLengthClaim
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaim
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaim
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaim        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaim          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaim = fmt.Errorf("proto: unexpected end of group")
)
