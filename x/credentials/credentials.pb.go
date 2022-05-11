// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: credentials/v1/credentials.proto

package credentials

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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

type CredentialDefinition struct {
	// the credential definition did
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// the did of the publisher of the credential
	PublisherId string `protobuf:"bytes,2,opt,name=publisherId,proto3" json:"publisherId,omitempty"`
	// the credential json-ld schema
	Schema string `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	// the credential vocabulary
	Vocab string `protobuf:"bytes,4,opt,name=vocab,proto3" json:"vocab,omitempty"`
	// the human readable name of the credential, should be included
	// in the type of the issued credential
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// the description of the credential, such as it's purpose
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// wherever the credential is intended for public use (on-chain) or not (off-chain)
	// if the value is false then the module will forbid the issuance of the credential on chain
	IsPublic bool `protobuf:"varint,7,opt,name=isPublic,proto3" json:"isPublic,omitempty"`
	// did of the credential should not be used anymore in favour of something else
	SupersededBy string `protobuf:"bytes,11,opt,name=supersededBy,proto3" json:"supersededBy,omitempty"`
	// the credential can be de-activated
	IsActive bool `protobuf:"varint,12,opt,name=isActive,proto3" json:"isActive,omitempty"`
}

func (m *CredentialDefinition) Reset()         { *m = CredentialDefinition{} }
func (m *CredentialDefinition) String() string { return proto.CompactTextString(m) }
func (*CredentialDefinition) ProtoMessage()    {}
func (*CredentialDefinition) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5d85b2b80c68f8, []int{0}
}
func (m *CredentialDefinition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CredentialDefinition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CredentialDefinition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CredentialDefinition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialDefinition.Merge(m, src)
}
func (m *CredentialDefinition) XXX_Size() int {
	return m.Size()
}
func (m *CredentialDefinition) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialDefinition.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialDefinition proto.InternalMessageInfo

func (m *CredentialDefinition) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CredentialDefinition) GetPublisherId() string {
	if m != nil {
		return m.PublisherId
	}
	return ""
}

func (m *CredentialDefinition) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *CredentialDefinition) GetVocab() string {
	if m != nil {
		return m.Vocab
	}
	return ""
}

func (m *CredentialDefinition) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CredentialDefinition) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CredentialDefinition) GetIsPublic() bool {
	if m != nil {
		return m.IsPublic
	}
	return false
}

func (m *CredentialDefinition) GetSupersededBy() string {
	if m != nil {
		return m.SupersededBy
	}
	return ""
}

func (m *CredentialDefinition) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

// DidMetadata defines metadata associated to a did document such as
// the status of the DID document
type PublicVerifiableCredential struct {
	// json-ld context
	Context string `protobuf:"bytes,1,opt,name=context,proto3" json:"@context,omitempty"`
	// the credential id
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// the definition
	Type []string `protobuf:"bytes,3,rep,name=type,proto3" json:"type,omitempty"`
	// the
	Issuer string `protobuf:"bytes,4,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// the date-time of issuance
	IssuanceDate *time.Time `protobuf:"bytes,5,opt,name=issuanceDate,proto3,stdtime" json:"issuanceDate,omitempty"`
	// the date-time of expiration
	ExpirationDate *time.Time `protobuf:"bytes,6,opt,name=expirationDate,proto3,stdtime" json:"expirationDate,omitempty"`
	// the subject of the credential
	CredentialSubject *types.Any `protobuf:"bytes,7,opt,name=credentialSubject,proto3" json:"credentialSubject,omitempty"`
}

func (m *PublicVerifiableCredential) Reset()         { *m = PublicVerifiableCredential{} }
func (m *PublicVerifiableCredential) String() string { return proto.CompactTextString(m) }
func (*PublicVerifiableCredential) ProtoMessage()    {}
func (*PublicVerifiableCredential) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5d85b2b80c68f8, []int{1}
}
func (m *PublicVerifiableCredential) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PublicVerifiableCredential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PublicVerifiableCredential.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PublicVerifiableCredential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicVerifiableCredential.Merge(m, src)
}
func (m *PublicVerifiableCredential) XXX_Size() int {
	return m.Size()
}
func (m *PublicVerifiableCredential) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicVerifiableCredential.DiscardUnknown(m)
}

var xxx_messageInfo_PublicVerifiableCredential proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CredentialDefinition)(nil), "elestodao.elesto.credentials.v1.CredentialDefinition")
	proto.RegisterType((*PublicVerifiableCredential)(nil), "elestodao.elesto.credentials.v1.PublicVerifiableCredential")
}

func init() { proto.RegisterFile("credentials/v1/credentials.proto", fileDescriptor_bc5d85b2b80c68f8) }

var fileDescriptor_bc5d85b2b80c68f8 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x31, 0x6f, 0xdb, 0x3c,
	0x10, 0xb5, 0x64, 0xc7, 0xf1, 0x47, 0x1b, 0x01, 0x3e, 0xc2, 0x08, 0x58, 0x0f, 0x92, 0xe1, 0x29,
	0x28, 0x5a, 0xa9, 0x49, 0xb7, 0x4e, 0x8d, 0xeb, 0xa1, 0xdd, 0x0a, 0xb5, 0xe8, 0xd0, 0x8d, 0xa2,
	0xce, 0x32, 0x0b, 0x49, 0x14, 0x44, 0xca, 0xb0, 0xfe, 0x41, 0xc6, 0xfc, 0x84, 0xfc, 0x9c, 0x8e,
	0x19, 0x3b, 0xb5, 0x85, 0xbd, 0x04, 0xfd, 0x15, 0x85, 0x48, 0xd9, 0x51, 0x92, 0xa5, 0xdb, 0xbd,
	0x77, 0xef, 0xdd, 0xd9, 0xef, 0x44, 0x34, 0x65, 0x05, 0x44, 0x90, 0x29, 0x4e, 0x13, 0xe9, 0xaf,
	0xcf, 0xfd, 0x16, 0xf4, 0xf2, 0x42, 0x28, 0x81, 0x5d, 0x48, 0x40, 0x2a, 0x11, 0x51, 0xe1, 0x99,
	0xca, 0x6b, 0x6b, 0xd6, 0xe7, 0x93, 0x71, 0x2c, 0x62, 0xa1, 0xb5, 0x7e, 0x5d, 0x19, 0xdb, 0xc4,
	0x8d, 0x85, 0x88, 0x13, 0xf0, 0x35, 0x0a, 0xcb, 0xa5, 0xaf, 0x78, 0x0a, 0x52, 0xd1, 0x34, 0x6f,
	0x04, 0xcf, 0x1e, 0x0b, 0x68, 0x56, 0x99, 0xd6, 0xec, 0xca, 0x46, 0xe3, 0x77, 0x87, 0x25, 0x0b,
	0x58, 0xf2, 0x8c, 0x2b, 0x2e, 0x32, 0x7c, 0x82, 0x6c, 0x1e, 0x11, 0x6b, 0x6a, 0x9d, 0xfd, 0x17,
	0xd8, 0x3c, 0xc2, 0x53, 0x34, 0xcc, 0xcb, 0x30, 0xe1, 0x72, 0x05, 0xc5, 0x87, 0x88, 0xd8, 0xba,
	0xd1, 0xa6, 0xf0, 0x29, 0xea, 0x4b, 0xb6, 0x82, 0x94, 0x92, 0xae, 0x6e, 0x36, 0x08, 0x8f, 0xd1,
	0xd1, 0x5a, 0x30, 0x1a, 0x92, 0x9e, 0xa6, 0x0d, 0xc0, 0x18, 0xf5, 0x32, 0x9a, 0x02, 0x39, 0xd2,
	0xa4, 0xae, 0xeb, 0x1d, 0x11, 0x48, 0x56, 0xf0, 0xbc, 0xfe, 0x09, 0xa4, 0x6f, 0x76, 0xb4, 0x28,
	0x3c, 0x41, 0x03, 0x2e, 0x3f, 0xd6, 0x4b, 0x19, 0x39, 0x9e, 0x5a, 0x67, 0x83, 0xe0, 0x80, 0xf1,
	0x0c, 0x8d, 0x64, 0x99, 0x43, 0x21, 0x21, 0x82, 0x68, 0x5e, 0x91, 0xa1, 0xb6, 0x3f, 0xe0, 0x8c,
	0xff, 0x92, 0x29, 0xbe, 0x06, 0x32, 0xda, 0xfb, 0x0d, 0x9e, 0xdd, 0xd9, 0x68, 0x62, 0x46, 0x7d,
	0x81, 0x82, 0x2f, 0x39, 0x0d, 0x13, 0xb8, 0x8f, 0x06, 0xbf, 0x42, 0xc7, 0x4c, 0x64, 0x0a, 0x36,
	0xca, 0xa4, 0x32, 0x3f, 0xfd, 0xf3, 0xd3, 0xc5, 0x6f, 0x1b, 0xee, 0x85, 0x48, 0xb9, 0x82, 0x34,
	0x57, 0x55, 0xb0, 0x97, 0x35, 0x11, 0xda, 0x87, 0x08, 0x31, 0xea, 0xa9, 0x2a, 0x07, 0xd2, 0x9d,
	0x76, 0xeb, 0xbf, 0x5c, 0xd7, 0x75, 0x68, 0x5c, 0xca, 0x12, 0x8a, 0x26, 0x9d, 0x06, 0xe1, 0x05,
	0x1a, 0xd5, 0x15, 0xcd, 0x18, 0x2c, 0xa8, 0x32, 0x31, 0x0d, 0x2f, 0x26, 0x9e, 0xb9, 0xa4, 0xb7,
	0xbf, 0xa4, 0xf7, 0x79, 0x7f, 0xea, 0x79, 0xef, 0xfa, 0x97, 0x6b, 0x05, 0x0f, 0x5c, 0xf8, 0x3d,
	0x3a, 0x81, 0x4d, 0xce, 0x0b, 0x5a, 0x87, 0xa7, 0xe7, 0xf4, 0xff, 0x71, 0xce, 0x23, 0x1f, 0x9e,
	0xa3, 0xff, 0xef, 0xbf, 0xc5, 0x4f, 0x65, 0xf8, 0x0d, 0x98, 0xd2, 0x17, 0x18, 0x5e, 0x8c, 0x9f,
	0x0c, 0xbb, 0xcc, 0xaa, 0xe0, 0xa9, 0xfc, 0xcd, 0xe0, 0xea, 0xc6, 0xed, 0xdc, 0xdd, 0xb8, 0xd6,
	0x7c, 0xf1, 0x7d, 0xeb, 0x58, 0xb7, 0x5b, 0xc7, 0xfa, 0xbd, 0x75, 0xac, 0xeb, 0x9d, 0xd3, 0xb9,
	0xdd, 0x39, 0x9d, 0x1f, 0x3b, 0xa7, 0xf3, 0xf5, 0x79, 0xcc, 0xd5, 0xaa, 0x0c, 0x3d, 0x26, 0x52,
	0xdf, 0xbc, 0x81, 0x97, 0x11, 0x15, 0x4d, 0xe9, 0x6f, 0xda, 0x8f, 0x26, 0xec, 0xeb, 0x85, 0xaf,
	0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x4e, 0x61, 0x4a, 0x59, 0x03, 0x00, 0x00,
}

func (this *PublicVerifiableCredential) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PublicVerifiableCredential)
	if !ok {
		that2, ok := that.(PublicVerifiableCredential)
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
	if this.Context != that1.Context {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if len(this.Type) != len(that1.Type) {
		return false
	}
	for i := range this.Type {
		if this.Type[i] != that1.Type[i] {
			return false
		}
	}
	if this.Issuer != that1.Issuer {
		return false
	}
	if that1.IssuanceDate == nil {
		if this.IssuanceDate != nil {
			return false
		}
	} else if !this.IssuanceDate.Equal(*that1.IssuanceDate) {
		return false
	}
	if that1.ExpirationDate == nil {
		if this.ExpirationDate != nil {
			return false
		}
	} else if !this.ExpirationDate.Equal(*that1.ExpirationDate) {
		return false
	}
	if !this.CredentialSubject.Equal(that1.CredentialSubject) {
		return false
	}
	return true
}
func (m *CredentialDefinition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CredentialDefinition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CredentialDefinition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsActive {
		i--
		if m.IsActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x60
	}
	if len(m.SupersededBy) > 0 {
		i -= len(m.SupersededBy)
		copy(dAtA[i:], m.SupersededBy)
		i = encodeVarintCredentials(dAtA, i, uint64(len(m.SupersededBy)))
		i--
		dAtA[i] = 0x5a
	}
	if m.IsPublic {
		i--
		if m.IsPublic {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintCredentials(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintCredentials(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Vocab) > 0 {
		i -= len(m.Vocab)
		copy(dAtA[i:], m.Vocab)
		i = encodeVarintCredentials(dAtA, i, uint64(len(m.Vocab)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Schema) > 0 {
		i -= len(m.Schema)
		copy(dAtA[i:], m.Schema)
		i = encodeVarintCredentials(dAtA, i, uint64(len(m.Schema)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PublisherId) > 0 {
		i -= len(m.PublisherId)
		copy(dAtA[i:], m.PublisherId)
		i = encodeVarintCredentials(dAtA, i, uint64(len(m.PublisherId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintCredentials(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PublicVerifiableCredential) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublicVerifiableCredential) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicVerifiableCredential) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CredentialSubject != nil {
		{
			size, err := m.CredentialSubject.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCredentials(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.ExpirationDate != nil {
		n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.ExpirationDate, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.ExpirationDate):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintCredentials(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0x32
	}
	if m.IssuanceDate != nil {
		n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.IssuanceDate, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.IssuanceDate):])
		if err3 != nil {
			return 0, err3
		}
		i -= n3
		i = encodeVarintCredentials(dAtA, i, uint64(n3))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintCredentials(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Type) > 0 {
		for iNdEx := len(m.Type) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Type[iNdEx])
			copy(dAtA[i:], m.Type[iNdEx])
			i = encodeVarintCredentials(dAtA, i, uint64(len(m.Type[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintCredentials(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Context) > 0 {
		i -= len(m.Context)
		copy(dAtA[i:], m.Context)
		i = encodeVarintCredentials(dAtA, i, uint64(len(m.Context)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCredentials(dAtA []byte, offset int, v uint64) int {
	offset -= sovCredentials(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CredentialDefinition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovCredentials(uint64(l))
	}
	l = len(m.PublisherId)
	if l > 0 {
		n += 1 + l + sovCredentials(uint64(l))
	}
	l = len(m.Schema)
	if l > 0 {
		n += 1 + l + sovCredentials(uint64(l))
	}
	l = len(m.Vocab)
	if l > 0 {
		n += 1 + l + sovCredentials(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCredentials(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovCredentials(uint64(l))
	}
	if m.IsPublic {
		n += 2
	}
	l = len(m.SupersededBy)
	if l > 0 {
		n += 1 + l + sovCredentials(uint64(l))
	}
	if m.IsActive {
		n += 2
	}
	return n
}

func (m *PublicVerifiableCredential) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Context)
	if l > 0 {
		n += 1 + l + sovCredentials(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovCredentials(uint64(l))
	}
	if len(m.Type) > 0 {
		for _, s := range m.Type {
			l = len(s)
			n += 1 + l + sovCredentials(uint64(l))
		}
	}
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovCredentials(uint64(l))
	}
	if m.IssuanceDate != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.IssuanceDate)
		n += 1 + l + sovCredentials(uint64(l))
	}
	if m.ExpirationDate != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.ExpirationDate)
		n += 1 + l + sovCredentials(uint64(l))
	}
	if m.CredentialSubject != nil {
		l = m.CredentialSubject.Size()
		n += 1 + l + sovCredentials(uint64(l))
	}
	return n
}

func sovCredentials(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCredentials(x uint64) (n int) {
	return sovCredentials(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CredentialDefinition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCredentials
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
			return fmt.Errorf("proto: CredentialDefinition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CredentialDefinition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
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
				return ErrInvalidLengthCredentials
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentials
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublisherId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
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
				return ErrInvalidLengthCredentials
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentials
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublisherId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Schema", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
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
				return ErrInvalidLengthCredentials
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentials
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Schema = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vocab", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
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
				return ErrInvalidLengthCredentials
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentials
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vocab = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
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
				return ErrInvalidLengthCredentials
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentials
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
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
				return ErrInvalidLengthCredentials
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentials
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsPublic", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
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
			m.IsPublic = bool(v != 0)
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupersededBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
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
				return ErrInvalidLengthCredentials
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentials
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SupersededBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
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
			m.IsActive = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCredentials(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCredentials
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
func (m *PublicVerifiableCredential) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCredentials
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
			return fmt.Errorf("proto: PublicVerifiableCredential: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PublicVerifiableCredential: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
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
				return ErrInvalidLengthCredentials
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentials
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Context = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
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
				return ErrInvalidLengthCredentials
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentials
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
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
				return ErrInvalidLengthCredentials
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentials
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = append(m.Type, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
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
				return ErrInvalidLengthCredentials
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredentials
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssuanceDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
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
				return ErrInvalidLengthCredentials
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCredentials
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.IssuanceDate == nil {
				m.IssuanceDate = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.IssuanceDate, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
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
				return ErrInvalidLengthCredentials
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCredentials
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExpirationDate == nil {
				m.ExpirationDate = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.ExpirationDate, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CredentialSubject", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
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
				return ErrInvalidLengthCredentials
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCredentials
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CredentialSubject == nil {
				m.CredentialSubject = &types.Any{}
			}
			if err := m.CredentialSubject.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCredentials(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCredentials
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
func skipCredentials(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCredentials
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
					return 0, ErrIntOverflowCredentials
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
					return 0, ErrIntOverflowCredentials
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
				return 0, ErrInvalidLengthCredentials
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCredentials
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCredentials
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCredentials        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCredentials          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCredentials = fmt.Errorf("proto: unexpected end of group")
)
