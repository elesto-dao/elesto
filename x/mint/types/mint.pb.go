// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mint/v1/mint.proto

package types

import (
	fmt "fmt"
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

// Params holds parameters for the mint module.
type Params struct {
	// mint_denom defines denomination of coin to be minted
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// inflation rates for each year, for example:
	// 1: first year, inflation 100%
	// 0.50: second year, inflation 50%
	InflationRates []string `protobuf:"bytes,3,rep,name=inflation_rates,json=inflationRates,proto3" json:"inflation_rates,omitempty"`
	// estimate number of blocks in a year
	BlocksPerYear int64 `protobuf:"varint,4,opt,name=blocks_per_year,json=blocksPerYear,proto3" json:"blocks_per_year,omitempty"`
	// total max supply
	MaxSupply int64 `protobuf:"varint,5,opt,name=max_supply,json=maxSupply,proto3" json:"max_supply,omitempty"`
	// team address
	TeamAddress string `protobuf:"bytes,6,opt,name=team_address,json=teamAddress,proto3" json:"team_address,omitempty"`
	// team fee
	TeamReward string `protobuf:"bytes,7,opt,name=team_reward,json=teamReward,proto3" json:"team_reward,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_7062018a8a770e0e, []int{0}
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

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetInflationRates() []string {
	if m != nil {
		return m.InflationRates
	}
	return nil
}

func (m *Params) GetBlocksPerYear() int64 {
	if m != nil {
		return m.BlocksPerYear
	}
	return 0
}

func (m *Params) GetMaxSupply() int64 {
	if m != nil {
		return m.MaxSupply
	}
	return 0
}

func (m *Params) GetTeamAddress() string {
	if m != nil {
		return m.TeamAddress
	}
	return ""
}

func (m *Params) GetTeamReward() string {
	if m != nil {
		return m.TeamReward
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "elestodao.elesto.mint.v1.Params")
}

func init() { proto.RegisterFile("mint/v1/mint.proto", fileDescriptor_7062018a8a770e0e) }

var fileDescriptor_7062018a8a770e0e = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xd0, 0x3d, 0x4e, 0xc3, 0x30,
	0x14, 0x07, 0xf0, 0x5a, 0x85, 0xa2, 0x98, 0x8f, 0x4a, 0x9e, 0xbc, 0x60, 0x0a, 0x03, 0x94, 0x81,
	0x44, 0x15, 0x27, 0xe0, 0xe3, 0x00, 0x55, 0x98, 0x60, 0xb1, 0x5e, 0xeb, 0x07, 0x44, 0xc4, 0x71,
	0x64, 0xbb, 0xa5, 0xb9, 0x05, 0xc7, 0x62, 0xec, 0x88, 0xc4, 0x82, 0x92, 0x8b, 0x20, 0x3b, 0x88,
	0xc9, 0xcf, 0x3f, 0xff, 0x25, 0x3f, 0xfd, 0x29, 0xd3, 0x45, 0xe5, 0xb3, 0xf5, 0x2c, 0x0b, 0x67,
	0x5a, 0x5b, 0xe3, 0x0d, 0xe3, 0x58, 0xa2, 0xf3, 0x46, 0x81, 0x49, 0xfb, 0x29, 0x8d, 0x8f, 0xeb,
	0xd9, 0xd9, 0x37, 0xa1, 0xa3, 0x39, 0x58, 0xd0, 0x8e, 0x1d, 0x53, 0x1a, 0x54, 0x2a, 0xac, 0x8c,
	0xe6, 0x64, 0x42, 0xa6, 0x49, 0x9e, 0x04, 0xb9, 0x0f, 0xc0, 0x2e, 0xe8, 0xb8, 0xa8, 0x9e, 0x4b,
	0xf0, 0x85, 0xa9, 0xa4, 0x05, 0x8f, 0x8e, 0x0f, 0x27, 0xc3, 0x69, 0x92, 0x1f, 0xfd, 0x73, 0x1e,
	0x94, 0x9d, 0xd3, 0xf1, 0xa2, 0x34, 0xcb, 0x37, 0x27, 0x6b, 0xb4, 0xb2, 0x41, 0xb0, 0x7c, 0x67,
	0x42, 0xa6, 0xc3, 0xfc, 0xb0, 0xe7, 0x39, 0xda, 0x47, 0x04, 0x1b, 0xff, 0x83, 0x8d, 0x74, 0xab,
	0xba, 0x2e, 0x1b, 0xbe, 0x1b, 0x23, 0x89, 0x86, 0xcd, 0x43, 0x04, 0x76, 0x4a, 0x0f, 0x3c, 0x82,
	0x96, 0xa0, 0x94, 0x45, 0xe7, 0xf8, 0x28, 0x2e, 0xb4, 0x1f, 0xec, 0xa6, 0x27, 0x76, 0x42, 0xe3,
	0x55, 0x5a, 0x7c, 0x07, 0xab, 0xf8, 0x5e, 0x4c, 0xd0, 0x40, 0x79, 0x94, 0xdb, 0xbb, 0xcf, 0x56,
	0x90, 0x6d, 0x2b, 0xc8, 0x4f, 0x2b, 0xc8, 0x47, 0x27, 0x06, 0xdb, 0x4e, 0x0c, 0xbe, 0x3a, 0x31,
	0x78, 0xba, 0x7c, 0x29, 0xfc, 0xeb, 0x6a, 0x91, 0x2e, 0x8d, 0xce, 0xfa, 0x4a, 0xae, 0x14, 0x98,
	0xbf, 0x31, 0xdb, 0xc4, 0xf2, 0x32, 0xdf, 0xd4, 0xe8, 0x16, 0xa3, 0xd8, 0xe1, 0xf5, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xc0, 0x8f, 0x44, 0x9d, 0x59, 0x01, 0x00, 0x00,
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
	if len(m.TeamReward) > 0 {
		i -= len(m.TeamReward)
		copy(dAtA[i:], m.TeamReward)
		i = encodeVarintMint(dAtA, i, uint64(len(m.TeamReward)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.TeamAddress) > 0 {
		i -= len(m.TeamAddress)
		copy(dAtA[i:], m.TeamAddress)
		i = encodeVarintMint(dAtA, i, uint64(len(m.TeamAddress)))
		i--
		dAtA[i] = 0x32
	}
	if m.MaxSupply != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.MaxSupply))
		i--
		dAtA[i] = 0x28
	}
	if m.BlocksPerYear != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.BlocksPerYear))
		i--
		dAtA[i] = 0x20
	}
	if len(m.InflationRates) > 0 {
		for iNdEx := len(m.InflationRates) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.InflationRates[iNdEx])
			copy(dAtA[i:], m.InflationRates[iNdEx])
			i = encodeVarintMint(dAtA, i, uint64(len(m.InflationRates[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
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
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	if len(m.InflationRates) > 0 {
		for _, s := range m.InflationRates {
			l = len(s)
			n += 1 + l + sovMint(uint64(l))
		}
	}
	if m.BlocksPerYear != 0 {
		n += 1 + sovMint(uint64(m.BlocksPerYear))
	}
	if m.MaxSupply != 0 {
		n += 1 + sovMint(uint64(m.MaxSupply))
	}
	l = len(m.TeamAddress)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = len(m.TeamReward)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationRates", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InflationRates = append(m.InflationRates, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksPerYear", wireType)
			}
			m.BlocksPerYear = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksPerYear |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSupply", wireType)
			}
			m.MaxSupply = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSupply |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TeamAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TeamAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TeamReward", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TeamReward = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)
