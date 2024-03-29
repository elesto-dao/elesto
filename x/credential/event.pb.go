// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: credential/v1/event.proto

package credential

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

// CredentialDefinitionPublishedEvent this event gets triggered when a credential definition is published
type CredentialDefinitionPublishedEvent struct {
	CredentialDefinitionID string `protobuf:"bytes,1,opt,name=credentialDefinitionID,proto3" json:"credentialDefinitionID,omitempty"`
	PublisherID            string `protobuf:"bytes,2,opt,name=publisherID,proto3" json:"publisherID,omitempty"`
}

func (m *CredentialDefinitionPublishedEvent) Reset()         { *m = CredentialDefinitionPublishedEvent{} }
func (m *CredentialDefinitionPublishedEvent) String() string { return proto.CompactTextString(m) }
func (*CredentialDefinitionPublishedEvent) ProtoMessage()    {}
func (*CredentialDefinitionPublishedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_d041b5503d3192a8, []int{0}
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

// CredentialDefinitionUpdatedEvent this event gets triggered when a definition gets updated
type CredentialDefinitionUpdatedEvent struct {
	CredentialDefinitionID string `protobuf:"bytes,1,opt,name=credentialDefinitionID,proto3" json:"credentialDefinitionID,omitempty"`
}

func (m *CredentialDefinitionUpdatedEvent) Reset()         { *m = CredentialDefinitionUpdatedEvent{} }
func (m *CredentialDefinitionUpdatedEvent) String() string { return proto.CompactTextString(m) }
func (*CredentialDefinitionUpdatedEvent) ProtoMessage()    {}
func (*CredentialDefinitionUpdatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_d041b5503d3192a8, []int{1}
}
func (m *CredentialDefinitionUpdatedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CredentialDefinitionUpdatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CredentialDefinitionUpdatedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CredentialDefinitionUpdatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialDefinitionUpdatedEvent.Merge(m, src)
}
func (m *CredentialDefinitionUpdatedEvent) XXX_Size() int {
	return m.Size()
}
func (m *CredentialDefinitionUpdatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialDefinitionUpdatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialDefinitionUpdatedEvent proto.InternalMessageInfo

// PublicCredentialIssuedEvent this event gets triggered when a public verifiable credential is issued on-chain
type PublicCredentialIssuedEvent struct {
	CredentialDefinitionID string `protobuf:"bytes,1,opt,name=credentialDefinitionID,proto3" json:"credentialDefinitionID,omitempty"`
	CredentialID           string `protobuf:"bytes,2,opt,name=credentialID,proto3" json:"credentialID,omitempty"`
	IssuerID               string `protobuf:"bytes,3,opt,name=issuerID,proto3" json:"issuerID,omitempty"`
}

func (m *PublicCredentialIssuedEvent) Reset()         { *m = PublicCredentialIssuedEvent{} }
func (m *PublicCredentialIssuedEvent) String() string { return proto.CompactTextString(m) }
func (*PublicCredentialIssuedEvent) ProtoMessage()    {}
func (*PublicCredentialIssuedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_d041b5503d3192a8, []int{2}
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
	proto.RegisterType((*CredentialDefinitionPublishedEvent)(nil), "elestodao.elesto.credential.v1.CredentialDefinitionPublishedEvent")
	proto.RegisterType((*CredentialDefinitionUpdatedEvent)(nil), "elestodao.elesto.credential.v1.CredentialDefinitionUpdatedEvent")
	proto.RegisterType((*PublicCredentialIssuedEvent)(nil), "elestodao.elesto.credential.v1.PublicCredentialIssuedEvent")
}

func init() { proto.RegisterFile("credential/v1/event.proto", fileDescriptor_d041b5503d3192a8) }

var fileDescriptor_d041b5503d3192a8 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x2e, 0x4a, 0x4d,
	0x49, 0xcd, 0x2b, 0xc9, 0x4c, 0xcc, 0xd1, 0x2f, 0x33, 0xd4, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4b, 0xcd, 0x49, 0x2d, 0x2e, 0xc9, 0x4f, 0x49, 0xcc,
	0xd7, 0x83, 0xb0, 0xf4, 0x10, 0x6a, 0xf5, 0xca, 0x0c, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1,
	0x4a, 0xf5, 0x41, 0x2c, 0x88, 0x2e, 0xa5, 0x0e, 0x46, 0x2e, 0x25, 0x67, 0xb8, 0x3a, 0x97, 0xd4,
	0xb4, 0xcc, 0xbc, 0xcc, 0x92, 0xcc, 0xfc, 0xbc, 0x80, 0xd2, 0xa4, 0x9c, 0xcc, 0xe2, 0x8c, 0xd4,
	0x14, 0x57, 0x90, 0x15, 0x42, 0x66, 0x5c, 0x62, 0xc9, 0x58, 0x54, 0x79, 0xba, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x06, 0xe1, 0x90, 0x15, 0x52, 0xe0, 0xe2, 0x2e, 0x80, 0x9a, 0x54, 0xe4, 0xe9,
	0x22, 0xc1, 0x04, 0x56, 0x8c, 0x2c, 0x64, 0xc5, 0xd1, 0xb1, 0x40, 0x9e, 0xe1, 0xc5, 0x02, 0x79,
	0x46, 0xa5, 0x14, 0x2e, 0x05, 0x6c, 0x2e, 0x09, 0x2d, 0x48, 0x49, 0x2c, 0xa1, 0xd0, 0x1d, 0x48,
	0xb6, 0xcc, 0x67, 0xe4, 0x92, 0x06, 0x7b, 0x2e, 0x19, 0x61, 0x99, 0x67, 0x71, 0x71, 0x29, 0xa5,
	0x3e, 0x55, 0xe2, 0xe2, 0x41, 0xc8, 0xc0, 0xbd, 0x8a, 0x22, 0x26, 0x24, 0xc5, 0xc5, 0x91, 0x09,
	0xb2, 0x0a, 0x14, 0x14, 0xcc, 0x60, 0x79, 0x38, 0x1f, 0xe1, 0x42, 0x27, 0xb7, 0x13, 0x8f, 0xe4,
	0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f,
	0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x49, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b,
	0xce, 0xcf, 0xd5, 0x87, 0xc4, 0xb1, 0x6e, 0x4a, 0x62, 0x3e, 0x94, 0xa9, 0x5f, 0x66, 0xa2, 0x5f,
	0xa1, 0x8f, 0xb0, 0x2f, 0x89, 0x0d, 0x1c, 0xc3, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc8,
	0x23, 0x92, 0xbc, 0x34, 0x02, 0x00, 0x00,
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
	if this.CredentialDefinitionID != that1.CredentialDefinitionID {
		return false
	}
	if this.PublisherID != that1.PublisherID {
		return false
	}
	return true
}
func (this *CredentialDefinitionUpdatedEvent) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CredentialDefinitionUpdatedEvent)
	if !ok {
		that2, ok := that.(CredentialDefinitionUpdatedEvent)
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
	if this.CredentialDefinitionID != that1.CredentialDefinitionID {
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
	if this.CredentialDefinitionID != that1.CredentialDefinitionID {
		return false
	}
	if this.CredentialID != that1.CredentialID {
		return false
	}
	if this.IssuerID != that1.IssuerID {
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
	if len(m.PublisherID) > 0 {
		i -= len(m.PublisherID)
		copy(dAtA[i:], m.PublisherID)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.PublisherID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CredentialDefinitionID) > 0 {
		i -= len(m.CredentialDefinitionID)
		copy(dAtA[i:], m.CredentialDefinitionID)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.CredentialDefinitionID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CredentialDefinitionUpdatedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CredentialDefinitionUpdatedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CredentialDefinitionUpdatedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CredentialDefinitionID) > 0 {
		i -= len(m.CredentialDefinitionID)
		copy(dAtA[i:], m.CredentialDefinitionID)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.CredentialDefinitionID)))
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
	if len(m.IssuerID) > 0 {
		i -= len(m.IssuerID)
		copy(dAtA[i:], m.IssuerID)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.IssuerID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.CredentialID) > 0 {
		i -= len(m.CredentialID)
		copy(dAtA[i:], m.CredentialID)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.CredentialID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CredentialDefinitionID) > 0 {
		i -= len(m.CredentialDefinitionID)
		copy(dAtA[i:], m.CredentialDefinitionID)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.CredentialDefinitionID)))
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
	l = len(m.CredentialDefinitionID)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.PublisherID)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *CredentialDefinitionUpdatedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CredentialDefinitionID)
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
	l = len(m.CredentialDefinitionID)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.CredentialID)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.IssuerID)
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
				return fmt.Errorf("proto: wrong wireType = %d for field CredentialDefinitionID", wireType)
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
			m.CredentialDefinitionID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublisherID", wireType)
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
			m.PublisherID = string(dAtA[iNdEx:postIndex])
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
func (m *CredentialDefinitionUpdatedEvent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: CredentialDefinitionUpdatedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CredentialDefinitionUpdatedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CredentialDefinitionID", wireType)
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
			m.CredentialDefinitionID = string(dAtA[iNdEx:postIndex])
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
				return fmt.Errorf("proto: wrong wireType = %d for field CredentialDefinitionID", wireType)
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
			m.CredentialDefinitionID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CredentialID", wireType)
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
			m.CredentialID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssuerID", wireType)
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
			m.IssuerID = string(dAtA[iNdEx:postIndex])
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
