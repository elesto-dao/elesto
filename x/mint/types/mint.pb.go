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
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0xd0, 0x31, 0x4e, 0xf3, 0x30,
	0x14, 0x07, 0xf0, 0x5a, 0xdf, 0x47, 0x51, 0x0c, 0x08, 0xc9, 0x93, 0x17, 0x4c, 0x61, 0x40, 0x1d,
	0x20, 0x56, 0xe1, 0x04, 0x20, 0xc4, 0x5c, 0x85, 0x09, 0x16, 0xeb, 0xa5, 0x7e, 0x82, 0x8a, 0xb8,
	0x8e, 0x6c, 0x37, 0x24, 0xb7, 0xe0, 0x2c, 0x9c, 0x82, 0xb1, 0x23, 0x23, 0x4a, 0x2e, 0x82, 0xec,
	0x30, 0xf9, 0xf9, 0xe7, 0xbf, 0x6c, 0xf9, 0x4f, 0x99, 0x59, 0x6f, 0x82, 0x6c, 0x16, 0x32, 0xae,
	0x79, 0xed, 0x6c, 0xb0, 0x8c, 0x63, 0x85, 0x3e, 0x58, 0x0d, 0x36, 0x1f, 0xa7, 0x3c, 0x1d, 0x36,
	0x8b, 0xf3, 0x4f, 0x42, 0xa7, 0x4b, 0x70, 0x60, 0x3c, 0x3b, 0xa1, 0x34, 0xaa, 0xd2, 0xb8, 0xb1,
	0x86, 0x93, 0x19, 0x99, 0x67, 0x45, 0x16, 0xe5, 0x3e, 0x02, 0xbb, 0xa0, 0xc7, 0x65, 0x65, 0x57,
	0x6f, 0x5e, 0xd5, 0xe8, 0x54, 0x87, 0xe0, 0xf8, 0xff, 0x19, 0x99, 0xff, 0x2b, 0x8e, 0x46, 0x5e,
	0xa2, 0x7b, 0x42, 0x70, 0xe9, 0x1a, 0x68, 0x95, 0xdf, 0xd6, 0x75, 0xd5, 0xf1, 0xbd, 0x14, 0xc9,
	0x0c, 0xb4, 0x8f, 0x09, 0xd8, 0x19, 0x3d, 0x0c, 0x08, 0x46, 0x81, 0xd6, 0x0e, 0xbd, 0xe7, 0xd3,
	0xf4, 0xce, 0x41, 0xb4, 0xdb, 0x91, 0xd8, 0x29, 0x4d, 0x5b, 0xe5, 0xf0, 0x1d, 0x9c, 0xe6, 0xfb,
	0x29, 0x41, 0x23, 0x15, 0x49, 0xee, 0x1e, 0xbe, 0x7a, 0x41, 0x76, 0xbd, 0x20, 0x3f, 0xbd, 0x20,
	0x1f, 0x83, 0x98, 0xec, 0x06, 0x31, 0xf9, 0x1e, 0xc4, 0xe4, 0xf9, 0xf2, 0x65, 0x1d, 0x5e, 0xb7,
	0x65, 0xbe, 0xb2, 0x46, 0x8e, 0x3f, 0xbd, 0xd2, 0x60, 0xff, 0x46, 0xd9, 0x5c, 0xcb, 0x36, 0xd5,
	0x22, 0x43, 0x57, 0xa3, 0x2f, 0xa7, 0xa9, 0x9d, 0x9b, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8,
	0x64, 0xc3, 0xa3, 0x33, 0x01, 0x00, 0x00,
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
