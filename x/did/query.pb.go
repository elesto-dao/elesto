// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: did/v1/query.proto

package did

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryDidDocumentsRequest is request type for Query/DidDocuments RPC method.
type QueryDidDocumentRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryDidDocumentRequest) Reset()         { *m = QueryDidDocumentRequest{} }
func (m *QueryDidDocumentRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDidDocumentRequest) ProtoMessage()    {}
func (*QueryDidDocumentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae1fa9bb626e2869, []int{0}
}
func (m *QueryDidDocumentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDidDocumentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDidDocumentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDidDocumentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDidDocumentRequest.Merge(m, src)
}
func (m *QueryDidDocumentRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDidDocumentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDidDocumentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDidDocumentRequest proto.InternalMessageInfo

func (m *QueryDidDocumentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// QueryDidDocumentsResponse is response type for the Query/DidDocuments RPC method
type QueryDidDocumentResponse struct {
	// Returns a did document
	DidDocument DidDocument `protobuf:"bytes,1,opt,name=didDocument,proto3" json:"didDocument"`
}

func (m *QueryDidDocumentResponse) Reset()         { *m = QueryDidDocumentResponse{} }
func (m *QueryDidDocumentResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDidDocumentResponse) ProtoMessage()    {}
func (*QueryDidDocumentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae1fa9bb626e2869, []int{1}
}
func (m *QueryDidDocumentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDidDocumentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDidDocumentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDidDocumentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDidDocumentResponse.Merge(m, src)
}
func (m *QueryDidDocumentResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDidDocumentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDidDocumentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDidDocumentResponse proto.InternalMessageInfo

func (m *QueryDidDocumentResponse) GetDidDocument() DidDocument {
	if m != nil {
		return m.DidDocument
	}
	return DidDocument{}
}

func init() {
	proto.RegisterType((*QueryDidDocumentRequest)(nil), "elestodao.elesto.did.v1.QueryDidDocumentRequest")
	proto.RegisterType((*QueryDidDocumentResponse)(nil), "elestodao.elesto.did.v1.QueryDidDocumentResponse")
}

func init() { proto.RegisterFile("did/v1/query.proto", fileDescriptor_ae1fa9bb626e2869) }

var fileDescriptor_ae1fa9bb626e2869 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x31, 0x4b, 0x03, 0x31,
	0x14, 0xc7, 0x2f, 0x45, 0x05, 0x53, 0x10, 0x09, 0x4a, 0x4b, 0xd1, 0x53, 0x0e, 0x07, 0x15, 0x4c,
	0xbc, 0xba, 0x3a, 0x95, 0x8e, 0x2e, 0x76, 0x74, 0xcb, 0x35, 0x21, 0x0d, 0xb4, 0x79, 0xd7, 0x26,
	0x57, 0x14, 0x71, 0xf1, 0x13, 0x28, 0x8e, 0x7e, 0xa1, 0x8e, 0x05, 0x17, 0x27, 0x91, 0xd6, 0x0f,
	0x22, 0xb9, 0x54, 0x2c, 0x48, 0x07, 0x87, 0xc0, 0xe3, 0xf1, 0xff, 0xfd, 0xc2, 0xff, 0x61, 0x22,
	0xb4, 0x60, 0xe3, 0x94, 0x0d, 0x0b, 0x39, 0xba, 0xa3, 0xf9, 0x08, 0x1c, 0x90, 0x9a, 0xec, 0x4b,
	0xeb, 0x40, 0x70, 0xa0, 0x61, 0xa2, 0x42, 0x0b, 0x3a, 0x4e, 0x1b, 0x7b, 0x0a, 0x40, 0xf5, 0x25,
	0xe3, 0xb9, 0x66, 0xdc, 0x18, 0x70, 0xdc, 0x69, 0x30, 0x36, 0x60, 0x8d, 0xd3, 0x2e, 0xd8, 0x01,
	0x58, 0x96, 0x71, 0x2b, 0x83, 0x8f, 0x8d, 0xd3, 0x4c, 0x3a, 0x9e, 0xb2, 0x9c, 0x2b, 0x6d, 0xca,
	0xf0, 0x22, 0xbb, 0xbd, 0xf8, 0xd6, 0x8b, 0xc3, 0x66, 0x47, 0x81, 0x82, 0x72, 0x64, 0x7e, 0x0a,
	0xdb, 0xe4, 0x04, 0xd7, 0xae, 0xbd, 0xa9, 0xad, 0x45, 0x1b, 0xba, 0xc5, 0x40, 0x1a, 0xd7, 0x91,
	0xc3, 0x42, 0x5a, 0x47, 0xb6, 0x70, 0x45, 0x8b, 0x3a, 0x3a, 0x44, 0xc7, 0x9b, 0x9d, 0x8a, 0x16,
	0x49, 0x0f, 0xd7, 0xff, 0x46, 0x6d, 0x0e, 0xc6, 0x4a, 0x72, 0x85, 0xab, 0xe2, 0x77, 0x5d, 0x42,
	0xd5, 0xe6, 0x11, 0x5d, 0xd1, 0x93, 0x2e, 0x29, 0x5a, 0x6b, 0x93, 0x8f, 0x83, 0xa8, 0xb3, 0x8c,
	0x37, 0x5f, 0x11, 0x5e, 0x2f, 0xbf, 0x22, 0xcf, 0x08, 0x57, 0x97, 0xc2, 0xe4, 0x7c, 0xa5, 0x72,
	0x45, 0x8b, 0x46, 0xfa, 0x0f, 0x22, 0x94, 0x49, 0xf6, 0x1f, 0xdf, 0xbe, 0x5e, 0x2a, 0x35, 0xb2,
	0xcb, 0x02, 0xe0, 0x8f, 0xe8, 0x9f, 0x65, 0xf7, 0x5a, 0x3c, 0xb4, 0x2e, 0x27, 0xb3, 0x18, 0x4d,
	0x67, 0x31, 0xfa, 0x9c, 0xc5, 0xe8, 0x69, 0x1e, 0x47, 0xd3, 0x79, 0x1c, 0xbd, 0xcf, 0xe3, 0xe8,
	0x26, 0x51, 0xda, 0xf5, 0x8a, 0x8c, 0x76, 0x61, 0xb0, 0x40, 0xcf, 0x04, 0x87, 0x1f, 0xcb, 0xad,
	0x77, 0x64, 0x1b, 0xe5, 0xdd, 0x2f, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xc2, 0xcb, 0x62,
	0x18, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// DidDocument queries a did documents with an id.
	DidDocument(ctx context.Context, in *QueryDidDocumentRequest, opts ...grpc.CallOption) (*QueryDidDocumentResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) DidDocument(ctx context.Context, in *QueryDidDocumentRequest, opts ...grpc.CallOption) (*QueryDidDocumentResponse, error) {
	out := new(QueryDidDocumentResponse)
	err := c.cc.Invoke(ctx, "/elestodao.elesto.did.v1.Query/DidDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// DidDocument queries a did documents with an id.
	DidDocument(context.Context, *QueryDidDocumentRequest) (*QueryDidDocumentResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) DidDocument(ctx context.Context, req *QueryDidDocumentRequest) (*QueryDidDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DidDocument not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_DidDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDidDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DidDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elestodao.elesto.did.v1.Query/DidDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DidDocument(ctx, req.(*QueryDidDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "elestodao.elesto.did.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DidDocument",
			Handler:    _Query_DidDocument_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "did/v1/query.proto",
}

func (m *QueryDidDocumentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDidDocumentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDidDocumentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDidDocumentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDidDocumentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDidDocumentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.DidDocument.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryDidDocumentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDidDocumentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.DidDocument.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryDidDocumentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDidDocumentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDidDocumentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryDidDocumentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDidDocumentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDidDocumentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidDocument", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DidDocument.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
