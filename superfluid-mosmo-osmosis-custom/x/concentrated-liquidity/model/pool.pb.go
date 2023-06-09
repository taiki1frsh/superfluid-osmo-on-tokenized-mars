// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/concentrated-liquidity/pool.proto

// This is a legacy package that requires additional migration logic
// in order to use the correct package. Decision made to use legacy package path
// until clear steps for migration logic and the unknowns for state breaking are
// investigated for changing proto package.

package model

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

type Pool struct {
	// pool's address holding all liquidity tokens.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// address holding the incentives liquidity.
	IncentivesAddress string `protobuf:"bytes,2,opt,name=incentives_address,json=incentivesAddress,proto3" json:"incentives_address,omitempty" yaml:"incentives_address"`
	Id                uint64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// Amount of total liquidity
	CurrentTickLiquidity github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=current_tick_liquidity,json=currentTickLiquidity,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"current_tick_liquidity" yaml:"current_tick_liquidity"`
	Token0               string                                 `protobuf:"bytes,5,opt,name=token0,proto3" json:"token0,omitempty"`
	Token1               string                                 `protobuf:"bytes,6,opt,name=token1,proto3" json:"token1,omitempty"`
	CurrentSqrtPrice     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=current_sqrt_price,json=currentSqrtPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"current_sqrt_price" yaml:"spot_price"`
	CurrentTick          github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=current_tick,json=currentTick,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"current_tick" yaml:"current_tick"`
	// tick_spacing must be one of the authorized_tick_spacing values set in the
	// concentrated-liquidity parameters
	TickSpacing        uint64                                 `protobuf:"varint,9,opt,name=tick_spacing,json=tickSpacing,proto3" json:"tick_spacing,omitempty" yaml:"tick_spacing"`
	ExponentAtPriceOne github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,10,opt,name=exponent_at_price_one,json=exponentAtPriceOne,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"exponent_at_price_one" yaml:"exponent_at_price_one"`
	// swap_fee is the ratio that is charged on the amount of token in.
	SwapFee github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,11,opt,name=swap_fee,json=swapFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"swap_fee" yaml:"swap_fee"`
	// last_liquidity_update is the last time either the pool liquidity or the
	// active tick changed
	LastLiquidityUpdate time.Time `protobuf:"bytes,12,opt,name=last_liquidity_update,json=lastLiquidityUpdate,proto3,stdtime" json:"last_liquidity_update" yaml:"last_liquidity_update"`
}

func (m *Pool) Reset()      { *m = Pool{} }
func (*Pool) ProtoMessage() {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_3526ea5373d96c9a, []int{0}
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

func init() {
	proto.RegisterType((*Pool)(nil), "osmosis.concentratedliquidity.v1beta1.Pool")
}

func init() {
	proto.RegisterFile("osmosis/concentrated-liquidity/pool.proto", fileDescriptor_3526ea5373d96c9a)
}

var fileDescriptor_3526ea5373d96c9a = []byte{
	// 626 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x4e, 0xdb, 0x40,
	0x10, 0xb6, 0x29, 0x10, 0xd8, 0x20, 0x5a, 0x96, 0x9f, 0x1a, 0x54, 0x6c, 0x64, 0xa9, 0x55, 0x2a,
	0x35, 0x76, 0xd3, 0xaa, 0x17, 0x6e, 0xa4, 0x3f, 0x12, 0x12, 0x2a, 0xc8, 0xd0, 0x4b, 0x85, 0x64,
	0x39, 0xf6, 0x62, 0x56, 0xb1, 0xbd, 0x8e, 0x77, 0x43, 0xc9, 0xb1, 0x87, 0x4a, 0x3d, 0x72, 0xec,
	0x91, 0x87, 0xe8, 0x43, 0xa0, 0x9e, 0x50, 0x4f, 0x55, 0x0f, 0x6e, 0x95, 0xbc, 0x41, 0x9e, 0xa0,
	0xf2, 0x7a, 0x9d, 0x58, 0x22, 0x3d, 0xe4, 0x64, 0xcf, 0x37, 0x33, 0xdf, 0x7c, 0x33, 0xbb, 0xb3,
	0xe0, 0x29, 0xa1, 0x21, 0xa1, 0x98, 0x9a, 0x2e, 0x89, 0x5c, 0x14, 0xb1, 0xc4, 0x61, 0xc8, 0xab,
	0x07, 0xb8, 0xd3, 0xc5, 0x1e, 0x66, 0x3d, 0x33, 0x26, 0x24, 0x30, 0xe2, 0x84, 0x30, 0x02, 0x1f,
	0x8b, 0x50, 0xa3, 0x1c, 0x3a, 0x8a, 0x34, 0x2e, 0x1a, 0x2d, 0xc4, 0x9c, 0xc6, 0xd6, 0xa6, 0xcb,
	0xe3, 0x6c, 0x9e, 0x64, 0xe6, 0x46, 0xce, 0xb0, 0xb5, 0xe6, 0x13, 0x9f, 0xe4, 0x78, 0xf6, 0x27,
	0x50, 0xcd, 0x27, 0xc4, 0x0f, 0x90, 0xc9, 0xad, 0x56, 0xf7, 0xcc, 0x64, 0x38, 0x44, 0x94, 0x39,
	0x61, 0x9c, 0x07, 0xe8, 0x3f, 0x2b, 0x60, 0xf6, 0x88, 0x90, 0x00, 0x3e, 0x03, 0x15, 0xc7, 0xf3,
	0x12, 0x44, 0xa9, 0x22, 0xef, 0xc8, 0xb5, 0xc5, 0x26, 0x1c, 0xa6, 0xda, 0x72, 0xcf, 0x09, 0x83,
	0x5d, 0x5d, 0x38, 0x74, 0xab, 0x08, 0x81, 0x07, 0x00, 0x62, 0x2e, 0x14, 0x5f, 0x20, 0x6a, 0x17,
	0x89, 0x33, 0x3c, 0x71, 0x7b, 0x98, 0x6a, 0x9b, 0x79, 0xe2, 0xdd, 0x18, 0xdd, 0x5a, 0x19, 0x83,
	0x7b, 0x82, 0x6d, 0x19, 0xcc, 0x60, 0x4f, 0xb9, 0xb7, 0x23, 0xd7, 0x66, 0xad, 0x19, 0xec, 0xc1,
	0x2f, 0x32, 0xd8, 0x70, 0xbb, 0x49, 0x82, 0x22, 0x66, 0x33, 0xec, 0xb6, 0xed, 0xd1, 0x24, 0x94,
	0x59, 0x5e, 0xe2, 0xf0, 0x26, 0xd5, 0xa4, 0xdf, 0xa9, 0xf6, 0xc4, 0xc7, 0xec, 0xbc, 0xdb, 0x32,
	0x5c, 0x12, 0x8a, 0x69, 0x88, 0x4f, 0x9d, 0x7a, 0x6d, 0x93, 0xf5, 0x62, 0x44, 0x8d, 0x37, 0xc8,
	0x1d, 0xa6, 0xda, 0x76, 0x2e, 0x68, 0x32, 0xab, 0x6e, 0xad, 0x09, 0xc7, 0x09, 0x76, 0xdb, 0x07,
	0x05, 0x0c, 0x37, 0xc0, 0x3c, 0x23, 0x6d, 0x14, 0x3d, 0x57, 0xe6, 0xb2, 0xb2, 0x96, 0xb0, 0x46,
	0x78, 0x43, 0x99, 0x2f, 0xe1, 0x0d, 0xd8, 0x01, 0xb0, 0x28, 0x40, 0x3b, 0x09, 0xb3, 0xe3, 0x04,
	0xbb, 0x48, 0xa9, 0x70, 0xc9, 0xaf, 0xa7, 0x96, 0xbc, 0x92, 0x4b, 0xa6, 0x31, 0x11, 0x4c, 0xba,
	0xf5, 0x40, 0xd0, 0x1f, 0x77, 0x12, 0x76, 0x94, 0x41, 0xf0, 0x1c, 0x2c, 0x95, 0x7b, 0x52, 0x16,
	0x78, 0xb1, 0xb7, 0x53, 0x14, 0xdb, 0x8f, 0xd8, 0x30, 0xd5, 0x56, 0xef, 0xce, 0x47, 0xb7, 0xaa,
	0xa5, 0xa9, 0xc0, 0x5d, 0xb0, 0xc4, 0xa7, 0x46, 0x63, 0xc7, 0xc5, 0x91, 0xaf, 0x2c, 0x66, 0xc7,
	0xd5, 0x7c, 0x38, 0xce, 0x2d, 0x7b, 0x75, 0xab, 0x9a, 0x99, 0xc7, 0xb9, 0x05, 0x3f, 0xcb, 0x60,
	0x1d, 0x5d, 0xc6, 0x24, 0xca, 0xb8, 0x1d, 0xd1, 0x8e, 0x4d, 0x22, 0xa4, 0x00, 0xae, 0xf7, 0xfd,
	0xd4, 0x7a, 0x1f, 0xe5, 0x35, 0x27, 0x92, 0xea, 0x16, 0x2c, 0xf0, 0xbd, 0x7c, 0x4c, 0x87, 0x11,
	0x82, 0xa7, 0x60, 0x81, 0x7e, 0x72, 0x62, 0xfb, 0x0c, 0x21, 0xa5, 0xca, 0xab, 0xee, 0x4d, 0x7d,
	0x24, 0xf7, 0xc5, 0x91, 0x08, 0x1e, 0xdd, 0xaa, 0x64, 0xbf, 0xef, 0x10, 0x82, 0x97, 0x60, 0x3d,
	0x70, 0x28, 0x1b, 0xdf, 0x29, 0xbb, 0x1b, 0x7b, 0x0e, 0x43, 0xca, 0xd2, 0x8e, 0x5c, 0xab, 0xbe,
	0xd8, 0x32, 0xf2, 0x45, 0x34, 0x8a, 0x45, 0x34, 0x4e, 0x8a, 0x45, 0x6c, 0xd6, 0x32, 0x19, 0xe3,
	0x96, 0x26, 0xd2, 0xe8, 0x57, 0x7f, 0x34, 0xd9, 0x5a, 0xcd, 0x7c, 0xa3, 0xeb, 0xf9, 0x81, 0x7b,
	0x76, 0x57, 0xbe, 0x5e, 0x6b, 0xd2, 0xb7, 0x6b, 0x4d, 0xfa, 0xf1, 0xbd, 0x3e, 0x97, 0xad, 0xf2,
	0x7e, 0xf3, 0xf4, 0xa6, 0xaf, 0xca, 0xb7, 0x7d, 0x55, 0xfe, 0xdb, 0x57, 0xe5, 0xab, 0x81, 0x2a,
	0xdd, 0x0e, 0x54, 0xe9, 0xd7, 0x40, 0x95, 0x3e, 0x36, 0x4b, 0xad, 0x8a, 0x27, 0xa7, 0x1e, 0x38,
	0x2d, 0x5a, 0x18, 0xe6, 0x45, 0xe3, 0x95, 0x79, 0xf9, 0xbf, 0x07, 0x2b, 0x24, 0x1e, 0x0a, 0x5a,
	0xf3, 0xbc, 0x87, 0x97, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdf, 0xcb, 0xdc, 0x32, 0xdf, 0x04,
	0x00, 0x00,
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
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastLiquidityUpdate, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastLiquidityUpdate):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintPool(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x62
	{
		size := m.SwapFee.Size()
		i -= size
		if _, err := m.SwapFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.ExponentAtPriceOne.Size()
		i -= size
		if _, err := m.ExponentAtPriceOne.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if m.TickSpacing != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.TickSpacing))
		i--
		dAtA[i] = 0x48
	}
	{
		size := m.CurrentTick.Size()
		i -= size
		if _, err := m.CurrentTick.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.CurrentSqrtPrice.Size()
		i -= size
		if _, err := m.CurrentSqrtPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.Token1) > 0 {
		i -= len(m.Token1)
		copy(dAtA[i:], m.Token1)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Token1)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Token0) > 0 {
		i -= len(m.Token0)
		copy(dAtA[i:], m.Token0)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Token0)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size := m.CurrentTickLiquidity.Size()
		i -= size
		if _, err := m.CurrentTickLiquidity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Id != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x18
	}
	if len(m.IncentivesAddress) > 0 {
		i -= len(m.IncentivesAddress)
		copy(dAtA[i:], m.IncentivesAddress)
		i = encodeVarintPool(dAtA, i, uint64(len(m.IncentivesAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.IncentivesAddress)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovPool(uint64(m.Id))
	}
	l = m.CurrentTickLiquidity.Size()
	n += 1 + l + sovPool(uint64(l))
	l = len(m.Token0)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.Token1)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = m.CurrentSqrtPrice.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.CurrentTick.Size()
	n += 1 + l + sovPool(uint64(l))
	if m.TickSpacing != 0 {
		n += 1 + sovPool(uint64(m.TickSpacing))
	}
	l = m.ExponentAtPriceOne.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.SwapFee.Size()
	n += 1 + l + sovPool(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastLiquidityUpdate)
	n += 1 + l + sovPool(uint64(l))
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncentivesAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IncentivesAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentTickLiquidity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentTickLiquidity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token0", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token0 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentSqrtPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentSqrtPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentTick", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentTick.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickSpacing", wireType)
			}
			m.TickSpacing = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TickSpacing |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExponentAtPriceOne", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExponentAtPriceOne.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastLiquidityUpdate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastLiquidityUpdate, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)
