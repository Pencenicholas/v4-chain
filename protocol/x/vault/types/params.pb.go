// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dydxprotocol/vault/params.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params stores `x/vault` parameters.
type Params struct {
	// The number of layers of orders a vault places. For example if
	// `layers=2`, a vault places 2 asks and 2 bids.
	Layers uint32 `protobuf:"varint,1,opt,name=layers,proto3" json:"layers,omitempty"`
	// The minimum base spread when a vault quotes around reservation price.
	SpreadMinPpm uint32 `protobuf:"varint,2,opt,name=spread_min_ppm,json=spreadMinPpm,proto3" json:"spread_min_ppm,omitempty"`
	// The buffer amount to add to min_price_change_ppm to arrive at `spread`
	// according to formula:
	// `spread = max(spread_min_ppm, min_price_change_ppm + spread_buffer_ppm)`.
	SpreadBufferPpm uint32 `protobuf:"varint,3,opt,name=spread_buffer_ppm,json=spreadBufferPpm,proto3" json:"spread_buffer_ppm,omitempty"`
	// The factor that determines how aggressive a vault skews its orders.
	SkewFactorPpm uint32 `protobuf:"varint,4,opt,name=skew_factor_ppm,json=skewFactorPpm,proto3" json:"skew_factor_ppm,omitempty"`
	// The percentage of vault equity that each order is sized at.
	OrderSizePctPpm uint32 `protobuf:"varint,5,opt,name=order_size_pct_ppm,json=orderSizePctPpm,proto3" json:"order_size_pct_ppm,omitempty"`
	// The duration that a vault's orders are valid for.
	OrderExpirationSeconds uint32 `protobuf:"varint,6,opt,name=order_expiration_seconds,json=orderExpirationSeconds,proto3" json:"order_expiration_seconds,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_6043e0b8bfdbca9f, []int{0}
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

func (m *Params) GetLayers() uint32 {
	if m != nil {
		return m.Layers
	}
	return 0
}

func (m *Params) GetSpreadMinPpm() uint32 {
	if m != nil {
		return m.SpreadMinPpm
	}
	return 0
}

func (m *Params) GetSpreadBufferPpm() uint32 {
	if m != nil {
		return m.SpreadBufferPpm
	}
	return 0
}

func (m *Params) GetSkewFactorPpm() uint32 {
	if m != nil {
		return m.SkewFactorPpm
	}
	return 0
}

func (m *Params) GetOrderSizePctPpm() uint32 {
	if m != nil {
		return m.OrderSizePctPpm
	}
	return 0
}

func (m *Params) GetOrderExpirationSeconds() uint32 {
	if m != nil {
		return m.OrderExpirationSeconds
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "dydxprotocol.vault.Params")
}

func init() { proto.RegisterFile("dydxprotocol/vault/params.proto", fileDescriptor_6043e0b8bfdbca9f) }

var fileDescriptor_6043e0b8bfdbca9f = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0xd0, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0xc6, 0xf1, 0xa6, 0xf7, 0x92, 0xc1, 0xa2, 0x54, 0x78, 0xa8, 0x32, 0x19, 0x84, 0x10, 0x42,
	0x20, 0x9a, 0x01, 0x24, 0x98, 0x2b, 0xc1, 0x86, 0x54, 0xe8, 0xc6, 0x62, 0xb9, 0x8e, 0x43, 0x2d,
	0x92, 0xd8, 0xb2, 0x1d, 0x48, 0xba, 0xf1, 0x06, 0x3c, 0x16, 0x63, 0x47, 0x46, 0x94, 0xbc, 0x08,
	0xca, 0x71, 0x54, 0xc1, 0x98, 0xff, 0xf7, 0x8b, 0x87, 0x83, 0x0e, 0x92, 0x3a, 0xa9, 0xb4, 0x51,
	0x4e, 0x71, 0x95, 0xc5, 0xaf, 0xac, 0xcc, 0x5c, 0xac, 0x99, 0x61, 0xb9, 0x9d, 0x42, 0xc5, 0xf8,
	0x37, 0x98, 0x02, 0x38, 0x7a, 0x1f, 0xa2, 0x70, 0x0e, 0x08, 0x4f, 0x50, 0x98, 0xb1, 0x5a, 0x18,
	0x1b, 0x05, 0x87, 0xc1, 0xe9, 0xe8, 0xb1, 0xff, 0xc2, 0xc7, 0x68, 0xcf, 0x6a, 0x23, 0x58, 0x42,
	0x73, 0x59, 0x50, 0xad, 0xf3, 0x68, 0x08, 0xfb, 0xae, 0xaf, 0xf7, 0xb2, 0x98, 0xeb, 0x1c, 0x9f,
	0xa1, 0xfd, 0x5e, 0x2d, 0xcb, 0x34, 0x15, 0x06, 0xe0, 0x3f, 0x80, 0x63, 0x3f, 0xcc, 0xa0, 0x77,
	0xf6, 0x04, 0x8d, 0xed, 0x8b, 0x78, 0xa3, 0x29, 0xe3, 0x4e, 0x79, 0xf9, 0x1f, 0xe4, 0xa8, 0xcb,
	0x77, 0x50, 0x3b, 0x77, 0x8e, 0xb0, 0x32, 0x89, 0x30, 0xd4, 0xca, 0xb5, 0xa0, 0x9a, 0x3b, 0xa0,
	0x3b, 0xfe, 0x51, 0x58, 0x16, 0x72, 0x2d, 0xe6, 0xdc, 0x75, 0xf8, 0x06, 0x45, 0x1e, 0x8b, 0x4a,
	0x4b, 0xc3, 0x9c, 0x54, 0x05, 0xb5, 0x82, 0xab, 0x22, 0xb1, 0x51, 0x08, 0xbf, 0x4c, 0x60, 0xbf,
	0xdd, 0xce, 0x0b, 0xbf, 0xce, 0x1e, 0x3e, 0x1b, 0x12, 0x6c, 0x1a, 0x12, 0x7c, 0x37, 0x24, 0xf8,
	0x68, 0xc9, 0x60, 0xd3, 0x92, 0xc1, 0x57, 0x4b, 0x06, 0x4f, 0xd7, 0xcf, 0xd2, 0xad, 0xca, 0xe5,
	0x94, 0xab, 0x3c, 0xfe, 0x7b, 0xdd, 0xab, 0x0b, 0xbe, 0x62, 0xb2, 0x88, 0xb7, 0xa5, 0xea, 0x2f,
	0xee, 0x6a, 0x2d, 0xec, 0x32, 0x84, 0x7e, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x9e, 0x6e,
	0xd5, 0x94, 0x01, 0x00, 0x00,
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
	if m.OrderExpirationSeconds != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.OrderExpirationSeconds))
		i--
		dAtA[i] = 0x30
	}
	if m.OrderSizePctPpm != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.OrderSizePctPpm))
		i--
		dAtA[i] = 0x28
	}
	if m.SkewFactorPpm != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SkewFactorPpm))
		i--
		dAtA[i] = 0x20
	}
	if m.SpreadBufferPpm != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SpreadBufferPpm))
		i--
		dAtA[i] = 0x18
	}
	if m.SpreadMinPpm != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SpreadMinPpm))
		i--
		dAtA[i] = 0x10
	}
	if m.Layers != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.Layers))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
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
	if m.Layers != 0 {
		n += 1 + sovParams(uint64(m.Layers))
	}
	if m.SpreadMinPpm != 0 {
		n += 1 + sovParams(uint64(m.SpreadMinPpm))
	}
	if m.SpreadBufferPpm != 0 {
		n += 1 + sovParams(uint64(m.SpreadBufferPpm))
	}
	if m.SkewFactorPpm != 0 {
		n += 1 + sovParams(uint64(m.SkewFactorPpm))
	}
	if m.OrderSizePctPpm != 0 {
		n += 1 + sovParams(uint64(m.OrderSizePctPpm))
	}
	if m.OrderExpirationSeconds != 0 {
		n += 1 + sovParams(uint64(m.OrderExpirationSeconds))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Layers", wireType)
			}
			m.Layers = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Layers |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpreadMinPpm", wireType)
			}
			m.SpreadMinPpm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpreadMinPpm |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpreadBufferPpm", wireType)
			}
			m.SpreadBufferPpm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpreadBufferPpm |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SkewFactorPpm", wireType)
			}
			m.SkewFactorPpm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SkewFactorPpm |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderSizePctPpm", wireType)
			}
			m.OrderSizePctPpm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderSizePctPpm |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderExpirationSeconds", wireType)
			}
			m.OrderExpirationSeconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderExpirationSeconds |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
