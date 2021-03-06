// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos.proto

package goctp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Snapshot struct {
	// /交易日
	TradingDay string `protobuf:"bytes,1,opt,name=TradingDay,proto3" json:"TradingDay,omitempty"`
	// /合约代码
	InstrumentID string `protobuf:"bytes,2,opt,name=InstrumentID,proto3" json:"InstrumentID,omitempty"`
	// /交易所代码
	ExchangeID string `protobuf:"bytes,3,opt,name=ExchangeID,proto3" json:"ExchangeID,omitempty"`
	// /合约在交易所的代码
	ExchangeInstID string `protobuf:"bytes,4,opt,name=ExchangeInstID,proto3" json:"ExchangeInstID,omitempty"`
	// /最新价
	LastPrice float64 `protobuf:"fixed64,5,opt,name=LastPrice,proto3" json:"LastPrice,omitempty"`
	// /上次结算价
	PreSettlementPrice float64 `protobuf:"fixed64,6,opt,name=PreSettlementPrice,proto3" json:"PreSettlementPrice,omitempty"`
	// /昨收盘
	PreClosePrice float64 `protobuf:"fixed64,7,opt,name=PreClosePrice,proto3" json:"PreClosePrice,omitempty"`
	// /昨持仓量
	PreOpenInterest float64 `protobuf:"fixed64,8,opt,name=PreOpenInterest,proto3" json:"PreOpenInterest,omitempty"`
	// /今开盘
	OpenPrice float64 `protobuf:"fixed64,9,opt,name=OpenPrice,proto3" json:"OpenPrice,omitempty"`
	// /最高价
	HighestPrice float64 `protobuf:"fixed64,10,opt,name=HighestPrice,proto3" json:"HighestPrice,omitempty"`
	// /最低价
	LowestPrice float64 `protobuf:"fixed64,11,opt,name=LowestPrice,proto3" json:"LowestPrice,omitempty"`
	// /数量
	Volume int64 `protobuf:"varint,12,opt,name=Volume,proto3" json:"Volume,omitempty"`
	// /成交金额
	Turnover float64 `protobuf:"fixed64,13,opt,name=Turnover,proto3" json:"Turnover,omitempty"`
	// /持仓量
	OpenInterest float64 `protobuf:"fixed64,14,opt,name=OpenInterest,proto3" json:"OpenInterest,omitempty"`
	// /今收盘
	ClosePrice float64 `protobuf:"fixed64,15,opt,name=ClosePrice,proto3" json:"ClosePrice,omitempty"`
	// /本次结算价
	SettlementPrice float64 `protobuf:"fixed64,16,opt,name=SettlementPrice,proto3" json:"SettlementPrice,omitempty"`
	// /涨停板价
	UpperLimitPrice float64 `protobuf:"fixed64,17,opt,name=UpperLimitPrice,proto3" json:"UpperLimitPrice,omitempty"`
	// /跌停板价
	LowerLimitPrice float64 `protobuf:"fixed64,18,opt,name=LowerLimitPrice,proto3" json:"LowerLimitPrice,omitempty"`
	// /昨虚实度
	PreDelta float64 `protobuf:"fixed64,19,opt,name=PreDelta,proto3" json:"PreDelta,omitempty"`
	// /今虚实度
	CurrDelta float64 `protobuf:"fixed64,20,opt,name=CurrDelta,proto3" json:"CurrDelta,omitempty"`
	// /最后修改时间
	UpdateTime string `protobuf:"bytes,21,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	// /最后修改毫秒
	UpdateMillisec int64 `protobuf:"varint,22,opt,name=UpdateMillisec,proto3" json:"UpdateMillisec,omitempty"`
	// /申买价一
	BidPrice1 float64 `protobuf:"fixed64,23,opt,name=BidPrice1,proto3" json:"BidPrice1,omitempty"`
	// /申买量一
	BidVolume1 int64 `protobuf:"varint,24,opt,name=BidVolume1,proto3" json:"BidVolume1,omitempty"`
	// /申卖价一
	AskPrice1 float64 `protobuf:"fixed64,25,opt,name=AskPrice1,proto3" json:"AskPrice1,omitempty"`
	// /申卖量一
	AskVolume1 int64 `protobuf:"varint,26,opt,name=AskVolume1,proto3" json:"AskVolume1,omitempty"`
	// /申买价二
	BidPrice2 float64 `protobuf:"fixed64,27,opt,name=BidPrice2,proto3" json:"BidPrice2,omitempty"`
	// /申买量二
	BidVolume2 int64 `protobuf:"varint,28,opt,name=BidVolume2,proto3" json:"BidVolume2,omitempty"`
	// /申卖价二
	AskPrice2 float64 `protobuf:"fixed64,29,opt,name=AskPrice2,proto3" json:"AskPrice2,omitempty"`
	// /申卖量二
	AskVolume2 int64 `protobuf:"varint,30,opt,name=AskVolume2,proto3" json:"AskVolume2,omitempty"`
	// /申买价三
	BidPrice3 float64 `protobuf:"fixed64,31,opt,name=BidPrice3,proto3" json:"BidPrice3,omitempty"`
	// /申买量三
	BidVolume3 int64 `protobuf:"varint,32,opt,name=BidVolume3,proto3" json:"BidVolume3,omitempty"`
	// /申卖价三
	AskPrice3 float64 `protobuf:"fixed64,33,opt,name=AskPrice3,proto3" json:"AskPrice3,omitempty"`
	// /申卖量三
	AskVolume3 int64 `protobuf:"varint,34,opt,name=AskVolume3,proto3" json:"AskVolume3,omitempty"`
	// /申买价四
	BidPrice4 float64 `protobuf:"fixed64,35,opt,name=BidPrice4,proto3" json:"BidPrice4,omitempty"`
	// /申买量四
	BidVolume4 int64 `protobuf:"varint,36,opt,name=BidVolume4,proto3" json:"BidVolume4,omitempty"`
	// /申卖价四
	AskPrice4 float64 `protobuf:"fixed64,37,opt,name=AskPrice4,proto3" json:"AskPrice4,omitempty"`
	// /申卖量四
	AskVolume4 int64 `protobuf:"varint,38,opt,name=AskVolume4,proto3" json:"AskVolume4,omitempty"`
	// /申买价五
	BidPrice5 float64 `protobuf:"fixed64,39,opt,name=BidPrice5,proto3" json:"BidPrice5,omitempty"`
	// /申买量五
	BidVolume5 int64 `protobuf:"varint,40,opt,name=BidVolume5,proto3" json:"BidVolume5,omitempty"`
	// /申卖价五
	AskPrice5 float64 `protobuf:"fixed64,41,opt,name=AskPrice5,proto3" json:"AskPrice5,omitempty"`
	// /申卖量五
	AskVolume5 int64 `protobuf:"varint,42,opt,name=AskVolume5,proto3" json:"AskVolume5,omitempty"`
	// /当日均价
	AveragePrice float64 `protobuf:"fixed64,43,opt,name=AveragePrice,proto3" json:"AveragePrice,omitempty"`
	// /业务日期
	ActionDay            string   `protobuf:"bytes,44,opt,name=ActionDay,proto3" json:"ActionDay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_protos_fe55b32bc3d6069e, []int{0}
}
func (m *Snapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Snapshot.Unmarshal(m, b)
}
func (m *Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Snapshot.Marshal(b, m, deterministic)
}
func (dst *Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snapshot.Merge(dst, src)
}
func (m *Snapshot) XXX_Size() int {
	return xxx_messageInfo_Snapshot.Size(m)
}
func (m *Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_Snapshot proto.InternalMessageInfo

func (m *Snapshot) GetTradingDay() string {
	if m != nil {
		return m.TradingDay
	}
	return ""
}

func (m *Snapshot) GetInstrumentID() string {
	if m != nil {
		return m.InstrumentID
	}
	return ""
}

func (m *Snapshot) GetExchangeID() string {
	if m != nil {
		return m.ExchangeID
	}
	return ""
}

func (m *Snapshot) GetExchangeInstID() string {
	if m != nil {
		return m.ExchangeInstID
	}
	return ""
}

func (m *Snapshot) GetLastPrice() float64 {
	if m != nil {
		return m.LastPrice
	}
	return 0
}

func (m *Snapshot) GetPreSettlementPrice() float64 {
	if m != nil {
		return m.PreSettlementPrice
	}
	return 0
}

func (m *Snapshot) GetPreClosePrice() float64 {
	if m != nil {
		return m.PreClosePrice
	}
	return 0
}

func (m *Snapshot) GetPreOpenInterest() float64 {
	if m != nil {
		return m.PreOpenInterest
	}
	return 0
}

func (m *Snapshot) GetOpenPrice() float64 {
	if m != nil {
		return m.OpenPrice
	}
	return 0
}

func (m *Snapshot) GetHighestPrice() float64 {
	if m != nil {
		return m.HighestPrice
	}
	return 0
}

func (m *Snapshot) GetLowestPrice() float64 {
	if m != nil {
		return m.LowestPrice
	}
	return 0
}

func (m *Snapshot) GetVolume() int64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *Snapshot) GetTurnover() float64 {
	if m != nil {
		return m.Turnover
	}
	return 0
}

func (m *Snapshot) GetOpenInterest() float64 {
	if m != nil {
		return m.OpenInterest
	}
	return 0
}

func (m *Snapshot) GetClosePrice() float64 {
	if m != nil {
		return m.ClosePrice
	}
	return 0
}

func (m *Snapshot) GetSettlementPrice() float64 {
	if m != nil {
		return m.SettlementPrice
	}
	return 0
}

func (m *Snapshot) GetUpperLimitPrice() float64 {
	if m != nil {
		return m.UpperLimitPrice
	}
	return 0
}

func (m *Snapshot) GetLowerLimitPrice() float64 {
	if m != nil {
		return m.LowerLimitPrice
	}
	return 0
}

func (m *Snapshot) GetPreDelta() float64 {
	if m != nil {
		return m.PreDelta
	}
	return 0
}

func (m *Snapshot) GetCurrDelta() float64 {
	if m != nil {
		return m.CurrDelta
	}
	return 0
}

func (m *Snapshot) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *Snapshot) GetUpdateMillisec() int64 {
	if m != nil {
		return m.UpdateMillisec
	}
	return 0
}

func (m *Snapshot) GetBidPrice1() float64 {
	if m != nil {
		return m.BidPrice1
	}
	return 0
}

func (m *Snapshot) GetBidVolume1() int64 {
	if m != nil {
		return m.BidVolume1
	}
	return 0
}

func (m *Snapshot) GetAskPrice1() float64 {
	if m != nil {
		return m.AskPrice1
	}
	return 0
}

func (m *Snapshot) GetAskVolume1() int64 {
	if m != nil {
		return m.AskVolume1
	}
	return 0
}

func (m *Snapshot) GetBidPrice2() float64 {
	if m != nil {
		return m.BidPrice2
	}
	return 0
}

func (m *Snapshot) GetBidVolume2() int64 {
	if m != nil {
		return m.BidVolume2
	}
	return 0
}

func (m *Snapshot) GetAskPrice2() float64 {
	if m != nil {
		return m.AskPrice2
	}
	return 0
}

func (m *Snapshot) GetAskVolume2() int64 {
	if m != nil {
		return m.AskVolume2
	}
	return 0
}

func (m *Snapshot) GetBidPrice3() float64 {
	if m != nil {
		return m.BidPrice3
	}
	return 0
}

func (m *Snapshot) GetBidVolume3() int64 {
	if m != nil {
		return m.BidVolume3
	}
	return 0
}

func (m *Snapshot) GetAskPrice3() float64 {
	if m != nil {
		return m.AskPrice3
	}
	return 0
}

func (m *Snapshot) GetAskVolume3() int64 {
	if m != nil {
		return m.AskVolume3
	}
	return 0
}

func (m *Snapshot) GetBidPrice4() float64 {
	if m != nil {
		return m.BidPrice4
	}
	return 0
}

func (m *Snapshot) GetBidVolume4() int64 {
	if m != nil {
		return m.BidVolume4
	}
	return 0
}

func (m *Snapshot) GetAskPrice4() float64 {
	if m != nil {
		return m.AskPrice4
	}
	return 0
}

func (m *Snapshot) GetAskVolume4() int64 {
	if m != nil {
		return m.AskVolume4
	}
	return 0
}

func (m *Snapshot) GetBidPrice5() float64 {
	if m != nil {
		return m.BidPrice5
	}
	return 0
}

func (m *Snapshot) GetBidVolume5() int64 {
	if m != nil {
		return m.BidVolume5
	}
	return 0
}

func (m *Snapshot) GetAskPrice5() float64 {
	if m != nil {
		return m.AskPrice5
	}
	return 0
}

func (m *Snapshot) GetAskVolume5() int64 {
	if m != nil {
		return m.AskVolume5
	}
	return 0
}

func (m *Snapshot) GetAveragePrice() float64 {
	if m != nil {
		return m.AveragePrice
	}
	return 0
}

func (m *Snapshot) GetActionDay() string {
	if m != nil {
		return m.ActionDay
	}
	return ""
}

func init() {
	proto.RegisterType((*Snapshot)(nil), "goctp.Snapshot")
}

func init() { proto.RegisterFile("protos.proto", fileDescriptor_protos_fe55b32bc3d6069e) }

var fileDescriptor_protos_fe55b32bc3d6069e = []byte{
	// 579 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0xed, 0x6e, 0x12, 0x4d,
	0x14, 0xc7, 0xc3, 0xd3, 0xa7, 0x15, 0x06, 0x5a, 0x74, 0xd4, 0x7a, 0xac, 0x15, 0x11, 0x6b, 0x5d,
	0x5f, 0xd2, 0xa4, 0xfb, 0x72, 0x01, 0xb4, 0x98, 0x48, 0x82, 0x91, 0x50, 0xea, 0xf7, 0x15, 0x4e,
	0x60, 0xd3, 0x65, 0x77, 0x33, 0x3b, 0x54, 0xbd, 0x39, 0xaf, 0xcd, 0xcc, 0x0c, 0x2c, 0x73, 0x26,
	0x7e, 0x22, 0xf3, 0xdb, 0xff, 0x99, 0xdf, 0x99, 0xc3, 0xce, 0xb2, 0x56, 0x21, 0x72, 0x99, 0x97,
	0x17, 0xfa, 0x87, 0xef, 0x2f, 0xf2, 0x99, 0x2c, 0x7a, 0x7f, 0x9a, 0xac, 0x7e, 0x93, 0xc5, 0x45,
	0xb9, 0xcc, 0x25, 0xef, 0x30, 0x36, 0x15, 0xf1, 0x3c, 0xc9, 0x16, 0x83, 0xf8, 0x37, 0xd4, 0xba,
	0x35, 0xaf, 0x31, 0xb1, 0x08, 0xef, 0xb1, 0xd6, 0x30, 0x2b, 0xa5, 0x58, 0xaf, 0x30, 0x93, 0xc3,
	0x01, 0xfc, 0xa7, 0x13, 0x84, 0xa9, 0x3d, 0x3e, 0xff, 0x9a, 0x2d, 0xe3, 0x6c, 0x81, 0xc3, 0x01,
	0xec, 0x99, 0x3d, 0x76, 0x84, 0x9f, 0xb3, 0xa3, 0x6a, 0x95, 0x95, 0x6a, 0x97, 0xff, 0x75, 0xc6,
	0xa1, 0xfc, 0x94, 0x35, 0x46, 0x71, 0x29, 0xc7, 0x22, 0x99, 0x21, 0xec, 0x77, 0x6b, 0x5e, 0x6d,
	0xb2, 0x03, 0xfc, 0x82, 0xf1, 0xb1, 0xc0, 0x1b, 0x94, 0x32, 0x45, 0x25, 0x36, 0xb1, 0x03, 0x1d,
	0xfb, 0xc7, 0x13, 0x7e, 0xc6, 0x0e, 0xc7, 0x02, 0xaf, 0xd3, 0xbc, 0x44, 0x13, 0x7d, 0xa0, 0xa3,
	0x14, 0x72, 0x8f, 0xb5, 0xc7, 0x02, 0xbf, 0x15, 0x98, 0x0d, 0x33, 0x89, 0x02, 0x4b, 0x09, 0x75,
	0x9d, 0x73, 0xb1, 0xea, 0x4e, 0xad, 0xcd, 0x5e, 0x0d, 0xd3, 0x5d, 0x05, 0xd4, 0x9c, 0xbe, 0x24,
	0x8b, 0x25, 0x6e, 0xdb, 0x67, 0x3a, 0x40, 0x18, 0xef, 0xb2, 0xe6, 0x28, 0xff, 0x59, 0x45, 0x9a,
	0x3a, 0x62, 0x23, 0x7e, 0xcc, 0x0e, 0xbe, 0xe7, 0xe9, 0x7a, 0x85, 0xd0, 0xea, 0xd6, 0xbc, 0xbd,
	0xc9, 0x66, 0xc5, 0x4f, 0x58, 0x7d, 0xba, 0x16, 0x59, 0x7e, 0x8f, 0x02, 0x0e, 0x75, 0x59, 0xb5,
	0x56, 0x66, 0xd2, 0xfe, 0x91, 0x31, 0x93, 0xde, 0x3b, 0x8c, 0x59, 0x83, 0x68, 0xeb, 0x04, 0xa3,
	0x53, 0x70, 0x07, 0xfb, 0xd0, 0x4c, 0xc1, 0x9d, 0xaa, 0xc7, 0xda, 0xb7, 0x45, 0x81, 0x62, 0x94,
	0xac, 0x92, 0x4d, 0xf2, 0x91, 0x49, 0x3a, 0x58, 0x25, 0xd5, 0xd1, 0xec, 0x24, 0x37, 0x49, 0x07,
	0xab, 0xd3, 0x8d, 0x05, 0x0e, 0x30, 0x95, 0x31, 0x3c, 0x36, 0xa7, 0xdb, 0xae, 0xd5, 0xd4, 0xaf,
	0xd7, 0x42, 0x98, 0x87, 0x4f, 0xcc, 0xd4, 0x2b, 0xa0, 0xce, 0x75, 0x5b, 0xcc, 0x63, 0x89, 0xd3,
	0x64, 0x85, 0xf0, 0xd4, 0xbc, 0x79, 0x3b, 0xa2, 0xde, 0x3c, 0xb3, 0xfa, 0x9a, 0xa4, 0x69, 0x52,
	0xe2, 0x0c, 0x8e, 0xf5, 0x5c, 0x1d, 0xaa, 0x2c, 0x57, 0xc9, 0x5c, 0x77, 0x73, 0x09, 0xcf, 0x8c,
	0xa5, 0x02, 0xca, 0x72, 0x95, 0xcc, 0xcd, 0x5f, 0x71, 0x09, 0xa0, 0x77, 0xb0, 0x88, 0xaa, 0xee,
	0x97, 0x77, 0x9b, 0xea, 0xe7, 0xa6, 0xba, 0x02, 0xaa, 0xba, 0x5f, 0xde, 0x6d, 0xab, 0x4f, 0x4c,
	0xf5, 0x8e, 0xd8, 0x6e, 0x1f, 0x5e, 0x50, 0xb7, 0x4f, 0xdc, 0x3e, 0x9c, 0x3a, 0x6e, 0xdf, 0x76,
	0xfb, 0xf0, 0x92, 0xba, 0x7d, 0xe2, 0xf6, 0xa1, 0xe3, 0xb8, 0x7d, 0xdb, 0x1d, 0xc0, 0x2b, 0xea,
	0x0e, 0x88, 0x3b, 0x80, 0xae, 0xe3, 0x0e, 0x6c, 0x77, 0x00, 0xaf, 0xa9, 0x3b, 0x20, 0xee, 0x00,
	0x7a, 0x8e, 0x3b, 0xb0, 0xdd, 0x21, 0xbc, 0xa1, 0xee, 0x90, 0xb8, 0x43, 0x38, 0x73, 0xdc, 0xa1,
	0xed, 0x0e, 0xe1, 0x2d, 0x75, 0x87, 0xc4, 0x1d, 0xc2, 0xb9, 0xe3, 0x0e, 0x6d, 0x77, 0x04, 0xef,
	0xa8, 0x3b, 0x22, 0xee, 0x08, 0x3c, 0xc7, 0x1d, 0xd9, 0xee, 0x08, 0xde, 0x53, 0x77, 0x44, 0xdc,
	0x11, 0x7c, 0x70, 0xdc, 0x91, 0xba, 0xaf, 0xfd, 0x7b, 0x14, 0xf1, 0x62, 0x73, 0x1b, 0x3f, 0x9a,
	0xfb, 0x6a, 0x33, 0x6d, 0x98, 0xc9, 0x24, 0xcf, 0xd4, 0x47, 0xf9, 0x93, 0x7e, 0xad, 0x77, 0xe0,
	0xc7, 0x81, 0xfe, 0x9c, 0x07, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x28, 0x38, 0x08, 0xde,
	0x05, 0x00, 0x00,
}
