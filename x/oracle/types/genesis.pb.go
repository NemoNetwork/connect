// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slinky/oracle/v1/genesis.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	types "github.com/skip-mev/connect/v2/pkg/types"
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

// QuotePrice is the representation of the aggregated prices for a CurrencyPair,
// where price represents the price of Base in terms of Quote
type QuotePrice struct {
	Price cosmossdk_io_math.Int `protobuf:"bytes,1,opt,name=price,proto3,customtype=cosmossdk.io/math.Int" json:"price"`
	// BlockTimestamp tracks the block height associated with this price update.
	// We include block timestamp alongside the price to ensure that smart
	// contracts and applications are not utilizing stale oracle prices
	BlockTimestamp time.Time `protobuf:"bytes,2,opt,name=block_timestamp,json=blockTimestamp,proto3,stdtime" json:"block_timestamp"`
	// BlockHeight is height of block mentioned above
	BlockHeight uint64 `protobuf:"varint,3,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (m *QuotePrice) Reset()         { *m = QuotePrice{} }
func (m *QuotePrice) String() string { return proto.CompactTextString(m) }
func (*QuotePrice) ProtoMessage()    {}
func (*QuotePrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_de36a97821ccc13b, []int{0}
}
func (m *QuotePrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuotePrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuotePrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuotePrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotePrice.Merge(m, src)
}
func (m *QuotePrice) XXX_Size() int {
	return m.Size()
}
func (m *QuotePrice) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotePrice.DiscardUnknown(m)
}

var xxx_messageInfo_QuotePrice proto.InternalMessageInfo

func (m *QuotePrice) GetBlockTimestamp() time.Time {
	if m != nil {
		return m.BlockTimestamp
	}
	return time.Time{}
}

func (m *QuotePrice) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

// CurrencyPairState represents the stateful information tracked by the x/oracle
// module per-currency-pair.
type CurrencyPairState struct {
	// QuotePrice is the latest price for a currency-pair, notice this value can
	// be null in the case that no price exists for the currency-pair
	Price *QuotePrice `protobuf:"bytes,1,opt,name=price,proto3" json:"price,omitempty"`
	// Nonce is the number of updates this currency-pair has received
	Nonce uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// ID is the ID of the CurrencyPair
	Id uint64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *CurrencyPairState) Reset()         { *m = CurrencyPairState{} }
func (m *CurrencyPairState) String() string { return proto.CompactTextString(m) }
func (*CurrencyPairState) ProtoMessage()    {}
func (*CurrencyPairState) Descriptor() ([]byte, []int) {
	return fileDescriptor_de36a97821ccc13b, []int{1}
}
func (m *CurrencyPairState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CurrencyPairState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CurrencyPairState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CurrencyPairState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrencyPairState.Merge(m, src)
}
func (m *CurrencyPairState) XXX_Size() int {
	return m.Size()
}
func (m *CurrencyPairState) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrencyPairState.DiscardUnknown(m)
}

var xxx_messageInfo_CurrencyPairState proto.InternalMessageInfo

func (m *CurrencyPairState) GetPrice() *QuotePrice {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *CurrencyPairState) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *CurrencyPairState) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// CurrencyPairGenesis is the information necessary for initialization of a
// CurrencyPair.
type CurrencyPairGenesis struct {
	// The CurrencyPair to be added to module state
	CurrencyPair types.CurrencyPair `protobuf:"bytes,1,opt,name=currency_pair,json=currencyPair,proto3" json:"currency_pair"`
	// A genesis price if one exists (note this will be empty, unless it results
	// from forking the state of this module)
	CurrencyPairPrice *QuotePrice `protobuf:"bytes,2,opt,name=currency_pair_price,json=currencyPairPrice,proto3" json:"currency_pair_price,omitempty"`
	// nonce is the nonce (number of updates) for the CP (same case as above,
	// likely 0 unless it results from fork of module)
	Nonce uint64 `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// id is the ID of the CurrencyPair
	Id uint64 `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *CurrencyPairGenesis) Reset()         { *m = CurrencyPairGenesis{} }
func (m *CurrencyPairGenesis) String() string { return proto.CompactTextString(m) }
func (*CurrencyPairGenesis) ProtoMessage()    {}
func (*CurrencyPairGenesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_de36a97821ccc13b, []int{2}
}
func (m *CurrencyPairGenesis) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CurrencyPairGenesis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CurrencyPairGenesis.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CurrencyPairGenesis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrencyPairGenesis.Merge(m, src)
}
func (m *CurrencyPairGenesis) XXX_Size() int {
	return m.Size()
}
func (m *CurrencyPairGenesis) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrencyPairGenesis.DiscardUnknown(m)
}

var xxx_messageInfo_CurrencyPairGenesis proto.InternalMessageInfo

func (m *CurrencyPairGenesis) GetCurrencyPair() types.CurrencyPair {
	if m != nil {
		return m.CurrencyPair
	}
	return types.CurrencyPair{}
}

func (m *CurrencyPairGenesis) GetCurrencyPairPrice() *QuotePrice {
	if m != nil {
		return m.CurrencyPairPrice
	}
	return nil
}

func (m *CurrencyPairGenesis) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *CurrencyPairGenesis) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// GenesisState is the genesis-state for the x/oracle module, it takes a set of
// predefined CurrencyPairGeneses
type GenesisState struct {
	// CurrencyPairGenesis is the set of CurrencyPairGeneses for the module. I.e
	// the starting set of CurrencyPairs for the module + information regarding
	// their latest update.
	CurrencyPairGenesis []CurrencyPairGenesis `protobuf:"bytes,1,rep,name=currency_pair_genesis,json=currencyPairGenesis,proto3" json:"currency_pair_genesis"`
	// NextID is the next ID to be used for a CurrencyPair
	NextId uint64 `protobuf:"varint,2,opt,name=next_id,json=nextId,proto3" json:"next_id,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_de36a97821ccc13b, []int{3}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetCurrencyPairGenesis() []CurrencyPairGenesis {
	if m != nil {
		return m.CurrencyPairGenesis
	}
	return nil
}

func (m *GenesisState) GetNextId() uint64 {
	if m != nil {
		return m.NextId
	}
	return 0
}

func init() {
	proto.RegisterType((*QuotePrice)(nil), "slinky.oracle.v1.QuotePrice")
	proto.RegisterType((*CurrencyPairState)(nil), "slinky.oracle.v1.CurrencyPairState")
	proto.RegisterType((*CurrencyPairGenesis)(nil), "slinky.oracle.v1.CurrencyPairGenesis")
	proto.RegisterType((*GenesisState)(nil), "slinky.oracle.v1.GenesisState")
}

func init() { proto.RegisterFile("slinky/oracle/v1/genesis.proto", fileDescriptor_de36a97821ccc13b) }

var fileDescriptor_de36a97821ccc13b = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x8f, 0x12, 0x3d,
	0x1c, 0xc7, 0x29, 0xb0, 0xfb, 0x3c, 0x16, 0x5c, 0xdd, 0x61, 0x37, 0x22, 0xd1, 0x01, 0x31, 0x26,
	0x24, 0x86, 0x36, 0x8b, 0x17, 0xaf, 0xe2, 0x61, 0x97, 0x83, 0x09, 0x8e, 0x9e, 0xbc, 0x4c, 0x86,
	0x52, 0x87, 0x06, 0xa6, 0x9d, 0x4c, 0x0b, 0x59, 0xde, 0x81, 0xc7, 0x7d, 0x31, 0xbe, 0x01, 0x6f,
	0x1c, 0x37, 0x9e, 0xd4, 0x03, 0x1a, 0x78, 0x23, 0x66, 0xda, 0x0e, 0x0e, 0xee, 0x1e, 0xbc, 0xcd,
	0xaf, 0xbf, 0x7f, 0xdf, 0xcf, 0xb7, 0x1d, 0xe8, 0xca, 0x19, 0xe3, 0xd3, 0x25, 0x16, 0x49, 0x40,
	0x66, 0x14, 0x2f, 0xce, 0x70, 0x48, 0x39, 0x95, 0x4c, 0xa2, 0x38, 0x11, 0x4a, 0x38, 0xf7, 0x4d,
	0x1e, 0x99, 0x3c, 0x5a, 0x9c, 0x35, 0x4e, 0x42, 0x11, 0x0a, 0x9d, 0xc4, 0xe9, 0x97, 0xa9, 0x6b,
	0x34, 0x43, 0x21, 0xc2, 0x19, 0xc5, 0x3a, 0x1a, 0xcd, 0x3f, 0x62, 0xc5, 0x22, 0x2a, 0x55, 0x10,
	0xc5, 0xb6, 0xe0, 0x21, 0x11, 0x32, 0x12, 0xd2, 0x37, 0x9d, 0x26, 0xb0, 0xa9, 0xa7, 0x56, 0x83,
	0x5a, 0xc6, 0x54, 0xa6, 0x12, 0xc8, 0x3c, 0x49, 0x28, 0x27, 0x4b, 0x3f, 0x0e, 0x58, 0x62, 0x8a,
	0xda, 0x5f, 0x00, 0x84, 0x6f, 0xe7, 0x42, 0xd1, 0x61, 0xc2, 0x08, 0x75, 0x5e, 0xc1, 0x83, 0x38,
	0xfd, 0xa8, 0x83, 0x16, 0xe8, 0xdc, 0xe9, 0x3f, 0x5f, 0xad, 0x9b, 0x85, 0x1f, 0xeb, 0xe6, 0xa9,
	0x19, 0x2c, 0xc7, 0x53, 0xc4, 0x04, 0x8e, 0x02, 0x35, 0x41, 0x03, 0xae, 0xbe, 0x7e, 0xee, 0x42,
	0xbb, 0x71, 0xc0, 0x95, 0x67, 0x3a, 0x9d, 0x37, 0xf0, 0xde, 0x68, 0x26, 0xc8, 0xd4, 0xdf, 0x49,
	0xad, 0x17, 0x5b, 0xa0, 0x53, 0xe9, 0x35, 0x90, 0x81, 0x41, 0x19, 0x0c, 0x7a, 0x9f, 0x55, 0xf4,
	0xff, 0x4f, 0x17, 0x5d, 0xfd, 0x6c, 0x02, 0xef, 0x48, 0x37, 0xef, 0x32, 0xce, 0x13, 0x58, 0x35,
	0xe3, 0x26, 0x94, 0x85, 0x13, 0x55, 0x2f, 0xb5, 0x40, 0xa7, 0xec, 0x55, 0xf4, 0xd9, 0x85, 0x3e,
	0x6a, 0x4b, 0x78, 0xfc, 0xda, 0xa2, 0x0d, 0x03, 0x96, 0xbc, 0x53, 0x81, 0xa2, 0xce, 0xcb, 0x3c,
	0x49, 0xa5, 0xf7, 0x08, 0xfd, 0xed, 0x38, 0xfa, 0x83, 0xdd, 0x2f, 0xaf, 0xd6, 0x4d, 0x90, 0x01,
	0x9c, 0xc0, 0x03, 0x2e, 0x38, 0xa1, 0x5a, 0x76, 0xd9, 0x33, 0x81, 0x73, 0x04, 0x8b, 0x6c, 0x6c,
	0xb7, 0x17, 0xd9, 0xb8, 0xfd, 0x1d, 0xc0, 0x5a, 0x7e, 0xeb, 0xb9, 0xb9, 0x5f, 0xe7, 0x02, 0xde,
	0xdd, 0xf3, 0xd9, 0xee, 0x7f, 0x9c, 0xed, 0xd7, 0xb7, 0x91, 0xae, 0xcf, 0x37, 0x6b, 0x01, 0x05,
	0xaf, 0x4a, 0x72, 0x67, 0x8e, 0x07, 0x6b, 0x7b, 0x93, 0x7c, 0xc3, 0x53, 0xfc, 0x67, 0x9e, 0xe3,
	0xfc, 0xb8, 0xe1, 0x3e, 0x5b, 0xe9, 0x26, 0x5b, 0x79, 0xc7, 0xf6, 0x09, 0xc0, 0xaa, 0xe5, 0x31,
	0x66, 0xfa, 0xf0, 0x74, 0x5f, 0x8a, 0x7d, 0xcd, 0x75, 0xd0, 0x2a, 0x75, 0x2a, 0xbd, 0x67, 0x37,
	0xc5, 0xdc, 0x62, 0x8d, 0x85, 0xac, 0x91, 0x5b, 0x5c, 0x7b, 0x00, 0xff, 0xe3, 0xf4, 0x52, 0xf9,
	0x6c, 0x6c, 0x5d, 0x3f, 0x4c, 0xc3, 0xc1, 0xb8, 0x7f, 0xbe, 0xda, 0xb8, 0xe0, 0x7a, 0xe3, 0x82,
	0x5f, 0x1b, 0x17, 0x5c, 0x6d, 0xdd, 0xc2, 0xf5, 0xd6, 0x2d, 0x7c, 0xdb, 0xba, 0x85, 0x0f, 0xdd,
	0x90, 0xa9, 0xc9, 0x7c, 0x84, 0x88, 0x88, 0xb0, 0x9c, 0xb2, 0xb8, 0x1b, 0xd1, 0x05, 0x26, 0x82,
	0x73, 0x4a, 0x14, 0x5e, 0xf4, 0xf0, 0x65, 0xf6, 0xf3, 0x69, 0xc7, 0x47, 0x87, 0xfa, 0xd5, 0xbd,
	0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x53, 0x7a, 0x86, 0xff, 0x9a, 0x03, 0x00, 0x00,
}

func (m *QuotePrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuotePrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuotePrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x18
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.BlockTimestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.BlockTimestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGenesis(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CurrencyPairState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CurrencyPairState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CurrencyPairState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x18
	}
	if m.Nonce != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x10
	}
	if m.Price != nil {
		{
			size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CurrencyPairGenesis) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CurrencyPairGenesis) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CurrencyPairGenesis) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x20
	}
	if m.Nonce != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x18
	}
	if m.CurrencyPairPrice != nil {
		{
			size, err := m.CurrencyPairPrice.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.CurrencyPair.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NextId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.CurrencyPairGenesis) > 0 {
		for iNdEx := len(m.CurrencyPairGenesis) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CurrencyPairGenesis[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QuotePrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Price.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.BlockTimestamp)
	n += 1 + l + sovGenesis(uint64(l))
	if m.BlockHeight != 0 {
		n += 1 + sovGenesis(uint64(m.BlockHeight))
	}
	return n
}

func (m *CurrencyPairState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Price != nil {
		l = m.Price.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovGenesis(uint64(m.Nonce))
	}
	if m.Id != 0 {
		n += 1 + sovGenesis(uint64(m.Id))
	}
	return n
}

func (m *CurrencyPairGenesis) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CurrencyPair.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.CurrencyPairPrice != nil {
		l = m.CurrencyPairPrice.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovGenesis(uint64(m.Nonce))
	}
	if m.Id != 0 {
		n += 1 + sovGenesis(uint64(m.Id))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CurrencyPairGenesis) > 0 {
		for _, e := range m.CurrencyPairGenesis {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.NextId != 0 {
		n += 1 + sovGenesis(uint64(m.NextId))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QuotePrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: QuotePrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuotePrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.BlockTimestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *CurrencyPairState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: CurrencyPairState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CurrencyPairState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Price == nil {
				m.Price = &QuotePrice{}
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *CurrencyPairGenesis) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: CurrencyPairGenesis: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CurrencyPairGenesis: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrencyPair", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrencyPair.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrencyPairPrice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CurrencyPairPrice == nil {
				m.CurrencyPairPrice = &QuotePrice{}
			}
			if err := m.CurrencyPairPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrencyPairGenesis", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrencyPairGenesis = append(m.CurrencyPairGenesis, CurrencyPairGenesis{})
			if err := m.CurrencyPairGenesis[len(m.CurrencyPairGenesis)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextId", wireType)
			}
			m.NextId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
