// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slinky/mm2/v1/market.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/skip-mev/slinky/pkg/types"
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

// AggregationType is the type of aggregation that will be used to aggregate the
// prices of the tickers.
type AggregationType int32

const (
	// UNKNOWN_AGGREGATION_TYPE is the default value for the aggregation type.
	AggregationType_UNKNOWN_AGGREGATION_TYPE AggregationType = 0
	// IndexPriceAggregation is the type of aggregation that will be used to
	// aggregate the prices of the tickers. Specifically, this converts the prices
	// either directly or using the index price to a common currency pair.
	AggregationType_INDEX_PRICE_AGGREGATION AggregationType = 1
	// StandardMedianAggregation is the type of aggregation that will be used to
	// aggregate the prices of the tickers. Specifically, this converts the prices
	// to a common currency pair and then takes the median of the prices. No
	// conversions are done if the prices are already in the common currency pair.
	AggregationType_STANDARD_MEDIAN_AGGREGATION AggregationType = 2
)

var AggregationType_name = map[int32]string{
	0: "UNKNOWN_AGGREGATION_TYPE",
	1: "INDEX_PRICE_AGGREGATION",
	2: "STANDARD_MEDIAN_AGGREGATION",
}

var AggregationType_value = map[string]int32{
	"UNKNOWN_AGGREGATION_TYPE":    0,
	"INDEX_PRICE_AGGREGATION":     1,
	"STANDARD_MEDIAN_AGGREGATION": 2,
}

func (x AggregationType) String() string {
	return proto.EnumName(AggregationType_name, int32(x))
}

func (AggregationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_446157c095b24f63, []int{0}
}

// Market encapsulates a Ticker and its provider-specific configuration.
type Market struct {
	// Ticker represents a price feed for a given asset pair i.e. BTC/USD. The
	// price feed is scaled to a number of decimal places and has a minimum number
	// of providers required to consider the ticker valid.
	Ticker *Ticker `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	// ProviderConfigs is the list of provider-specific configs for this Market.
	ProviderConfigs []ProviderConfig `protobuf:"bytes,2,rep,name=provider_configs,json=providerConfigs,proto3" json:"provider_configs"`
}

func (m *Market) Reset()      { *m = Market{} }
func (*Market) ProtoMessage() {}
func (*Market) Descriptor() ([]byte, []int) {
	return fileDescriptor_446157c095b24f63, []int{0}
}
func (m *Market) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Market) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Market.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Market) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Market.Merge(m, src)
}
func (m *Market) XXX_Size() int {
	return m.Size()
}
func (m *Market) XXX_DiscardUnknown() {
	xxx_messageInfo_Market.DiscardUnknown(m)
}

var xxx_messageInfo_Market proto.InternalMessageInfo

func (m *Market) GetTicker() *Ticker {
	if m != nil {
		return m.Ticker
	}
	return nil
}

func (m *Market) GetProviderConfigs() []ProviderConfig {
	if m != nil {
		return m.ProviderConfigs
	}
	return nil
}

// Ticker represents a price feed for a given asset pair i.e. BTC/USD. The price
// feed is scaled to a number of decimal places and has a minimum number of
// providers required to consider the ticker valid.
type Ticker struct {
	// CurrencyPair is the currency pair for this ticker.
	CurrencyPair types.CurrencyPair `protobuf:"bytes,1,opt,name=currency_pair,json=currencyPair,proto3" json:"currency_pair"`
	// Decimals is the number of decimal places for the ticker. The number of
	// decimal places is used to convert the price to a human-readable format.
	Decimals uint64 `protobuf:"varint,2,opt,name=decimals,proto3" json:"decimals,omitempty"`
	// MinProviderCount is the minimum number of providers required to consider
	// the ticker valid.
	MinProviderCount uint64 `protobuf:"varint,3,opt,name=min_provider_count,json=minProviderCount,proto3" json:"min_provider_count,omitempty"`
	// Enabled is the flag that denotes if the Ticker is enabled for price
	// fetching by an oracle.
	Enabled bool `protobuf:"varint,14,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// MetadataJSON is a string of JSON that encodes any extra configuration
	// for the given ticker.
	Metadata_JSON string `protobuf:"bytes,15,opt,name=metadata_JSON,json=metadataJSON,proto3" json:"metadata_JSON,omitempty"`
}

func (m *Ticker) Reset()      { *m = Ticker{} }
func (*Ticker) ProtoMessage() {}
func (*Ticker) Descriptor() ([]byte, []int) {
	return fileDescriptor_446157c095b24f63, []int{1}
}
func (m *Ticker) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ticker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ticker.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ticker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ticker.Merge(m, src)
}
func (m *Ticker) XXX_Size() int {
	return m.Size()
}
func (m *Ticker) XXX_DiscardUnknown() {
	xxx_messageInfo_Ticker.DiscardUnknown(m)
}

var xxx_messageInfo_Ticker proto.InternalMessageInfo

func (m *Ticker) GetCurrencyPair() types.CurrencyPair {
	if m != nil {
		return m.CurrencyPair
	}
	return types.CurrencyPair{}
}

func (m *Ticker) GetDecimals() uint64 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *Ticker) GetMinProviderCount() uint64 {
	if m != nil {
		return m.MinProviderCount
	}
	return 0
}

func (m *Ticker) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Ticker) GetMetadata_JSON() string {
	if m != nil {
		return m.Metadata_JSON
	}
	return ""
}

type ProviderConfig struct {
	// Name corresponds to the name of the provider for which the configuration is
	// being set.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// OffChainTicker is the off-chain representation of the ticker i.e. BTC/USD.
	// The off-chain ticker is unique to a given provider and is used to fetch the
	// price of the ticker from the provider.
	OffChainTicker string `protobuf:"bytes,2,opt,name=off_chain_ticker,json=offChainTicker,proto3" json:"off_chain_ticker,omitempty"`
	Index          string `protobuf:"bytes,3,opt,name=index,proto3" json:"index,omitempty"`
	Invert         bool   `protobuf:"varint,4,opt,name=invert,proto3" json:"invert,omitempty"`
	// MetadataJSON is a string of JSON that encodes any extra configuration
	// for the given provider config.
	Metadata_JSON string `protobuf:"bytes,15,opt,name=metadata_JSON,json=metadataJSON,proto3" json:"metadata_JSON,omitempty"`
}

func (m *ProviderConfig) Reset()         { *m = ProviderConfig{} }
func (m *ProviderConfig) String() string { return proto.CompactTextString(m) }
func (*ProviderConfig) ProtoMessage()    {}
func (*ProviderConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_446157c095b24f63, []int{2}
}
func (m *ProviderConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProviderConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProviderConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProviderConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderConfig.Merge(m, src)
}
func (m *ProviderConfig) XXX_Size() int {
	return m.Size()
}
func (m *ProviderConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderConfig proto.InternalMessageInfo

func (m *ProviderConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProviderConfig) GetOffChainTicker() string {
	if m != nil {
		return m.OffChainTicker
	}
	return ""
}

func (m *ProviderConfig) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *ProviderConfig) GetInvert() bool {
	if m != nil {
		return m.Invert
	}
	return false
}

func (m *ProviderConfig) GetMetadata_JSON() string {
	if m != nil {
		return m.Metadata_JSON
	}
	return ""
}

// MarketMap maps ticker strings to their Markets.
type MarketMap struct {
	// Markets is the full list of tickers and their associated configurations
	// to be stored on-chain.
	Markets map[string]Market `protobuf:"bytes,1,rep,name=markets,proto3" json:"markets" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// AggregationType is the type of aggregation that will be used to aggregate
	// the prices of the tickers.
	AggregationType AggregationType `protobuf:"varint,4,opt,name=aggregation_type,json=aggregationType,proto3,enum=slinky.mm2.v1.AggregationType" json:"aggregation_type,omitempty"`
}

func (m *MarketMap) Reset()      { *m = MarketMap{} }
func (*MarketMap) ProtoMessage() {}
func (*MarketMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_446157c095b24f63, []int{3}
}
func (m *MarketMap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarketMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MarketMap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MarketMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketMap.Merge(m, src)
}
func (m *MarketMap) XXX_Size() int {
	return m.Size()
}
func (m *MarketMap) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketMap.DiscardUnknown(m)
}

var xxx_messageInfo_MarketMap proto.InternalMessageInfo

func (m *MarketMap) GetMarkets() map[string]Market {
	if m != nil {
		return m.Markets
	}
	return nil
}

func (m *MarketMap) GetAggregationType() AggregationType {
	if m != nil {
		return m.AggregationType
	}
	return AggregationType_UNKNOWN_AGGREGATION_TYPE
}

func init() {
	proto.RegisterEnum("slinky.mm2.v1.AggregationType", AggregationType_name, AggregationType_value)
	proto.RegisterType((*Market)(nil), "slinky.mm2.v1.Market")
	proto.RegisterType((*Ticker)(nil), "slinky.mm2.v1.Ticker")
	proto.RegisterType((*ProviderConfig)(nil), "slinky.mm2.v1.ProviderConfig")
	proto.RegisterType((*MarketMap)(nil), "slinky.mm2.v1.MarketMap")
	proto.RegisterMapType((map[string]Market)(nil), "slinky.mm2.v1.MarketMap.MarketsEntry")
}

func init() { proto.RegisterFile("slinky/mm2/v1/market.proto", fileDescriptor_446157c095b24f63) }

var fileDescriptor_446157c095b24f63 = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xf6, 0xa6, 0x69, 0xda, 0x4c, 0xdb, 0xc4, 0x5a, 0xf5, 0x7d, 0xb1, 0x52, 0x70, 0xa3, 0x56,
	0xa0, 0x08, 0x68, 0xa2, 0x86, 0x0b, 0xea, 0x89, 0x34, 0xb1, 0x4a, 0x40, 0x75, 0x83, 0x1b, 0xc4,
	0xc7, 0xc5, 0xda, 0x3a, 0x1b, 0x77, 0x95, 0x78, 0x6d, 0xd9, 0x4e, 0xd4, 0xdc, 0xf8, 0x09, 0x70,
	0xe3, 0x08, 0xff, 0xa6, 0xc7, 0x1e, 0x39, 0x20, 0x84, 0x5a, 0xf1, 0x27, 0x38, 0x21, 0xaf, 0xb7,
	0x6d, 0x12, 0x2e, 0xdc, 0xe6, 0xe3, 0xd9, 0x99, 0x67, 0xe6, 0xd9, 0x5d, 0x28, 0x45, 0x43, 0xc6,
	0x07, 0x93, 0x9a, 0xe7, 0xd5, 0x6b, 0xe3, 0xdd, 0x9a, 0x47, 0xc2, 0x01, 0x8d, 0xab, 0x41, 0xe8,
	0xc7, 0x3e, 0x5e, 0x4b, 0x73, 0x55, 0xcf, 0xab, 0x57, 0xc7, 0xbb, 0xa5, 0x75, 0xd7, 0x77, 0x7d,
	0x91, 0xa9, 0x25, 0x56, 0x0a, 0x2a, 0x6d, 0xcb, 0x02, 0xf1, 0x24, 0xa0, 0x51, 0x52, 0xc2, 0x19,
	0x85, 0x21, 0xe5, 0xce, 0xc4, 0x0e, 0x08, 0x0b, 0x53, 0xd0, 0xd6, 0x27, 0x04, 0xb9, 0x43, 0x51,
	0x1a, 0xef, 0x40, 0x2e, 0x66, 0xce, 0x80, 0x86, 0x1a, 0x2a, 0xa3, 0xca, 0x4a, 0xfd, 0xbf, 0xea,
	0x4c, 0x97, 0x6a, 0x57, 0x24, 0x2d, 0x09, 0xc2, 0x26, 0xa8, 0x41, 0xe8, 0x8f, 0x59, 0x8f, 0x86,
	0xb6, 0xe3, 0xf3, 0x3e, 0x73, 0x23, 0x2d, 0x53, 0x5e, 0xa8, 0xac, 0xd4, 0xef, 0xcd, 0x1d, 0xec,
	0x48, 0x58, 0x53, 0xa0, 0xf6, 0xb3, 0xe7, 0x3f, 0x36, 0x15, 0xab, 0x18, 0xcc, 0x44, 0xa3, 0xbd,
	0xe5, 0xcf, 0x5f, 0x36, 0x95, 0x0f, 0xdf, 0xcb, 0xca, 0xd6, 0x2f, 0x04, 0xb9, 0xb4, 0x19, 0x7e,
	0x0e, 0x6b, 0x33, 0xac, 0x25, 0xb5, 0x9b, 0x0e, 0x62, 0xb6, 0xa4, 0x47, 0x53, 0xa2, 0x3a, 0x84,
	0x85, 0xb2, 0xc3, 0xaa, 0x33, 0x15, 0xc3, 0x25, 0x58, 0xee, 0x51, 0x87, 0x79, 0x64, 0x98, 0xd0,
	0x44, 0x95, 0xac, 0x75, 0xe3, 0xe3, 0xc7, 0x80, 0x3d, 0xc6, 0xed, 0xa9, 0x71, 0x46, 0x3c, 0xd6,
	0x16, 0x04, 0x4a, 0xf5, 0x18, 0xbf, 0x1d, 0x60, 0xc4, 0x63, 0xac, 0xc1, 0x12, 0xe5, 0xe4, 0x64,
	0x48, 0x7b, 0x5a, 0xa1, 0x8c, 0x2a, 0xcb, 0xd6, 0xb5, 0x8b, 0xb7, 0x61, 0xcd, 0xa3, 0x31, 0xe9,
	0x91, 0x98, 0xd8, 0x2f, 0x8e, 0x8f, 0x4c, 0xad, 0x58, 0x46, 0x95, 0xbc, 0xb5, 0x7a, 0x1d, 0x4c,
	0x62, 0x53, 0x73, 0x7e, 0x45, 0x50, 0x98, 0xdd, 0x0d, 0xc6, 0x90, 0xe5, 0xc4, 0xa3, 0x62, 0xcc,
	0xbc, 0x25, 0x6c, 0x5c, 0x01, 0xd5, 0xef, 0xf7, 0x6d, 0xe7, 0x94, 0x30, 0x6e, 0x4b, 0x85, 0x32,
	0x22, 0x5f, 0xf0, 0xfb, 0xfd, 0x66, 0x12, 0x96, 0xdb, 0x5a, 0x87, 0x45, 0xc6, 0x7b, 0xf4, 0x4c,
	0x50, 0xcf, 0x5b, 0xa9, 0x83, 0xff, 0x87, 0x1c, 0xe3, 0x63, 0x1a, 0xc6, 0x5a, 0x56, 0xd0, 0x95,
	0xde, 0x3f, 0xb1, 0xdd, 0xfa, 0x8d, 0x20, 0x9f, 0xde, 0x8f, 0x43, 0x12, 0x60, 0x03, 0x96, 0xd2,
	0x7b, 0x18, 0x69, 0x48, 0x48, 0x7d, 0x7f, 0x4e, 0xea, 0x1b, 0xa8, 0xb4, 0x22, 0x83, 0xc7, 0xe1,
	0x44, 0x0a, 0x72, 0x7d, 0x16, 0xb7, 0x41, 0x25, 0xae, 0x1b, 0x52, 0x97, 0xc4, 0xcc, 0xe7, 0x76,
	0x22, 0xa2, 0xe0, 0x56, 0xa8, 0xeb, 0x73, 0xf5, 0x1a, 0xb7, 0xb0, 0xee, 0x24, 0xa0, 0x56, 0x91,
	0xcc, 0x06, 0x4a, 0xaf, 0x60, 0x75, 0xba, 0x13, 0x56, 0x61, 0x61, 0x40, 0x27, 0x72, 0x7f, 0x89,
	0x89, 0x1f, 0xc1, 0xe2, 0x98, 0x0c, 0x47, 0x54, 0xec, 0xec, 0xef, 0x5b, 0x9d, 0x9e, 0xb6, 0x52,
	0xcc, 0x5e, 0xe6, 0x29, 0xba, 0x15, 0xe8, 0xa1, 0x07, 0xc5, 0x39, 0x02, 0xf8, 0x2e, 0x68, 0xaf,
	0xcd, 0x97, 0xe6, 0xd1, 0x1b, 0xd3, 0x6e, 0x1c, 0x1c, 0x58, 0xc6, 0x41, 0xa3, 0xdb, 0x3e, 0x32,
	0xed, 0xee, 0xbb, 0x8e, 0xa1, 0x2a, 0x78, 0x03, 0xee, 0xb4, 0xcd, 0x96, 0xf1, 0xd6, 0xee, 0x58,
	0xed, 0xa6, 0x31, 0x8d, 0x50, 0x11, 0xde, 0x84, 0x8d, 0xe3, 0x6e, 0xc3, 0x6c, 0x35, 0xac, 0x96,
	0x7d, 0x68, 0xb4, 0xda, 0x8d, 0x99, 0x12, 0x6a, 0x66, 0xff, 0xd9, 0xf9, 0xa5, 0x8e, 0x2e, 0x2e,
	0x75, 0xf4, 0xf3, 0x52, 0x47, 0x1f, 0xaf, 0x74, 0xe5, 0xe2, 0x4a, 0x57, 0xbe, 0x5d, 0xe9, 0xca,
	0xfb, 0x07, 0x2e, 0x8b, 0x4f, 0x47, 0x27, 0x55, 0xc7, 0xf7, 0x6a, 0xd1, 0x80, 0x05, 0x3b, 0x1e,
	0x1d, 0xd7, 0xe4, 0xf3, 0x3e, 0x13, 0x3f, 0x84, 0x78, 0x08, 0x27, 0x39, 0xf1, 0xa8, 0x9f, 0xfc,
	0x09, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x88, 0x59, 0xe8, 0x3c, 0x04, 0x00, 0x00,
}

func (m *Market) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Market) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Market) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ProviderConfigs) > 0 {
		for iNdEx := len(m.ProviderConfigs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProviderConfigs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMarket(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Ticker != nil {
		{
			size, err := m.Ticker.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMarket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Ticker) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ticker) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ticker) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metadata_JSON) > 0 {
		i -= len(m.Metadata_JSON)
		copy(dAtA[i:], m.Metadata_JSON)
		i = encodeVarintMarket(dAtA, i, uint64(len(m.Metadata_JSON)))
		i--
		dAtA[i] = 0x7a
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x70
	}
	if m.MinProviderCount != 0 {
		i = encodeVarintMarket(dAtA, i, uint64(m.MinProviderCount))
		i--
		dAtA[i] = 0x18
	}
	if m.Decimals != 0 {
		i = encodeVarintMarket(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.CurrencyPair.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMarket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ProviderConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProviderConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProviderConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metadata_JSON) > 0 {
		i -= len(m.Metadata_JSON)
		copy(dAtA[i:], m.Metadata_JSON)
		i = encodeVarintMarket(dAtA, i, uint64(len(m.Metadata_JSON)))
		i--
		dAtA[i] = 0x7a
	}
	if m.Invert {
		i--
		if m.Invert {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintMarket(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OffChainTicker) > 0 {
		i -= len(m.OffChainTicker)
		copy(dAtA[i:], m.OffChainTicker)
		i = encodeVarintMarket(dAtA, i, uint64(len(m.OffChainTicker)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMarket(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MarketMap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarketMap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarketMap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AggregationType != 0 {
		i = encodeVarintMarket(dAtA, i, uint64(m.AggregationType))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Markets) > 0 {
		for k := range m.Markets {
			v := m.Markets[k]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMarket(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintMarket(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintMarket(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintMarket(dAtA []byte, offset int, v uint64) int {
	offset -= sovMarket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Market) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ticker != nil {
		l = m.Ticker.Size()
		n += 1 + l + sovMarket(uint64(l))
	}
	if len(m.ProviderConfigs) > 0 {
		for _, e := range m.ProviderConfigs {
			l = e.Size()
			n += 1 + l + sovMarket(uint64(l))
		}
	}
	return n
}

func (m *Ticker) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CurrencyPair.Size()
	n += 1 + l + sovMarket(uint64(l))
	if m.Decimals != 0 {
		n += 1 + sovMarket(uint64(m.Decimals))
	}
	if m.MinProviderCount != 0 {
		n += 1 + sovMarket(uint64(m.MinProviderCount))
	}
	if m.Enabled {
		n += 2
	}
	l = len(m.Metadata_JSON)
	if l > 0 {
		n += 1 + l + sovMarket(uint64(l))
	}
	return n
}

func (m *ProviderConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMarket(uint64(l))
	}
	l = len(m.OffChainTicker)
	if l > 0 {
		n += 1 + l + sovMarket(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovMarket(uint64(l))
	}
	if m.Invert {
		n += 2
	}
	l = len(m.Metadata_JSON)
	if l > 0 {
		n += 1 + l + sovMarket(uint64(l))
	}
	return n
}

func (m *MarketMap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Markets) > 0 {
		for k, v := range m.Markets {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovMarket(uint64(len(k))) + 1 + l + sovMarket(uint64(l))
			n += mapEntrySize + 1 + sovMarket(uint64(mapEntrySize))
		}
	}
	if m.AggregationType != 0 {
		n += 1 + sovMarket(uint64(m.AggregationType))
	}
	return n
}

func sovMarket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMarket(x uint64) (n int) {
	return sovMarket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Market) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarket
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
			return fmt.Errorf("proto: Market: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Market: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ticker", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
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
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ticker == nil {
				m.Ticker = &Ticker{}
			}
			if err := m.Ticker.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderConfigs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
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
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProviderConfigs = append(m.ProviderConfigs, ProviderConfig{})
			if err := m.ProviderConfigs[len(m.ProviderConfigs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMarket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMarket
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
func (m *Ticker) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarket
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
			return fmt.Errorf("proto: Ticker: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ticker: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrencyPair", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
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
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrencyPair.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinProviderCount", wireType)
			}
			m.MinProviderCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinProviderCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
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
			m.Enabled = bool(v != 0)
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata_JSON", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
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
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata_JSON = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMarket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMarket
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
func (m *ProviderConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarket
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
			return fmt.Errorf("proto: ProviderConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProviderConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
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
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OffChainTicker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
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
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OffChainTicker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
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
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Invert", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
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
			m.Invert = bool(v != 0)
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata_JSON", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
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
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata_JSON = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMarket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMarket
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
func (m *MarketMap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarket
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
			return fmt.Errorf("proto: MarketMap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarketMap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Markets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
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
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Markets == nil {
				m.Markets = make(map[string]Market)
			}
			var mapkey string
			mapvalue := &Market{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMarket
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMarket
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthMarket
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthMarket
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMarket
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthMarket
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthMarket
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Market{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipMarket(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthMarket
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Markets[mapkey] = *mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregationType", wireType)
			}
			m.AggregationType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AggregationType |= AggregationType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMarket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMarket
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
func skipMarket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMarket
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
					return 0, ErrIntOverflowMarket
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
					return 0, ErrIntOverflowMarket
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
				return 0, ErrInvalidLengthMarket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMarket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMarket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMarket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMarket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMarket = fmt.Errorf("proto: unexpected end of group")
)
