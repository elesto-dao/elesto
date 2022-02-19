// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: credentials/v1/event.proto

package credentials

import (
	fmt "fmt"
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

type CredentialDefinitionPublishedEvent struct {
	CredentialDefinitionId string `protobuf:"bytes,1,opt,name=credentialDefinitionId,proto3" json:"credentialDefinitionId,omitempty"`
	PublisherId            string `protobuf:"bytes,2,opt,name=publisherId,proto3" json:"publisherId,omitempty"`
}

func (m *CredentialDefinitionPublishedEvent) Reset()         { *m = CredentialDefinitionPublishedEvent{} }
func (m *CredentialDefinitionPublishedEvent) String() string { return proto.CompactTextString(m) }
func (*CredentialDefinitionPublishedEvent) ProtoMessage()    {}
func (*CredentialDefinitionPublishedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_be91971557aa138f, []int{0}
}
func (m *CredentialDefinitionPublishedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CredentialDefinitionPublishedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CredentialDefinitionPublishedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CredentialDefinitionPublishedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialDefinitionPublishedEvent.Merge(m, src)
}
func (m *CredentialDefinitionPublishedEvent) XXX_Size() int {
	return m.Size()
}
func (m *CredentialDefinitionPublishedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialDefinitionPublishedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialDefinitionPublishedEvent proto.InternalMessageInfo

type CredentialIssuerRegisteredEvent struct {
	IssuerId string `protobuf:"bytes,1,opt,name=issuerId,proto3" json:"issuerId,omitempty"`
}

func (m *CredentialIssuerRegisteredEvent) Reset()         { *m = CredentialIssuerRegisteredEvent{} }
func (m *CredentialIssuerRegisteredEvent) String() string { return proto.CompactTextString(m) }
func (*CredentialIssuerRegisteredEvent) ProtoMessage()    {}
func (*CredentialIssuerRegisteredEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_be91971557aa138f, []int{1}
}
func (m *CredentialIssuerRegisteredEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CredentialIssuerRegisteredEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CredentialIssuerRegisteredEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CredentialIssuerRegisteredEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialIssuerRegisteredEvent.Merge(m, src)
}
func (m *CredentialIssuerRegisteredEvent) XXX_Size() int {
	return m.Size()
}
func (m *CredentialIssuerRegisteredEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialIssuerRegisteredEvent.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialIssuerRegisteredEvent proto.InternalMessageInfo

type PublicCredentialIssuedEvent struct {
	CredentialDefinitionId string `protobuf:"bytes,1,opt,name=credentialDefinitionId,proto3" json:"credentialDefinitionId,omitempty"`
	CredentialId           string `protobuf:"bytes,2,opt,name=credentialId,proto3" json:"credentialId,omitempty"`
	IssuerId               string `protobuf:"bytes,3,opt,name=issuerId,proto3" json:"issuerId,omitempty"`
	HolderId               string `protobuf:"bytes,4,opt,name=holderId,proto3" json:"holderId,omitempty"`
}

func (m *PublicCredentialIssuedEvent) Reset()         { *m = PublicCredentialIssuedEvent{} }
func (m *PublicCredentialIssuedEvent) String() string { return proto.CompactTextString(m) }
func (*PublicCredentialIssuedEvent) ProtoMessage()    {}
func (*PublicCredentialIssuedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_be91971557aa138f, []int{2}
}
func (m *PublicCredentialIssuedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PublicCredentialIssuedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PublicCredentialIssuedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PublicCredentialIssuedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicCredentialIssuedEvent.Merge(m, src)
}
func (m *PublicCredentialIssuedEvent) XXX_Size() int {
	return m.Size()
}
func (m *PublicCredentialIssuedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicCredentialIssuedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_PublicCredentialIssuedEvent proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CredentialDefinitionPublishedEvent)(nil), "elestodao.elesto.credentials.v1.CredentialDefinitionPublishedEvent")
	proto.RegisterType((*CredentialIssuerRegisteredEvent)(nil), "elestodao.elesto.credentials.v1.CredentialIssuerRegisteredEvent")
	proto.RegisterType((*PublicCredentialIssuedEvent)(nil), "elestodao.elesto.credentials.v1.PublicCredentialIssuedEvent")
}

func init() { proto.RegisterFile("credentials/v1/event.proto", fileDescriptor_be91971557aa138f) }

var fileDescriptor_be91971557aa138f = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x2e, 0x4a, 0x4d,
	0x49, 0xcd, 0x2b, 0xc9, 0x4c, 0xcc, 0x29, 0xd6, 0x2f, 0x33, 0xd4, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4f, 0xcd, 0x49, 0x2d, 0x2e, 0xc9, 0x4f, 0x49,
	0xcc, 0xd7, 0x83, 0xb0, 0xf4, 0x90, 0x14, 0xeb, 0x95, 0x19, 0x4a, 0x89, 0xa4, 0xe7, 0xa7, 0xe7,
	0x83, 0xd5, 0xea, 0x83, 0x58, 0x10, 0x6d, 0x4a, 0x1d, 0x8c, 0x5c, 0x4a, 0xce, 0x70, 0x85, 0x2e,
	0xa9, 0x69, 0x99, 0x79, 0x99, 0x25, 0x99, 0xf9, 0x79, 0x01, 0xa5, 0x49, 0x39, 0x99, 0xc5, 0x19,
	0xa9, 0x29, 0xae, 0x20, 0x3b, 0x84, 0xcc, 0xb8, 0xc4, 0x92, 0xb1, 0xa8, 0xf2, 0x4c, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xc2, 0x21, 0x2b, 0xa4, 0xc0, 0xc5, 0x5d, 0x00, 0x35, 0xa9, 0xc8,
	0x33, 0x45, 0x82, 0x09, 0xac, 0x18, 0x59, 0xc8, 0x8a, 0xa3, 0x63, 0x81, 0x3c, 0xc3, 0x8b, 0x05,
	0xf2, 0x8c, 0x4a, 0xee, 0x5c, 0xf2, 0x08, 0x97, 0x78, 0x16, 0x17, 0x97, 0xa6, 0x16, 0x05, 0xa5,
	0xa6, 0x67, 0x16, 0x97, 0xa4, 0x16, 0xc1, 0x9c, 0x21, 0xc5, 0xc5, 0x91, 0x09, 0x96, 0x80, 0x5b,
	0x0c, 0xe7, 0x23, 0x19, 0xb4, 0x9b, 0x91, 0x4b, 0x1a, 0xec, 0xfe, 0x64, 0x34, 0xf3, 0x28, 0xf4,
	0x8c, 0x12, 0x17, 0x0f, 0x42, 0x06, 0xee, 0x1b, 0x14, 0x31, 0x14, 0x17, 0x32, 0xa3, 0xba, 0x10,
	0x24, 0x97, 0x91, 0x9f, 0x93, 0x02, 0x96, 0x63, 0x81, 0xc8, 0xc1, 0xf8, 0x08, 0xd7, 0x3b, 0xb9,
	0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb,
	0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x56, 0x7a, 0x66, 0x49, 0x46,
	0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0x24, 0x8e, 0x75, 0x53, 0x12, 0xf3, 0xa1, 0x4c, 0xfd,
	0x0a, 0x7d, 0xa4, 0x08, 0x4f, 0x62, 0x03, 0x47, 0xaf, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xe8,
	0xc3, 0xe8, 0x04, 0x33, 0x02, 0x00, 0x00,
}

func (this *CredentialDefinitionPublishedEvent) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CredentialDefinitionPublishedEvent)
	if !ok {
		that2, ok := that.(CredentialDefinitionPublishedEvent)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.CredentialDefinitionId != that1.CredentialDefinitionId {
		return false
	}
	if this.PublisherId != that1.PublisherId {
		return false
	}
	return true
}
func (this *CredentialIssuerRegisteredEvent) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CredentialIssuerRegisteredEvent)
	if !ok {
		that2, ok := that.(CredentialIssuerRegisteredEvent)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.IssuerId != that1.IssuerId {
		return false
	}
	return true
}
func (this *PublicCredentialIssuedEvent) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PublicCredentialIssuedEvent)
	if !ok {
		that2, ok := that.(PublicCredentialIssuedEvent)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.CredentialDefinitionId != that1.CredentialDefinitionId {
		return false
	}
	if this.CredentialId != that1.CredentialId {
		return false
	}
	if this.IssuerId != that1.IssuerId {
		return false
	}
	if this.HolderId != that1.HolderId {
		return false
	}
	return true
}
func (m *CredentialDefinitionPublishedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CredentialDefinitionPublishedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CredentialDefinitionPublishedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PublisherId) > 0 {
		i -= len(m.PublisherId)
		copy(dAtA[i:], m.PublisherId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.PublisherId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CredentialDefinitionId) > 0 {
		i -= len(m.CredentialDefinitionId)
		copy(dAtA[i:], m.CredentialDefinitionId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.CredentialDefinitionId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CredentialIssuerRegisteredEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CredentialIssuerRegisteredEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CredentialIssuerRegisteredEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IssuerId) > 0 {
		i -= len(m.IssuerId)
		copy(dAtA[i:], m.IssuerId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.IssuerId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PublicCredentialIssuedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublicCredentialIssuedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicCredentialIssuedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.HolderId) > 0 {
		i -= len(m.HolderId)
		copy(dAtA[i:], m.HolderId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.HolderId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.IssuerId) > 0 {
		i -= len(m.IssuerId)
		copy(dAtA[i:], m.IssuerId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.IssuerId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.CredentialId) > 0 {
		i -= len(m.CredentialId)
		copy(dAtA[i:], m.CredentialId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.CredentialId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CredentialDefinitionId) > 0 {
		i -= len(m.CredentialDefinitionId)
		copy(dAtA[i:], m.CredentialDefinitionId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.CredentialDefinitionId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CredentialDefinitionPublishedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CredentialDefinitionId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.PublisherId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *CredentialIssuerRegisteredEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IssuerId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *PublicCredentialIssuedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CredentialDefinitionId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.CredentialId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.IssuerId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.HolderId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CredentialDefinitionPublishedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: CredentialDefinitionPublishedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CredentialDefinitionPublishedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CredentialDefinitionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CredentialDefinitionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublisherId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublisherId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *CredentialIssuerRegisteredEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: CredentialIssuerRegisteredEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CredentialIssuerRegisteredEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssuerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IssuerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *PublicCredentialIssuedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: PublicCredentialIssuedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PublicCredentialIssuedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CredentialDefinitionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CredentialDefinitionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CredentialId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CredentialId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssuerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IssuerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HolderId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HolderId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
